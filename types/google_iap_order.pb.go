// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/google_iap_order.proto

package types

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

type GoogleInAppPurchaseOrder struct {
	Creator           string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ProductID         string `protobuf:"bytes,2,opt,name=productID,proto3" json:"productID,omitempty"`
	PurchaseToken     string `protobuf:"bytes,3,opt,name=purchaseToken,proto3" json:"purchaseToken,omitempty"`
	ReceiptDataBase64 string `protobuf:"bytes,4,opt,name=receiptDataBase64,proto3" json:"receiptDataBase64,omitempty"`
	Signature         string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *GoogleInAppPurchaseOrder) Reset()         { *m = GoogleInAppPurchaseOrder{} }
func (m *GoogleInAppPurchaseOrder) String() string { return proto.CompactTextString(m) }
func (*GoogleInAppPurchaseOrder) ProtoMessage()    {}
func (*GoogleInAppPurchaseOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a9266864181088e, []int{0}
}
func (m *GoogleInAppPurchaseOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoogleInAppPurchaseOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoogleInAppPurchaseOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoogleInAppPurchaseOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleInAppPurchaseOrder.Merge(m, src)
}
func (m *GoogleInAppPurchaseOrder) XXX_Size() int {
	return m.Size()
}
func (m *GoogleInAppPurchaseOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleInAppPurchaseOrder.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleInAppPurchaseOrder proto.InternalMessageInfo

func (m *GoogleInAppPurchaseOrder) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *GoogleInAppPurchaseOrder) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *GoogleInAppPurchaseOrder) GetPurchaseToken() string {
	if m != nil {
		return m.PurchaseToken
	}
	return ""
}

func (m *GoogleInAppPurchaseOrder) GetReceiptDataBase64() string {
	if m != nil {
		return m.ReceiptDataBase64
	}
	return ""
}

func (m *GoogleInAppPurchaseOrder) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*GoogleInAppPurchaseOrder)(nil), "Pylonstech.pylons.pylons.GoogleInAppPurchaseOrder")
}

func init() { proto.RegisterFile("pylons/google_iap_order.proto", fileDescriptor_1a9266864181088e) }

var fileDescriptor_1a9266864181088e = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xa8, 0xcc, 0xc9,
	0xcf, 0x2b, 0xd6, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0x8d, 0xcf, 0x4c, 0x2c, 0x88, 0xcf, 0x2f,
	0x4a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x08, 0x00, 0x4b, 0x97, 0xa4,
	0x26, 0x67, 0xe8, 0x41, 0x54, 0x42, 0x29, 0xa5, 0x23, 0x8c, 0x5c, 0x12, 0xee, 0x60, 0x4d, 0x9e,
	0x79, 0x8e, 0x05, 0x05, 0x01, 0xa5, 0x45, 0xc9, 0x19, 0x89, 0xc5, 0xa9, 0xfe, 0x20, 0xcd, 0x42,
	0x12, 0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x30, 0xae, 0x90, 0x0c, 0x17, 0x67, 0x41, 0x51, 0x7e, 0x4a, 0x69, 0x72, 0x89, 0xa7, 0x8b,
	0x04, 0x13, 0x58, 0x0e, 0x21, 0x20, 0xa4, 0xc2, 0xc5, 0x5b, 0x00, 0x35, 0x28, 0x24, 0x3f, 0x3b,
	0x35, 0x4f, 0x82, 0x19, 0xac, 0x02, 0x55, 0x50, 0x48, 0x87, 0x4b, 0xb0, 0x28, 0x35, 0x39, 0x35,
	0xb3, 0xa0, 0xc4, 0x25, 0xb1, 0x24, 0xd1, 0x29, 0xb1, 0x38, 0xd5, 0xcc, 0x44, 0x82, 0x05, 0xac,
	0x12, 0x53, 0x02, 0x64, 0x63, 0x71, 0x66, 0x7a, 0x5e, 0x62, 0x49, 0x69, 0x51, 0xaa, 0x04, 0x2b,
	0xc4, 0x46, 0xb8, 0x80, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x42, 0x41, 0x17,
	0x14, 0x0c, 0xfa, 0xd0, 0x00, 0xab, 0x80, 0x31, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0,
	0xe1, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x1a, 0xa1, 0xf0, 0x50, 0x01, 0x00, 0x00,
}

func (m *GoogleInAppPurchaseOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoogleInAppPurchaseOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoogleInAppPurchaseOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintGoogleIapOrder(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ReceiptDataBase64) > 0 {
		i -= len(m.ReceiptDataBase64)
		copy(dAtA[i:], m.ReceiptDataBase64)
		i = encodeVarintGoogleIapOrder(dAtA, i, uint64(len(m.ReceiptDataBase64)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PurchaseToken) > 0 {
		i -= len(m.PurchaseToken)
		copy(dAtA[i:], m.PurchaseToken)
		i = encodeVarintGoogleIapOrder(dAtA, i, uint64(len(m.PurchaseToken)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProductID) > 0 {
		i -= len(m.ProductID)
		copy(dAtA[i:], m.ProductID)
		i = encodeVarintGoogleIapOrder(dAtA, i, uint64(len(m.ProductID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGoogleIapOrder(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGoogleIapOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovGoogleIapOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoogleInAppPurchaseOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGoogleIapOrder(uint64(l))
	}
	l = len(m.ProductID)
	if l > 0 {
		n += 1 + l + sovGoogleIapOrder(uint64(l))
	}
	l = len(m.PurchaseToken)
	if l > 0 {
		n += 1 + l + sovGoogleIapOrder(uint64(l))
	}
	l = len(m.ReceiptDataBase64)
	if l > 0 {
		n += 1 + l + sovGoogleIapOrder(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovGoogleIapOrder(uint64(l))
	}
	return n
}

func sovGoogleIapOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGoogleIapOrder(x uint64) (n int) {
	return sovGoogleIapOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoogleInAppPurchaseOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoogleIapOrder
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
			return fmt.Errorf("proto: GoogleInAppPurchaseOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoogleInAppPurchaseOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoogleIapOrder
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
				return ErrInvalidLengthGoogleIapOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGoogleIapOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoogleIapOrder
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
				return ErrInvalidLengthGoogleIapOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGoogleIapOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoogleIapOrder
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
				return ErrInvalidLengthGoogleIapOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGoogleIapOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PurchaseToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptDataBase64", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoogleIapOrder
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
				return ErrInvalidLengthGoogleIapOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGoogleIapOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiptDataBase64 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoogleIapOrder
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
				return ErrInvalidLengthGoogleIapOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGoogleIapOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGoogleIapOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGoogleIapOrder
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
func skipGoogleIapOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGoogleIapOrder
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
					return 0, ErrIntOverflowGoogleIapOrder
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
					return 0, ErrIntOverflowGoogleIapOrder
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
				return 0, ErrInvalidLengthGoogleIapOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGoogleIapOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGoogleIapOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGoogleIapOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGoogleIapOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGoogleIapOrder = fmt.Errorf("proto: unexpected end of group")
)
