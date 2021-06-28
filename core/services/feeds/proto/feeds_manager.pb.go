// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: feeds_manager.proto

package proto

import (
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

type ApproveJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ApproveJobRequest) Reset() {
	*x = ApproveJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feeds_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveJobRequest) ProtoMessage() {}

func (x *ApproveJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveJobRequest.ProtoReflect.Descriptor instead.
func (*ApproveJobRequest) Descriptor() ([]byte, []int) {
	return file_feeds_manager_proto_rawDescGZIP(), []int{0}
}

func (x *ApproveJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ApproveJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApproveJobResponse) Reset() {
	*x = ApproveJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feeds_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveJobResponse) ProtoMessage() {}

func (x *ApproveJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveJobResponse.ProtoReflect.Descriptor instead.
func (*ApproveJobResponse) Descriptor() ([]byte, []int) {
	return file_feeds_manager_proto_rawDescGZIP(), []int{1}
}

type ProposeJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec string `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *ProposeJobRequest) Reset() {
	*x = ProposeJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feeds_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeJobRequest) ProtoMessage() {}

func (x *ProposeJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeJobRequest.ProtoReflect.Descriptor instead.
func (*ProposeJobRequest) Descriptor() ([]byte, []int) {
	return file_feeds_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ProposeJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProposeJobRequest) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

type ProposeJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProposeJobResponse) Reset() {
	*x = ProposeJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feeds_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeJobResponse) ProtoMessage() {}

func (x *ProposeJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeJobResponse.ProtoReflect.Descriptor instead.
func (*ProposeJobResponse) Descriptor() ([]byte, []int) {
	return file_feeds_manager_proto_rawDescGZIP(), []int{3}
}

func (x *ProposeJobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_feeds_manager_proto protoreflect.FileDescriptor

var file_feeds_manager_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x65, 0x65, 0x64, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x66, 0x6d, 0x22, 0x23, 0x0a, 0x11, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x24,
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x4d, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x73, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4a,
	0x6f, 0x62, 0x12, 0x16, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x66, 0x6d,
	0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x4c, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62,
	0x12, 0x16, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x66, 0x6d, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feeds_manager_proto_rawDescOnce sync.Once
	file_feeds_manager_proto_rawDescData = file_feeds_manager_proto_rawDesc
)

func file_feeds_manager_proto_rawDescGZIP() []byte {
	file_feeds_manager_proto_rawDescOnce.Do(func() {
		file_feeds_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_feeds_manager_proto_rawDescData)
	})
	return file_feeds_manager_proto_rawDescData
}

var file_feeds_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_feeds_manager_proto_goTypes = []interface{}{
	(*ApproveJobRequest)(nil),  // 0: cfm.ApproveJobRequest
	(*ApproveJobResponse)(nil), // 1: cfm.ApproveJobResponse
	(*ProposeJobRequest)(nil),  // 2: cfm.ProposeJobRequest
	(*ProposeJobResponse)(nil), // 3: cfm.ProposeJobResponse
}
var file_feeds_manager_proto_depIdxs = []int32{
	0, // 0: cfm.FeedsManager.ApproveJob:input_type -> cfm.ApproveJobRequest
	2, // 1: cfm.NodeService.ProposeJob:input_type -> cfm.ProposeJobRequest
	1, // 2: cfm.FeedsManager.ApproveJob:output_type -> cfm.ApproveJobResponse
	3, // 3: cfm.NodeService.ProposeJob:output_type -> cfm.ProposeJobResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feeds_manager_proto_init() }
func file_feeds_manager_proto_init() {
	if File_feeds_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feeds_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveJobRequest); i {
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
		file_feeds_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveJobResponse); i {
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
		file_feeds_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeJobRequest); i {
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
		file_feeds_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeJobResponse); i {
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
			RawDescriptor: file_feeds_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_feeds_manager_proto_goTypes,
		DependencyIndexes: file_feeds_manager_proto_depIdxs,
		MessageInfos:      file_feeds_manager_proto_msgTypes,
	}.Build()
	File_feeds_manager_proto = out.File
	file_feeds_manager_proto_rawDesc = nil
	file_feeds_manager_proto_goTypes = nil
	file_feeds_manager_proto_depIdxs = nil
}
