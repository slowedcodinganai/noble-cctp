// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/signature_threshold.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// *
// SignatureThreshold is the minimum amount of signatures required to attest to
// a message (the 'm' in a m/n multisig)
// @param amount the number of signatures required
type SignatureThreshold struct {
	Amount uint32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *SignatureThreshold) Reset()         { *m = SignatureThreshold{} }
func (m *SignatureThreshold) String() string { return proto.CompactTextString(m) }
func (*SignatureThreshold) ProtoMessage()    {}
func (*SignatureThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b3271fb2b49b51, []int{0}
}
func (m *SignatureThreshold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureThreshold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureThreshold.Merge(m, src)
}
func (m *SignatureThreshold) XXX_Size() int {
	return m.Size()
}
func (m *SignatureThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureThreshold proto.InternalMessageInfo

func (m *SignatureThreshold) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*SignatureThreshold)(nil), "circle.cctp.v1.SignatureThreshold")
}

func init() {
	proto.RegisterFile("circle/cctp/v1/signature_threshold.proto", fileDescriptor_26b3271fb2b49b51)
}

var fileDescriptor_26b3271fb2b49b51 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xce, 0x2c, 0x4a,
	0xce, 0x49, 0xd5, 0x4f, 0x4e, 0x2e, 0x29, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0xce, 0x4c, 0xcf, 0x4b,
	0x2c, 0x29, 0x2d, 0x4a, 0x8d, 0x2f, 0xc9, 0x28, 0x4a, 0x2d, 0xce, 0xc8, 0xcf, 0x49, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xa8, 0xd4, 0x03, 0xa9, 0xd4, 0x2b, 0x33, 0x54, 0xd2,
	0xe1, 0x12, 0x0a, 0x86, 0x29, 0x0e, 0x81, 0xa9, 0x15, 0x12, 0xe3, 0x62, 0x4b, 0xcc, 0xcd, 0x2f,
	0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0xf2, 0x9c, 0xdc, 0x4e, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x27, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x1f, 0x62, 0x45, 0x5a, 0x66, 0x9e, 0x7e, 0x5e, 0x7e, 0x52, 0x4e, 0xaa, 0x2e,
	0xd8, 0x55, 0x15, 0x10, 0xc7, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x1d, 0x63, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x25, 0xe9, 0xb1, 0xd4, 0xb8, 0x00, 0x00, 0x00,
}

func (m *SignatureThreshold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureThreshold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureThreshold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintSignatureThreshold(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSignatureThreshold(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignatureThreshold(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignatureThreshold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Amount != 0 {
		n += 1 + sovSignatureThreshold(uint64(m.Amount))
	}
	return n
}

func sovSignatureThreshold(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignatureThreshold(x uint64) (n int) {
	return sovSignatureThreshold(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignatureThreshold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignatureThreshold
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
			return fmt.Errorf("proto: SignatureThreshold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureThreshold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignatureThreshold
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignatureThreshold(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignatureThreshold
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
func skipSignatureThreshold(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignatureThreshold
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
					return 0, ErrIntOverflowSignatureThreshold
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
					return 0, ErrIntOverflowSignatureThreshold
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
				return 0, ErrInvalidLengthSignatureThreshold
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignatureThreshold
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignatureThreshold
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignatureThreshold        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignatureThreshold          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignatureThreshold = fmt.Errorf("proto: unexpected end of group")
)
