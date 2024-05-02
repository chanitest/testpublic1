// protoc scan_status_service.proto -Iapi -Ithird_party --go_out=plugins=grpc:./pkg/
// protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --proto_path=./api/ scan_status_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: scan_status_service.proto

package scan_status_clinet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetScanStatusDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetScanStatusDetailsRequest) Reset() {
	*x = GetScanStatusDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scan_status_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScanStatusDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScanStatusDetailsRequest) ProtoMessage() {}

func (x *GetScanStatusDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scan_status_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScanStatusDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetScanStatusDetailsRequest) Descriptor() ([]byte, []int) {
	return file_scan_status_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetScanStatusDetailsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetScanStatusDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	StepStatus    []*StepStatus          `protobuf:"bytes,3,rep,name=stepStatus,proto3" json:"stepStatus,omitempty"`
	CorrelationId string                 `protobuf:"bytes,5,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *GetScanStatusDetailsResponse) Reset() {
	*x = GetScanStatusDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scan_status_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScanStatusDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScanStatusDetailsResponse) ProtoMessage() {}

func (x *GetScanStatusDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scan_status_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScanStatusDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetScanStatusDetailsResponse) Descriptor() ([]byte, []int) {
	return file_scan_status_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetScanStatusDetailsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetScanStatusDetailsResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetScanStatusDetailsResponse) GetStepStatus() []*StepStatus {
	if x != nil {
		return x.StepStatus
	}
	return nil
}

func (x *GetScanStatusDetailsResponse) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *GetScanStatusDetailsResponse) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type UpdateScanStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CorrelationId string                 `protobuf:"bytes,5,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *UpdateScanStatusRequest) Reset() {
	*x = UpdateScanStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scan_status_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScanStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScanStatusRequest) ProtoMessage() {}

func (x *UpdateScanStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scan_status_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScanStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateScanStatusRequest) Descriptor() ([]byte, []int) {
	return file_scan_status_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateScanStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateScanStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateScanStatusRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *UpdateScanStatusRequest) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type UpdateScanStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateStatus string `protobuf:"bytes,1,opt,name=updateStatus,proto3" json:"updateStatus,omitempty"`
	Error        string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateScanStatusResponse) Reset() {
	*x = UpdateScanStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scan_status_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScanStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScanStatusResponse) ProtoMessage() {}

func (x *UpdateScanStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scan_status_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScanStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateScanStatusResponse) Descriptor() ([]byte, []int) {
	return file_scan_status_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateScanStatusResponse) GetUpdateStatus() string {
	if x != nil {
		return x.UpdateStatus
	}
	return ""
}

func (x *UpdateScanStatusResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateScanStatusDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StepName      string                 `protobuf:"bytes,2,opt,name=stepName,proto3" json:"stepName,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CorrelationId string                 `protobuf:"bytes,4,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *UpdateScanStatusDetailsRequest) Reset() {
	*x = UpdateScanStatusDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scan_status_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScanStatusDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScanStatusDetailsRequest) ProtoMessage() {}

func (x *UpdateScanStatusDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scan_status_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScanStatusDetailsRequest.ProtoReflect.Descriptor instead.
func (*UpdateScanStatusDetailsRequest) Descriptor() ([]byte, []int) {
	return file_scan_status_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateScanStatusDetailsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateScanStatusDetailsRequest) GetStepName() string {
	if x != nil {
		return x.StepName
	}
	return ""
}

func (x *UpdateScanStatusDetailsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateScanStatusDetailsRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *UpdateScanStatusDetailsRequest) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type StepStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status     string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *StepStatus) Reset() {
	*x = StepStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scan_status_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepStatus) ProtoMessage() {}

func (x *StepStatus) ProtoReflect() protoreflect.Message {
	mi := &file_scan_status_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepStatus.ProtoReflect.Descriptor instead.
func (*StepStatus) Descriptor() ([]byte, []int) {
	return file_scan_status_service_proto_rawDescGZIP(), []int{5}
}

func (x *StepStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StepStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StepStatus) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_scan_status_service_proto protoreflect.FileDescriptor

var file_scan_status_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x63, 0x61,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe1, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a,
	0x73, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa3, 0x01, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x54, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xc6, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x74, 0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x74, 0x0a, 0x0a, 0x53, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xd0, 0x02, 0x0a, 0x11, 0x53, 0x63, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x24, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2b, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x28, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scan_status_service_proto_rawDescOnce sync.Once
	file_scan_status_service_proto_rawDescData = file_scan_status_service_proto_rawDesc
)

func file_scan_status_service_proto_rawDescGZIP() []byte {
	file_scan_status_service_proto_rawDescOnce.Do(func() {
		file_scan_status_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_scan_status_service_proto_rawDescData)
	})
	return file_scan_status_service_proto_rawDescData
}

var file_scan_status_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_scan_status_service_proto_goTypes = []interface{}{
	(*GetScanStatusDetailsRequest)(nil),    // 0: scan_status.GetScanStatusDetailsRequest
	(*GetScanStatusDetailsResponse)(nil),   // 1: scan_status.GetScanStatusDetailsResponse
	(*UpdateScanStatusRequest)(nil),        // 2: scan_status.UpdateScanStatusRequest
	(*UpdateScanStatusResponse)(nil),       // 3: scan_status.UpdateScanStatusResponse
	(*UpdateScanStatusDetailsRequest)(nil), // 4: scan_status.UpdateScanStatusDetailsRequest
	(*StepStatus)(nil),                     // 5: scan_status.StepStatus
	(*timestamppb.Timestamp)(nil),          // 6: google.protobuf.Timestamp
}
var file_scan_status_service_proto_depIdxs = []int32{
	5, // 0: scan_status.GetScanStatusDetailsResponse.stepStatus:type_name -> scan_status.StepStatus
	6, // 1: scan_status.GetScanStatusDetailsResponse.updateTime:type_name -> google.protobuf.Timestamp
	6, // 2: scan_status.UpdateScanStatusRequest.updateTime:type_name -> google.protobuf.Timestamp
	6, // 3: scan_status.UpdateScanStatusDetailsRequest.updateTime:type_name -> google.protobuf.Timestamp
	6, // 4: scan_status.StepStatus.updateTime:type_name -> google.protobuf.Timestamp
	2, // 5: scan_status.ScanStatusService.UpdateScanStatus:input_type -> scan_status.UpdateScanStatusRequest
	4, // 6: scan_status.ScanStatusService.UpdateScanStatusDetails:input_type -> scan_status.UpdateScanStatusDetailsRequest
	0, // 7: scan_status.ScanStatusService.GetScanStatusDetails:input_type -> scan_status.GetScanStatusDetailsRequest
	3, // 8: scan_status.ScanStatusService.UpdateScanStatus:output_type -> scan_status.UpdateScanStatusResponse
	3, // 9: scan_status.ScanStatusService.UpdateScanStatusDetails:output_type -> scan_status.UpdateScanStatusResponse
	1, // 10: scan_status.ScanStatusService.GetScanStatusDetails:output_type -> scan_status.GetScanStatusDetailsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_scan_status_service_proto_init() }
func file_scan_status_service_proto_init() {
	if File_scan_status_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scan_status_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScanStatusDetailsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scan_status_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScanStatusDetailsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scan_status_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScanStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scan_status_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScanStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scan_status_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScanStatusDetailsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scan_status_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scan_status_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scan_status_service_proto_goTypes,
		DependencyIndexes: file_scan_status_service_proto_depIdxs,
		MessageInfos:      file_scan_status_service_proto_msgTypes,
	}.Build()
	File_scan_status_service_proto = out.File
	file_scan_status_service_proto_rawDesc = nil
	file_scan_status_service_proto_goTypes = nil
	file_scan_status_service_proto_depIdxs = nil
}
