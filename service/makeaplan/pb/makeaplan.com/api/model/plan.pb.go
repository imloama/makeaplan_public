// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: plan.proto

package model

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PlanType int32

const (
	PlanType_UNKNOWN_TYPE PlanType = 0
	// 不分阶段的任务列表计划
	PlanType_TASK_LIST PlanType = 1
	// 分阶段计划
	PlanType_PHASED PlanType = 2
)

// Enum value maps for PlanType.
var (
	PlanType_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "TASK_LIST",
		2: "PHASED",
	}
	PlanType_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"TASK_LIST":    1,
		"PHASED":       2,
	}
)

func (x PlanType) Enum() *PlanType {
	p := new(PlanType)
	*p = x
	return p
}

func (x PlanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanType) Descriptor() protoreflect.EnumDescriptor {
	return file_plan_proto_enumTypes[0].Descriptor()
}

func (PlanType) Type() protoreflect.EnumType {
	return &file_plan_proto_enumTypes[0]
}

func (x PlanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanType.Descriptor instead.
func (PlanType) EnumDescriptor() ([]byte, []int) {
	return file_plan_proto_rawDescGZIP(), []int{0}
}

type PlanCompleteStatus int32

const (
	PlanCompleteStatus_UNKNOWN_PLAN_COMPLETE_STATUS PlanCompleteStatus = 0
	// 完成
	PlanCompleteStatus_PLAN_DONE PlanCompleteStatus = 1
	// 成功
	PlanCompleteStatus_PLAN_SUCCEED PlanCompleteStatus = 2
	// 失败
	PlanCompleteStatus_PLAN_FAILED PlanCompleteStatus = 3
)

// Enum value maps for PlanCompleteStatus.
var (
	PlanCompleteStatus_name = map[int32]string{
		0: "UNKNOWN_PLAN_COMPLETE_STATUS",
		1: "PLAN_DONE",
		2: "PLAN_SUCCEED",
		3: "PLAN_FAILED",
	}
	PlanCompleteStatus_value = map[string]int32{
		"UNKNOWN_PLAN_COMPLETE_STATUS": 0,
		"PLAN_DONE":                    1,
		"PLAN_SUCCEED":                 2,
		"PLAN_FAILED":                  3,
	}
)

func (x PlanCompleteStatus) Enum() *PlanCompleteStatus {
	p := new(PlanCompleteStatus)
	*p = x
	return p
}

func (x PlanCompleteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanCompleteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_plan_proto_enumTypes[1].Descriptor()
}

func (PlanCompleteStatus) Type() protoreflect.EnumType {
	return &file_plan_proto_enumTypes[1]
}

func (x PlanCompleteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanCompleteStatus.Descriptor instead.
func (PlanCompleteStatus) EnumDescriptor() ([]byte, []int) {
	return file_plan_proto_rawDescGZIP(), []int{1}
}

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title          string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code           string               `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Note           string               `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	Type           PlanType             `protobuf:"varint,5,opt,name=type,proto3,enum=service.PlanType" json:"type,omitempty"`
	TotalIndex     int32                `protobuf:"varint,7,opt,name=total_index,json=totalIndex,proto3" json:"total_index,omitempty"`
	MinIndex       int32                `protobuf:"varint,8,opt,name=min_index,json=minIndex,proto3" json:"min_index,omitempty"`
	CompletedIndex int32                `protobuf:"varint,9,opt,name=completed_index,json=completedIndex,proto3" json:"completed_index,omitempty"`
	StartTime      *timestamp.Timestamp `protobuf:"bytes,10,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime        *timestamp.Timestamp `protobuf:"bytes,11,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	UpdateTime     *timestamp.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_plan_proto_rawDescGZIP(), []int{0}
}

func (x *Plan) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Plan) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Plan) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Plan) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Plan) GetType() PlanType {
	if x != nil {
		return x.Type
	}
	return PlanType_UNKNOWN_TYPE
}

func (x *Plan) GetTotalIndex() int32 {
	if x != nil {
		return x.TotalIndex
	}
	return 0
}

func (x *Plan) GetMinIndex() int32 {
	if x != nil {
		return x.MinIndex
	}
	return 0
}

func (x *Plan) GetCompletedIndex() int32 {
	if x != nil {
		return x.CompletedIndex
	}
	return 0
}

func (x *Plan) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Plan) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Plan) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type PlanSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan        *Plan        `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	Status      CommonStatus `protobuf:"varint,2,opt,name=status,proto3,enum=service.CommonStatus" json:"status,omitempty"`
	TaskStatis  *Statistics  `protobuf:"bytes,3,opt,name=taskStatis,proto3" json:"taskStatis,omitempty"`
	PhaseStatis *Statistics  `protobuf:"bytes,4,opt,name=phaseStatis,proto3" json:"phaseStatis,omitempty"`
}

func (x *PlanSummary) Reset() {
	*x = PlanSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanSummary) ProtoMessage() {}

func (x *PlanSummary) ProtoReflect() protoreflect.Message {
	mi := &file_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanSummary.ProtoReflect.Descriptor instead.
func (*PlanSummary) Descriptor() ([]byte, []int) {
	return file_plan_proto_rawDescGZIP(), []int{1}
}

func (x *PlanSummary) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *PlanSummary) GetStatus() CommonStatus {
	if x != nil {
		return x.Status
	}
	return CommonStatus_UNKNOWN_COMMON_STATUS
}

func (x *PlanSummary) GetTaskStatis() *Statistics {
	if x != nil {
		return x.TaskStatis
	}
	return nil
}

func (x *PlanSummary) GetPhaseStatis() *Statistics {
	if x != nil {
		return x.PhaseStatis
	}
	return nil
}

var File_plan_proto protoreflect.FileDescriptor

var file_plan_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x0b, 0x50,
	0x6c, 0x61, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x6c,
	0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x2d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0a,
	0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x12, 0x35, 0x0a, 0x0b, 0x70, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0b, 0x70, 0x68, 0x61,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x2a, 0x37, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x4c,
	0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x48, 0x41, 0x53, 0x45, 0x44, 0x10,
	0x02, 0x2a, 0x68, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4c, 0x41,
	0x4e, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x4e,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4c,
	0x41, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x19, 0x5a, 0x17, 0x6d,
	0x61, 0x6b, 0x65, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plan_proto_rawDescOnce sync.Once
	file_plan_proto_rawDescData = file_plan_proto_rawDesc
)

func file_plan_proto_rawDescGZIP() []byte {
	file_plan_proto_rawDescOnce.Do(func() {
		file_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_plan_proto_rawDescData)
	})
	return file_plan_proto_rawDescData
}

var file_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_plan_proto_goTypes = []interface{}{
	(PlanType)(0),               // 0: service.PlanType
	(PlanCompleteStatus)(0),     // 1: service.PlanCompleteStatus
	(*Plan)(nil),                // 2: service.Plan
	(*PlanSummary)(nil),         // 3: service.PlanSummary
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(CommonStatus)(0),           // 5: service.CommonStatus
	(*Statistics)(nil),          // 6: service.Statistics
}
var file_plan_proto_depIdxs = []int32{
	0, // 0: service.Plan.type:type_name -> service.PlanType
	4, // 1: service.Plan.start_time:type_name -> google.protobuf.Timestamp
	4, // 2: service.Plan.end_time:type_name -> google.protobuf.Timestamp
	4, // 3: service.Plan.update_time:type_name -> google.protobuf.Timestamp
	2, // 4: service.PlanSummary.plan:type_name -> service.Plan
	5, // 5: service.PlanSummary.status:type_name -> service.CommonStatus
	6, // 6: service.PlanSummary.taskStatis:type_name -> service.Statistics
	6, // 7: service.PlanSummary.phaseStatis:type_name -> service.Statistics
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_plan_proto_init() }
func file_plan_proto_init() {
	if File_plan_proto != nil {
		return
	}
	file_common_proto_init()
	file_statistics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
		file_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanSummary); i {
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
			RawDescriptor: file_plan_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plan_proto_goTypes,
		DependencyIndexes: file_plan_proto_depIdxs,
		EnumInfos:         file_plan_proto_enumTypes,
		MessageInfos:      file_plan_proto_msgTypes,
	}.Build()
	File_plan_proto = out.File
	file_plan_proto_rawDesc = nil
	file_plan_proto_goTypes = nil
	file_plan_proto_depIdxs = nil
}
