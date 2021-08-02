// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/splits/queries/ownable/v1beta1/request.proto

package ownable

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

type queryRequest struct {
	OwnableID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=ownableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"ownableID" valid:"required~required field ownableID missing"`
}

func (m *queryRequest) Reset()         { *m = queryRequest{} }
func (m *queryRequest) String() string { return proto.CompactTextString(m) }
func (*queryRequest) ProtoMessage()    {}
func (*queryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ccbfb9712316af5, []int{0}
}
func (m *queryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *queryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *queryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *queryRequest) XXX_Size() int {
	return m.Size()
}
func (m *queryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*queryRequest)(nil), "persistence.assets.splits.queries.ownable.v1beta1.queryRequest")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/splits/queries/ownable/v1beta1/request.proto", fileDescriptor_2ccbfb9712316af5)
}

var fileDescriptor_2ccbfb9712316af5 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x48, 0x2d, 0x2a, 0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x4f, 0x2c,
	0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d,
	0xca, 0x4c, 0x2d, 0xd6, 0xcf, 0x2f, 0xcf, 0x4b, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x03, 0xeb, 0x17, 0x32,
	0x44, 0x32, 0x40, 0x0f, 0x62, 0x80, 0x1e, 0xc4, 0x00, 0x3d, 0xa8, 0x01, 0x7a, 0x50, 0x03, 0xf4,
	0xa0, 0x06, 0x48, 0xa9, 0x95, 0x64, 0x64, 0x16, 0xa5, 0xc4, 0x17, 0x24, 0x16, 0x95, 0x54, 0xea,
	0x43, 0x5c, 0x91, 0x9e, 0x9f, 0x9e, 0x8f, 0x60, 0x41, 0x8c, 0x56, 0xaa, 0xe0, 0xe2, 0x01, 0x19,
	0x51, 0x19, 0x04, 0xb1, 0x50, 0x28, 0x8e, 0x8b, 0x13, 0x6a, 0x94, 0xa7, 0x8b, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xa7, 0x93, 0xc3, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x5b, 0xa4, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x22, 0xfb, 0xc8, 0x3f, 0x2f, 0x15, 0x99, 0x1b, 0xec,
	0xe2, 0xad, 0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0xa8, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0xac, 0xe7,
	0xe9, 0x12, 0x84, 0x30, 0xd2, 0x8a, 0xa3, 0x63, 0x81, 0x3c, 0xc3, 0x8b, 0x05, 0xf2, 0x0c, 0x4e,
	0x8e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9e, 0x9b, 0x9f, 0x52,
	0x9a, 0x93, 0x0a, 0x0f, 0xa4, 0xcc, 0xbc, 0x92, 0xd4, 0xa2, 0xbc, 0xc4, 0x1c, 0xf4, 0xd0, 0x4a,
	0x62, 0x03, 0xfb, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xae, 0x43, 0xd4, 0x67, 0x01,
	0x00, 0x00,
}

func (m *queryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *queryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *queryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OwnableID.Size() > 0 {
		i -= m.OwnableID.Size()
		if _, err := m.OwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRequest(dAtA, i, uint64(m.OwnableID.Size()))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *queryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OwnableID.Size()
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *queryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: queryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: queryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)
