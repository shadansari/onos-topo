// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/topo/e2interface.proto

// Package topo.interfaces defines api for managing sd-ran interfaces.

package topo

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

// E2Interface contains information about a E2 Interface
type E2Interface struct {
}

func (m *E2Interface) Reset()         { *m = E2Interface{} }
func (m *E2Interface) String() string { return proto.CompactTextString(m) }
func (*E2Interface) ProtoMessage()    {}
func (*E2Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_74435e3cdd38c79e, []int{0}
}
func (m *E2Interface) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *E2Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_E2Interface.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *E2Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E2Interface.Merge(m, src)
}
func (m *E2Interface) XXX_Size() int {
	return m.Size()
}
func (m *E2Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_E2Interface.DiscardUnknown(m)
}

var xxx_messageInfo_E2Interface proto.InternalMessageInfo

func init() {
	proto.RegisterType((*E2Interface)(nil), "topo.E2Interface")
}

func init() { proto.RegisterFile("api/topo/e2interface.proto", fileDescriptor_74435e3cdd38c79e) }

var fileDescriptor_74435e3cdd38c79e = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2c, 0xc8, 0xd4,
	0x2f, 0xc9, 0x2f, 0xc8, 0xd7, 0x4f, 0x35, 0xca, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89, 0x2b, 0xf1, 0x72, 0x71, 0xbb, 0x1a,
	0x79, 0xc2, 0xa4, 0x9c, 0x24, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89,
	0x0d, 0xac, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x04, 0xa2, 0x2c, 0x53, 0x00, 0x00,
	0x00,
}

func (m *E2Interface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *E2Interface) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *E2Interface) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintE2Interface(dAtA []byte, offset int, v uint64) int {
	offset -= sovE2Interface(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *E2Interface) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovE2Interface(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozE2Interface(x uint64) (n int) {
	return sovE2Interface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *E2Interface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowE2Interface
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
			return fmt.Errorf("proto: E2Interface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: E2Interface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipE2Interface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthE2Interface
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthE2Interface
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
func skipE2Interface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowE2Interface
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
					return 0, ErrIntOverflowE2Interface
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
					return 0, ErrIntOverflowE2Interface
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
				return 0, ErrInvalidLengthE2Interface
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupE2Interface
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthE2Interface
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthE2Interface        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowE2Interface          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupE2Interface = fmt.Errorf("proto: unexpected end of group")
)
