// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aura/smartaccount/v1beta1/tx.proto

package v1beta1

import (
	context "context"
	fmt "fmt"
	github_com_CosmWasm_wasmd_x_wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgRecover struct {
	// Sender is the actor who signs the message
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// smart-account address that need update
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// New PubKey using for signature verification of this account
	PubKey *types.Any `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Credentials
	Credentials string `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (m *MsgRecover) Reset()         { *m = MsgRecover{} }
func (m *MsgRecover) String() string { return proto.CompactTextString(m) }
func (*MsgRecover) ProtoMessage()    {}
func (*MsgRecover) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3865272ae9f517a, []int{0}
}
func (m *MsgRecover) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRecover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRecover.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRecover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRecover.Merge(m, src)
}
func (m *MsgRecover) XXX_Size() int {
	return m.Size()
}
func (m *MsgRecover) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRecover.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRecover proto.InternalMessageInfo

type MsgRecoverResponse struct {
}

func (m *MsgRecoverResponse) Reset()         { *m = MsgRecoverResponse{} }
func (m *MsgRecoverResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRecoverResponse) ProtoMessage()    {}
func (*MsgRecoverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3865272ae9f517a, []int{1}
}
func (m *MsgRecoverResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRecoverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRecoverResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRecoverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRecoverResponse.Merge(m, src)
}
func (m *MsgRecoverResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRecoverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRecoverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRecoverResponse proto.InternalMessageInfo

type MsgActivateAccount struct {
	// AccountAddress is the actor who signs the message
	AccountAddress string `protobuf:"bytes,1,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	// CodeID indicates which wasm binary code is to be used for this contract
	CodeID uint64 `protobuf:"varint,3,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	// an arbitrary value provided by the sender. Size can be 1 to 64.
	Salt []byte `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	// InitMsg is the JSON-encoded instantiate message for the contract
	InitMsg github_com_CosmWasm_wasmd_x_wasm_types.RawContractMessage `protobuf:"bytes,5,opt,name=init_msg,json=initMsg,proto3,casttype=github.com/CosmWasm/wasmd/x/wasm/types.RawContractMessage" json:"init_msg,omitempty"`
	// Public key of smart account
	PubKey *types.Any `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (m *MsgActivateAccount) Reset()         { *m = MsgActivateAccount{} }
func (m *MsgActivateAccount) String() string { return proto.CompactTextString(m) }
func (*MsgActivateAccount) ProtoMessage()    {}
func (*MsgActivateAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3865272ae9f517a, []int{2}
}
func (m *MsgActivateAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgActivateAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgActivateAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgActivateAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgActivateAccount.Merge(m, src)
}
func (m *MsgActivateAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgActivateAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgActivateAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgActivateAccount proto.InternalMessageInfo

func (m *MsgActivateAccount) GetAccountAddress() string {
	if m != nil {
		return m.AccountAddress
	}
	return ""
}

func (m *MsgActivateAccount) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *MsgActivateAccount) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *MsgActivateAccount) GetInitMsg() github_com_CosmWasm_wasmd_x_wasm_types.RawContractMessage {
	if m != nil {
		return m.InitMsg
	}
	return nil
}

func (m *MsgActivateAccount) GetPubKey() *types.Any {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type MsgActivateAccountResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgActivateAccountResponse) Reset()         { *m = MsgActivateAccountResponse{} }
func (m *MsgActivateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgActivateAccountResponse) ProtoMessage()    {}
func (*MsgActivateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3865272ae9f517a, []int{3}
}
func (m *MsgActivateAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgActivateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgActivateAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgActivateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgActivateAccountResponse.Merge(m, src)
}
func (m *MsgActivateAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgActivateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgActivateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgActivateAccountResponse proto.InternalMessageInfo

func (m *MsgActivateAccountResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRecover)(nil), "aura.smartaccount.v1beta1.MsgRecover")
	proto.RegisterType((*MsgRecoverResponse)(nil), "aura.smartaccount.v1beta1.MsgRecoverResponse")
	proto.RegisterType((*MsgActivateAccount)(nil), "aura.smartaccount.v1beta1.MsgActivateAccount")
	proto.RegisterType((*MsgActivateAccountResponse)(nil), "aura.smartaccount.v1beta1.MsgActivateAccountResponse")
}

func init() {
	proto.RegisterFile("aura/smartaccount/v1beta1/tx.proto", fileDescriptor_b3865272ae9f517a)
}

var fileDescriptor_b3865272ae9f517a = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x4f, 0xb6, 0xd2, 0x6c, 0x1e, 0xda, 0x24, 0xab, 0x87, 0x2c, 0x87, 0xa4, 0x0a, 0x42, 0xda,
	0x65, 0xb6, 0x36, 0xfe, 0x48, 0xc0, 0x01, 0xb5, 0xe5, 0x40, 0x35, 0x15, 0xa1, 0x5c, 0x90, 0xb8,
	0x44, 0x8e, 0x63, 0x42, 0x44, 0x13, 0x57, 0xb1, 0xd3, 0x36, 0x6f, 0xc0, 0x91, 0x47, 0xd8, 0x0b,
	0x20, 0xf1, 0x18, 0x1c, 0x77, 0xe4, 0x54, 0x50, 0x7a, 0xe1, 0x19, 0x38, 0xa1, 0x38, 0x0d, 0x6b,
	0x8b, 0x10, 0xdb, 0xc9, 0xfe, 0xbe, 0xef, 0xe7, 0xef, 0xfb, 0xf9, 0xf7, 0xb3, 0x81, 0x4b, 0xf2,
	0x8c, 0x60, 0x91, 0x90, 0x4c, 0x12, 0x4a, 0x79, 0x9e, 0x4a, 0x3c, 0x3d, 0x0b, 0x98, 0x24, 0x67,
	0x58, 0xce, 0xd1, 0x24, 0xe3, 0x92, 0xc3, 0xe3, 0x0a, 0x83, 0xd6, 0x31, 0x68, 0x85, 0xb1, 0x3a,
	0x11, 0x8f, 0xb8, 0x42, 0xe1, 0x6a, 0x57, 0x1f, 0xb0, 0x8e, 0x23, 0xce, 0xa3, 0x31, 0xc3, 0x2a,
	0x0a, 0xf2, 0x77, 0x98, 0xa4, 0x45, 0x5d, 0x72, 0x3f, 0xeb, 0x00, 0x8c, 0x44, 0xe4, 0x31, 0xca,
	0xa7, 0x2c, 0x83, 0x26, 0x30, 0x68, 0xc6, 0x88, 0xe4, 0x99, 0xa9, 0x77, 0xf5, 0x93, 0x7d, 0xaf,
	0x09, 0xab, 0x0a, 0x09, 0xc3, 0x8c, 0x09, 0x61, 0xee, 0xd4, 0x95, 0x55, 0x08, 0x9f, 0x03, 0x30,
	0xc9, 0x83, 0x71, 0x4c, 0xfd, 0x0f, 0xac, 0x30, 0x77, 0xbb, 0xfa, 0xc9, 0xc1, 0x79, 0x07, 0xd5,
	0x23, 0x51, 0x33, 0x12, 0xf5, 0xd2, 0xa2, 0x0f, 0xca, 0x85, 0xd3, 0x7e, 0x9d, 0x07, 0x17, 0xac,
	0xf0, 0xf6, 0xeb, 0x33, 0x17, 0xac, 0x80, 0x5d, 0x70, 0x40, 0x33, 0x16, 0xb2, 0x54, 0xc6, 0x64,
	0x2c, 0xcc, 0x96, 0x6a, 0xbf, 0x9e, 0x7a, 0xba, 0xf7, 0xf1, 0xd2, 0xd1, 0x7e, 0x5e, 0x3a, 0x9a,
	0xdb, 0x01, 0xf0, 0x9a, 0xae, 0xc7, 0xc4, 0x84, 0xa7, 0x82, 0xb9, 0x5f, 0x76, 0x54, 0xba, 0x47,
	0x65, 0x3c, 0x25, 0x92, 0xf5, 0x6a, 0x55, 0xe0, 0x33, 0x70, 0xb4, 0x12, 0xc8, 0x6f, 0xb8, 0xab,
	0x5b, 0xf5, 0x61, 0xb9, 0x70, 0x0e, 0x57, 0xa8, 0x5e, 0x5d, 0xf1, 0x0e, 0xc9, 0x46, 0x0c, 0xef,
	0x01, 0x83, 0xf2, 0x90, 0xf9, 0x71, 0xa8, 0xee, 0xd4, 0xaa, 0xd9, 0x0f, 0x78, 0xc8, 0x86, 0x2f,
	0xbc, 0x76, 0x55, 0x1a, 0x86, 0x10, 0x82, 0x96, 0x20, 0x63, 0xa9, 0x24, 0xb9, 0xeb, 0xa9, 0x3d,
	0xa4, 0x60, 0x2f, 0x4e, 0x63, 0xe9, 0x27, 0x22, 0x32, 0xef, 0x54, 0xf9, 0xfe, 0xcb, 0x72, 0xe1,
	0x18, 0xc3, 0x34, 0x96, 0x23, 0x11, 0xfd, 0x5a, 0x38, 0x4f, 0xa2, 0x58, 0xbe, 0xcf, 0x03, 0x44,
	0x79, 0x82, 0x07, 0x5c, 0x24, 0x6f, 0x88, 0x48, 0xf0, 0x8c, 0x88, 0x24, 0xc4, 0x73, 0xb5, 0x62,
	0x59, 0x4c, 0x98, 0x40, 0x1e, 0x99, 0x0d, 0x78, 0x2a, 0x33, 0x42, 0xe5, 0x88, 0x09, 0x41, 0x22,
	0xe6, 0x19, 0x71, 0xdd, 0x65, 0x4b, 0xf4, 0xd6, 0xad, 0x45, 0x77, 0x1f, 0x03, 0xeb, 0x6f, 0xc5,
	0x1a, 0x41, 0xd7, 0xdd, 0xd6, 0x37, 0xdc, 0x3e, 0xff, 0xae, 0x83, 0xdd, 0x8a, 0x80, 0x0f, 0x8c,
	0xe6, 0xd1, 0xdc, 0x47, 0xff, 0x7c, 0x90, 0xe8, 0xda, 0x2c, 0xeb, 0xf4, 0x46, 0xb0, 0x3f, 0x14,
	0x66, 0xe0, 0x68, 0xdb, 0xcf, 0xff, 0x74, 0xd8, 0x82, 0x5b, 0x8f, 0x6e, 0x05, 0x6f, 0x06, 0xf7,
	0x5f, 0x7d, 0x2d, 0x6d, 0xfd, 0xaa, 0xb4, 0xf5, 0x1f, 0xa5, 0xad, 0x7f, 0x5a, 0xda, 0xda, 0xd5,
	0xd2, 0xd6, 0xbe, 0x2d, 0x6d, 0xed, 0xed, 0xc3, 0x35, 0xe3, 0xaa, 0xd6, 0xa7, 0xe9, 0x4c, 0xad,
	0x78, 0xbe, 0xf9, 0x63, 0x95, 0x7d, 0xcd, 0xbf, 0x0d, 0xda, 0xca, 0x8e, 0x07, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0xdd, 0x4f, 0x4c, 0xdb, 0x03, 0x00, 0x00,
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
	Recover(ctx context.Context, in *MsgRecover, opts ...grpc.CallOption) (*MsgRecoverResponse, error)
	ActivateAccount(ctx context.Context, in *MsgActivateAccount, opts ...grpc.CallOption) (*MsgActivateAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Recover(ctx context.Context, in *MsgRecover, opts ...grpc.CallOption) (*MsgRecoverResponse, error) {
	out := new(MsgRecoverResponse)
	err := c.cc.Invoke(ctx, "/aura.smartaccount.v1beta1.Msg/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ActivateAccount(ctx context.Context, in *MsgActivateAccount, opts ...grpc.CallOption) (*MsgActivateAccountResponse, error) {
	out := new(MsgActivateAccountResponse)
	err := c.cc.Invoke(ctx, "/aura.smartaccount.v1beta1.Msg/ActivateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Recover(context.Context, *MsgRecover) (*MsgRecoverResponse, error)
	ActivateAccount(context.Context, *MsgActivateAccount) (*MsgActivateAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Recover(ctx context.Context, req *MsgRecover) (*MsgRecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (*UnimplementedMsgServer) ActivateAccount(ctx context.Context, req *MsgActivateAccount) (*MsgActivateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRecover)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aura.smartaccount.v1beta1.Msg/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Recover(ctx, req.(*MsgRecover))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ActivateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgActivateAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ActivateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aura.smartaccount.v1beta1.Msg/ActivateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ActivateAccount(ctx, req.(*MsgActivateAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aura.smartaccount.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recover",
			Handler:    _Msg_Recover_Handler,
		},
		{
			MethodName: "ActivateAccount",
			Handler:    _Msg_ActivateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aura/smartaccount/v1beta1/tx.proto",
}

func (m *MsgRecover) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRecover) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRecover) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Credentials) > 0 {
		i -= len(m.Credentials)
		copy(dAtA[i:], m.Credentials)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Credentials)))
		i--
		dAtA[i] = 0x22
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRecoverResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRecoverResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRecoverResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgActivateAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgActivateAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgActivateAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InitMsg) > 0 {
		i -= len(m.InitMsg)
		copy(dAtA[i:], m.InitMsg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InitMsg)))
		i--
		dAtA[i] = 0x2a
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.CodeID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountAddress) > 0 {
		i -= len(m.AccountAddress)
		copy(dAtA[i:], m.AccountAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AccountAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgActivateAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgActivateAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgActivateAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRecover) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Credentials)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRecoverResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgActivateAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CodeID != 0 {
		n += 1 + sovTx(uint64(m.CodeID))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InitMsg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgActivateAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRecover) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRecover: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRecover: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credentials = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRecoverResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRecoverResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRecoverResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgActivateAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgActivateAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgActivateAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = append(m.Salt[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt == nil {
				m.Salt = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitMsg = append(m.InitMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.InitMsg == nil {
				m.InitMsg = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgActivateAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgActivateAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgActivateAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
