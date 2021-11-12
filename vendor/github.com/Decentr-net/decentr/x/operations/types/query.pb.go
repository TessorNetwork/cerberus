// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operations/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MinGasPriceRequest struct {
}

func (m *MinGasPriceRequest) Reset()         { *m = MinGasPriceRequest{} }
func (m *MinGasPriceRequest) String() string { return proto.CompactTextString(m) }
func (*MinGasPriceRequest) ProtoMessage()    {}
func (*MinGasPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2686d2c3e5e5e856, []int{0}
}
func (m *MinGasPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinGasPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinGasPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinGasPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinGasPriceRequest.Merge(m, src)
}
func (m *MinGasPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *MinGasPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MinGasPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MinGasPriceRequest proto.InternalMessageInfo

type MinGasPriceResponse struct {
	MinGasPrice types.DecCoin `protobuf:"bytes,1,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price"`
}

func (m *MinGasPriceResponse) Reset()         { *m = MinGasPriceResponse{} }
func (m *MinGasPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MinGasPriceResponse) ProtoMessage()    {}
func (*MinGasPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2686d2c3e5e5e856, []int{1}
}
func (m *MinGasPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinGasPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinGasPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinGasPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinGasPriceResponse.Merge(m, src)
}
func (m *MinGasPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MinGasPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MinGasPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MinGasPriceResponse proto.InternalMessageInfo

func (m *MinGasPriceResponse) GetMinGasPrice() types.DecCoin {
	if m != nil {
		return m.MinGasPrice
	}
	return types.DecCoin{}
}

type IsAccountBannedRequest struct {
	// address is the address to query balances for.
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
}

func (m *IsAccountBannedRequest) Reset()         { *m = IsAccountBannedRequest{} }
func (m *IsAccountBannedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAccountBannedRequest) ProtoMessage()    {}
func (*IsAccountBannedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2686d2c3e5e5e856, []int{2}
}
func (m *IsAccountBannedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IsAccountBannedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IsAccountBannedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IsAccountBannedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAccountBannedRequest.Merge(m, src)
}
func (m *IsAccountBannedRequest) XXX_Size() int {
	return m.Size()
}
func (m *IsAccountBannedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAccountBannedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAccountBannedRequest proto.InternalMessageInfo

type IsAccountBannedResponse struct {
	IsBanned bool `protobuf:"varint,1,opt,name=is_banned,json=isBanned,proto3" json:"is_banned,omitempty"`
}

func (m *IsAccountBannedResponse) Reset()         { *m = IsAccountBannedResponse{} }
func (m *IsAccountBannedResponse) String() string { return proto.CompactTextString(m) }
func (*IsAccountBannedResponse) ProtoMessage()    {}
func (*IsAccountBannedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2686d2c3e5e5e856, []int{3}
}
func (m *IsAccountBannedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IsAccountBannedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IsAccountBannedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IsAccountBannedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAccountBannedResponse.Merge(m, src)
}
func (m *IsAccountBannedResponse) XXX_Size() int {
	return m.Size()
}
func (m *IsAccountBannedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAccountBannedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAccountBannedResponse proto.InternalMessageInfo

func (m *IsAccountBannedResponse) GetIsBanned() bool {
	if m != nil {
		return m.IsBanned
	}
	return false
}

func init() {
	proto.RegisterType((*MinGasPriceRequest)(nil), "operations.MinGasPriceRequest")
	proto.RegisterType((*MinGasPriceResponse)(nil), "operations.MinGasPriceResponse")
	proto.RegisterType((*IsAccountBannedRequest)(nil), "operations.IsAccountBannedRequest")
	proto.RegisterType((*IsAccountBannedResponse)(nil), "operations.IsAccountBannedResponse")
}

func init() { proto.RegisterFile("operations/query.proto", fileDescriptor_2686d2c3e5e5e856) }

var fileDescriptor_2686d2c3e5e5e856 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xbf, 0x8b, 0xd4, 0x40,
	0x14, 0x4e, 0x0e, 0x7f, 0xac, 0xb3, 0x8a, 0x30, 0xca, 0x29, 0xeb, 0x31, 0xd1, 0x5c, 0xa3, 0x45,
	0x66, 0xc8, 0x09, 0x0a, 0x76, 0xbb, 0x1e, 0x8a, 0x1c, 0x82, 0x6e, 0x29, 0xc8, 0x32, 0x99, 0x0c,
	0x71, 0xd0, 0xcc, 0xcb, 0xe5, 0xcd, 0x8a, 0x8b, 0xd8, 0x58, 0x59, 0x0a, 0xd7, 0x58, 0xde, 0x9f,
	0x73, 0xe5, 0x81, 0x8d, 0xd5, 0x21, 0xbb, 0x16, 0xfe, 0x0d, 0x56, 0xb2, 0x99, 0xe8, 0x46, 0xef,
	0xb4, 0x4a, 0x78, 0xdf, 0x9b, 0xef, 0xfb, 0xde, 0xf7, 0x1e, 0x59, 0x87, 0x4a, 0xd7, 0xd2, 0x19,
	0xb0, 0x28, 0x76, 0xa7, 0xba, 0x9e, 0xf1, 0xaa, 0x06, 0x07, 0x94, 0xac, 0xea, 0x83, 0xcb, 0x05,
	0x14, 0xd0, 0x94, 0xc5, 0xf2, 0xcf, 0x77, 0x0c, 0x36, 0x0a, 0x80, 0xe2, 0x95, 0x16, 0xb2, 0x32,
	0x42, 0x5a, 0x0b, 0xce, 0x77, 0xb7, 0x28, 0x53, 0x80, 0x25, 0xa0, 0xc8, 0x24, 0x6a, 0xf1, 0x3a,
	0xcd, 0xb4, 0x93, 0xa9, 0x50, 0x60, 0xac, 0xc7, 0x63, 0x46, 0xe8, 0x63, 0x63, 0x1f, 0x4a, 0x7c,
	0x52, 0x1b, 0xa5, 0xc7, 0x7a, 0x77, 0xaa, 0xd1, 0xdd, 0xeb, 0x7d, 0xd8, 0x8f, 0x82, 0xef, 0xfb,
	0x51, 0x10, 0x3f, 0x27, 0x97, 0xfe, 0xc0, 0xb1, 0x02, 0x8b, 0x9a, 0x3e, 0x20, 0x17, 0x4a, 0x63,
	0x27, 0x85, 0xc4, 0x49, 0xb5, 0x04, 0xae, 0x86, 0xd7, 0xc3, 0x9b, 0xfd, 0xad, 0x0d, 0xee, 0xe5,
	0xf8, 0x52, 0x8e, 0xb7, 0x72, 0x7c, 0x5b, 0xab, 0xfb, 0x60, 0xec, 0xe8, 0xd4, 0xc1, 0x51, 0x14,
	0x8c, 0xfb, 0xe5, 0x8a, 0x2f, 0x06, 0xb2, 0xfe, 0x08, 0x87, 0x4a, 0xc1, 0xd4, 0xba, 0x91, 0xb4,
	0x56, 0xe7, 0xad, 0x05, 0xba, 0x43, 0xce, 0xca, 0x3c, 0xaf, 0x35, 0x62, 0xc3, 0x7d, 0x7e, 0x94,
	0xfe, 0x38, 0x8a, 0x92, 0xc2, 0xb8, 0x17, 0xd3, 0x8c, 0x2b, 0x28, 0x45, 0x3b, 0x98, 0xff, 0x24,
	0x98, 0xbf, 0x14, 0x6e, 0x56, 0x69, 0xe4, 0x43, 0xa5, 0x86, 0xfe, 0xe1, 0xf8, 0x17, 0x43, 0x67,
	0x9e, 0x3b, 0xe4, 0xca, 0x31, 0xc1, 0x76, 0xa6, 0x6b, 0xe4, 0x9c, 0xc1, 0x49, 0xd6, 0x14, 0x1b,
	0xcd, 0xde, 0xb8, 0x67, 0xd0, 0x37, 0x6d, 0x7d, 0x5a, 0x23, 0xa7, 0x9f, 0x2e, 0xf7, 0x42, 0x67,
	0xa4, 0xdf, 0x49, 0x84, 0x32, 0xbe, 0xda, 0x10, 0x3f, 0x1e, 0xe5, 0x20, 0xfa, 0x27, 0xee, 0x65,
	0xe3, 0x5b, 0xef, 0x3f, 0x7f, 0xdb, 0x5b, 0xdb, 0xa4, 0x37, 0x44, 0xae, 0x95, 0xb6, 0xae, 0x16,
	0x9d, 0x53, 0x28, 0x8d, 0x4d, 0x0a, 0x89, 0x49, 0x13, 0x32, 0xdd, 0x0b, 0xc9, 0xc5, 0xbf, 0xdc,
	0xd3, 0xb8, 0xcb, 0x7f, 0x72, 0x96, 0x83, 0xcd, 0xff, 0xf6, 0xb4, 0x3e, 0xee, 0x36, 0x3e, 0x52,
	0x2a, 0x4e, 0xf2, 0x61, 0x30, 0x91, 0xfe, 0x55, 0xe2, 0x03, 0x12, 0x6f, 0xdb, 0x6c, 0xdf, 0x8d,
	0x76, 0x0e, 0xe6, 0x2c, 0x3c, 0x9c, 0xb3, 0xf0, 0xeb, 0x9c, 0x85, 0x1f, 0x17, 0x2c, 0x38, 0x5c,
	0xb0, 0xe0, 0xcb, 0x82, 0x05, 0xcf, 0xd2, 0xce, 0xba, 0xb6, 0x3d, 0x69, 0x62, 0xb5, 0xfb, 0x2d,
	0xf0, 0xa6, 0x2b, 0xd1, 0x6c, 0x2f, 0x3b, 0xd3, 0x9c, 0xe5, 0xed, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x17, 0x28, 0x94, 0xd5, 0x10, 0x03, 0x00, 0x00,
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
	// MinGasPrice returns set gas price
	MinGasPrice(ctx context.Context, in *MinGasPriceRequest, opts ...grpc.CallOption) (*MinGasPriceResponse, error)
	// IsAccountBanned
	IsAccountBanned(ctx context.Context, in *IsAccountBannedRequest, opts ...grpc.CallOption) (*IsAccountBannedResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) MinGasPrice(ctx context.Context, in *MinGasPriceRequest, opts ...grpc.CallOption) (*MinGasPriceResponse, error) {
	out := new(MinGasPriceResponse)
	err := c.cc.Invoke(ctx, "/operations.Query/MinGasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IsAccountBanned(ctx context.Context, in *IsAccountBannedRequest, opts ...grpc.CallOption) (*IsAccountBannedResponse, error) {
	out := new(IsAccountBannedResponse)
	err := c.cc.Invoke(ctx, "/operations.Query/IsAccountBanned", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// MinGasPrice returns set gas price
	MinGasPrice(context.Context, *MinGasPriceRequest) (*MinGasPriceResponse, error)
	// IsAccountBanned
	IsAccountBanned(context.Context, *IsAccountBannedRequest) (*IsAccountBannedResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) MinGasPrice(ctx context.Context, req *MinGasPriceRequest) (*MinGasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinGasPrice not implemented")
}
func (*UnimplementedQueryServer) IsAccountBanned(ctx context.Context, req *IsAccountBannedRequest) (*IsAccountBannedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAccountBanned not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_MinGasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinGasPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinGasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operations.Query/MinGasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinGasPrice(ctx, req.(*MinGasPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IsAccountBanned_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAccountBannedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IsAccountBanned(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operations.Query/IsAccountBanned",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IsAccountBanned(ctx, req.(*IsAccountBannedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "operations.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MinGasPrice",
			Handler:    _Query_MinGasPrice_Handler,
		},
		{
			MethodName: "IsAccountBanned",
			Handler:    _Query_IsAccountBanned_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operations/query.proto",
}

func (m *MinGasPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinGasPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinGasPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MinGasPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinGasPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinGasPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinGasPrice.MarshalToSizedBuffer(dAtA[:i])
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

func (m *IsAccountBannedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IsAccountBannedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IsAccountBannedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IsAccountBannedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IsAccountBannedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IsAccountBannedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsBanned {
		i--
		if m.IsBanned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
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
func (m *MinGasPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MinGasPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinGasPrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *IsAccountBannedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *IsAccountBannedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsBanned {
		n += 2
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MinGasPriceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MinGasPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinGasPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MinGasPriceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MinGasPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinGasPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
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
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *IsAccountBannedRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IsAccountBannedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IsAccountBannedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
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
func (m *IsAccountBannedResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IsAccountBannedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IsAccountBannedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBanned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsBanned = bool(v != 0)
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
