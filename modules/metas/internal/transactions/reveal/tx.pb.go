// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/metas/transactions/reveal/v1beta1/tx.proto

package reveal

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type message struct {
	From     github_com_cosmos_cosmos_sdk_types.AccAddress                  `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from" valid:"required~required field from missing"`
	MetaFact github_com_persistenceOne_persistenceSDK_schema_types.MetaFact `protobuf:"bytes,2,opt,name=metaFact,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaFact" json:"metaFact" valid:"required~required field metaFact missing"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_68111efdafda7831, []int{0}
}
func (m *message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *message) XXX_Size() int {
	return m.Size()
}
func (m *message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*message)(nil), "persistence.assets.metas.transactions.reveal.v1beta1.message")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/metas/transactions/reveal/v1beta1/tx.proto", fileDescriptor_68111efdafda7831)
}

var fileDescriptor_68111efdafda7831 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x13, 0x11, 0xad, 0x19, 0x3b, 0x15, 0x87, 0x8b, 0x38, 0x88, 0x08, 0xbd, 0xa3, 0xa8,
	0x8b, 0x83, 0xd0, 0x5a, 0x0a, 0x22, 0x22, 0xd4, 0xcd, 0x45, 0x2e, 0x97, 0x67, 0x1a, 0xec, 0xdd,
	0x85, 0x7b, 0xaf, 0xc5, 0x7e, 0x03, 0x47, 0x3f, 0x42, 0x3f, 0x8b, 0x53, 0xc7, 0x8e, 0xe2, 0x50,
	0xa4, 0x59, 0xfc, 0x18, 0x62, 0x12, 0x35, 0x83, 0xd3, 0xdd, 0x1b, 0x7e, 0xbf, 0xf7, 0x7f, 0xff,
	0xe0, 0x22, 0x73, 0x96, 0xac, 0xc8, 0xc0, 0x61, 0x8a, 0x04, 0x46, 0x81, 0x90, 0x88, 0x40, 0x28,
	0x34, 0x90, 0x44, 0x41, 0x4e, 0x1a, 0x94, 0x8a, 0x52, 0x6b, 0x50, 0x38, 0x98, 0x82, 0x1c, 0x8b,
	0x69, 0x27, 0x02, 0x92, 0x1d, 0x41, 0x4f, 0xbc, 0xa0, 0x9b, 0x27, 0x35, 0x9c, 0x97, 0x38, 0x2f,
	0x70, 0x5e, 0xc7, 0x79, 0x89, 0xf3, 0x0a, 0xdf, 0x3d, 0xa0, 0x51, 0xea, 0xe2, 0xfb, 0x4c, 0x3a,
	0x9a, 0x89, 0x32, 0x46, 0x62, 0x13, 0xfb, 0xf7, 0x2b, 0xed, 0xfb, 0xaf, 0x7e, 0xb0, 0xad, 0x01,
	0x51, 0x26, 0xd0, 0xbc, 0x0c, 0x36, 0x1f, 0x9c, 0xd5, 0x2d, 0x7f, 0xcf, 0x3f, 0xdc, 0xe9, 0x9d,
	0x2e, 0x56, 0xa1, 0xf7, 0xbe, 0x0a, 0xdb, 0x49, 0x4a, 0xa3, 0x49, 0xc4, 0x95, 0xd5, 0x42, 0x59,
	0xd4, 0x16, 0xab, 0xa7, 0x8d, 0xf1, 0xa3, 0xa0, 0x59, 0x06, 0xc8, 0xbb, 0x4a, 0x75, 0xe3, 0xd8,
	0x01, 0xe2, 0xb0, 0x50, 0x34, 0xa3, 0xa0, 0xf1, 0x9d, 0x71, 0x20, 0x15, 0xb5, 0x36, 0x0a, 0xdd,
	0xa0, 0xd2, 0x9d, 0xd7, 0x74, 0xb5, 0xcb, 0x6e, 0x0c, 0xd4, 0xc7, 0xdb, 0xfe, 0x95, 0x40, 0x35,
	0x02, 0x2d, 0xab, 0x0d, 0xd7, 0x95, 0x6d, 0xf8, 0xeb, 0x3d, 0x6b, 0x3c, 0xcf, 0x43, 0xef, 0x73,
	0x1e, 0x7a, 0xbd, 0xfe, 0x62, 0xcd, 0xfc, 0xe5, 0x9a, 0xf9, 0x1f, 0x6b, 0xe6, 0xbf, 0xe4, 0xcc,
	0x5b, 0xe6, 0xcc, 0x7b, 0xcb, 0x99, 0x77, 0x77, 0xa4, 0x6d, 0x3c, 0x19, 0xc3, 0x4f, 0xe1, 0xa9,
	0x21, 0x70, 0x46, 0x8e, 0xff, 0x6b, 0x3e, 0xda, 0x2a, 0x1a, 0x39, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x83, 0x88, 0x30, 0x5b, 0xb6, 0x01, 0x00, 0x00,
}

func (m *message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MetaFact.Size() > 0 {
		i -= m.MetaFact.Size()
		if _, err := m.MetaFact.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(m.MetaFact.Size()))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
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
func (m *message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MetaFact.Size()
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
func (m *message) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaFact", wireType)
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
			if err := m.MetaFact.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
