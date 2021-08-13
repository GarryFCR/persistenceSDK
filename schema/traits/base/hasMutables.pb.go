// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/traits/base/hasMutables.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
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

type HasMutables struct {
	Properties github_com_persistenceOne_persistenceSDK_schema_types.Properties `protobuf:"bytes,1,opt,name=properties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"properties"`
}

func (m *HasMutables) Reset()         { *m = HasMutables{} }
func (m *HasMutables) String() string { return proto.CompactTextString(m) }
func (*HasMutables) ProtoMessage()    {}
func (*HasMutables) Descriptor() ([]byte, []int) {
	return fileDescriptor_d231908cd16c0dd7, []int{0}
}
func (m *HasMutables) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HasMutables) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HasMutables.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HasMutables) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasMutables.Merge(m, src)
}
func (m *HasMutables) XXX_Size() int {
	return m.Size()
}
func (m *HasMutables) XXX_DiscardUnknown() {
	xxx_messageInfo_HasMutables.DiscardUnknown(m)
}

var xxx_messageInfo_HasMutables proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HasMutables)(nil), "schema.traits.base.HasMutables")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/traits/base/hasMutables.proto", fileDescriptor_d231908cd16c0dd7)
}

var fileDescriptor_d231908cd16c0dd7 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x29, 0x4a, 0xcc, 0x2c, 0x29, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf,
	0x48, 0x2c, 0xf6, 0x2d, 0x2d, 0x49, 0x4c, 0xca, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x82, 0xa8, 0xd2, 0x83, 0xa8, 0xd2, 0x03, 0xa9, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x4a, 0xe5, 0x5c, 0xdc, 0x1e, 0x08, 0xed, 0x42, 0x19,
	0x5c, 0x5c, 0x05, 0x45, 0xf9, 0x05, 0xa9, 0x45, 0x25, 0x99, 0xa9, 0xc5, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x4e, 0x1e, 0x27, 0xee, 0xc9, 0x33, 0xdc, 0xba, 0x27, 0xef, 0x90, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0xe4, 0x2a, 0xff, 0xbc, 0x54, 0x64, 0x6e, 0xb0,
	0x8b, 0x37, 0xdc, 0x8d, 0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x01, 0x70, 0xf3, 0x82, 0x90, 0xcc, 0x76,
	0x0a, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x6b, 0x92, 0xed, 0x41,
	0x84, 0x45, 0x12, 0x1b, 0xd8, 0x5b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xfd, 0x73,
	0xb7, 0x38, 0x01, 0x00, 0x00,
}

func (m *HasMutables) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasMutables) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HasMutables) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Properties.Size()
		i -= size
		if _, err := m.Properties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintHasMutables(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintHasMutables(dAtA []byte, offset int, v uint64) int {
	offset -= sovHasMutables(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HasMutables) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Properties.Size()
	n += 1 + l + sovHasMutables(uint64(l))
	return n
}

func sovHasMutables(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHasMutables(x uint64) (n int) {
	return sovHasMutables(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HasMutables) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHasMutables
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
			return fmt.Errorf("proto: HasMutables: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasMutables: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHasMutables
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
				return ErrInvalidLengthHasMutables
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHasMutables
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Properties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHasMutables(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHasMutables
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
func skipHasMutables(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHasMutables
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
					return 0, ErrIntOverflowHasMutables
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
					return 0, ErrIntOverflowHasMutables
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
				return 0, ErrInvalidLengthHasMutables
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHasMutables
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHasMutables
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHasMutables        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHasMutables          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHasMutables = fmt.Errorf("proto: unexpected end of group")
)
