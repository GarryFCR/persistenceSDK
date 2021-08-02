// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/orders/queries/orders/v1beta1/response.proto

package order

import (
	"errors"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	github_com_persistenceOne_persistenceSDK_schema_helpers "github.com/persistenceOne/persistenceSDK/schema/helpers"
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

type queryResponse struct {
	Success bool                                                               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   error                                                              `protobuf:"bytes,2,opt,name=error,proto3,customtype=error" json:"error"`
	List    []github_com_persistenceOne_persistenceSDK_schema_helpers.Mappable `protobuf:"bytes,3,rep,name=list,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/helpers.Mappable" json:"list"`
}

func (m *queryResponse) Reset()         { *m = queryResponse{} }
func (m *queryResponse) String() string { return proto.CompactTextString(m) }
func (*queryResponse) ProtoMessage()    {}
func (*queryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e675ce6d6798364, []int{0}
}
func (m *queryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *queryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *queryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *queryResponse) XXX_Size() int {
	return m.Size()
}
func (m *queryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *queryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*queryResponse)(nil), "persistence.assets.orders.queries.order.v1beta1.queryResponse")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/orders/queries/orders/v1beta1/response.proto", fileDescriptor_6e675ce6d6798364)
}

var fileDescriptor_6e675ce6d6798364 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0xc7, 0x37, 0x5f, 0xfb, 0xa9, 0x5d, 0xe8, 0x65, 0x4f, 0x8b, 0x87, 0xb4, 0x28, 0x4a, 0x4f,
	0x09, 0xc5, 0x9b, 0x17, 0xa5, 0x28, 0x08, 0x22, 0xc2, 0x7a, 0x13, 0x41, 0xb2, 0xdb, 0xa1, 0x1b,
	0xd8, 0x6e, 0xe2, 0x4c, 0x56, 0xe8, 0x5b, 0xf8, 0x08, 0x3e, 0x83, 0x4f, 0xd1, 0x63, 0x8f, 0xe2,
	0xa1, 0x48, 0x7b, 0xf1, 0x31, 0xa4, 0xcd, 0x6a, 0xeb, 0x29, 0xf9, 0x25, 0xf9, 0xff, 0x32, 0x33,
	0xe1, 0xa5, 0x45, 0xe3, 0x8c, 0xb4, 0x80, 0xa4, 0xc9, 0x41, 0x99, 0x81, 0x54, 0x44, 0xe0, 0x48,
	0x1a, 0x1c, 0x02, 0x92, 0x7c, 0xaa, 0x00, 0x35, 0xfc, 0xe2, 0x73, 0x3f, 0x05, 0xa7, 0xfa, 0x12,
	0x81, 0xac, 0x29, 0x09, 0xc4, 0x3a, 0x1f, 0x6d, 0x0b, 0x84, 0x17, 0x08, 0x9f, 0x10, 0xb5, 0xc0,
	0xa3, 0xa8, 0xf3, 0xfb, 0xc7, 0x2e, 0xd7, 0x38, 0x7c, 0xb4, 0x0a, 0xdd, 0x44, 0xfa, 0x1a, 0x46,
	0x66, 0x64, 0x36, 0x3b, 0x2f, 0x3e, 0x78, 0x63, 0x61, 0x7b, 0x65, 0x98, 0x24, 0xf5, 0x87, 0x51,
	0x1c, 0xee, 0x52, 0x95, 0x65, 0x40, 0x14, 0xb3, 0x2e, 0xeb, 0xed, 0x25, 0x3f, 0x18, 0x1d, 0x86,
	0xff, 0x01, 0xd1, 0x60, 0xfc, 0xaf, 0xcb, 0x7a, 0xad, 0x41, 0x7b, 0x3a, 0xef, 0x04, 0x1f, 0xf3,
	0x8e, 0x3f, 0x4c, 0xfc, 0x12, 0x3d, 0x84, 0xcd, 0x42, 0x93, 0x8b, 0x1b, 0xdd, 0x46, 0xaf, 0x35,
	0xb8, 0xaa, 0xdf, 0x9c, 0x8f, 0xb4, 0xcb, 0xab, 0x54, 0x64, 0x66, 0xbc, 0xdd, 0xca, 0x6d, 0x09,
	0xdb, 0x78, 0x77, 0x71, 0x2d, 0x29, 0xcb, 0x61, 0xac, 0x64, 0x0e, 0xc5, 0xea, 0x42, 0xdc, 0x28,
	0x6b, 0x55, 0x5a, 0x40, 0xb2, 0xb6, 0x9e, 0x36, 0xbf, 0x5e, 0x3b, 0xc1, 0xe0, 0x6c, 0xba, 0xe0,
	0x6c, 0xb6, 0xe0, 0xec, 0x73, 0xc1, 0xd9, 0xcb, 0x92, 0x07, 0xb3, 0x25, 0x0f, 0xde, 0x97, 0x3c,
	0xb8, 0x3f, 0x1a, 0x9b, 0x61, 0x55, 0x6c, 0xc6, 0xa9, 0x4b, 0x07, 0x58, 0xaa, 0xe2, 0xef, 0x98,
	0xd3, 0x9d, 0x75, 0xf3, 0x27, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0x2d, 0x38, 0x96, 0x9e,
	0x01, 0x00, 0x00,
}

func (m *queryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *queryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *queryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			i -= m.List[iNdEx].Size()
			if _, err := m.List[iNdEx].MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintResponse(dAtA, i, uint64(m.List[iNdEx].Size()))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Error.Error()) > 0 {
		i -= len(m.Error.Error())
		copy(dAtA[i:], m.Error.Error())
		i = encodeVarintResponse(dAtA, i, uint64(len(m.Error.Error())))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *queryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Error.Error())
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	if len(m.List) > 0 {
		for _, s := range m.List {
			l = s.Size()
			n += 1 + l + sovResponse(uint64(l))
		}
	}
	return n
}

func sovResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponse(x uint64) (n int) {
	return sovResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *queryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponse
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
			return fmt.Errorf("proto: queryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: queryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = error(errors.New(string(dAtA[iNdEx:postIndex])))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, mappable.Order{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResponse
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
func skipResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
				return 0, ErrInvalidLengthResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResponse = fmt.Errorf("proto: unexpected end of group")
)
