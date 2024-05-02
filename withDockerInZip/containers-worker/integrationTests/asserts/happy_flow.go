package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
	g "tests/grpc"
)

const (
	HappyFlowScanId        = "3799f608-4c75-4d8f-8792-abdbb2db217a"
	HappyFlowCorrelationId = "123455"
)

func CheckHappyFlow(client worker.JobClient, job entities.Job) {
	defer GinkgoRecover()
	jobKey := job.GetKey()
	log.Println("Fetching items for happy", jobKey)

	var input Input
	err := json.Unmarshal([]byte(job.Variables), &input)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Extract the required fields
	engineObject := input.Engine
	containersField := engineObject.Containers

	// Create a new struct with the extracted values
	resultObject := Engine{
		Containers: containersField,
	}

	fmt.Printf("happy flow result Object = %+v\n", resultObject)

	_, err = client.NewCompleteJobCommand().JobKey(jobKey).Send(context.Background())
	if err != nil {
		return
	}

	res := Expect(resultObject.Containers.Finished).To(Equal(true))

	if res {
		fmt.Println("Passed happy flow")
		PublishTestResultToRedis(HappyFlowScanId, ContainersScanStatusCompleted)
	} else {
		fmt.Println("Failed happy flow")
		PublishTestResultToRedis(HappyFlowScanId, ContainersScanStatusFailed)
	}

}

func HandleRabbitMQMessageHappyFlow(body []byte) {

	var message map[string]interface{}
	err := json.Unmarshal(body, &message)
	if err != nil {
		log.Println("Failed to parse RabbitMQ message:", err)
		return
	}

	fmt.Println("Received RabbitMQ message:", message)

	err = g.SendStatusUpdate(context.Background(), HappyFlowScanId, ContainersScanStatusCompleted, HappyFlowCorrelationId)
	if err != nil {
		return
	}
	log.Println("Successfully send success status update")
}
