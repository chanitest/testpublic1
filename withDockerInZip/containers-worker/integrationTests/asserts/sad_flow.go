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
	SadFlowScanId        = "51d9f608-7z75-438f-8592-abdbb1db217b"
	SadFlowCorrelationId = "12123"
)

func CheckSadFlow(client worker.JobClient, job entities.Job) {
	defer GinkgoRecover()
	jobKey := job.GetKey()
	log.Println("Fetching items for sad", jobKey)

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

	fmt.Printf("sad flow result Object = %+v\n", resultObject)

	_, err = client.NewCompleteJobCommand().JobKey(jobKey).Send(context.Background())
	if err != nil {
		return
	}

	res := Expect(resultObject.Containers.Finished).To(Equal(false))
	if res {
		fmt.Println("Passed sad flow")
		PublishTestResultToRedis(SadFlowScanId, ContainersScanStatusCompleted)
	} else {
		fmt.Println("Failed sad flow")
		PublishTestResultToRedis(SadFlowScanId, ContainersScanStatusFailed)
	}
}

func HandleRabbitMQMessageSadFlow(body []byte) {
	var message map[string]interface{}
	err := json.Unmarshal(body, &message)
	if err != nil {
		log.Println("Failed to parse RabbitMQ message:", err)
		return
	}

	fmt.Println("Received RabbitMQ message:", message)

	err = g.SendStatusUpdate(context.Background(), SadFlowScanId, ContainersScanStatusFailed, SadFlowCorrelationId)
	if err != nil {
		return
	}
	log.Println("Successfully sent failed status update")
}
