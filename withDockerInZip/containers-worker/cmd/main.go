package main

import (
	"context"
	"github.com/Checkmarx-Containers/containers-worker/internal/cache"
	rabbitMq2 "github.com/Checkmarx-Containers/containers-worker/internal/messages"
	"github.com/Checkmarx-Containers/containers-worker/internal/scanStatus/scan_status_service"
	scanStatusService "github.com/Checkmarx-Containers/containers-worker/internal/scanStatus/service"
	redisV8 "github.com/go-redis/redis/v8"
	"net/http"

	"github.com/Checkmarx-Containers/containers-worker/internal/handlers"
	"github.com/Checkmarx-Containers/containers-worker/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	zbWorker "github.com/checkmarxDev/scans-flow/v8/pkg/zbAbstractionLayer/workflow/worker"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	NewScanWorkerName    = "Start Containers Scan"
	ScanStatusWorkerName = "Containers Scan Status"
)

var (
	zbClient             zbc.Client
	scanStatusGrpcServer scan_status_service.Server
	redisClient          redisV8.UniversalClient
	rabbitMq             *rabbitMq2.RabbitMq
	statusService        *scanStatusService.ScanStatusService
	newScanHandler       *handlers.ScanHandler
	newScanWorker        *worker.Worker
	scanStatusHandler    *handlers.ScanStatusHandler
	scanStatusWorker     *worker.Worker
	healthServer         *http.Server
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed loading configuration")
	}
	ctx := context.Background()

	logger := extendedlogger.GetExtendedLogger()
	zbClient, err = zbWorker.GetZbClient(
		cfg.ZeebeParams.BrokerAddress, cfg.ZeebeParams.AuthServerURL, cfg.ZeebeParams.AuthClientID,
		cfg.ZeebeParams.AuthClientSecret, cfg.ZeebeParams.TLSEnable, cfg.ZeebeParams.TLSCaPath)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to establish connection with Zeebe")
	}
	rabbitMq, err = rabbitMq2.NewRabbitMq(logger, cfg.ContainersRabbit.URL, cfg.ContainersRabbit.Exchange, cfg.ContainersRabbit.RoutingKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create messages publisher")
	}

	redisClient, err = cache.New(cache.Config{
		IsClusterMode: cfg.Redis.IsClusterMode,
		Addr:          cfg.Redis.Addr,
		Password:      cfg.Redis.Password,
		MaxRetries:    cfg.Redis.MaxRetries,
		TLSEnabled:    cfg.Redis.TLSEnabled,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Redis")
	}
	statusService = scanStatusService.NewScanStatusService(logger, redisClient)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create scan status service")
	}

	newScanHandler, err = handlers.NewScanHandler(logger, rabbitMq, redisClient, statusService)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new scan handler")
	}

	newScanWorker, err = worker.NewWorker(ctx, zbClient, NewScanWorkerName,
		cfg.NewContainersScanWorkerConfig.TimeoutMinutes, cfg.NewContainersScanWorkerConfig.MaxConcurrentJobs,
		cfg.NewContainersScanWorkerConfig.JobType, cfg.ZeebeParams.Rabbit.URL, cfg.ZeebeParams.Rabbit.Exchange, cfg.ZeebeParams.Rabbit.SkipTLSCertVerify,
		cfg.ZeebeParams.Rabbit.Retries, time.Minute, newScanHandler)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Zeebe worker")
	}

	scanStatusHandler, err = handlers.NewScanStatusHandler(logger, redisClient, statusService)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new scan handler")
	}

	scanStatusWorker, err = worker.NewWorker(ctx, zbClient, ScanStatusWorkerName,
		cfg.ScanStatusWorkerConfig.TimeoutMinutes, cfg.ScanStatusWorkerConfig.MaxConcurrentJobs,
		cfg.ScanStatusWorkerConfig.JobType, cfg.ZeebeParams.Rabbit.URL, cfg.ZeebeParams.Rabbit.Exchange, cfg.ZeebeParams.Rabbit.SkipTLSCertVerify,
		cfg.ZeebeParams.Rabbit.Retries, time.Minute, scanStatusHandler)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Zeebe worker")
	}

	scanStatusGrpcServer = scan_status_service.NewScanStatusGRPCServer(
		cfg.Ports.ScanStatus,
		statusService, logger,
	)

	healthServer = InitializeHealthServer(*cfg)
	log.Info().Msg("Containers Worker Started...")

	listenForErrors(ctx)
	closeResourcesGracefully()
}

func listenForErrors(ctx context.Context) {
	errChan := make(chan error, 1)
	go func() {
		errChan <- newScanWorker.ListenAndServe()
	}()

	go func() {
		errChan <- scanStatusWorker.ListenAndServe()
	}()

	go func() {
		errChan <- scanStatusGrpcServer.ListenAndServe()
	}()

	go func() {
		errChan <- healthServer.ListenAndServe()
	}()

	wait(ctx, errChan)
}

func wait(ctx context.Context, errChan chan error) {
	signalCh := make(chan os.Signal, 1)                                                       // os signals channel
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM) // reload, ctrl+c, quit
	// wait for program to finish
	select {
	case err := <-errChan:
		log.Err(err).Msg("Error while running containers worker HTTP web services")
	case sig := <-signalCh:
		log.Warn().Msgf("Got signal from the operation system. peacefully exiting. signal=%s", sig.String())
	case <-ctx.Done():
		log.Err(ctx.Err()).Msg("main context ended. peacefully exiting")
	}
}

// closeResourcesGracefully gracefully closes the resources that need to be closed before exiting
func closeResourcesGracefully() {
	err := zbClient.Close()
	if err != nil {
		log.Err(err).Msg("error closing zeebe client")
	}

	err = scanStatusGrpcServer.Close()
	if err != nil {
		log.Err(err).Msg("error closing scan status grpc server")
	}

	err = redisClient.Close()
	if err != nil {
		log.Err(err).Msg("error closing redis client")
	}

	err = newScanWorker.Close()
	if err != nil {
		log.Err(err).Msg("error closing new scan worker")
	}

	err = scanStatusWorker.Close()
	if err != nil {
		log.Err(err).Msg("error closing scan status worker")
	}

	err = healthServer.Close()
	if err != nil {
		log.Err(err).Msg("error closing health server")
	}
}

func LoadConfig() (*Config, error) {
	err := setConfigDefaults()
	if err != nil {
		return nil, errors.Wrap(err, "setting config values")
	}

	logLevel, err := zerolog.ParseLevel(strings.ToLower(viper.GetString(logLevelEnvField)))
	if err != nil {
		return nil, errors.Wrap(err, "error parsing log level")
	}

	return &Config{
		LogLevel: logLevel,
		Ports: ports{
			ScanStatus: viper.GetUint(scanStatusServicePortEnv),
			Health:     viper.GetUint(healthServicePortEnv),
			Server:     viper.GetUint(serverPortEnvField),
		},
		ZeebeParams: zeebeParams{
			BrokerAddress:    viper.GetString(workflowZeebeBrokerAddressEnv),
			TLSEnable:        viper.GetBool(workflowZeebeTLSEnableEnv),
			TLSCaPath:        viper.GetString(workflowZeebeTLSCaPathEnv),
			AuthServerURL:    viper.GetString(workflowZeebeAuthServerURLEnv),
			AuthClientID:     viper.GetString(workflowZeebeAuthClientIDEnv),
			AuthClientSecret: viper.GetString(workflowZeebeAuthClientSecretEnv),
			Namespace:        viper.GetString(workflowZeebeNamespace),
			Rabbit: rabbit{
				URL:               viper.GetString(rabbitURLEnv),
				SkipTLSCertVerify: viper.GetBool(rabbitTLSSkipVerifyEnv),
				Exchange:          viper.GetString(rabbitExchangeEnv),
				Retries:           viper.GetInt(rabbitReconnectRetriesEnv),
				RetryWaitDuration: viper.GetDuration(rabbitRetryWaitDurationEnv),
			},
		},
		Redis: redis{
			Addr:               viper.GetString(redisAddressEnv),
			Password:           viper.GetString(redisPasswordEnv),
			KeyExpirationHours: viper.GetUint(redisKeyExpirationHoursEnv),
			MaxRetries:         viper.GetInt(redisMaxRetriesEnv),
			IsClusterMode:      viper.GetBool(redisIsClusterModeEnv),
			TLSEnabled:         viper.GetBool(redisTlsEnabledEnv),
			SkipVerify:         false,
		},
		NewContainersScanWorkerConfig: workerConfig{
			JobType:                 viper.GetString(newContainersScanJobTypeEnv),
			TimeoutMinutes:          viper.GetUint(newContainersScanTimeoutMinutesEnv),
			MaxConcurrentJobs:       viper.GetUint(newContainersScanMaxConcurrentJobsEnv),
			ReturnJobTimeoutMinutes: viper.GetUint(newContainersScanReturnJobTimeoutMinutesEnv),
		},
		ScanStatusWorkerConfig: workerConfig{
			JobType:                 viper.GetString(scanStatusJobTypeEnv),
			TimeoutMinutes:          viper.GetUint(scanStatusTimeoutMinutesEnv),
			MaxConcurrentJobs:       viper.GetUint(scanStatusMaxConcurrentJobsEnv),
			ReturnJobTimeoutMinutes: viper.GetUint(scanStatusReturnJobTimeoutMinutesEnv),
		},
		Maintainer: maintainerConfig{},
		ContainersRabbit: rabbit{
			URL:               viper.GetString(rabbitURLEnv),
			SkipTLSCertVerify: viper.GetBool(rabbitTLSSkipVerifyEnv),
			Exchange:          viper.GetString(containersRabbitExchangeEnv),
			QueueName:         viper.GetString(containersRabbitContainersScanQueueNameEnv),
			RoutingKey:        viper.GetString(containersRabbitContainersScansRoutingKeyNameEnv),
			Retries:           viper.GetInt(containersRabbitReconnectRetriesEnv),
			RetryWaitDuration: viper.GetDuration(containersRabbitRetryWaitDurationEnv),
		},
	}, nil
}
