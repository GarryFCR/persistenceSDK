// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/orders/transactions/immediate/v1beta1/response.proto

package immediate

import (
	"errors"
	fmt "fmt"
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

type transactionResponse struct {
	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   error `protobuf:"bytes,2,opt,name=error,proto3,customtype=error" json:"error"`
}

func (m *transactionResponse) Reset()         { *m = transactionResponse{} }
func (m *transactionResponse) String() string { return proto.CompactTextString(m) }
func (*transactionResponse) ProtoMessage()    {}
func (*transactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c73a7aa3ada50bc, []int{0}
}
func (m *transactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *transactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *transactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *transactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *transactionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*transactionResponse)(nil), "persistence.assets.orders.transactions.immediate.v1beta1.transactionResponse")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/orders/transactions/immediate/v1beta1/response.proto", fileDescriptor_0c73a7aa3ada50bc)
}

var fileDescriptor_0c73a7aa3ada50bc = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xbd, 0x4a, 0x03, 0x41,
	0x14, 0x05, 0xe0, 0x19, 0xf1, 0x77, 0xc1, 0x66, 0x6d, 0x16, 0x8b, 0xd9, 0xa0, 0x20, 0xa9, 0x66,
	0x08, 0x36, 0x62, 0x99, 0xca, 0xc6, 0x66, 0xcb, 0x34, 0x32, 0xd9, 0xbd, 0xc4, 0x81, 0xec, 0xcc,
	0x72, 0xef, 0x8d, 0xe0, 0x5b, 0xf8, 0x08, 0x3e, 0x4e, 0xca, 0x94, 0x62, 0x11, 0x64, 0xb7, 0xf1,
	0x31, 0x84, 0x4c, 0x5c, 0xb7, 0xb1, 0x9a, 0x9f, 0xe2, 0x3b, 0x87, 0x93, 0x3c, 0x36, 0x18, 0x38,
	0x98, 0x06, 0x90, 0x1c, 0x31, 0xf8, 0x12, 0x8c, 0x25, 0x02, 0x26, 0x13, 0xb0, 0x02, 0x24, 0xc3,
	0x68, 0x3d, 0xd9, 0x92, 0x5d, 0xf0, 0x64, 0x5c, 0x5d, 0x43, 0xe5, 0x2c, 0x83, 0x79, 0x99, 0xcc,
	0x81, 0xed, 0xc4, 0x20, 0x50, 0x13, 0x3c, 0x81, 0xde, 0x39, 0xe9, 0xdd, 0x00, 0xd2, 0x11, 0xd2,
	0x11, 0xd2, 0x43, 0x48, 0xf7, 0x90, 0xde, 0x43, 0x97, 0x37, 0xfc, 0xec, 0xb0, 0x7a, 0x6a, 0x2c,
	0xf2, 0xab, 0x89, 0xa5, 0x16, 0x61, 0x11, 0xfe, 0x6e, 0x31, 0xe1, 0x6a, 0x96, 0x5c, 0x0c, 0xa4,
	0x62, 0x1f, 0x9f, 0x66, 0xc9, 0x09, 0xad, 0xca, 0x12, 0x88, 0x32, 0x39, 0x92, 0xe3, 0xd3, 0xe2,
	0xf7, 0x99, 0x5e, 0x27, 0x47, 0x80, 0x18, 0x30, 0x3b, 0x18, 0xc9, 0xf1, 0xd9, 0xf4, 0x7c, 0xbd,
	0xcd, 0xc5, 0xe7, 0x36, 0x8f, 0x9f, 0x45, 0x3c, 0xee, 0x0f, 0xbf, 0xdf, 0x73, 0x31, 0x7d, 0x58,
	0xb7, 0x4a, 0x6e, 0x5a, 0x25, 0xbf, 0x5a, 0x25, 0xdf, 0x3a, 0x25, 0x36, 0x9d, 0x12, 0x1f, 0x9d,
	0x12, 0x33, 0x5d, 0x87, 0x6a, 0xb5, 0x84, 0x7e, 0x15, 0xe7, 0x19, 0xd0, 0xdb, 0xe5, 0x3f, 0xf3,
	0xcc, 0x8f, 0x77, 0x65, 0x6f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x53, 0x2d, 0x33, 0x86, 0x5f,
	0x01, 0x00, 0x00,
}

func (m *transactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *transactionResponse) Size() (n int) {
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
	return n
}

func sovResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponse(x uint64) (n int) {
	return sovResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *transactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: transactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
