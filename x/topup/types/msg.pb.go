// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/topup/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the topup module's genesis state.
type MsgTopup struct {
	FromAddress string                                 `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	User        string                                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Fee         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"varint,3,opt,name=fee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee,omitempty"`
	TxHash      string                                 `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"tx_hash"`
	LogIndex    uint64                                 `protobuf:"varint,5,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty" yaml:"log_index"`
	BlockNumber uint64                                 `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty" yaml:"block_number"`
}

func (m *MsgTopup) Reset()         { *m = MsgTopup{} }
func (m *MsgTopup) String() string { return proto.CompactTextString(m) }
func (*MsgTopup) ProtoMessage()    {}
func (*MsgTopup) Descriptor() ([]byte, []int) {
	return fileDescriptor_944a1ef1f3e2d8aa, []int{0}
}
func (m *MsgTopup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTopup.Unmarshal(m, b)
}
func (m *MsgTopup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTopup.Marshal(b, m, deterministic)
}
func (m *MsgTopup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTopup.Merge(m, src)
}
func (m *MsgTopup) XXX_Size() int {
	return xxx_messageInfo_MsgTopup.Size(m)
}
func (m *MsgTopup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTopup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTopup proto.InternalMessageInfo

type MsgTopupResponse struct {
}

func (m *MsgTopupResponse) Reset()         { *m = MsgTopupResponse{} }
func (m *MsgTopupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTopupResponse) ProtoMessage()    {}
func (*MsgTopupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_944a1ef1f3e2d8aa, []int{1}
}
func (m *MsgTopupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTopupResponse.Unmarshal(m, b)
}
func (m *MsgTopupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTopupResponse.Marshal(b, m, deterministic)
}
func (m *MsgTopupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTopupResponse.Merge(m, src)
}
func (m *MsgTopupResponse) XXX_Size() int {
	return xxx_messageInfo_MsgTopupResponse.Size(m)
}
func (m *MsgTopupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTopupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTopupResponse proto.InternalMessageInfo

type MsgWithdrawFee struct {
	UserAddress string                                 `protobuf:"bytes,1,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty" yaml:"user_address"`
	Amount      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"varint,2,opt,name=amount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount,omitempty"`
}

func (m *MsgWithdrawFee) Reset()         { *m = MsgWithdrawFee{} }
func (m *MsgWithdrawFee) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFee) ProtoMessage()    {}
func (*MsgWithdrawFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_944a1ef1f3e2d8aa, []int{2}
}
func (m *MsgWithdrawFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithdrawFee.Unmarshal(m, b)
}
func (m *MsgWithdrawFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithdrawFee.Marshal(b, m, deterministic)
}
func (m *MsgWithdrawFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFee.Merge(m, src)
}
func (m *MsgWithdrawFee) XXX_Size() int {
	return xxx_messageInfo_MsgWithdrawFee.Size(m)
}
func (m *MsgWithdrawFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFee proto.InternalMessageInfo

type MsgWithdrawFeeResponse struct {
}

func (m *MsgWithdrawFeeResponse) Reset()         { *m = MsgWithdrawFeeResponse{} }
func (m *MsgWithdrawFeeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFeeResponse) ProtoMessage()    {}
func (*MsgWithdrawFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_944a1ef1f3e2d8aa, []int{3}
}
func (m *MsgWithdrawFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithdrawFeeResponse.Unmarshal(m, b)
}
func (m *MsgWithdrawFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithdrawFeeResponse.Marshal(b, m, deterministic)
}
func (m *MsgWithdrawFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFeeResponse.Merge(m, src)
}
func (m *MsgWithdrawFeeResponse) XXX_Size() int {
	return xxx_messageInfo_MsgWithdrawFeeResponse.Size(m)
}
func (m *MsgWithdrawFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFeeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTopup)(nil), "heimdall.topup.v1beta1.MsgTopup")
	proto.RegisterType((*MsgTopupResponse)(nil), "heimdall.topup.v1beta1.MsgTopupResponse")
	proto.RegisterType((*MsgWithdrawFee)(nil), "heimdall.topup.v1beta1.MsgWithdrawFee")
	proto.RegisterType((*MsgWithdrawFeeResponse)(nil), "heimdall.topup.v1beta1.MsgWithdrawFeeResponse")
}

func init() { proto.RegisterFile("heimdall/topup/v1beta1/msg.proto", fileDescriptor_944a1ef1f3e2d8aa) }

var fileDescriptor_944a1ef1f3e2d8aa = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xed, 0x26, 0x0d, 0xed, 0x05, 0x55, 0xd5, 0x51, 0x15, 0x2b, 0x83, 0x1d, 0xdd, 0x50,
	0x45, 0x20, 0xce, 0x0a, 0x6c, 0x11, 0x0b, 0x19, 0x50, 0x3b, 0x84, 0xc1, 0x20, 0x21, 0xb1, 0x44,
	0xe7, 0xf8, 0x7a, 0xb6, 0xe2, 0xf3, 0x59, 0xbe, 0x33, 0x4d, 0xff, 0x03, 0x46, 0x26, 0x66, 0xfe,
	0x0e, 0x36, 0x36, 0xc6, 0x8e, 0x4c, 0x11, 0x4a, 0xfe, 0x83, 0x8c, 0x4c, 0xe8, 0xce, 0x76, 0x65,
	0x24, 0x7e, 0x4e, 0x7e, 0x4f, 0xef, 0xf3, 0xfc, 0xbe, 0xef, 0xab, 0x7b, 0x60, 0x18, 0xd3, 0x84,
	0x47, 0x24, 0x4d, 0x7d, 0x25, 0xf2, 0x32, 0xf7, 0xdf, 0x8e, 0x43, 0xaa, 0xc8, 0xd8, 0xe7, 0x92,
	0xe1, 0xbc, 0x10, 0x4a, 0xc0, 0xd3, 0x86, 0xc0, 0x86, 0xc0, 0x35, 0x31, 0x38, 0x61, 0x82, 0x09,
	0x83, 0xf8, 0x3a, 0xaa, 0x68, 0xf4, 0x69, 0x0f, 0x1c, 0xcc, 0x24, 0x7b, 0xa5, 0x51, 0x38, 0x01,
	0x77, 0x2f, 0x0b, 0xc1, 0xe7, 0x24, 0x8a, 0x0a, 0x2a, 0xa5, 0x63, 0x0f, 0xed, 0xd1, 0xe1, 0xf4,
	0xfe, 0x6e, 0xed, 0xdd, 0xbb, 0x26, 0x3c, 0x9d, 0xa0, 0x76, 0x15, 0x05, 0x7d, 0x9d, 0x3e, 0xab,
	0x32, 0x08, 0x41, 0xb7, 0x94, 0xb4, 0x70, 0xf6, 0x74, 0x4f, 0x60, 0x62, 0xf8, 0x14, 0x74, 0x2e,
	0x29, 0x75, 0x3a, 0x43, 0x7b, 0xd4, 0x9d, 0x3e, 0xf8, 0xbe, 0xf6, 0xce, 0x58, 0xa2, 0xe2, 0x32,
	0xc4, 0x0b, 0xc1, 0xfd, 0x85, 0x90, 0x5c, 0xc8, 0xfa, 0xf3, 0x48, 0x46, 0x4b, 0x5f, 0x5d, 0xe7,
	0x54, 0xe2, 0x8b, 0x4c, 0x05, 0xba, 0x0d, 0x3e, 0x04, 0x77, 0xd4, 0x6a, 0x1e, 0x13, 0x19, 0x3b,
	0x5d, 0x23, 0x04, 0xee, 0xd6, 0xde, 0x51, 0x25, 0xa4, 0x2e, 0xa0, 0xa0, 0xa7, 0x56, 0xe7, 0x44,
	0xc6, 0x70, 0x0c, 0x0e, 0x53, 0xc1, 0xe6, 0x49, 0x16, 0xd1, 0x95, 0xb3, 0x6f, 0x06, 0x9e, 0xec,
	0xd6, 0xde, 0x71, 0x85, 0xdf, 0x96, 0x50, 0x70, 0x90, 0x0a, 0x76, 0xa1, 0x43, 0xbd, 0x6d, 0x98,
	0x8a, 0xc5, 0x72, 0x9e, 0x95, 0x3c, 0xa4, 0x85, 0xd3, 0x33, 0x5d, 0xad, 0x6d, 0xdb, 0x55, 0x14,
	0xf4, 0x4d, 0xfa, 0xc2, 0x64, 0x93, 0xee, 0xbb, 0x8f, 0x9e, 0x85, 0x20, 0x38, 0x6e, 0xbc, 0x0b,
	0xa8, 0xcc, 0x45, 0x26, 0x29, 0xfa, 0x60, 0x83, 0xa3, 0x99, 0x64, 0xaf, 0x13, 0x15, 0x47, 0x05,
	0xb9, 0x7a, 0x4e, 0xa9, 0x1e, 0xa4, 0xed, 0xf8, 0xbd, 0xad, 0xed, 0x2a, 0x0a, 0xfa, 0x3a, 0x6d,
	0x6c, 0x9d, 0x82, 0x1e, 0xe1, 0xa2, 0xcc, 0x94, 0x31, 0xf6, 0xff, 0x5c, 0xac, 0x3b, 0x6b, 0xb1,
	0x0e, 0x38, 0xfd, 0x59, 0x57, 0x23, 0xf9, 0xf1, 0x67, 0x1b, 0x74, 0x66, 0x92, 0xc1, 0x97, 0x60,
	0xbf, 0x7a, 0x07, 0x43, 0xfc, 0xeb, 0x37, 0x84, 0x9b, 0x6d, 0x07, 0xa3, 0xbf, 0x11, 0xcd, 0xcf,
	0x21, 0x05, 0xfd, 0xb6, 0x17, 0x67, 0x7f, 0x68, 0x6c, 0x71, 0x03, 0xfc, 0x6f, 0x5c, 0x33, 0x66,
	0x7a, 0xfe, 0x65, 0xe3, 0x5a, 0x37, 0x1b, 0xd7, 0xfa, 0xb6, 0x71, 0xad, 0xf7, 0x5b, 0xd7, 0xba,
	0xd9, 0xba, 0xd6, 0xd7, 0xad, 0x6b, 0xbd, 0xc1, 0x2d, 0xb7, 0x38, 0x51, 0xc9, 0x22, 0xa3, 0xea,
	0x4a, 0x14, 0x4b, 0xff, 0xf6, 0x92, 0x56, 0xf5, 0x2d, 0x19, 0xe7, 0xc2, 0x9e, 0x39, 0x8c, 0x27,
	0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xce, 0xa5, 0x13, 0x6a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Topup defines a method to topup validator account with additional tokens.
	Topup(ctx context.Context, in *MsgTopup, opts ...grpc.CallOption) (*MsgTopupResponse, error)
	// WithdrawFee defines a method for validators to withdraw tokens from
	// heimdall.
	WithdrawFee(ctx context.Context, in *MsgWithdrawFee, opts ...grpc.CallOption) (*MsgWithdrawFeeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Topup(ctx context.Context, in *MsgTopup, opts ...grpc.CallOption) (*MsgTopupResponse, error) {
	out := new(MsgTopupResponse)
	err := c.cc.Invoke(ctx, "/heimdall.topup.v1beta1.Msg/Topup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawFee(ctx context.Context, in *MsgWithdrawFee, opts ...grpc.CallOption) (*MsgWithdrawFeeResponse, error) {
	out := new(MsgWithdrawFeeResponse)
	err := c.cc.Invoke(ctx, "/heimdall.topup.v1beta1.Msg/WithdrawFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Topup defines a method to topup validator account with additional tokens.
	Topup(context.Context, *MsgTopup) (*MsgTopupResponse, error)
	// WithdrawFee defines a method for validators to withdraw tokens from
	// heimdall.
	WithdrawFee(context.Context, *MsgWithdrawFee) (*MsgWithdrawFeeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Topup(ctx context.Context, req *MsgTopup) (*MsgTopupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topup not implemented")
}
func (*UnimplementedMsgServer) WithdrawFee(ctx context.Context, req *MsgWithdrawFee) (*MsgWithdrawFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFee not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Topup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTopup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Topup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.topup.v1beta1.Msg/Topup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Topup(ctx, req.(*MsgTopup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.topup.v1beta1.Msg/WithdrawFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawFee(ctx, req.(*MsgWithdrawFee))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.topup.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Topup",
			Handler:    _Msg_Topup_Handler,
		},
		{
			MethodName: "WithdrawFee",
			Handler:    _Msg_WithdrawFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heimdall/topup/v1beta1/msg.proto",
}
