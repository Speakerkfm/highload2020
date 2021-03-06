// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: serviceA/adapters/service_b/phone_sevice.proto

package service_b

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type CheckPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CheckPhoneRequest) Reset() {
	*x = CheckPhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPhoneRequest) ProtoMessage() {}

func (x *CheckPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPhoneRequest.ProtoReflect.Descriptor instead.
func (*CheckPhoneRequest) Descriptor() ([]byte, []int) {
	return file_serviceA_adapters_service_b_phone_sevice_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CheckPhoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *CheckPhoneResponse) Reset() {
	*x = CheckPhoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPhoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPhoneResponse) ProtoMessage() {}

func (x *CheckPhoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPhoneResponse.ProtoReflect.Descriptor instead.
func (*CheckPhoneResponse) Descriptor() ([]byte, []int) {
	return file_serviceA_adapters_service_b_phone_sevice_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPhoneResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_serviceA_adapters_service_b_phone_sevice_proto protoreflect.FileDescriptor

var file_serviceA_adapters_service_b_phone_sevice_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x2f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x29, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0x45, 0x0a, 0x0c, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1d, 0x5a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x2f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serviceA_adapters_service_b_phone_sevice_proto_rawDescOnce sync.Once
	file_serviceA_adapters_service_b_phone_sevice_proto_rawDescData = file_serviceA_adapters_service_b_phone_sevice_proto_rawDesc
)

func file_serviceA_adapters_service_b_phone_sevice_proto_rawDescGZIP() []byte {
	file_serviceA_adapters_service_b_phone_sevice_proto_rawDescOnce.Do(func() {
		file_serviceA_adapters_service_b_phone_sevice_proto_rawDescData = protoimpl.X.CompressGZIP(file_serviceA_adapters_service_b_phone_sevice_proto_rawDescData)
	})
	return file_serviceA_adapters_service_b_phone_sevice_proto_rawDescData
}

var file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_serviceA_adapters_service_b_phone_sevice_proto_goTypes = []interface{}{
	(*CheckPhoneRequest)(nil),  // 0: CheckPhoneRequest
	(*CheckPhoneResponse)(nil), // 1: CheckPhoneResponse
}
var file_serviceA_adapters_service_b_phone_sevice_proto_depIdxs = []int32{
	0, // 0: PhoneService.CheckPhone:input_type -> CheckPhoneRequest
	1, // 1: PhoneService.CheckPhone:output_type -> CheckPhoneResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serviceA_adapters_service_b_phone_sevice_proto_init() }
func file_serviceA_adapters_service_b_phone_sevice_proto_init() {
	if File_serviceA_adapters_service_b_phone_sevice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPhoneRequest); i {
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
		file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPhoneResponse); i {
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
			RawDescriptor: file_serviceA_adapters_service_b_phone_sevice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serviceA_adapters_service_b_phone_sevice_proto_goTypes,
		DependencyIndexes: file_serviceA_adapters_service_b_phone_sevice_proto_depIdxs,
		MessageInfos:      file_serviceA_adapters_service_b_phone_sevice_proto_msgTypes,
	}.Build()
	File_serviceA_adapters_service_b_phone_sevice_proto = out.File
	file_serviceA_adapters_service_b_phone_sevice_proto_rawDesc = nil
	file_serviceA_adapters_service_b_phone_sevice_proto_goTypes = nil
	file_serviceA_adapters_service_b_phone_sevice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PhoneServiceClient is the client API for PhoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhoneServiceClient interface {
	CheckPhone(ctx context.Context, in *CheckPhoneRequest, opts ...grpc.CallOption) (*CheckPhoneResponse, error)
}

type phoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhoneServiceClient(cc grpc.ClientConnInterface) PhoneServiceClient {
	return &phoneServiceClient{cc}
}

func (c *phoneServiceClient) CheckPhone(ctx context.Context, in *CheckPhoneRequest, opts ...grpc.CallOption) (*CheckPhoneResponse, error) {
	out := new(CheckPhoneResponse)
	err := c.cc.Invoke(ctx, "/PhoneService/CheckPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhoneServiceServer is the server API for PhoneService service.
type PhoneServiceServer interface {
	CheckPhone(context.Context, *CheckPhoneRequest) (*CheckPhoneResponse, error)
}

// UnimplementedPhoneServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPhoneServiceServer struct {
}

func (*UnimplementedPhoneServiceServer) CheckPhone(context.Context, *CheckPhoneRequest) (*CheckPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPhone not implemented")
}

func RegisterPhoneServiceServer(s *grpc.Server, srv PhoneServiceServer) {
	s.RegisterService(&_PhoneService_serviceDesc, srv)
}

func _PhoneService_CheckPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneServiceServer).CheckPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PhoneService/CheckPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneServiceServer).CheckPhone(ctx, req.(*CheckPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhoneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PhoneService",
	HandlerType: (*PhoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPhone",
			Handler:    _PhoneService_CheckPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceA/adapters/service_b/phone_sevice.proto",
}
