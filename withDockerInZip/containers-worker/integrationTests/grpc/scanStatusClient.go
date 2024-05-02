package scan_status_clinet

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	a "tests/grpc/scan_status"
)

var (
	grpcClient a.ScanStatusServiceClient
	grpcConn   *grpc.ClientConn
)

func InitializeGrpcClient() {
	grpcServerAddress := viper.GetString("GRPC_SERVER_ADDRESS")

	log.Println("GRPC_SERVER_ADDRESS: " + grpcServerAddress)

	conn, err := grpc.Dial(grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to grpc server: %s", err)
	}
	grpcConn = conn

	// Create a gRPC client
	grpcClient = a.NewScanStatusServiceClient(conn)
}

func SendStatusUpdate(ctx context.Context, scanId, status, correlationId string) error {
	scanStatusObject := &a.UpdateScanStatusRequest{
		Id:            scanId,
		Status:        status,
		CorrelationId: correlationId,
		UpdateTime:    timestamppb.Now(),
	}

	scanStatus, err := grpcClient.UpdateScanStatus(ctx, scanStatusObject)
	if err != nil {
		log.Fatalf("Could not update scan status. error: %s", err)
		return err
	}

	if scanStatus.GetUpdateStatus() != "SUCCESS" {
		return fmt.Errorf("could not update scan status. returned status: %s", scanStatus.GetUpdateStatus())
	}
	return nil
}
