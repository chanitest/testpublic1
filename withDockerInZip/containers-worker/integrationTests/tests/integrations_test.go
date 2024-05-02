package tests

import (
	"context"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"log"
	"sync"
	"testing"
	a "tests/asserts"
	g "tests/grpc"
	q "tests/queue"
	"time"
)

var (
	client zbc.Client
	wg     sync.WaitGroup
)

func TestContainersScan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containers Suite")

	viper.SetDefault("REDIS_ADDRESS", "localhost:6379")
	viper.SetDefault("BPMN_FILE_LOCATION", "/Users/danielgreenspan/GolandProjects/containers-worker/integrationTests/containers-scan.bpmn")
	viper.SetDefault("WORKFLOW_ZEEBE_BROKER_ADDRESS", "localhost:26500")
	viper.SetDefault("RABBITMQ_ADDRESS", q.RabbitMQURL)
	viper.SetDefault("REDIS_ADDRESS", "localhost:6379")
	viper.SetDefault("GRPC_SERVER_ADDRESS", "localhost:5000")
}

var _ = Describe("Containers", func() {

	Describe("Running containers scan", Ordered, func() {

		BeforeAll(func() {
			viper.AutomaticEnv()
			a.InitializeRedisClient()
			g.InitializeGrpcClient()

			fileLocation := viper.GetString("BPMN_FILE_LOCATION")

			log.Println("BPMN_FILE_LOCATION: " + fileLocation)

			client, _ = a.DeployZeebeWorkflow(context.Background(), fileLocation)
		})

		ReportAfterEach(func(report SpecReport) {
			customFormat := fmt.Sprintf("%s | %s", report.State, report.FullText())
			log.Printf(customFormat)
		})

		Context("valid input", func() {
			It("should finish scan successfully", func() {

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				wg.Add(1)

				go func() {
					defer wg.Done()
					q.ListenToContainersEngineQueue(ctx, a.HandleRabbitMQMessageHappyFlow)
				}()

				AddAssertionWorker(client, "asserter", "asserter-worker", a.CheckHappyFlow)
				a.PublishTestResultToRedis(a.HappyFlowScanId, a.ContainersScanStatusQueued)

				err := a.SendMessageToZeebe(client, ctx, a.GetMessageData(a.HappyFlowScanId, a.HappyFlowCorrelationId))
				if err != nil {
					panic(err)
				}

				PollForStatus(a.HappyFlowScanId)

				cancel()
				wg.Wait()
			})
		})
		Context("invalid input", func() {
			It("should not finish scan successfully", func() {

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				wg.Add(1)

				go func() {
					defer wg.Done()
					q.ListenToContainersEngineQueue(ctx, a.HandleRabbitMQMessageSadFlow)
				}()

				AddAssertionWorker(client, "asserter", "asserter-worker", a.CheckSadFlow)
				a.PublishTestResultToRedis(a.SadFlowScanId, a.ContainersScanStatusQueued)

				err := a.SendMessageToZeebe(client, ctx, a.GetMessageData(a.SadFlowScanId, a.SadFlowCorrelationId))
				if err != nil {
					panic(err)
				}

				PollForStatus(a.SadFlowScanId)

				cancel()
				wg.Wait()
			})
		})
	})
})

func PollForStatus(scanId string) {
	status := a.ContainersScanStatusQueued
	var err error

	for time.Since(time.Now()) < time.Minute {
		status, err = a.GetStatusFromRedis(scanId)
		if err != nil {
			log.Println("Error getting status:", err)
			time.Sleep(time.Second) // Sleep before retrying
			continue
		}

		if status != a.ContainersScanStatusQueued {
			log.Printf("Final status for scan ID %s: %s", scanId, status)
			break
		}
		time.Sleep(time.Second)
	}

	Expect(status).To(Equal(a.ContainersScanStatusCompleted))
}

func AddAssertionWorker(client zbc.Client, workerType, workerName string,
	assertFunc func(client worker.JobClient, job entities.Job)) worker.JobWorker {
	return client.NewJobWorker().JobType(workerType).Handler(assertFunc).Name(workerName).Open()
}
