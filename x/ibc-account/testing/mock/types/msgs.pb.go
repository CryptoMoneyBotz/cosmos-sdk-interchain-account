// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/account/testing/msgs.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type MsgTryRegisterIBCAccount struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// the salt that will be transfered to counterparty chain.
	Salt []byte `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64                                        `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	Sender           github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
}

func (m *MsgTryRegisterIBCAccount) Reset()         { *m = MsgTryRegisterIBCAccount{} }
func (m *MsgTryRegisterIBCAccount) String() string { return proto.CompactTextString(m) }
func (*MsgTryRegisterIBCAccount) ProtoMessage()    {}
func (*MsgTryRegisterIBCAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e829767197ba51, []int{0}
}
func (m *MsgTryRegisterIBCAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTryRegisterIBCAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTryRegisterIBCAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTryRegisterIBCAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTryRegisterIBCAccount.Merge(m, src)
}
func (m *MsgTryRegisterIBCAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgTryRegisterIBCAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTryRegisterIBCAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTryRegisterIBCAccount proto.InternalMessageInfo

func (m *MsgTryRegisterIBCAccount) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *MsgTryRegisterIBCAccount) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *MsgTryRegisterIBCAccount) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *MsgTryRegisterIBCAccount) GetTimeoutHeight() types.Height {
	if m != nil {
		return m.TimeoutHeight
	}
	return types.Height{}
}

func (m *MsgTryRegisterIBCAccount) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgTryRegisterIBCAccount) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgTryRegisterIBCAccount)(nil), "ibc.account.testing.MsgTryRegisterIBCAccount")
}

func init() { proto.RegisterFile("ibc/account/testing/msgs.proto", fileDescriptor_48e829767197ba51) }

var fileDescriptor_48e829767197ba51 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0xae, 0x54, 0xc2, 0xc7, 0x9d, 0xc0, 0xfc, 0x33, 0x27, 0x48, 0xaa, 0x4c, 0x5d,
	0x92, 0x08, 0x18, 0x90, 0x98, 0x68, 0x6e, 0xa1, 0x03, 0x12, 0x8a, 0x3a, 0x20, 0x96, 0x53, 0xe2,
	0x5a, 0x8e, 0x45, 0x63, 0x47, 0x7e, 0xdf, 0x4a, 0xf4, 0x5b, 0xf0, 0x61, 0xf8, 0x10, 0x37, 0xde,
	0xc8, 0x14, 0xa1, 0xf6, 0x1b, 0x74, 0x64, 0x42, 0x49, 0x5c, 0x68, 0x85, 0x74, 0xd3, 0x1b, 0x3f,
	0xbf, 0xe7, 0x79, 0x14, 0x5b, 0x2f, 0xf1, 0x55, 0xc1, 0x93, 0x9c, 0x73, 0xb3, 0xd2, 0x98, 0xa0,
	0x00, 0x54, 0x5a, 0x26, 0x15, 0x48, 0x88, 0x6b, 0x6b, 0xd0, 0xd0, 0x47, 0xaa, 0xe0, 0xb1, 0xe3,
	0xb1, 0xe3, 0x17, 0x8f, 0xa5, 0x91, 0xa6, 0xe3, 0x49, 0xfb, 0xd5, 0x5b, 0x2f, 0x9e, 0xb5, 0x55,
	0x7c, 0xa9, 0x84, 0x46, 0x37, 0x7a, 0x10, 0xfe, 0x38, 0x21, 0xec, 0x23, 0xc8, 0xb9, 0x5d, 0x67,
	0x42, 0x2a, 0x40, 0x61, 0x67, 0xe9, 0xe5, 0xb4, 0xef, 0xa4, 0x6f, 0xc9, 0x29, 0x98, 0x95, 0xe5,
	0xe2, 0xaa, 0x36, 0x16, 0x99, 0x37, 0xf6, 0x26, 0xf7, 0xd2, 0xa7, 0xbb, 0x26, 0xa0, 0xeb, 0xbc,
	0x5a, 0xbe, 0x0b, 0x0f, 0x60, 0x98, 0x91, 0xfe, 0xf4, 0xc9, 0x58, 0xa4, 0xef, 0xc9, 0xb9, 0x63,
	0xbc, 0xcc, 0xb5, 0x16, 0x4b, 0x76, 0xa7, 0xcb, 0x3e, 0xdf, 0x35, 0xc1, 0x93, 0xa3, 0xac, 0xe3,
	0x61, 0x76, 0xd6, 0x0b, 0x97, 0xfd, 0x99, 0x46, 0x64, 0x08, 0xf9, 0x12, 0xd9, 0xc9, 0xd8, 0x9b,
	0xdc, 0xbf, 0x2d, 0xd7, 0xd9, 0xe8, 0x67, 0x72, 0x8e, 0xaa, 0x12, 0x66, 0x85, 0x57, 0xa5, 0x50,
	0xb2, 0x44, 0x36, 0x1c, 0x7b, 0x93, 0xd3, 0xd7, 0x34, 0x6e, 0xdf, 0xc8, 0xdd, 0xf8, 0x43, 0x47,
	0xd2, 0x97, 0xd7, 0x4d, 0x30, 0xf8, 0x57, 0x78, 0x9c, 0x0b, 0xb3, 0x33, 0x27, 0xf4, 0x6e, 0x3a,
	0x23, 0x0f, 0xf7, 0x8e, 0x76, 0x02, 0xe6, 0x55, 0xcd, 0xee, 0x8e, 0xbd, 0xc9, 0x30, 0x7d, 0xb1,
	0x6b, 0x02, 0x76, 0x5c, 0xf2, 0xd7, 0x12, 0x66, 0x0f, 0x9c, 0x36, 0xdf, 0x4b, 0x74, 0x46, 0x46,
	0x20, 0xf4, 0x42, 0x58, 0x36, 0xea, 0x6e, 0xf5, 0xea, 0x77, 0x13, 0x44, 0x52, 0x61, 0xb9, 0x2a,
	0x62, 0x6e, 0xaa, 0x84, 0x1b, 0xa8, 0x0c, 0xb8, 0x11, 0xc1, 0xe2, 0x6b, 0x82, 0xeb, 0x5a, 0x40,
	0x3c, 0xe5, 0x7c, 0xba, 0x58, 0x58, 0x01, 0x90, 0xb9, 0x82, 0x54, 0x5f, 0x6f, 0x7c, 0xef, 0x66,
	0xe3, 0x7b, 0xbf, 0x36, 0xbe, 0xf7, 0x7d, 0xeb, 0x0f, 0x6e, 0xb6, 0xfe, 0xe0, 0xe7, 0xd6, 0x1f,
	0x7c, 0x99, 0x1f, 0x16, 0x96, 0xb9, 0xd2, 0x79, 0x0d, 0xea, 0xb0, 0x34, 0x52, 0x1a, 0x85, 0xed,
	0x50, 0xb4, 0x5f, 0xae, 0x6f, 0x89, 0x2a, 0x78, 0xf4, 0xdf, 0xaa, 0x19, 0xee, 0x7e, 0xa1, 0x18,
	0x75, 0xdb, 0xf2, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x7a, 0x85, 0x28, 0x93, 0x02,
	0x00, 0x00,
}

func (m *MsgTryRegisterIBCAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTryRegisterIBCAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTryRegisterIBCAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgTryRegisterIBCAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovMsgs(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsgs(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTryRegisterIBCAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgTryRegisterIBCAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTryRegisterIBCAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = append(m.Salt[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt == nil {
				m.Salt = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
