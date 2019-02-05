// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Todo struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateTodoRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(m, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type CreateTodoResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Todo                 *Todo    `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(m, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *CreateTodoResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type GetTodoRequest struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (m *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(m, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type GetTodoResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResponse) Reset()         { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()    {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *GetTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResponse.Unmarshal(m, b)
}
func (m *GetTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResponse.Marshal(b, m, deterministic)
}
func (m *GetTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResponse.Merge(m, src)
}
func (m *GetTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoResponse.Size(m)
}
func (m *GetTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResponse proto.InternalMessageInfo

func (m *GetTodoResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.CreateTodoResponse")
	proto.RegisterType((*GetTodoRequest)(nil), "todo.GetTodoRequest")
	proto.RegisterType((*GetTodoResponse)(nil), "todo.GetTodoResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0x4f, 0xc9,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xa4, 0xb8, 0x58, 0x42, 0xf2,
	0x53, 0xf2, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x25, 0x63, 0x2e, 0x41, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0x90, 0x8a, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x39, 0x2e, 0xb0, 0x46, 0xb0, 0x42, 0x6e, 0x23, 0x2e,
	0x3d, 0xb0, 0x89, 0x60, 0x05, 0x10, 0x03, 0x7d, 0xb8, 0x84, 0x90, 0x35, 0x15, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x83, 0xf5, 0x71, 0x04,
	0x41, 0x79, 0x70, 0xd3, 0x98, 0x70, 0x98, 0x66, 0xc0, 0xc5, 0xe7, 0x9e, 0x5a, 0x42, 0x8a, 0xfd,
	0x86, 0x5c, 0xfc, 0x70, 0x1d, 0x50, 0xcb, 0x09, 0x68, 0x31, 0x6a, 0x67, 0xe4, 0xe2, 0x06, 0x71,
	0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x6c, 0xb9, 0xd8, 0x20, 0x5e, 0x10, 0x12, 0x87,
	0xa8, 0xc5, 0x08, 0x05, 0x29, 0x09, 0x4c, 0x09, 0x88, 0x65, 0x4a, 0x0c, 0x42, 0x96, 0x5c, 0x6c,
	0xee, 0xa9, 0x25, 0x8e, 0x39, 0x39, 0x42, 0x22, 0x10, 0x55, 0xa8, 0x3e, 0x90, 0x12, 0x45, 0x13,
	0x85, 0x69, 0x34, 0x60, 0x74, 0xe2, 0x88, 0x62, 0x03, 0xc9, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x23,
	0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x57, 0xd8, 0x4b, 0xb2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	// unary
	Create(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	// server streaming
	GetAll(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (TodoService_GetAllClient, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Create(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetAll(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (TodoService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TodoService_serviceDesc.Streams[0], "/todo.TodoService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_GetAllClient interface {
	Recv() (*GetTodoResponse, error)
	grpc.ClientStream
}

type todoServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *todoServiceGetAllClient) Recv() (*GetTodoResponse, error) {
	m := new(GetTodoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	// unary
	Create(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	// server streaming
	GetAll(*GetTodoRequest, TodoService_GetAllServer) error
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Create(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTodoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).GetAll(m, &todoServiceGetAllServer{stream})
}

type TodoService_GetAllServer interface {
	Send(*GetTodoResponse) error
	grpc.ServerStream
}

type todoServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *todoServiceGetAllServer) Send(m *GetTodoResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _TodoService_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}
