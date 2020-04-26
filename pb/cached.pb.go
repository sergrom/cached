// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.6.1
// source: cached.proto

package cached

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cached_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cached_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cached_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIdReq) Reset() {
	*x = GetUserByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cached_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdReq) ProtoMessage() {}

func (x *GetUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_cached_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdReq.ProtoReflect.Descriptor instead.
func (*GetUserByIdReq) Descriptor() ([]byte, []int) {
	return file_cached_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIdReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByIdResp) Reset() {
	*x = GetUserByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cached_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdResp) ProtoMessage() {}

func (x *GetUserByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_cached_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdResp.ProtoReflect.Descriptor instead.
func (*GetUserByIdResp) Descriptor() ([]byte, []int) {
	return file_cached_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIdResp) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserListReq) Reset() {
	*x = GetUserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cached_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListReq) ProtoMessage() {}

func (x *GetUserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_cached_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListReq.ProtoReflect.Descriptor instead.
func (*GetUserListReq) Descriptor() ([]byte, []int) {
	return file_cached_proto_rawDescGZIP(), []int{3}
}

type GetUserListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*User `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *GetUserListResp) Reset() {
	*x = GetUserListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cached_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResp) ProtoMessage() {}

func (x *GetUserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_cached_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResp.ProtoReflect.Descriptor instead.
func (*GetUserListResp) Descriptor() ([]byte, []int) {
	return file_cached_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserListResp) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

var File_cached_proto protoreflect.FileDescriptor

var file_cached_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x35, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x22, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0x70, 0x0a, 0x06, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x32, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cached_proto_rawDescOnce sync.Once
	file_cached_proto_rawDescData = file_cached_proto_rawDesc
)

func file_cached_proto_rawDescGZIP() []byte {
	file_cached_proto_rawDescOnce.Do(func() {
		file_cached_proto_rawDescData = protoimpl.X.CompressGZIP(file_cached_proto_rawDescData)
	})
	return file_cached_proto_rawDescData
}

var file_cached_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cached_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: User
	(*GetUserByIdReq)(nil),  // 1: GetUserByIdReq
	(*GetUserByIdResp)(nil), // 2: GetUserByIdResp
	(*GetUserListReq)(nil),  // 3: GetUserListReq
	(*GetUserListResp)(nil), // 4: GetUserListResp
}
var file_cached_proto_depIdxs = []int32{
	0, // 0: GetUserByIdResp.user:type_name -> User
	0, // 1: GetUserListResp.user_list:type_name -> User
	1, // 2: Cached.GetUserById:input_type -> GetUserByIdReq
	3, // 3: Cached.GetUserList:input_type -> GetUserListReq
	2, // 4: Cached.GetUserById:output_type -> GetUserByIdResp
	4, // 5: Cached.GetUserList:output_type -> GetUserListResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cached_proto_init() }
func file_cached_proto_init() {
	if File_cached_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cached_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_cached_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdReq); i {
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
		file_cached_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdResp); i {
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
		file_cached_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListReq); i {
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
		file_cached_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListResp); i {
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
			RawDescriptor: file_cached_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cached_proto_goTypes,
		DependencyIndexes: file_cached_proto_depIdxs,
		MessageInfos:      file_cached_proto_msgTypes,
	}.Build()
	File_cached_proto = out.File
	file_cached_proto_rawDesc = nil
	file_cached_proto_goTypes = nil
	file_cached_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CachedClient is the client API for Cached service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CachedClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error)
}

type cachedClient struct {
	cc grpc.ClientConnInterface
}

func NewCachedClient(cc grpc.ClientConnInterface) CachedClient {
	return &cachedClient{cc}
}

func (c *cachedClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	out := new(GetUserByIdResp)
	err := c.cc.Invoke(ctx, "/Cached/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cachedClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error) {
	out := new(GetUserListResp)
	err := c.cc.Invoke(ctx, "/Cached/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CachedServer is the server API for Cached service.
type CachedServer interface {
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error)
	GetUserList(context.Context, *GetUserListReq) (*GetUserListResp, error)
}

// UnimplementedCachedServer can be embedded to have forward compatible implementations.
type UnimplementedCachedServer struct {
}

func (*UnimplementedCachedServer) GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedCachedServer) GetUserList(context.Context, *GetUserListReq) (*GetUserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}

func RegisterCachedServer(s *grpc.Server, srv CachedServer) {
	s.RegisterService(&_Cached_serviceDesc, srv)
}

func _Cached_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CachedServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cached/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CachedServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cached_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CachedServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cached/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CachedServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cached_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cached",
	HandlerType: (*CachedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _Cached_GetUserById_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _Cached_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cached.proto",
}
