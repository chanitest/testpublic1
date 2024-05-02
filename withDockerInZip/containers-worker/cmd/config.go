package main

import (
	"errors"
	"github.com/checkmarxDev/scans-flow/v8/pkg/zbAbstractionLayer/workflow/worker"
	"github.com/checkmarxDev/scans/pkg/events/cloudevents/sender"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"time"
)

type zeebeParams struct {
	BrokerAddress    string
	TLSEnable        bool
	TLSCaPath        string
	AuthServerURL    string
	AuthClientID     string
	AuthClientSecret string
	Namespace        string
	Rabbit           rabbit
}

type ports struct {
	ScanStatus uint
	Health     uint
	Server     uint
}

type rabbit struct {
	URL               string
	SkipTLSCertVerify bool
	Exchange          string
	QueueName         string
	RoutingKey        string
	Retries           int
	RetryWaitDuration time.Duration
}

type redis struct {
	Addr               string
	Password           string
	KeyExpirationHours uint
	MaxRetries         int
	IsClusterMode      bool
	TLSEnabled         bool
	SkipVerify         bool
}

type workerConfig struct {
	JobType                 string
	TimeoutMinutes          uint
	MaxConcurrentJobs       uint
	ReturnJobTimeoutMinutes uint
}

type maintainerConfig struct {
	Disable           bool
	KeyExpiration     time.Duration
	JobMaxDuplication int
	Retries           int
	RetriesInterval   time.Duration
}

type Config struct {
	LogLevel                      zerolog.Level
	Ports                         ports
	ZeebeParams                   zeebeParams
	Redis                         redis
	NewContainersScanWorkerConfig workerConfig
	ScanStatusWorkerConfig        workerConfig
	Maintainer                    maintainerConfig
	ContainersRabbit              rabbit
}

func setConfigDefaults() error {
	setDefaultConfigValues()
	viper.AutomaticEnv()
	viper.AddConfigPath(".")

	viper.SetDefault(configFileNameEnvField, configFileNameDefault)
	configFileName := viper.GetString(configFileNameEnvField)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Str("filename", configFileName).Msg("Config file no set")
		} else {

		}
	}

	workflowZeebeAddr := viper.GetString(workflowZeebeBrokerAddressEnv)
	if workflowZeebeAddr == "" {
		return errors.New("empty name")
	}
	workflowZeebeStartScanJobType := viper.GetString(newContainersScanJobTypeEnv)
	if workflowZeebeStartScanJobType == "" {
		return errors.New("empty name")
	}

	workflowZeebeCheckStatusJobType := viper.GetString(scanStatusJobTypeEnv)
	if workflowZeebeCheckStatusJobType == "" {
		return errors.New("empty name")
	}

	rabbitConnectionString := viper.GetString(rabbitURLEnv)
	if rabbitConnectionString == "" {
		return errors.New("empty name")
	}

	return nil
}

func setDefaultConfigValues() {
	viper.SetDefault(logLevelEnvField, logLevelDefault)

	viper.SetDefault(workflowZeebeBrokerAddressEnv, "localhost:26500")
	viper.SetDefault(workflowZeebeProcessIDEnv, "containersWorker")
	viper.SetDefault(workflowZeebeTLSEnableEnv, false)
	viper.SetDefault(workflowZeebeMaxJobEnv, worker.MaxConcurrentJobsDefault)
	viper.SetDefault(workflowZeebeNamespace, "default")
	viper.SetDefault(workflowZeebeTimeoutMinutesEnv, workTimeoutMinutesDefault)
	viper.SetDefault(workflowZeebeMaxStatusJobEnv, worker.MaxConcurrentJobsDefault)

	viper.SetDefault(scanStatusWorkflowZeebeTimeoutMinutesEnv, scanStatusWorkTimeoutMinutesDefault)
	viper.SetDefault(rabbitURLEnv, containersRabbitURLDefault)
	viper.SetDefault(rabbitTLSSkipVerifyEnv, rabbitTLSSkipVerifyDefault)
	viper.SetDefault(rabbitExchangeEnv, rabbitExchangeDefault)
	viper.SetDefault(rabbitReconnectRetriesEnv, 25) //nolint
	viper.SetDefault(rabbitRetryWaitDurationEnv, sender.DefaultRetryWaitDuration)

	viper.SetDefault(redisAddressEnv, RedisAddressDefault)
	viper.SetDefault(redisPasswordEnv, RedisPasswordDefault)
	viper.SetDefault(redisKeyExpirationHoursEnv, RedisKeyExpirationHoursDefault)
	viper.SetDefault(redisMaxRetriesEnv, RedisMaxRetriesDefault)
	viper.SetDefault(redisIsClusterModeEnv, RedisIsClusterModeDefault)
	viper.SetDefault(redisTlsEnabledEnv, RedisTlsEnabledDefault)

	viper.SetDefault(scanStatusServicePortEnv, scanStatusServicePortDefault)
	viper.SetDefault(healthServicePortEnv, healthServicePortDefault)
	viper.SetDefault(serverPortEnvField, serverPortDefault)

	viper.SetDefault(newContainersScanJobTypeEnv, newContainersScanJobTypeDefault)
	viper.SetDefault(newContainersScanTimeoutMinutesEnv, newContainersScanTimeoutMinutesDefault)
	viper.SetDefault(newContainersScanMaxConcurrentJobsEnv, newContainersScanMaxConcurrentJobsDefault)
	viper.SetDefault(newContainersScanReturnJobTimeoutMinutesEnv, newContainersScanReturnJobTimeoutMinutesDefault)

	viper.SetDefault(scanStatusJobTypeEnv, scanStatusJobTypeDefault)
	viper.SetDefault(scanStatusTimeoutMinutesEnv, scanStatusTimeoutMinutesDefault)
	viper.SetDefault(scanStatusMaxConcurrentJobsEnv, scanStatusMaxConcurrentJobsDefault)
	viper.SetDefault(scanStatusReturnJobTimeoutMinutesEnv, scanStatusReturnJobTimeoutMinutesDefault)

	viper.SetDefault(containersRabbitURLEnv, containersRabbitURLDefault)
	viper.SetDefault(containersRabbitTLSSkipVerifyEnv, containersRabbitTLSSkipVerifyDefault)
	viper.SetDefault(containersRabbitExchangeEnv, containersRabbitExchangeDefault)

	viper.SetDefault(containersRabbitContainersScanQueueNameEnv, containersRabbitContainersScanQueueNameDefault)
	viper.SetDefault(scanStatusReturnJobTimeoutMinutesEnv, scanStatusReturnJobTimeoutMinutesDefault)
	viper.SetDefault(containersRabbitContainersScansRoutingKeyNameEnv, containersRabbitContainersScansRoutingKeyNameDefault)
	viper.SetDefault(containersRabbitReconnectRetriesEnv, containersRabbitReconnectRetriesDefault)
	viper.SetDefault(containersRabbitRetryWaitDurationEnv, containersRabbitRetryWaitDurationDefault)
	viper.SetDefault(serviceNameEnv, "containers-worker")
}

const (
	serviceNameEnv         = "ServiceName"
	logLevelEnvField       = "LOG_LEVEL"
	configFileNameEnvField = "CONFIG_ENV_FILE_NAME"

	//ast redis config
	redisAddressEnv            = "REDIS_ADDRESS"
	redisPasswordEnv           = "REDIS_PASSWORD"
	redisKeyExpirationHoursEnv = "REDIS_TTL_IN_HOURS"
	redisMaxRetriesEnv         = "REDIS_MAX_RETRIES"
	redisIsClusterModeEnv      = "REDIS_IS_CLUSTER_MODE"
	redisTlsEnabledEnv         = "REDIS_TLS_ENABLED"

	//ast rabbitmq config
	rabbitURLEnv               = "RABBIT_CONNECTION_STRING"
	rabbitTLSSkipVerifyEnv     = "RABBIT_TLS_SKIP_VERIFY"
	rabbitExchangeEnv          = "RABBIT_EXCHANGE_TOPIC"
	rabbitReconnectRetriesEnv  = "RABBIT_RECONNECT_RETRIES"
	rabbitRetryWaitDurationEnv = "RABBIT_RETRY_WAIT_DURATION"

	//zeebe config
	workflowZeebeBrokerAddressEnv            = "WORKFLOW_ZEEBE_BROKER_ADDRESS"
	workflowZeebeTimeoutMinutesEnv           = "WORKFLOW_TIMEOUT_MIN"
	scanStatusWorkflowZeebeTimeoutMinutesEnv = "SCAN_STATUS_WORKFLOW_TIMEOUT_MIN"
	workflowZeebeTLSEnableEnv                = "WORKFLOW_ZEEBE_TLS_ENABLE"
	workflowZeebeTLSCaPathEnv                = "WORKFLOW_ZEEBE_TLS_CA_PATH"
	workflowZeebeAuthServerURLEnv            = "WORKFLOW_ZEEBE_AUTH_SERVER_URL"
	workflowZeebeAuthClientIDEnv             = "WORKFLOW_ZEEBE_AUTH_CLIENT_ID"
	workflowZeebeAuthClientSecretEnv         = "WORKFLOW_ZEEBE_AUTH_CLIENT_SECRET" //nolint
	workflowZeebeMaxJobEnv                   = "WORKFLOW_ZEEBE_MAX_CONCURRENT_JOBS"
	workflowZeebeMaxStatusJobEnv             = "WORKFLOW_ZEEBE_MAX_CONCURRENT_STATUS_JOBS"
	workflowZeebeNamespace                   = "WORKFLOW_ZEEBE_NAMESPACE"
	workflowZeebeProcessIDEnv                = "WORKFLOW_ZEEBE_PROCESS_ID"

	//NewContainersScan Worker Config
	newContainersScanJobTypeEnv                 = "CONTAINERS_NEW_SCAN_JOB_TYPE"
	newContainersScanTimeoutMinutesEnv          = "CONTAINERS_NEW_SCAN_TIMEOUT_IN_MINUTES"
	newContainersScanMaxConcurrentJobsEnv       = "CONTAINERS_NEW_SCAN_MAX_CONCURRENT_JOBS"
	newContainersScanReturnJobTimeoutMinutesEnv = "CONTAINERS_NEW_SCAN_RETURN_JOB_TIMEOUT_IN_MINUTES"

	//ScanStatus Worker Config
	scanStatusJobTypeEnv                 = "CONTAINERS_SCAN_STATUS_JOB_TYPE"
	scanStatusTimeoutMinutesEnv          = "CONTAINERS_SCAN_STATUS_TIMEOUT_IN_MINUTES"
	scanStatusMaxConcurrentJobsEnv       = "CONTAINERS_SCAN_STATUS_MAX_CONCURRENT_JOBS"
	scanStatusReturnJobTimeoutMinutesEnv = "CONTAINERS_SCAN_STATUS_RETURN_JOB_TIMEOUT_IN_MINUTES"

	//ports
	scanStatusServicePortEnv = "SCAN_STATUS_SERVICE_PORT"
	healthServicePortEnv     = "HEALTH_SERVICE_PORT"
	serverPortEnvField       = "SERVER_PORT"

	//containers engine messages configs
	containersRabbitURLEnv                           = "CONTAINERS_RABBIT_CONNECTION_STRING"
	containersRabbitTLSSkipVerifyEnv                 = "CONTAINERS_RABBIT_TLS_SKIP_VERIFY"
	containersRabbitExchangeEnv                      = "CONTAINERS_RABBIT_EXCHANGE_TOPIC"
	containersRabbitContainersScanQueueNameEnv       = "CONTAINERS_RABBIT_INITIALIZE_SCAN_QUEUE"
	containersRabbitContainersScansRoutingKeyNameEnv = "CONTAINERS_RABBIT_ROUTING_KEY"
	containersRabbitReconnectRetriesEnv              = "CONTAINERS_RABBIT_RECONNECT_RETRIES"
	containersRabbitRetryWaitDurationEnv             = "CONTAINERS_RABBIT_RETRY_WAIT_DURATION"
)

// defaults
const (
	logLevelDefault       = "debug"
	configFileNameDefault = "config"

	//zeebe
	workTimeoutMinutesDefault           = "10"
	scanStatusWorkTimeoutMinutesDefault = "1"
	rabbitTLSSkipVerifyDefault          = true
	rabbitExchangeDefault               = "/exchange/amq.topic/"

	//cache
	RedisAddressDefault            = "localhost:6379"
	RedisPasswordDefault           = ""
	RedisKeyExpirationHoursDefault = 1
	RedisMaxRetriesDefault         = 1
	RedisIsClusterModeDefault      = false
	RedisTlsEnabledDefault         = false

	//ports
	scanStatusServicePortDefault = 4568
	healthServicePortDefault     = 4321
	serverPortDefault            = "80"

	//NewContainersScan Worker Config
	newContainersScanJobTypeDefault                 = "containers-scan-initiator"
	newContainersScanTimeoutMinutesDefault          = 2
	newContainersScanMaxConcurrentJobsDefault       = 10
	newContainersScanReturnJobTimeoutMinutesDefault = 2

	//ScanStatus Worker Config
	scanStatusJobTypeDefault                 = "containers-scan-status"
	scanStatusTimeoutMinutesDefault          = 2
	scanStatusMaxConcurrentJobsDefault       = 10
	scanStatusReturnJobTimeoutMinutesDefault = 2

	//containers rabbitmq defaults
	containersRabbitURLDefault                           = "amqp://guest:guest@localhost:5672/"
	containersRabbitTLSSkipVerifyDefault                 = true
	containersRabbitExchangeDefault                      = "containers.topic"
	containersRabbitContainersScanQueueNameDefault       = "initialize-scan"
	containersRabbitContainersScansRoutingKeyNameDefault = "containers.scan.initialize-scan"
	containersRabbitReconnectRetriesDefault              = 25
	containersRabbitRetryWaitDurationDefault             = sender.DefaultRetryWaitDuration
)
