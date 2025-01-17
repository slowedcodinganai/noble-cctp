// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/remote_token_messenger.proto

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
// @param domain_id
// @param address
type RemoteTokenMessenger struct {
	DomainId uint32 `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Address  []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *RemoteTokenMessenger) Reset()         { *m = RemoteTokenMessenger{} }
func (m *RemoteTokenMessenger) String() string { return proto.CompactTextString(m) }
func (*RemoteTokenMessenger) ProtoMessage()    {}
func (*RemoteTokenMessenger) Descriptor() ([]byte, []int) {
	return fileDescriptor_1046e968a0e9ecf4, []int{0}
}
func (m *RemoteTokenMessenger) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteTokenMessenger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoteTokenMessenger.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoteTokenMessenger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteTokenMessenger.Merge(m, src)
}
func (m *RemoteTokenMessenger) XXX_Size() int {
	return m.Size()
}
func (m *RemoteTokenMessenger) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteTokenMessenger.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteTokenMessenger proto.InternalMessageInfo

func (m *RemoteTokenMessenger) GetDomainId() uint32 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *RemoteTokenMessenger) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteTokenMessenger)(nil), "circle.cctp.v1.RemoteTokenMessenger")
}

func init() {
	proto.RegisterFile("circle/cctp/v1/remote_token_messenger.proto", fileDescriptor_1046e968a0e9ecf4)
}

var fileDescriptor_1046e968a0e9ecf4 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xce, 0x2c, 0x4a,
	0xce, 0x49, 0xd5, 0x4f, 0x4e, 0x2e, 0x29, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0x4a, 0xcd, 0xcd, 0x2f,
	0x49, 0x8d, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0x8b, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0xcd, 0x4b, 0x4f,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x28, 0xd6, 0x03, 0x29, 0xd6, 0x2b,
	0x33, 0x54, 0xf2, 0xe5, 0x12, 0x09, 0x02, 0xab, 0x0f, 0x01, 0x29, 0xf7, 0x85, 0xa9, 0x16, 0x92,
	0xe6, 0xe2, 0x4c, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x8b, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0d, 0xe2, 0x80, 0x08, 0x78, 0xa6, 0x08, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5,
	0x16, 0x17, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x4e, 0x6e, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x93, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x0f, 0x71, 0x43, 0x5a, 0x66, 0x9e, 0x7e, 0x5e, 0x7e, 0x52, 0x4e, 0xaa, 0x2e,
	0xd8, 0xe5, 0x15, 0x10, 0x0f, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x5d, 0x6b, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xab, 0xf9, 0x18, 0x48, 0xdc, 0x00, 0x00, 0x00,
}

func (m *RemoteTokenMessenger) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteTokenMessenger) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoteTokenMessenger) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRemoteTokenMessenger(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.DomainId != 0 {
		i = encodeVarintRemoteTokenMessenger(dAtA, i, uint64(m.DomainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRemoteTokenMessenger(dAtA []byte, offset int, v uint64) int {
	offset -= sovRemoteTokenMessenger(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RemoteTokenMessenger) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DomainId != 0 {
		n += 1 + sovRemoteTokenMessenger(uint64(m.DomainId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRemoteTokenMessenger(uint64(l))
	}
	return n
}

func sovRemoteTokenMessenger(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRemoteTokenMessenger(x uint64) (n int) {
	return sovRemoteTokenMessenger(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteTokenMessenger) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoteTokenMessenger
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
			return fmt.Errorf("proto: RemoteTokenMessenger: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteTokenMessenger: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainId", wireType)
			}
			m.DomainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTokenMessenger
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DomainId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteTokenMessenger
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
				return ErrInvalidLengthRemoteTokenMessenger
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRemoteTokenMessenger
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
			skippy, err := skipRemoteTokenMessenger(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRemoteTokenMessenger
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
func skipRemoteTokenMessenger(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRemoteTokenMessenger
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
					return 0, ErrIntOverflowRemoteTokenMessenger
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
					return 0, ErrIntOverflowRemoteTokenMessenger
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
				return 0, ErrInvalidLengthRemoteTokenMessenger
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRemoteTokenMessenger
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRemoteTokenMessenger
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRemoteTokenMessenger        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRemoteTokenMessenger          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRemoteTokenMessenger = fmt.Errorf("proto: unexpected end of group")
)
