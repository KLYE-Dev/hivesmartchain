// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acm.proto

package acm

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	github_com_klye-dev_hsc-main_binary "github.com/klye-dev/hsc-main/binary"
	crypto "github.com/klye-dev/hsc-main/crypto"
	github_com_klye-dev_hsc-main_crypto "github.com/klye-dev/hsc-main/crypto"
	permission "github.com/klye-dev/hsc-main/permission"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Account struct {
	Address   github_com_klye-dev_hsc-main_crypto.Address `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/klye-dev/hsc-main/crypto.Address" json:"Address"`
	PublicKey *crypto.PublicKey                            `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	// Sequence counts the number of transactions that have been accepted from this account
	Sequence uint64 `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	// The account's current native token balance
	Balance uint64 `protobuf:"varint,4,opt,name=Balance,proto3" json:"Balance,omitempty"`
	// We expect exactly one of EVMCode, WASMCode, and NativeName to be non-empty
	// EVM bytecode
	EVMCode     Bytecode                      `protobuf:"bytes,5,opt,name=EVMCode,proto3,customtype=Bytecode" json:"EVMCode"`
	Permissions permission.AccountPermissions `protobuf:"bytes,6,opt,name=Permissions,proto3" json:"Permissions"`
	// WASM bytecode
	WASMCode Bytecode `protobuf:"bytes,7,opt,name=WASMCode,proto3,customtype=Bytecode" json:",omitempty"`
	// Fully qualified (`<contract name>.<function name>`) name of native contract this for which this account object
	// is a sentinel value. Which is to say this account object is a pointer to compiled code and does not contain
	// the contract logic in its entirety
	NativeName string `protobuf:"bytes,11,opt,name=NativeName,proto3" json:",omitempty"`
	// The sha3 hash of the code associated with the account
	CodeHash github_com_klye-dev_hsc-main_binary.HexBytes `protobuf:"bytes,8,opt,name=CodeHash,proto3,customtype=github.com/klye-dev/hsc-main/binary.HexBytes" json:"-"`
	// Pointer to the Metadata associated with this account
	ContractMeta []*ContractMeta `protobuf:"bytes,9,rep,name=ContractMeta,proto3" json:"ContractMeta,omitempty"`
	// The metadata is stored in the deployed account. When the deployed account creates new account
	// (from Solidity/EVM), they point to the original deployed account where the metadata is stored.
	// This original account is called the forebear.
	Forebear             *github_com_klye-dev_hsc-main_crypto.Address `protobuf:"bytes,10,opt,name=Forebear,proto3,customtype=github.com/klye-dev/hsc-main/crypto.Address" json:"Forebear,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *Account) Reset()      { *m = Account{} }
func (*Account) ProtoMessage() {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ed775bc0a6adf6, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetPublicKey() *crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Account) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetPermissions() permission.AccountPermissions {
	if m != nil {
		return m.Permissions
	}
	return permission.AccountPermissions{}
}

func (m *Account) GetNativeName() string {
	if m != nil {
		return m.NativeName
	}
	return ""
}

func (m *Account) GetContractMeta() []*ContractMeta {
	if m != nil {
		return m.ContractMeta
	}
	return nil
}

func (*Account) XXX_MessageName() string {
	return "acm.Account"
}

type ContractMeta struct {
	CodeHash     github_com_klye-dev_hsc-main_binary.HexBytes `protobuf:"bytes,1,opt,name=CodeHash,proto3,customtype=github.com/klye-dev/hsc-main/binary.HexBytes" json:"CodeHash"`
	MetadataHash github_com_klye-dev_hsc-main_binary.HexBytes `protobuf:"bytes,2,opt,name=MetadataHash,proto3,customtype=github.com/klye-dev/hsc-main/binary.HexBytes" json:"MetadataHash"`
	// In the dump format we would like the ABI rather than its hash
	Metadata             string   `protobuf:"bytes,3,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractMeta) Reset()         { *m = ContractMeta{} }
func (m *ContractMeta) String() string { return proto.CompactTextString(m) }
func (*ContractMeta) ProtoMessage()    {}
func (*ContractMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ed775bc0a6adf6, []int{1}
}
func (m *ContractMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ContractMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractMeta.Merge(m, src)
}
func (m *ContractMeta) XXX_Size() int {
	return m.Size()
}
func (m *ContractMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ContractMeta proto.InternalMessageInfo

func (m *ContractMeta) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (*ContractMeta) XXX_MessageName() string {
	return "acm.ContractMeta"
}
func init() {
	proto.RegisterType((*Account)(nil), "acm.Account")
	golang_proto.RegisterType((*Account)(nil), "acm.Account")
	proto.RegisterType((*ContractMeta)(nil), "acm.ContractMeta")
	golang_proto.RegisterType((*ContractMeta)(nil), "acm.ContractMeta")
}

func init() { proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }
func init() { golang_proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }

var fileDescriptor_49ed775bc0a6adf6 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xee, 0x35, 0x69, 0xe3, 0x5c, 0x22, 0x54, 0x4e, 0x0c, 0xa7, 0x0c, 0x8e, 0xe9, 0x14, 0xa1,
	0xd6, 0x46, 0x40, 0x97, 0x30, 0xc5, 0x15, 0x55, 0x25, 0x68, 0x54, 0x5c, 0xa9, 0x08, 0xb6, 0xf3,
	0xf9, 0x29, 0xb1, 0x14, 0xfb, 0xcc, 0xf9, 0x0c, 0x78, 0xe6, 0x4f, 0x30, 0xf2, 0x53, 0x18, 0x33,
	0x32, 0x56, 0x0c, 0x11, 0x4a, 0xb7, 0xfe, 0x0a, 0xe4, 0x8b, 0x63, 0x9c, 0x0e, 0x95, 0xa0, 0x5b,
	0x5e, 0xbe, 0xef, 0x7d, 0xdf, 0xf7, 0xde, 0x3d, 0xe3, 0x36, 0xe3, 0x91, 0x9d, 0x48, 0xa1, 0x04,
	0x69, 0x30, 0x1e, 0xf5, 0x1e, 0x4d, 0xc4, 0x44, 0xe8, 0xda, 0x29, 0x7e, 0xad, 0xa0, 0xde, 0x5e,
	0x02, 0x32, 0x0a, 0xd3, 0x34, 0x14, 0x71, 0xf9, 0x4f, 0x97, 0xcb, 0x3c, 0x51, 0x25, 0xbe, 0xff,
	0x75, 0x07, 0xb7, 0x46, 0x9c, 0x8b, 0x2c, 0x56, 0x64, 0x8c, 0x5b, 0xa3, 0x20, 0x90, 0x90, 0xa6,
	0x14, 0x59, 0x68, 0xd0, 0x75, 0x5f, 0xcc, 0x17, 0xfd, 0xad, 0x5f, 0x8b, 0xfe, 0xc1, 0x24, 0x54,
	0xd3, 0xcc, 0xb7, 0xb9, 0x88, 0x9c, 0x69, 0x9e, 0x80, 0x9c, 0x41, 0x30, 0x01, 0xe9, 0xf8, 0x99,
	0x94, 0xe2, 0xb3, 0x53, 0x0a, 0x96, 0xbd, 0xde, 0x5a, 0x84, 0x38, 0xb8, 0x7d, 0x9e, 0xf9, 0xb3,
	0x90, 0xbf, 0x86, 0x9c, 0x6e, 0x5b, 0x68, 0xd0, 0x79, 0xf6, 0xd0, 0x2e, 0xc9, 0x15, 0xe0, 0xfd,
	0xe5, 0x90, 0x1e, 0x36, 0x2e, 0xe0, 0x63, 0x06, 0x31, 0x07, 0xda, 0xb0, 0xd0, 0xa0, 0xe9, 0x55,
	0x35, 0xa1, 0xb8, 0xe5, 0xb2, 0x19, 0x2b, 0xa0, 0xa6, 0x86, 0xd6, 0x25, 0x79, 0x82, 0x5b, 0xaf,
	0x2e, 0xcf, 0x8e, 0x45, 0x00, 0x74, 0x47, 0xc7, 0xde, 0x2b, 0x63, 0x1b, 0x6e, 0xae, 0x80, 0x8b,
	0x00, 0xbc, 0x35, 0x81, 0x9c, 0xe0, 0xce, 0x79, 0xb5, 0x90, 0x94, 0xee, 0xea, 0x50, 0xa6, 0x5d,
	0x5b, 0x52, 0xb9, 0x8c, 0x1a, 0xcb, 0x6d, 0x16, 0x7a, 0x5e, 0xbd, 0x91, 0x0c, 0xb1, 0xf1, 0x6e,
	0x74, 0xb1, 0x32, 0x6d, 0x69, 0x53, 0xf3, 0xb6, 0xe9, 0xcd, 0xa2, 0x8f, 0x0f, 0x44, 0x14, 0x2a,
	0x88, 0x12, 0x95, 0x7b, 0x15, 0x9f, 0xd8, 0x18, 0x8f, 0x99, 0x0a, 0x3f, 0xc1, 0x98, 0x45, 0x40,
	0x3b, 0x16, 0x1a, 0xb4, 0xdd, 0x07, 0xb7, 0xd8, 0x35, 0x06, 0xb9, 0xc4, 0x46, 0xd1, 0x77, 0xca,
	0xd2, 0x29, 0x35, 0xb4, 0xd7, 0xb0, 0xf4, 0x3a, 0xbc, 0xfb, 0x5d, 0xfc, 0x30, 0x66, 0x32, 0xb7,
	0x4f, 0xe1, 0x4b, 0x91, 0x29, 0xbd, 0x59, 0xf4, 0xd1, 0xa1, 0x57, 0x69, 0x91, 0x23, 0xdc, 0x3d,
	0x16, 0xb1, 0x92, 0x8c, 0xab, 0x33, 0x50, 0x8c, 0xb6, 0xad, 0x86, 0x7e, 0xa1, 0xe2, 0xae, 0xea,
	0x80, 0xb7, 0x41, 0x23, 0x6f, 0xb0, 0x71, 0x22, 0x24, 0xf8, 0xc0, 0x24, 0xc5, 0x3a, 0xce, 0xd3,
	0x7f, 0x3e, 0x91, 0x4a, 0x61, 0xd8, 0xfc, 0xf6, 0xbd, 0xbf, 0xb5, 0x7f, 0x85, 0x36, 0xb3, 0x90,
	0xb7, 0xb5, 0x99, 0x57, 0xb7, 0x78, 0xf4, 0x5f, 0x33, 0xd7, 0xc6, 0x7d, 0x8f, 0xbb, 0x85, 0x74,
	0xc0, 0x14, 0xd3, 0xb2, 0xdb, 0xf7, 0x91, 0xdd, 0x90, 0x2a, 0xee, 0x76, 0x5d, 0xeb, 0xbb, 0x6d,
	0x7b, 0x55, 0xed, 0xbe, 0x9c, 0x2f, 0x4d, 0xf4, 0x73, 0x69, 0xa2, 0xab, 0xa5, 0x89, 0x7e, 0x2f,
	0x4d, 0xf4, 0xe3, 0xda, 0x44, 0xf3, 0x6b, 0x13, 0x7d, 0x78, 0x7c, 0xb7, 0x25, 0xe3, 0x91, 0xbf,
	0xab, 0x3f, 0xd2, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x37, 0xf6, 0x93, 0x98, 0xec, 0x03,
	0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NativeName) > 0 {
		i -= len(m.NativeName)
		copy(dAtA[i:], m.NativeName)
		i = encodeVarintAcm(dAtA, i, uint64(len(m.NativeName)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Forebear != nil {
		{
			size := m.Forebear.Size()
			i -= size
			if _, err := m.Forebear.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintAcm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.ContractMeta) > 0 {
		for iNdEx := len(m.ContractMeta) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractMeta[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAcm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	{
		size := m.CodeHash.Size()
		i -= size
		if _, err := m.CodeHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.WASMCode.Size()
		i -= size
		if _, err := m.WASMCode.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Permissions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.EVMCode.Size()
		i -= size
		if _, err := m.EVMCode.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Balance != 0 {
		i = encodeVarintAcm(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x20
	}
	if m.Sequence != 0 {
		i = encodeVarintAcm(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAcm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Address.Size()
		i -= size
		if _, err := m.Address.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContractMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintAcm(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.MetadataHash.Size()
		i -= size
		if _, err := m.MetadataHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.CodeHash.Size()
		i -= size
		if _, err := m.CodeHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAcm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAcm(dAtA []byte, offset int, v uint64) int {
	offset -= sovAcm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Address.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovAcm(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovAcm(uint64(m.Sequence))
	}
	if m.Balance != 0 {
		n += 1 + sovAcm(uint64(m.Balance))
	}
	l = m.EVMCode.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.Permissions.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.WASMCode.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.CodeHash.Size()
	n += 1 + l + sovAcm(uint64(l))
	if len(m.ContractMeta) > 0 {
		for _, e := range m.ContractMeta {
			l = e.Size()
			n += 1 + l + sovAcm(uint64(l))
		}
	}
	if m.Forebear != nil {
		l = m.Forebear.Size()
		n += 1 + l + sovAcm(uint64(l))
	}
	l = len(m.NativeName)
	if l > 0 {
		n += 1 + l + sovAcm(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContractMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CodeHash.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.MetadataHash.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovAcm(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAcm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAcm(x uint64) (n int) {
	return sovAcm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcm
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &crypto.PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EVMCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EVMCode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WASMCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WASMCode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CodeHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractMeta = append(m.ContractMeta, &ContractMeta{})
			if err := m.ContractMeta[len(m.ContractMeta)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Forebear", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_klye-dev_hsc-main_crypto.Address
			m.Forebear = &v
			if err := m.Forebear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAcm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContractMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcm
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
			return fmt.Errorf("proto: ContractMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CodeHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MetadataHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAcm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAcm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAcm
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
					return 0, ErrIntOverflowAcm
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
					return 0, ErrIntOverflowAcm
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
				return 0, ErrInvalidLengthAcm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAcm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAcm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAcm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAcm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAcm = fmt.Errorf("proto: unexpected end of group")
)
