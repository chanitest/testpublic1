package scan_status_service

import (
	"context"
	"fmt"
	scanStatusService "github.com/Checkmarx-Containers/containers-worker/internal/scanStatus/service"
	pb "github.com/Checkmarx-Containers/containers-worker/pkg/api/scan_status"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	ListenAndServe() error
	Close() error

	UpdateScanStatus(ctx context.Context, in *pb.UpdateScanStatusRequest) (*pb.UpdateScanStatusResponse, error)
	UpdateScanStatusDetails(ctx context.Context, in *pb.UpdateScanStatusDetailsRequest) (*pb.UpdateScanStatusResponse, error)
	GetScanStatusDetails(ctx context.Context, in *pb.GetScanStatusDetailsRequest) (*pb.GetScanStatusDetailsResponse, error)
}

type grpcServer struct {
	port   uint
	server *grpc.Server
	pb.UnimplementedScanStatusServiceServer
	scanStatusService scanStatusService.ScanStatusServiceInt
	logger            extendedlogger.ExtendedLogger
}

func NewScanStatusGRPCServer(port uint, scanStatusService scanStatusService.ScanStatusServiceInt, logger *extendedlogger.ExtendedLogger) Server {
	s := &grpcServer{
		port:              port,
		server:            grpc.NewServer(),
		scanStatusService: scanStatusService,
		logger:            *logger,
	}
	pb.RegisterScanStatusServiceServer(s.server, s)
	return s
}

func (s *grpcServer) ListenAndServe() error {
	s.logger.Info("starting scans grpc server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	return s.server.Serve(lis)
}

func (s *grpcServer) Close() error {
	s.logger.Info("stopping scans grpc server")
	s.server.GracefulStop()
	return nil
}

func (s *grpcServer) UpdateScanStatus(ctx context.Context, req *pb.UpdateScanStatusRequest) (*pb.UpdateScanStatusResponse, error) {
	s.logger.Info(fmt.Sprintf("Got UpdateScanStatus request for scanId=%s, status=%s", req.GetId(), req.GetStatus()))

	err := s.scanStatusService.UpdateScanStatus(ctx, req.Id, req.Status, req.UpdateTime)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to update scan status for scanId=%s, status=%s ", req.GetId(), req.GetStatus()), err)
		return &pb.UpdateScanStatusResponse{
			UpdateStatus: "FAIL",
			Error:        err.Error(),
		}, nil
	}

	s.logger.Info(fmt.Sprintf("Successfully updated scan status for scanId=%s, status=%s", req.GetId(), req.GetStatus()))
	return &pb.UpdateScanStatusResponse{
		UpdateStatus: "SUCCESS",
	}, nil
}

func (s *grpcServer) UpdateScanStatusDetails(ctx context.Context, req *pb.UpdateScanStatusDetailsRequest) (*pb.UpdateScanStatusResponse, error) {
	s.logger.Info(fmt.Sprintf("Got UpdateScanStatusDetails request for scanId=%s, stepName=%s, status=%s", req.GetId(), req.GetStepName(), req.GetStatus()))

	err := s.scanStatusService.UpdateScanStatusDetails(ctx, req.Id, req.StepName, req.Status, req.UpdateTime)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to update scan status details for scanId=%s, stepName=%s, status=%s", req.GetId(), req.GetStepName(), req.GetStatus()), err)
		return &pb.UpdateScanStatusResponse{
			UpdateStatus: "FAIL",
			Error:        err.Error(),
		}, nil
	}

	s.logger.Info(fmt.Sprintf("Successfully updated scan status details for scanId=%s, stepName=%s, status=%s", req.GetId(), req.GetStepName(), req.GetStatus()))
	return &pb.UpdateScanStatusResponse{
		UpdateStatus: "SUCCESS",
	}, nil
}

func (s *grpcServer) GetScanStatusDetails(ctx context.Context, req *pb.GetScanStatusDetailsRequest) (*pb.GetScanStatusDetailsResponse, error) {
	s.logger.Info(fmt.Sprintf("Got GetScanStatusDetails request for scanId=%s", req.Id))

	scanDetails, err := s.scanStatusService.GetScanStatusDetails(ctx, req.Id)
	if err != nil || scanDetails == nil {
		s.logger.Error(fmt.Sprintf("Failed to get scan status details for scanId=%s", req.Id), err)
		return &pb.GetScanStatusDetailsResponse{}, err
	}

	s.logger.Info(fmt.Sprintf("Successfully got scan status details for scanId=%s", req.Id))
	return scanStatusService.ConvertScanStatusToProto(scanDetails), nil
}
