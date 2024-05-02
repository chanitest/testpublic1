package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type Input struct {
	Engine  Engine  `json:"engine"`
	Tenant  string  `json:"tenant"`
	General General `json:"general"`
	// ... other fields as needed
}

type Engine struct {
	Containers Containers `json:"containers"`
	// ... other fields as needed
}

type Containers struct {
	Step          Step   `json:"step"`
	FailureCode   int    `json:"failureCode"`
	FailureReason string `json:"failureReason"`
	Failed        bool   `json:"failed"`
	Finished      bool   `json:"finished"`
	// ... other fields as needed
}

type Step struct {
	StartContainersScan  ContainersScanStep `json:"start_containers_scan"`
	ContainersScanStatus ContainersScanStep `json:"containers_scan_status"`
	// ... other fields as needed
}

type ContainersScanStep struct {
	Finished bool `json:"finished"`
	// ... other fields as needed
}

type General struct {
	Finished bool `json:"finished"`
	Step     Step `json:"step"`
	// ... other fields as needed
}

type ScanStatus struct {
	Id                     string                 `json:"ScanId"`
	Status                 string                 `json:"Status"`
	StepStatus             []ScanStepStatus       `json:"StepStatus"`
	WorkflowCorrelationKey string                 `json:"WorkflowCorrelationKey"`
	CorrelationId          string                 `json:"CorrelationId"`
	UpdateTime             *timestamppb.Timestamp `json:"UpdateTime"`
	Seen                   bool                   `json:"Seen"`
	Error                  string                 `json:"Error"`
}

type ScanStepStatus struct {
	Status     string                 `json:"Status"`
	StepName   string                 `json:"StepName"`
	UpdateTime *timestamppb.Timestamp `json:"UpdateTime"`
	Seen       bool                   `json:"Seen"`
}

type TestResult struct {
	Result string `json:"result"`
}

// containers engine status
const (
	ContainersTestResultKeyPrefix = "ContainersTestResult"

	ContainersScanStatusQueued    = "Queued"
	ContainersScanStatusWorking   = "Working"
	ContainersScanStatusCompleted = "Success"
	ContainersScanStatusFailed    = "Failed"

	ContainersScanStatusCanceled = "Canceled"
	ContainersScanStatusTimeOut  = "Timeout"
)

var (
	redisClient redis.Client
)

func SendMessageToZeebe(client zbc.Client, ctx context.Context, message map[string]interface{}) error {
	request, err := client.NewCreateInstanceCommand().BPMNProcessId("containers-scan").LatestVersion().VariablesFromMap(message)
	if err != nil {
		return err
	}

	ctx = context.Background()

	msg, err := request.Send(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Sent Message to zeebe: %+v", msg)
	return nil
}

func DeployZeebeWorkflow(ctx context.Context, bpmnFilePath string) (zbc.Client, error) {

	add := viper.GetString("WORKFLOW_ZEEBE_BROKER_ADDRESS")

	log.Println("WORKFLOW_ZEEBE_BROKER_ADDRESS: " + add)

	var err error

	client, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         add,
		UsePlaintextConnection: true,
	})
	if err != nil {
		panic(err)
	}

	response1, err := client.NewDeployResourceCommand().AddResourceFile(bpmnFilePath).Send(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Zeebe deployment response message: %+v", response1))

	return client, err
}

func InitializeRedisClient() {

	redisAddress := viper.GetString("REDIS_ADDRESS")

	log.Println("REDIS_ADDRESS: " + redisAddress)

	redisClient = *redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Could not reach redis server")
	}
}

func GetStatusFromRedis(scanID string) (string, error) {
	key := fmt.Sprintf("%s:%s", ContainersTestResultKeyPrefix, scanID)

	// Retrieve the value from Redis
	result, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("Failed to get value from Redis:", err)
		return "", err
	}

	// Unmarshal the JSON value into a ScanStatus struct
	var status TestResult
	err = json.Unmarshal([]byte(result), &status)
	if err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return "", err
	}

	if status.Result != ContainersScanStatusQueued {
		log.Printf("Status for scan ID %s indicates test has finished. Current status: %s", scanID, status.Result)
	} else {
		log.Printf("Status for scan ID %s is %s", scanID, status.Result)
	}
	return status.Result, nil
}

func PublishTestResultToRedis(scanId, testStatus string) {
	key := fmt.Sprintf("%s:%s", ContainersTestResultKeyPrefix, scanId)

	result := TestResult{
		Result: testStatus,
	}

	// Marshal the extended struct to JSON
	jsonMessage, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed to marshal message to JSON:", err)
		return
	}

	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis client can not be reached")
	}

	log.Printf("key: %s, pong: %s", key, pong)

	_, err = redisClient.Set(context.Background(), key, jsonMessage, 0).Result()
	if err != nil {
		log.Println("Failed to publish message to Redis:", err)
		return
	}

	fmt.Println("Published message to Redis:", string(jsonMessage))
}

func GetMessageData(scanId, workflowCorrelationId string) map[string]interface{} {
	jsonData := fmt.Sprintf(`{
	    "workflowCorrelationKey": "%s",
	    "correlationId":"%s",
	    "tenantName": "testing_na",
	    "tenant": "b4627f2b-80da-4b6c-a806-2e3a34794bbb",
	    "runtimeVars": {
	        "runContainersEnginesScan": true,
	        "workflowTimeout": "PT2880M",
	        "jobTimeout": "PT0M",
	        "longJobTimeout": "PT0M",
	        "shortJobTimeout": "PT15M",
	        "errorMessage": "",
	        "jobRetries": 3,
	        "longJobRetries": 0,
	        "shortJobRetries": 3,
	        "sastLoc": 0
	    },
	    "licenseMetadata": {
	        "maxConcurrentScans": "30",
	        "allowedEngines": {
	            "containers": true
	        }
	    },
	    "input": {
	        "@type": "scans.Scan",
	        "id": "%s",
	        "project": {
	            "id": "37b5301a-c85c-4074-9525-6cb0b7c2d2ab",
	            "name": "",
	            "tags": {}
	        },
	        "configs": [
	            {
	                "type": "containers",
	                "value": {
	                    "someKey": "someValue"
	                }
	            }
	        ],
	        "tags": {},
	        "createdAt": "2024-01-05T16:57:42.508387079Z",
	        "type": "upload",
	        "uploadHandler": {
	            "repoUrl": "",
	            "branch": "master",
	            "uploadUrl": "https://ast.checkmarx.net/storage/uploads.n-virginia.us-east-1-602005780816/b4627f2b-80da-4b6c-a806-2e3a34794bbb/blabla"
	        }
	    },
	    "general": {
	        "step": {
	            "fetch_sources": {
	                "finished": true
	            }
	        },
	        "finished": true
	    },
	    "engine": {}
	}`, workflowCorrelationId, workflowCorrelationId, scanId)

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Printf("Message Sent to Zeebe: %#v", data)
	return data
}
