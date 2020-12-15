// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: phase.proto

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

type PhaseCompleteStatus int32

const (
	PhaseCompleteStatus_UNKNOWN_PHASE_COMPLETE_STATUS PhaseCompleteStatus = 0
	// 完成
	PhaseCompleteStatus_PHASE_DONE PhaseCompleteStatus = 1
)

// Enum value maps for PhaseCompleteStatus.
var (
	PhaseCompleteStatus_name = map[int32]string{
		0: "UNKNOWN_PHASE_COMPLETE_STATUS",
		1: "PHASE_DONE",
	}
	PhaseCompleteStatus_value = map[string]int32{
		"UNKNOWN_PHASE_COMPLETE_STATUS": 0,
		"PHASE_DONE":                    1,
	}
)

func (x PhaseCompleteStatus) Enum() *PhaseCompleteStatus {
	p := new(PhaseCompleteStatus)
	*p = x
	return p
}

func (x PhaseCompleteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhaseCompleteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_phase_proto_enumTypes[0].Descriptor()
}

func (PhaseCompleteStatus) Type() protoreflect.EnumType {
	return &file_phase_proto_enumTypes[0]
}

func (x PhaseCompleteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhaseCompleteStatus.Descriptor instead.
func (PhaseCompleteStatus) EnumDescriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{0}
}

type Phase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Note      string               `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	StartTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Phase) Reset() {
	*x = Phase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Phase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phase) ProtoMessage() {}

func (x *Phase) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phase.ProtoReflect.Descriptor instead.
func (*Phase) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{0}
}

func (x *Phase) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Phase) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Phase) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Phase) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Phase) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type PhaseSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phase      *Phase       `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase,omitempty"`
	Status     CommonStatus `protobuf:"varint,2,opt,name=status,proto3,enum=service.CommonStatus" json:"status,omitempty"`
	TaskStatis *Statistics  `protobuf:"bytes,3,opt,name=taskStatis,proto3" json:"taskStatis,omitempty"`
}

func (x *PhaseSummary) Reset() {
	*x = PhaseSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhaseSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhaseSummary) ProtoMessage() {}

func (x *PhaseSummary) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhaseSummary.ProtoReflect.Descriptor instead.
func (*PhaseSummary) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{1}
}

func (x *PhaseSummary) GetPhase() *Phase {
	if x != nil {
		return x.Phase
	}
	return nil
}

func (x *PhaseSummary) GetStatus() CommonStatus {
	if x != nil {
		return x.Status
	}
	return CommonStatus_UNKNOWN_COMMON_STATUS
}

func (x *PhaseSummary) GetTaskStatis() *Statistics {
	if x != nil {
		return x.TaskStatis
	}
	return nil
}

var File_phase_proto protoreflect.FileDescriptor

var file_phase_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x05, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x98, 0x01,
	0x0a, 0x0c, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x24,
	0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x05, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x74, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x2a, 0x48, 0x0a, 0x13, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x1d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x6d, 0x61, 0x6b, 0x65, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_phase_proto_rawDescOnce sync.Once
	file_phase_proto_rawDescData = file_phase_proto_rawDesc
)

func file_phase_proto_rawDescGZIP() []byte {
	file_phase_proto_rawDescOnce.Do(func() {
		file_phase_proto_rawDescData = protoimpl.X.CompressGZIP(file_phase_proto_rawDescData)
	})
	return file_phase_proto_rawDescData
}

var file_phase_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_phase_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_phase_proto_goTypes = []interface{}{
	(PhaseCompleteStatus)(0),    // 0: service.PhaseCompleteStatus
	(*Phase)(nil),               // 1: service.Phase
	(*PhaseSummary)(nil),        // 2: service.PhaseSummary
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(CommonStatus)(0),           // 4: service.CommonStatus
	(*Statistics)(nil),          // 5: service.Statistics
}
var file_phase_proto_depIdxs = []int32{
	3, // 0: service.Phase.start_time:type_name -> google.protobuf.Timestamp
	3, // 1: service.Phase.end_time:type_name -> google.protobuf.Timestamp
	1, // 2: service.PhaseSummary.phase:type_name -> service.Phase
	4, // 3: service.PhaseSummary.status:type_name -> service.CommonStatus
	5, // 4: service.PhaseSummary.taskStatis:type_name -> service.Statistics
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_phase_proto_init() }
func file_phase_proto_init() {
	if File_phase_proto != nil {
		return
	}
	file_common_proto_init()
	file_statistics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_phase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Phase); i {
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
		file_phase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhaseSummary); i {
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
			RawDescriptor: file_phase_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_phase_proto_goTypes,
		DependencyIndexes: file_phase_proto_depIdxs,
		EnumInfos:         file_phase_proto_enumTypes,
		MessageInfos:      file_phase_proto_msgTypes,
	}.Build()
	File_phase_proto = out.File
	file_phase_proto_rawDesc = nil
	file_phase_proto_goTypes = nil
	file_phase_proto_depIdxs = nil
}
