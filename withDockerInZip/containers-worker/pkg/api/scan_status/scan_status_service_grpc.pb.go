// protoc scan_status_service.proto -Iapi -Ithird_party --go_out=plugins=grpc:./pkg/
// protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --proto_path=./api/ scan_status_service.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: scan_status_service.proto

package scan_status

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ScanStatusService_UpdateScanStatus_FullMethodName        = "/scan_status.ScanStatusService/UpdateScanStatus"
	ScanStatusService_UpdateScanStatusDetails_FullMethodName = "/scan_status.ScanStatusService/UpdateScanStatusDetails"
	ScanStatusService_GetScanStatusDetails_FullMethodName    = "/scan_status.ScanStatusService/GetScanStatusDetails"
)

// ScanStatusServiceClient is the client API for ScanStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScanStatusServiceClient interface {
	UpdateScanStatus(ctx context.Context, in *UpdateScanStatusRequest, opts ...grpc.CallOption) (*UpdateScanStatusResponse, error)
	UpdateScanStatusDetails(ctx context.Context, in *UpdateScanStatusDetailsRequest, opts ...grpc.CallOption) (*UpdateScanStatusResponse, error)
	GetScanStatusDetails(ctx context.Context, in *GetScanStatusDetailsRequest, opts ...grpc.CallOption) (*GetScanStatusDetailsResponse, error)
}

type scanStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScanStatusServiceClient(cc grpc.ClientConnInterface) ScanStatusServiceClient {
	return &scanStatusServiceClient{cc}
}

func (c *scanStatusServiceClient) UpdateScanStatus(ctx context.Context, in *UpdateScanStatusRequest, opts ...grpc.CallOption) (*UpdateScanStatusResponse, error) {
	out := new(UpdateScanStatusResponse)
	err := c.cc.Invoke(ctx, ScanStatusService_UpdateScanStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scanStatusServiceClient) UpdateScanStatusDetails(ctx context.Context, in *UpdateScanStatusDetailsRequest, opts ...grpc.CallOption) (*UpdateScanStatusResponse, error) {
	out := new(UpdateScanStatusResponse)
	err := c.cc.Invoke(ctx, ScanStatusService_UpdateScanStatusDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scanStatusServiceClient) GetScanStatusDetails(ctx context.Context, in *GetScanStatusDetailsRequest, opts ...grpc.CallOption) (*GetScanStatusDetailsResponse, error) {
	out := new(GetScanStatusDetailsResponse)
	err := c.cc.Invoke(ctx, ScanStatusService_GetScanStatusDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanStatusServiceServer is the server API for ScanStatusService service.
// All implementations must embed UnimplementedScanStatusServiceServer
// for forward compatibility
type ScanStatusServiceServer interface {
	UpdateScanStatus(context.Context, *UpdateScanStatusRequest) (*UpdateScanStatusResponse, error)
	UpdateScanStatusDetails(context.Context, *UpdateScanStatusDetailsRequest) (*UpdateScanStatusResponse, error)
	GetScanStatusDetails(context.Context, *GetScanStatusDetailsRequest) (*GetScanStatusDetailsResponse, error)
	mustEmbedUnimplementedScanStatusServiceServer()
}

// UnimplementedScanStatusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScanStatusServiceServer struct {
}

func (UnimplementedScanStatusServiceServer) UpdateScanStatus(context.Context, *UpdateScanStatusRequest) (*UpdateScanStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScanStatus not implemented")
}
func (UnimplementedScanStatusServiceServer) UpdateScanStatusDetails(context.Context, *UpdateScanStatusDetailsRequest) (*UpdateScanStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScanStatusDetails not implemented")
}
func (UnimplementedScanStatusServiceServer) GetScanStatusDetails(context.Context, *GetScanStatusDetailsRequest) (*GetScanStatusDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScanStatusDetails not implemented")
}
func (UnimplementedScanStatusServiceServer) mustEmbedUnimplementedScanStatusServiceServer() {}

// UnsafeScanStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScanStatusServiceServer will
// result in compilation errors.
type UnsafeScanStatusServiceServer interface {
	mustEmbedUnimplementedScanStatusServiceServer()
}

func RegisterScanStatusServiceServer(s grpc.ServiceRegistrar, srv ScanStatusServiceServer) {
	s.RegisterService(&ScanStatusService_ServiceDesc, srv)
}

func _ScanStatusService_UpdateScanStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScanStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanStatusServiceServer).UpdateScanStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanStatusService_UpdateScanStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanStatusServiceServer).UpdateScanStatus(ctx, req.(*UpdateScanStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScanStatusService_UpdateScanStatusDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScanStatusDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanStatusServiceServer).UpdateScanStatusDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanStatusService_UpdateScanStatusDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanStatusServiceServer).UpdateScanStatusDetails(ctx, req.(*UpdateScanStatusDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScanStatusService_GetScanStatusDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScanStatusDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanStatusServiceServer).GetScanStatusDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanStatusService_GetScanStatusDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanStatusServiceServer).GetScanStatusDetails(ctx, req.(*GetScanStatusDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScanStatusService_ServiceDesc is the grpc.ServiceDesc for ScanStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScanStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scan_status.ScanStatusService",
	HandlerType: (*ScanStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateScanStatus",
			Handler:    _ScanStatusService_UpdateScanStatus_Handler,
		},
		{
			MethodName: "UpdateScanStatusDetails",
			Handler:    _ScanStatusService_UpdateScanStatusDetails_Handler,
		},
		{
			MethodName: "GetScanStatusDetails",
			Handler:    _ScanStatusService_GetScanStatusDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scan_status_service.proto",
}
