// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: video-rpc.proto

package video

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

type StreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StreamReq) Reset() {
	*x = StreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReq) ProtoMessage() {}

func (x *StreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReq.ProtoReflect.Descriptor instead.
func (*StreamReq) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *StreamReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StreamResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greet string `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
}

func (x *StreamResp) Reset() {
	*x = StreamResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResp) ProtoMessage() {}

func (x *StreamResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResp.ProtoReflect.Descriptor instead.
func (*StreamResp) Descriptor() ([]byte, []int) {
	return file_video_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResp) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

var File_video_rpc_proto protoreflect.FileDescriptor

var file_video_rpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x1f, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x32, 0x35, 0x0a,
	0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x12,
	0x10, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_rpc_proto_rawDescOnce sync.Once
	file_video_rpc_proto_rawDescData = file_video_rpc_proto_rawDesc
)

func file_video_rpc_proto_rawDescGZIP() []byte {
	file_video_rpc_proto_rawDescOnce.Do(func() {
		file_video_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_rpc_proto_rawDescData)
	})
	return file_video_rpc_proto_rawDescData
}

var file_video_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_video_rpc_proto_goTypes = []interface{}{
	(*StreamReq)(nil),  // 0: video.StreamReq
	(*StreamResp)(nil), // 1: video.StreamResp
}
var file_video_rpc_proto_depIdxs = []int32{
	0, // 0: video.Video.greet:input_type -> video.StreamReq
	1, // 1: video.Video.greet:output_type -> video.StreamResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_rpc_proto_init() }
func file_video_rpc_proto_init() {
	if File_video_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReq); i {
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
		file_video_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResp); i {
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
			RawDescriptor: file_video_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_rpc_proto_goTypes,
		DependencyIndexes: file_video_rpc_proto_depIdxs,
		MessageInfos:      file_video_rpc_proto_msgTypes,
	}.Build()
	File_video_rpc_proto = out.File
	file_video_rpc_proto_rawDesc = nil
	file_video_rpc_proto_goTypes = nil
	file_video_rpc_proto_depIdxs = nil
}
