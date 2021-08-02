// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/splits/queries/split/v1beta1/response.proto

package split

import (
	"errors"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
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
	return fileDescriptor_3c2cedadf96794ce, []int{0}
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
	proto.RegisterType((*queryResponse)(nil), "persistence.assets.splits.queries.split.v1beta1.queryResponse")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/splits/queries/split/v1beta1/response.proto", fileDescriptor_3c2cedadf96794ce)
}

var fileDescriptor_3c2cedadf96794ce = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4a, 0x03, 0x31,
	0x14, 0x87, 0x27, 0xb6, 0xfe, 0xe9, 0x40, 0x37, 0xb3, 0x1a, 0x5c, 0xa4, 0x45, 0x51, 0xba, 0x4a,
	0x28, 0xee, 0xdc, 0x28, 0xa5, 0x0b, 0x41, 0x44, 0x18, 0x77, 0x22, 0x48, 0x66, 0xfa, 0xe8, 0x04,
	0xd2, 0x49, 0xcc, 0xcb, 0x08, 0xbd, 0x85, 0x47, 0xf0, 0x0c, 0x9e, 0xa2, 0xcb, 0x2e, 0xc5, 0x45,
	0x91, 0x76, 0xe3, 0x31, 0xa4, 0x93, 0x51, 0xc7, 0x55, 0xf2, 0x25, 0xf9, 0x7e, 0x79, 0xef, 0x85,
	0x63, 0x63, 0xb5, 0xd3, 0xdc, 0x80, 0x45, 0x89, 0x0e, 0x8a, 0x0c, 0xb8, 0x40, 0x04, 0x87, 0x1c,
	0x8d, 0x92, 0x0e, 0xf9, 0x53, 0x09, 0x56, 0x42, 0x8d, 0xfc, 0x79, 0x98, 0x82, 0x13, 0x43, 0x6e,
	0x01, 0x8d, 0x2e, 0x10, 0x58, 0xa5, 0x47, 0x4d, 0x9f, 0x79, 0x9f, 0x79, 0x9f, 0xd5, 0xbe, 0x47,
	0x56, 0xfb, 0x87, 0xa7, 0x2e, 0x97, 0x76, 0xf2, 0x68, 0x84, 0x75, 0x73, 0xee, 0x4b, 0x98, 0xea,
	0xa9, 0xfe, 0xdb, 0xf9, 0xe0, 0xa3, 0x37, 0x12, 0x76, 0xb7, 0x09, 0xf3, 0xa4, 0xfe, 0x30, 0x8a,
	0xc3, 0x7d, 0x2c, 0xb3, 0x0c, 0x10, 0x63, 0xd2, 0x27, 0x83, 0x83, 0xe4, 0x07, 0xa3, 0xe3, 0x70,
	0x17, 0xac, 0xd5, 0x36, 0xde, 0xe9, 0x93, 0x41, 0x67, 0xd4, 0x5d, 0xac, 0x7a, 0xc1, 0xc7, 0xaa,
	0xe7, 0x0f, 0x13, 0xbf, 0x44, 0x0f, 0x61, 0x5b, 0x49, 0x74, 0x71, 0xab, 0xdf, 0x1a, 0x74, 0x46,
	0x57, 0xf5, 0x9b, 0xcb, 0xa9, 0x74, 0x79, 0x99, 0xb2, 0x4c, 0xcf, 0x9a, 0xad, 0xdc, 0x16, 0xd0,
	0xc4, 0xbb, 0xf1, 0x35, 0xc7, 0x2c, 0x87, 0x99, 0xe0, 0x39, 0xa8, 0xed, 0x05, 0xbb, 0x11, 0xc6,
	0x88, 0x54, 0x41, 0x52, 0xa5, 0x9e, 0xb7, 0xbf, 0x5e, 0x7b, 0xc1, 0xe8, 0x62, 0xb1, 0xa6, 0x64,
	0xb9, 0xa6, 0xe4, 0x73, 0x4d, 0xc9, 0xcb, 0x86, 0x06, 0xcb, 0x0d, 0x0d, 0xde, 0x37, 0x34, 0xb8,
	0x3f, 0x99, 0xe9, 0x49, 0xa9, 0xe0, 0x77, 0xb8, 0xb2, 0x70, 0x60, 0x0b, 0xa1, 0xfe, 0x4f, 0x39,
	0xdd, 0xab, 0x9a, 0x3f, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x63, 0x73, 0xa8, 0x9d, 0x01,
	0x00, 0x00,
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
			m.List = append(m.List, mappable.Split{})
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
