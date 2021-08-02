// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/persistence/assets/metas/transactions/reveal/v1beta1/request.proto

package reveal

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	test_types "github.com/persistenceOne/persistenceSDK/schema/test_types"
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

type transactionRequest struct {
	BaseReq  test_types.BaseReq `protobuf:"bytes,1,opt,name=baseReq,proto3" json:"baseReq"`
	MetaFact string             `protobuf:"bytes,2,opt,name=metaFact,proto3" json:"metaFact,omitempty" valid:"required~required field metaFact missing, matches(^[DHIS]{1}[|]{1}.*$)"`
}

func (m *transactionRequest) Reset()         { *m = transactionRequest{} }
func (m *transactionRequest) String() string { return proto.CompactTextString(m) }
func (*transactionRequest) ProtoMessage()    {}
func (*transactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_641b8cccb7bc96b3, []int{0}
}
func (m *transactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *transactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *transactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *transactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*transactionRequest)(nil), "persistence.assets.metas.transactions.reveal.v1beta1.transactionRequest")
}

func init() {
	proto.RegisterFile("proto/persistence/assets/metas/transactions/reveal/v1beta1/request.proto", fileDescriptor_641b8cccb7bc96b3)
}

var fileDescriptor_641b8cccb7bc96b3 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x6d, 0x84, 0xa0, 0x84, 0x2d, 0x53, 0xd5, 0xc1, 0xad, 0x18, 0x50, 0xc5, 0x60, 0xab,
	0x94, 0x89, 0x31, 0x42, 0xa8, 0x73, 0x46, 0x16, 0xe4, 0xa4, 0xa7, 0x12, 0x29, 0x89, 0x13, 0xdf,
	0xb5, 0x12, 0x3b, 0x03, 0x23, 0x8f, 0xd0, 0xc7, 0xe9, 0xd8, 0x91, 0x09, 0xa1, 0x64, 0xe1, 0x31,
	0x50, 0x62, 0x52, 0x22, 0xc1, 0xe6, 0x93, 0xfe, 0xfb, 0xfc, 0xdd, 0xef, 0x2d, 0x0a, 0x6b, 0xc8,
	0xa8, 0x02, 0x2c, 0x26, 0x48, 0x90, 0xc7, 0xa0, 0x34, 0x22, 0x10, 0xaa, 0x0c, 0x48, 0xa3, 0x22,
	0xab, 0x73, 0xd4, 0x31, 0x25, 0x26, 0x47, 0x65, 0x61, 0x03, 0x3a, 0x55, 0x9b, 0x59, 0x04, 0xa4,
	0x67, 0xca, 0x42, 0xb9, 0x06, 0x24, 0xd9, 0x22, 0xfc, 0x9b, 0x1e, 0x43, 0x3a, 0x86, 0x6c, 0x19,
	0xb2, 0xcf, 0x90, 0x8e, 0x21, 0x7f, 0x18, 0xa3, 0x4b, 0x7a, 0x4a, 0xec, 0xf2, 0xb1, 0xd0, 0x96,
	0x9e, 0x95, 0x73, 0x59, 0x99, 0x95, 0xf9, 0x7d, 0x39, 0xfa, 0x68, 0xfe, 0xd7, 0x33, 0x36, 0x98,
	0x19, 0x54, 0x91, 0x46, 0x38, 0x28, 0x35, 0x83, 0x85, 0xd2, 0x2d, 0x5d, 0xbc, 0x70, 0xcf, 0xef,
	0x7d, 0x1e, 0x3a, 0x5f, 0x7f, 0xe1, 0x9d, 0x36, 0xb9, 0x10, 0xca, 0x21, 0x9f, 0xf0, 0xe9, 0xf9,
	0xf5, 0x54, 0xf6, 0xdd, 0x1d, 0x57, 0x36, 0x91, 0x4e, 0x53, 0x06, 0x2e, 0x1f, 0x1c, 0xef, 0x3e,
	0xc6, 0x2c, 0xec, 0xd6, 0xfd, 0x91, 0x37, 0x68, 0x4e, 0xbc, 0xd7, 0x31, 0x0d, 0x8f, 0x26, 0x7c,
	0x7a, 0x16, 0x1e, 0xe6, 0xdb, 0xc1, 0xeb, 0x76, 0xcc, 0xbe, 0xb6, 0x63, 0x16, 0xdc, 0xed, 0x2a,
	0xc1, 0xf7, 0x95, 0xe0, 0x9f, 0x95, 0xe0, 0x6f, 0xb5, 0x60, 0xfb, 0x5a, 0xb0, 0xf7, 0x5a, 0xb0,
	0x87, 0xab, 0xcc, 0x2c, 0xd7, 0x29, 0x74, 0x65, 0x27, 0x39, 0x81, 0xcd, 0x75, 0xfa, 0x5f, 0xeb,
	0xd1, 0x49, 0x7b, 0xd3, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x76, 0x70, 0xa1, 0xb1, 0xb2, 0x01,
	0x00, 0x00,
}

func (m *transactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MetaFact) > 0 {
		i -= len(m.MetaFact)
		copy(dAtA[i:], m.MetaFact)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.MetaFact)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.BaseReq.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *transactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseReq.Size()
	n += 1 + l + sovRequest(uint64(l))
	l = len(m.MetaFact)
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
func (m *transactionRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: transactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaFact", wireType)
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
			m.MetaFact = string(dAtA[iNdEx:postIndex])
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
