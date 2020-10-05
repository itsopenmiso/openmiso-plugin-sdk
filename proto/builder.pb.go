// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: builder.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_builder_proto_rawDescGZIP(), []int{0}
}

type Build_Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *any.Any          `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Build_Resp) Reset() {
	*x = Build_Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build_Resp) ProtoMessage() {}

func (x *Build_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build_Resp.ProtoReflect.Descriptor instead.
func (*Build_Resp) Descriptor() ([]byte, []int) {
	return file_builder_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Build_Resp) GetResult() *any.Any {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Build_Resp) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_builder_proto protoreflect.FileDescriptor

var file_builder_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a,
	0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x1a, 0xa6, 0x01, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0xde, 0x04, 0x0a, 0x07, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0f, 0x49,
	0x73, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x75,
	0x6e, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a, 0x0c, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x72, 0x67, 0x73,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e,
	0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x40, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0d,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x30,
	0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x75, 0x6e, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_builder_proto_rawDescOnce sync.Once
	file_builder_proto_rawDescData = file_builder_proto_rawDesc
)

func file_builder_proto_rawDescGZIP() []byte {
	file_builder_proto_rawDescOnce.Do(func() {
		file_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_builder_proto_rawDescData)
	})
	return file_builder_proto_rawDescData
}

var file_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_builder_proto_goTypes = []interface{}{
	(*Build)(nil),                   // 0: proto.Build
	(*Build_Resp)(nil),              // 1: proto.Build.Resp
	nil,                             // 2: proto.Build.Resp.LabelsEntry
	(*any.Any)(nil),                 // 3: google.protobuf.Any
	(*empty.Empty)(nil),             // 4: google.protobuf.Empty
	(*FuncSpec_Args)(nil),           // 5: proto.FuncSpec.Args
	(*Config_ConfigureRequest)(nil), // 6: proto.Config.ConfigureRequest
	(*Empty)(nil),                   // 7: proto.Empty
	(*ImplementsResp)(nil),          // 8: proto.ImplementsResp
	(*Auth_AuthResponse)(nil),       // 9: proto.Auth.AuthResponse
	(*FuncSpec)(nil),                // 10: proto.FuncSpec
	(*Config_StructResp)(nil),       // 11: proto.Config.StructResp
	(*Config_Documentation)(nil),    // 12: proto.Config.Documentation
}
var file_builder_proto_depIdxs = []int32{
	3,  // 0: proto.Build.Resp.result:type_name -> google.protobuf.Any
	2,  // 1: proto.Build.Resp.labels:type_name -> proto.Build.Resp.LabelsEntry
	4,  // 2: proto.Builder.IsAuthenticator:input_type -> google.protobuf.Empty
	5,  // 3: proto.Builder.Auth:input_type -> proto.FuncSpec.Args
	4,  // 4: proto.Builder.AuthSpec:input_type -> google.protobuf.Empty
	5,  // 5: proto.Builder.ValidateAuth:input_type -> proto.FuncSpec.Args
	4,  // 6: proto.Builder.ValidateAuthSpec:input_type -> google.protobuf.Empty
	4,  // 7: proto.Builder.ConfigStruct:input_type -> google.protobuf.Empty
	6,  // 8: proto.Builder.Configure:input_type -> proto.Config.ConfigureRequest
	4,  // 9: proto.Builder.Documentation:input_type -> google.protobuf.Empty
	7,  // 10: proto.Builder.BuildSpec:input_type -> proto.Empty
	5,  // 11: proto.Builder.Build:input_type -> proto.FuncSpec.Args
	8,  // 12: proto.Builder.IsAuthenticator:output_type -> proto.ImplementsResp
	9,  // 13: proto.Builder.Auth:output_type -> proto.Auth.AuthResponse
	10, // 14: proto.Builder.AuthSpec:output_type -> proto.FuncSpec
	4,  // 15: proto.Builder.ValidateAuth:output_type -> google.protobuf.Empty
	10, // 16: proto.Builder.ValidateAuthSpec:output_type -> proto.FuncSpec
	11, // 17: proto.Builder.ConfigStruct:output_type -> proto.Config.StructResp
	4,  // 18: proto.Builder.Configure:output_type -> google.protobuf.Empty
	12, // 19: proto.Builder.Documentation:output_type -> proto.Config.Documentation
	10, // 20: proto.Builder.BuildSpec:output_type -> proto.FuncSpec
	1,  // 21: proto.Builder.Build:output_type -> proto.Build.Resp
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_builder_proto_init() }
func file_builder_proto_init() {
	if File_builder_proto != nil {
		return
	}
	file_plugin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_builder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_builder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build_Resp); i {
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
			RawDescriptor: file_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_builder_proto_goTypes,
		DependencyIndexes: file_builder_proto_depIdxs,
		MessageInfos:      file_builder_proto_msgTypes,
	}.Build()
	File_builder_proto = out.File
	file_builder_proto_rawDesc = nil
	file_builder_proto_goTypes = nil
	file_builder_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BuilderClient is the client API for Builder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuilderClient interface {
	IsAuthenticator(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error)
	Auth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Auth_AuthResponse, error)
	AuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	ValidateAuth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
	ValidateAuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error)
	Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Documentation(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_Documentation, error)
	BuildSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	Build(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Build_Resp, error)
}

type builderClient struct {
	cc grpc.ClientConnInterface
}

func NewBuilderClient(cc grpc.ClientConnInterface) BuilderClient {
	return &builderClient{cc}
}

func (c *builderClient) IsAuthenticator(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error) {
	out := new(ImplementsResp)
	err := c.cc.Invoke(ctx, "/proto.Builder/IsAuthenticator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) Auth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Auth_AuthResponse, error) {
	out := new(Auth_AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Builder/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) AuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Builder/AuthSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) ValidateAuth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Builder/ValidateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) ValidateAuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Builder/ValidateAuthSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error) {
	out := new(Config_StructResp)
	err := c.cc.Invoke(ctx, "/proto.Builder/ConfigStruct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Builder/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) Documentation(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_Documentation, error) {
	out := new(Config_Documentation)
	err := c.cc.Invoke(ctx, "/proto.Builder/Documentation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) BuildSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Builder/BuildSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) Build(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Build_Resp, error) {
	out := new(Build_Resp)
	err := c.cc.Invoke(ctx, "/proto.Builder/Build", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuilderServer is the server API for Builder service.
type BuilderServer interface {
	IsAuthenticator(context.Context, *empty.Empty) (*ImplementsResp, error)
	Auth(context.Context, *FuncSpec_Args) (*Auth_AuthResponse, error)
	AuthSpec(context.Context, *empty.Empty) (*FuncSpec, error)
	ValidateAuth(context.Context, *FuncSpec_Args) (*empty.Empty, error)
	ValidateAuthSpec(context.Context, *empty.Empty) (*FuncSpec, error)
	ConfigStruct(context.Context, *empty.Empty) (*Config_StructResp, error)
	Configure(context.Context, *Config_ConfigureRequest) (*empty.Empty, error)
	Documentation(context.Context, *empty.Empty) (*Config_Documentation, error)
	BuildSpec(context.Context, *Empty) (*FuncSpec, error)
	Build(context.Context, *FuncSpec_Args) (*Build_Resp, error)
}

// UnimplementedBuilderServer can be embedded to have forward compatible implementations.
type UnimplementedBuilderServer struct {
}

func (*UnimplementedBuilderServer) IsAuthenticator(context.Context, *empty.Empty) (*ImplementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthenticator not implemented")
}
func (*UnimplementedBuilderServer) Auth(context.Context, *FuncSpec_Args) (*Auth_AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedBuilderServer) AuthSpec(context.Context, *empty.Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthSpec not implemented")
}
func (*UnimplementedBuilderServer) ValidateAuth(context.Context, *FuncSpec_Args) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuth not implemented")
}
func (*UnimplementedBuilderServer) ValidateAuthSpec(context.Context, *empty.Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuthSpec not implemented")
}
func (*UnimplementedBuilderServer) ConfigStruct(context.Context, *empty.Empty) (*Config_StructResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigStruct not implemented")
}
func (*UnimplementedBuilderServer) Configure(context.Context, *Config_ConfigureRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedBuilderServer) Documentation(context.Context, *empty.Empty) (*Config_Documentation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Documentation not implemented")
}
func (*UnimplementedBuilderServer) BuildSpec(context.Context, *Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSpec not implemented")
}
func (*UnimplementedBuilderServer) Build(context.Context, *FuncSpec_Args) (*Build_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}

func RegisterBuilderServer(s *grpc.Server, srv BuilderServer) {
	s.RegisterService(&_Builder_serviceDesc, srv)
}

func _Builder_IsAuthenticator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).IsAuthenticator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/IsAuthenticator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).IsAuthenticator(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).Auth(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_AuthSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).AuthSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/AuthSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).AuthSpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_ValidateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).ValidateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/ValidateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).ValidateAuth(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_ValidateAuthSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).ValidateAuthSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/ValidateAuthSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).ValidateAuthSpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_ConfigStruct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).ConfigStruct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/ConfigStruct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).ConfigStruct(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config_ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).Configure(ctx, req.(*Config_ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_Documentation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).Documentation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/Documentation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).Documentation(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_BuildSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).BuildSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/BuildSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).BuildSpec(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).Build(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Builder",
	HandlerType: (*BuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthenticator",
			Handler:    _Builder_IsAuthenticator_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Builder_Auth_Handler,
		},
		{
			MethodName: "AuthSpec",
			Handler:    _Builder_AuthSpec_Handler,
		},
		{
			MethodName: "ValidateAuth",
			Handler:    _Builder_ValidateAuth_Handler,
		},
		{
			MethodName: "ValidateAuthSpec",
			Handler:    _Builder_ValidateAuthSpec_Handler,
		},
		{
			MethodName: "ConfigStruct",
			Handler:    _Builder_ConfigStruct_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Builder_Configure_Handler,
		},
		{
			MethodName: "Documentation",
			Handler:    _Builder_Documentation_Handler,
		},
		{
			MethodName: "BuildSpec",
			Handler:    _Builder_BuildSpec_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _Builder_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "builder.proto",
}
