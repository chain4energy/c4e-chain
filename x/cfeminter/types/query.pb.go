// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chain4energy/c4echain/cfeminter/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05b42f8a8fd1aaf3, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05b42f8a8fd1aaf3, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryInflationRequest struct {
}

func (m *QueryInflationRequest) Reset()         { *m = QueryInflationRequest{} }
func (m *QueryInflationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryInflationRequest) ProtoMessage()    {}
func (*QueryInflationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05b42f8a8fd1aaf3, []int{2}
}
func (m *QueryInflationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInflationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInflationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInflationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInflationRequest.Merge(m, src)
}
func (m *QueryInflationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryInflationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInflationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInflationRequest proto.InternalMessageInfo

type QueryInflationResponse struct {
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
}

func (m *QueryInflationResponse) Reset()         { *m = QueryInflationResponse{} }
func (m *QueryInflationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryInflationResponse) ProtoMessage()    {}
func (*QueryInflationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05b42f8a8fd1aaf3, []int{3}
}
func (m *QueryInflationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInflationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInflationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInflationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInflationResponse.Merge(m, src)
}
func (m *QueryInflationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryInflationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInflationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInflationResponse proto.InternalMessageInfo

type QueryStateRequest struct {
}

func (m *QueryStateRequest) Reset()         { *m = QueryStateRequest{} }
func (m *QueryStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStateRequest) ProtoMessage()    {}
func (*QueryStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05b42f8a8fd1aaf3, []int{4}
}
func (m *QueryStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStateRequest.Merge(m, src)
}
func (m *QueryStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStateRequest proto.InternalMessageInfo

type QueryStateResponse struct {
	MinterState  MinterState    `protobuf:"bytes,1,opt,name=minter_state,json=minterState,proto3" json:"minter_state"`
	StateHistory []*MinterState `protobuf:"bytes,2,rep,name=state_history,json=stateHistory,proto3" json:"state_history,omitempty"`
}

func (m *QueryStateResponse) Reset()         { *m = QueryStateResponse{} }
func (m *QueryStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStateResponse) ProtoMessage()    {}
func (*QueryStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05b42f8a8fd1aaf3, []int{5}
}
func (m *QueryStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStateResponse.Merge(m, src)
}
func (m *QueryStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStateResponse proto.InternalMessageInfo

func (m *QueryStateResponse) GetMinterState() MinterState {
	if m != nil {
		return m.MinterState
	}
	return MinterState{}
}

func (m *QueryStateResponse) GetStateHistory() []*MinterState {
	if m != nil {
		return m.StateHistory
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "chain4energy.c4echain.cfeminter.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "chain4energy.c4echain.cfeminter.QueryParamsResponse")
	proto.RegisterType((*QueryInflationRequest)(nil), "chain4energy.c4echain.cfeminter.QueryInflationRequest")
	proto.RegisterType((*QueryInflationResponse)(nil), "chain4energy.c4echain.cfeminter.QueryInflationResponse")
	proto.RegisterType((*QueryStateRequest)(nil), "chain4energy.c4echain.cfeminter.QueryStateRequest")
	proto.RegisterType((*QueryStateResponse)(nil), "chain4energy.c4echain.cfeminter.QueryStateResponse")
}

func init() {
	proto.RegisterFile("chain4energy/c4echain/cfeminter/query.proto", fileDescriptor_05b42f8a8fd1aaf3)
}

var fileDescriptor_05b42f8a8fd1aaf3 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0xb4, 0x36, 0x90, 0x69, 0x3d, 0x38, 0xad, 0x5a, 0xd7, 0xba, 0xa9, 0x0b, 0xfe, 0x41,
	0xdb, 0x1d, 0x9a, 0x84, 0x7a, 0x0f, 0x0a, 0x0a, 0x0a, 0x6d, 0xc4, 0x8b, 0x08, 0x65, 0xb2, 0x4e,
	0x36, 0x83, 0xdd, 0x99, 0xed, 0xce, 0xa4, 0x98, 0xab, 0x9f, 0x40, 0xe8, 0xc1, 0x0f, 0xe0, 0xf7,
	0xf0, 0xdc, 0x63, 0xc1, 0x8b, 0x08, 0x16, 0x49, 0xfc, 0x20, 0xb2, 0x6f, 0x26, 0xe9, 0x1a, 0x0b,
	0x49, 0x4e, 0x93, 0xbc, 0xfd, 0xbd, 0xdf, 0xef, 0xf7, 0xde, 0xbc, 0x37, 0xf8, 0x71, 0xd4, 0x65,
	0x42, 0x36, 0xb8, 0xe4, 0x59, 0xdc, 0xa7, 0x51, 0x83, 0xc3, 0x7f, 0x1a, 0x75, 0x78, 0x22, 0xa4,
	0xe1, 0x19, 0x3d, 0xea, 0xf1, 0xac, 0x1f, 0xa6, 0x99, 0x32, 0x8a, 0x54, 0x8b, 0xe0, 0x70, 0x04,
	0x0e, 0xc7, 0x60, 0x6f, 0x2d, 0x56, 0xb1, 0x02, 0x2c, 0xcd, 0x7f, 0xd9, 0x34, 0x6f, 0x23, 0x56,
	0x2a, 0x3e, 0xe4, 0x94, 0xa5, 0x82, 0x32, 0x29, 0x95, 0x61, 0x46, 0x28, 0xa9, 0xdd, 0xd7, 0x47,
	0x91, 0xd2, 0x89, 0xd2, 0xb4, 0xcd, 0x34, 0xb7, 0x6a, 0xf4, 0x78, 0xa7, 0xcd, 0x0d, 0xdb, 0xa1,
	0x29, 0x8b, 0x85, 0x04, 0xb0, 0xc3, 0x6e, 0x4d, 0x73, 0x9b, 0xb2, 0x8c, 0x25, 0x7a, 0x56, 0xb4,
	0x3d, 0x2c, 0x3a, 0x58, 0xc3, 0x64, 0x3f, 0x57, 0xdf, 0x03, 0x8a, 0x16, 0x3f, 0xea, 0x71, 0x6d,
	0x82, 0x77, 0x78, 0xf5, 0x9f, 0xa8, 0x4e, 0x95, 0xd4, 0x9c, 0x3c, 0xc3, 0x65, 0x2b, 0xb5, 0x8e,
	0x36, 0xd1, 0xc3, 0xe5, 0xda, 0x83, 0x70, 0x4a, 0x6b, 0x42, 0x4b, 0xd0, 0xbc, 0x72, 0x7a, 0x5e,
	0x2d, 0xb5, 0x5c, 0x72, 0x70, 0x13, 0x5f, 0x07, 0xf6, 0x17, 0xb2, 0x73, 0x08, 0x75, 0x8e, 0x64,
	0x3b, 0xf8, 0xc6, 0xe4, 0x07, 0xa7, 0xfc, 0x12, 0x57, 0xc4, 0x28, 0x08, 0xe2, 0x95, 0x66, 0x98,
	0x73, 0xfe, 0x3c, 0xaf, 0xde, 0x8f, 0x85, 0xe9, 0xf6, 0xda, 0x61, 0xa4, 0x12, 0xea, 0x9a, 0x6a,
	0x8f, 0x6d, 0xfd, 0xfe, 0x03, 0x35, 0xfd, 0x94, 0xeb, 0xf0, 0x29, 0x8f, 0x5a, 0x17, 0x04, 0xc1,
	0x2a, 0xbe, 0x06, 0x3a, 0xaf, 0x0d, 0x33, 0x7c, 0x24, 0xfe, 0x0d, 0xb9, 0x56, 0xb8, 0xa8, 0x53,
	0x7e, 0x83, 0x57, 0x6c, 0x2d, 0x07, 0x3a, 0x8f, 0xbb, 0xca, 0xb7, 0xa6, 0x56, 0xfe, 0x0a, 0x0e,
	0xe0, 0x72, 0xe5, 0x2f, 0x27, 0x17, 0x21, 0xb2, 0x8f, 0xaf, 0x02, 0xdf, 0x41, 0x57, 0x68, 0xa3,
	0xb2, 0xfe, 0xfa, 0xc2, 0xe6, 0xe2, 0xbc, 0xbc, 0xad, 0x15, 0xa0, 0x78, 0x6e, 0x19, 0x6a, 0xbf,
	0x16, 0xf1, 0x12, 0x14, 0x40, 0xbe, 0x20, 0x5c, 0xb6, 0x9d, 0x27, 0xf5, 0xa9, 0x84, 0xff, 0x5f,
	0xbf, 0xd7, 0x98, 0x2f, 0xc9, 0x76, 0x2a, 0x08, 0x3e, 0x7d, 0xff, 0x73, 0xb2, 0xb0, 0x41, 0xbc,
	0x7c, 0xe8, 0xdc, 0x94, 0x15, 0x86, 0x1a, 0xec, 0x7c, 0x45, 0xb8, 0x32, 0xbe, 0x5d, 0xb2, 0x3b,
	0x9b, 0xce, 0xe4, 0x9c, 0x78, 0x4f, 0xe6, 0xce, 0x73, 0x16, 0xef, 0x81, 0xc5, 0x2a, 0xb9, 0x73,
	0x99, 0xc5, 0xf1, 0x7c, 0x90, 0x13, 0x84, 0x97, 0xec, 0x35, 0xd5, 0x66, 0x53, 0x2a, 0x0e, 0x92,
	0x57, 0x9f, 0x2b, 0xc7, 0x39, 0xbb, 0x0b, 0xce, 0x6e, 0x93, 0x5b, 0x97, 0x39, 0x83, 0x6b, 0x6e,
	0xee, 0x9d, 0x0e, 0x7c, 0x74, 0x36, 0xf0, 0xd1, 0xef, 0x81, 0x8f, 0x3e, 0x0f, 0xfd, 0xd2, 0xd9,
	0xd0, 0x2f, 0xfd, 0x18, 0xfa, 0xa5, 0xb7, 0xbb, 0xc5, 0x15, 0x98, 0xd8, 0xfe, 0x6d, 0xbb, 0xfe,
	0xc7, 0x35, 0xfa, 0xb1, 0xf0, 0x06, 0xc0, 0x5a, 0xb4, 0xcb, 0xf0, 0x06, 0xd4, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xa6, 0x60, 0x3b, 0xf5, 0x0f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Inflation items.
	Inflation(ctx context.Context, in *QueryInflationRequest, opts ...grpc.CallOption) (*QueryInflationResponse, error)
	// Queries a list of State items.
	State(ctx context.Context, in *QueryStateRequest, opts ...grpc.CallOption) (*QueryStateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/chain4energy.c4echain.cfeminter.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Inflation(ctx context.Context, in *QueryInflationRequest, opts ...grpc.CallOption) (*QueryInflationResponse, error) {
	out := new(QueryInflationResponse)
	err := c.cc.Invoke(ctx, "/chain4energy.c4echain.cfeminter.Query/Inflation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) State(ctx context.Context, in *QueryStateRequest, opts ...grpc.CallOption) (*QueryStateResponse, error) {
	out := new(QueryStateResponse)
	err := c.cc.Invoke(ctx, "/chain4energy.c4echain.cfeminter.Query/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Inflation items.
	Inflation(context.Context, *QueryInflationRequest) (*QueryInflationResponse, error)
	// Queries a list of State items.
	State(context.Context, *QueryStateRequest) (*QueryStateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Inflation(ctx context.Context, req *QueryInflationRequest) (*QueryInflationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inflation not implemented")
}
func (*UnimplementedQueryServer) State(ctx context.Context, req *QueryStateRequest) (*QueryStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain4energy.c4echain.cfeminter.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Inflation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInflationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Inflation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain4energy.c4echain.cfeminter.Query/Inflation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Inflation(ctx, req.(*QueryInflationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain4energy.c4echain.cfeminter.Query/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).State(ctx, req.(*QueryStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chain4energy.c4echain.cfeminter.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Inflation",
			Handler:    _Query_Inflation_Handler,
		},
		{
			MethodName: "State",
			Handler:    _Query_State_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chain4energy/c4echain/cfeminter/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryInflationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInflationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInflationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryInflationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInflationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInflationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StateHistory) > 0 {
		for iNdEx := len(m.StateHistory) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StateHistory[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.MinterState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryInflationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryInflationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinterState.Size()
	n += 1 + l + sovQuery(uint64(l))
	if len(m.StateHistory) > 0 {
		for _, e := range m.StateHistory {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryInflationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInflationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInflationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryInflationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInflationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInflationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinterState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateHistory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateHistory = append(m.StateHistory, &MinterState{})
			if err := m.StateHistory[len(m.StateHistory)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
