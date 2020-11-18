// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: peggy/v1/batch.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// OutgoingTxBatch represents a batch of transactions going from Peggy to ETH
type OutgoingTxBatch struct {
	// The nonce of...?
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The individual operation of the transaction batch
	Elements []*OutgoingTransferTx `protobuf:"bytes,2,rep,name=elements,proto3" json:"elements,omitempty"`
	// The fee to be paid on ETH for the transaction
	Erc20Fee *ERC20Token `protobuf:"bytes,4,opt,name=erc20_fee,json=erc20Fee,proto3" json:"erc20_fee,omitempty"`
	// The peggy side representation of the ERC20 token
	BridgedDenominator *BridgedDenominator `protobuf:"bytes,5,opt,name=bridged_denominator,json=bridgedDenominator,proto3" json:"bridged_denominator,omitempty"`
	// The validator set on Peggy that will sign the transactions
	Valset *Valset `protobuf:"bytes,7,opt,name=valset,proto3" json:"valset,omitempty"`
	// The token contract on ETH were...?
	TokenContract string `protobuf:"bytes,8,opt,name=token_contract,json=tokenContract,proto3" json:"token_contract,omitempty"`
}

func (m *OutgoingTxBatch) Reset()         { *m = OutgoingTxBatch{} }
func (m *OutgoingTxBatch) String() string { return proto.CompactTextString(m) }
func (*OutgoingTxBatch) ProtoMessage()    {}
func (*OutgoingTxBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_398e85e0d69cec73, []int{0}
}
func (m *OutgoingTxBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTxBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTxBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTxBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTxBatch.Merge(m, src)
}
func (m *OutgoingTxBatch) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTxBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTxBatch.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTxBatch proto.InternalMessageInfo

func (m *OutgoingTxBatch) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *OutgoingTxBatch) GetElements() []*OutgoingTransferTx {
	if m != nil {
		return m.Elements
	}
	return nil
}

func (m *OutgoingTxBatch) GetErc20Fee() *ERC20Token {
	if m != nil {
		return m.Erc20Fee
	}
	return nil
}

func (m *OutgoingTxBatch) GetBridgedDenominator() *BridgedDenominator {
	if m != nil {
		return m.BridgedDenominator
	}
	return nil
}

func (m *OutgoingTxBatch) GetValset() *Valset {
	if m != nil {
		return m.Valset
	}
	return nil
}

func (m *OutgoingTxBatch) GetTokenContract() string {
	if m != nil {
		return m.TokenContract
	}
	return ""
}

// OutgoingTransferTx represents an individual send from Peggy to ETH
type OutgoingTransferTx struct {
	// The nonce on the peggy side for this individual transfer
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The peggy address of the sender of this transaction
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// The address on ETH where the transfer is bound
	DestAddress string `protobuf:"bytes,3,opt,name=dest_address,json=destAddress,proto3" json:"dest_address,omitempty"`
	// The ERC20 token amount to be sent back to ETH
	Erc20Token *ERC20Token `protobuf:"bytes,4,opt,name=erc20_token,json=erc20Token,proto3" json:"erc20_token,omitempty"`
	// The ERC20 fee paid to the bridge
	Erc20Fee *ERC20Token `protobuf:"bytes,5,opt,name=erc20_fee,json=erc20Fee,proto3" json:"erc20_fee,omitempty"`
}

func (m *OutgoingTransferTx) Reset()         { *m = OutgoingTransferTx{} }
func (m *OutgoingTransferTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTransferTx) ProtoMessage()    {}
func (*OutgoingTransferTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_398e85e0d69cec73, []int{1}
}
func (m *OutgoingTransferTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTransferTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTransferTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTransferTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTransferTx.Merge(m, src)
}
func (m *OutgoingTransferTx) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTransferTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTransferTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTransferTx proto.InternalMessageInfo

func (m *OutgoingTransferTx) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OutgoingTransferTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTransferTx) GetDestAddress() string {
	if m != nil {
		return m.DestAddress
	}
	return ""
}

func (m *OutgoingTransferTx) GetErc20Token() *ERC20Token {
	if m != nil {
		return m.Erc20Token
	}
	return nil
}

func (m *OutgoingTransferTx) GetErc20Fee() *ERC20Token {
	if m != nil {
		return m.Erc20Fee
	}
	return nil
}

func init() {
	proto.RegisterType((*OutgoingTxBatch)(nil), "peggy.v1.OutgoingTxBatch")
	proto.RegisterType((*OutgoingTransferTx)(nil), "peggy.v1.OutgoingTransferTx")
}

func init() { proto.RegisterFile("peggy/v1/batch.proto", fileDescriptor_398e85e0d69cec73) }

var fileDescriptor_398e85e0d69cec73 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xd9, 0x6d, 0xed, 0x4e, 0x75, 0x95, 0xd9, 0x22, 0x43, 0x91, 0x50, 0x17, 0x84,
	0x5c, 0x4c, 0xda, 0x8a, 0xe0, 0xd5, 0xae, 0x7a, 0x10, 0x44, 0x08, 0xc5, 0x83, 0x97, 0x32, 0xc9,
	0xbc, 0x4d, 0x83, 0xe9, 0x4c, 0x99, 0x79, 0x2d, 0xdd, 0x6f, 0xe1, 0x87, 0xf1, 0x43, 0xe8, 0x6d,
	0x8f, 0x1e, 0xa5, 0xfd, 0x22, 0xd2, 0x49, 0x4c, 0x94, 0x82, 0xec, 0x2d, 0xef, 0xf7, 0xfe, 0xf3,
	0x7f, 0x2f, 0x7f, 0x1e, 0xed, 0xaf, 0x20, 0xcb, 0x6e, 0xa2, 0xcd, 0x38, 0x4a, 0x04, 0xa6, 0x8b,
	0x70, 0x65, 0x34, 0x6a, 0xd6, 0x75, 0x34, 0xdc, 0x8c, 0x07, 0x83, 0xba, 0x2f, 0x10, 0xc1, 0xa2,
	0xc0, 0x5c, 0xab, 0x52, 0x35, 0xb8, 0xa8, 0x7b, 0x2b, 0xad, 0x8b, 0x0a, 0x36, 0x86, 0x78, 0xb3,
	0x02, 0x5b, 0xd2, 0xcb, 0x6f, 0x1e, 0x7d, 0xf8, 0x71, 0x8d, 0x99, 0xce, 0x55, 0x36, 0xdb, 0x4e,
	0x0f, 0xa3, 0x58, 0x9f, 0xb6, 0x95, 0x56, 0x29, 0x70, 0x32, 0x24, 0xc1, 0x69, 0x5c, 0x16, 0xec,
	0x15, 0xed, 0x42, 0x01, 0x4b, 0x50, 0x68, 0xb9, 0x37, 0x3c, 0x09, 0x7a, 0x93, 0x27, 0xe1, 0x9f,
	0x6d, 0xc2, 0xda, 0xc2, 0x08, 0x65, 0xaf, 0xc1, 0xcc, 0xb6, 0x71, 0xad, 0x66, 0x63, 0x7a, 0x06,
	0x26, 0x9d, 0x8c, 0xe6, 0xd7, 0x00, 0xfc, 0x74, 0x48, 0x82, 0xde, 0xa4, 0xdf, 0x3c, 0x7d, 0x1b,
	0x5f, 0x4d, 0x46, 0x33, 0xfd, 0x05, 0x54, 0xdc, 0x75, 0xb2, 0x77, 0x00, 0xec, 0x03, 0xbd, 0x48,
	0x4c, 0x2e, 0x33, 0x90, 0x73, 0x09, 0x4a, 0x2f, 0x73, 0x25, 0x50, 0x1b, 0xde, 0x76, 0x8f, 0xff,
	0x9a, 0x3b, 0x2d, 0x45, 0x6f, 0x1a, 0x4d, 0xcc, 0x92, 0x23, 0xc6, 0x02, 0xda, 0xd9, 0x88, 0xc2,
	0x02, 0xf2, 0x7b, 0xce, 0xe1, 0x51, 0xe3, 0xf0, 0xc9, 0xf1, 0xb8, 0xea, 0xb3, 0x67, 0xf4, 0x1c,
	0x0f, 0xbb, 0xcc, 0x53, 0xad, 0xd0, 0x88, 0x14, 0x79, 0x77, 0x48, 0x82, 0xb3, 0xf8, 0x81, 0xa3,
	0x57, 0x15, 0xbc, 0xfc, 0x41, 0x28, 0x3b, 0xfe, 0x67, 0x76, 0x4e, 0xbd, 0x5c, 0x56, 0xb1, 0x79,
	0xb9, 0x64, 0x8f, 0x69, 0xc7, 0x82, 0x92, 0x60, 0xb8, 0xe7, 0x5c, 0xaa, 0x8a, 0x3d, 0xa5, 0xf7,
	0x25, 0x58, 0x9c, 0x0b, 0x29, 0x0d, 0x58, 0xcb, 0x4f, 0x5c, 0xb7, 0x77, 0x60, 0xaf, 0x4b, 0xc4,
	0x5e, 0xd2, 0x5e, 0x19, 0x9a, 0x1b, 0xfc, 0xdf, 0xd8, 0xa8, 0x13, 0xba, 0xef, 0x7f, 0xb3, 0x6e,
	0xdf, 0x25, 0xeb, 0xe9, 0xfb, 0xef, 0x3b, 0x9f, 0xdc, 0xee, 0x7c, 0xf2, 0x6b, 0xe7, 0x93, 0xaf,
	0x7b, 0xbf, 0x75, 0xbb, 0xf7, 0x5b, 0x3f, 0xf7, 0x7e, 0xeb, 0xf3, 0x28, 0xcb, 0x71, 0xb1, 0x4e,
	0xc2, 0x54, 0x2f, 0x23, 0x51, 0xe0, 0x02, 0xc4, 0x73, 0x05, 0x18, 0x95, 0x87, 0xb4, 0xd4, 0x72,
	0x5d, 0x40, 0xb4, 0xad, 0x4a, 0x77, 0x54, 0x49, 0xc7, 0x5d, 0xd5, 0x8b, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x51, 0xc0, 0xb7, 0x27, 0xbe, 0x02, 0x00, 0x00,
}

func (m *OutgoingTxBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTxBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTxBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenContract) > 0 {
		i -= len(m.TokenContract)
		copy(dAtA[i:], m.TokenContract)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.TokenContract)))
		i--
		dAtA[i] = 0x42
	}
	if m.Valset != nil {
		{
			size, err := m.Valset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.BridgedDenominator != nil {
		{
			size, err := m.BridgedDenominator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Erc20Fee != nil {
		{
			size, err := m.Erc20Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Elements) > 0 {
		for iNdEx := len(m.Elements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Elements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Nonce != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutgoingTransferTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTransferTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTransferTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Erc20Fee != nil {
		{
			size, err := m.Erc20Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Erc20Token != nil {
		{
			size, err := m.Erc20Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.DestAddress) > 0 {
		i -= len(m.DestAddress)
		copy(dAtA[i:], m.DestAddress)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.DestAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovBatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutgoingTxBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovBatch(uint64(m.Nonce))
	}
	if len(m.Elements) > 0 {
		for _, e := range m.Elements {
			l = e.Size()
			n += 1 + l + sovBatch(uint64(l))
		}
	}
	if m.Erc20Fee != nil {
		l = m.Erc20Fee.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.BridgedDenominator != nil {
		l = m.BridgedDenominator.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Valset != nil {
		l = m.Valset.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.TokenContract)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	return n
}

func (m *OutgoingTransferTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBatch(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.DestAddress)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Erc20Token != nil {
		l = m.Erc20Token.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.Erc20Fee != nil {
		l = m.Erc20Fee.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	return n
}

func sovBatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBatch(x uint64) (n int) {
	return sovBatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutgoingTxBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
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
			return fmt.Errorf("proto: OutgoingTxBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTxBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elements = append(m.Elements, &OutgoingTransferTx{})
			if err := m.Elements[len(m.Elements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Erc20Fee == nil {
				m.Erc20Fee = &ERC20Token{}
			}
			if err := m.Erc20Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgedDenominator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BridgedDenominator == nil {
				m.BridgedDenominator = &BridgedDenominator{}
			}
			if err := m.BridgedDenominator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Valset == nil {
				m.Valset = &Valset{}
			}
			if err := m.Valset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBatch
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBatch
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
func (m *OutgoingTransferTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
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
			return fmt.Errorf("proto: OutgoingTransferTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTransferTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Erc20Token == nil {
				m.Erc20Token = &ERC20Token{}
			}
			if err := m.Erc20Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Erc20Fee == nil {
				m.Erc20Fee = &ERC20Token{}
			}
			if err := m.Erc20Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBatch
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBatch
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
func skipBatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
				return 0, ErrInvalidLengthBatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBatch = fmt.Errorf("proto: unexpected end of group")
)