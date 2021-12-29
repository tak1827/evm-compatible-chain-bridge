// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

package pb

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConfirmedBlock struct {
	Hash                 string     `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Number               string     `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	UpdatedAt            *time.Time `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConfirmedBlock) Reset()      { *m = ConfirmedBlock{} }
func (*ConfirmedBlock) ProtoMessage() {}
func (*ConfirmedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}
func (m *ConfirmedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfirmedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfirmedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfirmedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmedBlock.Merge(m, src)
}
func (m *ConfirmedBlock) XXX_Size() int {
	return m.Size()
}
func (m *ConfirmedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmedBlock proto.InternalMessageInfo

func (m *ConfirmedBlock) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ConfirmedBlock) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *ConfirmedBlock) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfirmedBlock)(nil), "tak1827.transactionconfirmer.sample.ConfirmedBlock")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xca, 0xc9, 0x4f,
	0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2e, 0x49, 0xcc, 0x36, 0xb4, 0x30, 0x32,
	0xd7, 0x2b, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x4b, 0xce, 0xcf, 0x4b,
	0xcb, 0x2c, 0xca, 0x4d, 0x2d, 0xd2, 0x2b, 0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0x95, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0xab, 0xd7, 0x07, 0xb1, 0x20, 0x5a, 0xa5, 0xe4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73,
	0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0x92, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc,
	0x02, 0x88, 0x02, 0xa5, 0x5a, 0x2e, 0x3e, 0x67, 0xa8, 0x51, 0x29, 0x4e, 0x20, 0x3b, 0x85, 0x84,
	0xb8, 0x58, 0x32, 0x12, 0x8b, 0x33, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21,
	0x31, 0x2e, 0xb6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x26, 0xb0, 0x28, 0x94, 0x27, 0x64,
	0xcf, 0xc5, 0x55, 0x5a, 0x90, 0x92, 0x58, 0x92, 0x9a, 0x12, 0x9f, 0x58, 0x22, 0xc1, 0xac, 0xc0,
	0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0xb1, 0x53, 0x0f, 0x66, 0xa7, 0x5e, 0x08, 0xcc, 0x4e, 0x27,
	0x96, 0x09, 0xf7, 0xe5, 0x19, 0x83, 0x38, 0xa1, 0x7a, 0x1c, 0x4b, 0x9c, 0xdc, 0x6e, 0x3c, 0x94,
	0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51,
	0x2a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xd0, 0x70, 0xd0, 0x4f,
	0x2d, 0xcb, 0xd5, 0x4d, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0xd5, 0x4f, 0xce, 0xc9, 0xd4, 0x2f, 0x48,
	0x4a, 0x62, 0x03, 0x5b, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x4a, 0x6c, 0xde, 0x38,
	0x01, 0x00, 0x00,
}

func (this *ConfirmedBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfirmedBlock)
	if !ok {
		that2, ok := that.(ConfirmedBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Hash != that1.Hash {
		return false
	}
	if this.Number != that1.Number {
		return false
	}
	if that1.UpdatedAt == nil {
		if this.UpdatedAt != nil {
			return false
		}
	} else if !this.UpdatedAt.Equal(*that1.UpdatedAt) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ConfirmedBlock) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pb.ConfirmedBlock{")
	s = append(s, "Hash: "+fmt.Sprintf("%#v", this.Hash)+",\n")
	s = append(s, "Number: "+fmt.Sprintf("%#v", this.Number)+",\n")
	s = append(s, "UpdatedAt: "+fmt.Sprintf("%#v", this.UpdatedAt)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringBlock(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ConfirmedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfirmedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfirmedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdatedAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.UpdatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintBlock(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Number) > 0 {
		i -= len(m.Number)
		copy(dAtA[i:], m.Number)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Number)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfirmedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Number)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.UpdatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt)
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ConfirmedBlock) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConfirmedBlock{`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`UpdatedAt:` + strings.Replace(fmt.Sprintf("%v", this.UpdatedAt), "Timestamp", "timestamppb.Timestamp", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlock(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ConfirmedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: ConfirmedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfirmedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Number = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedAt == nil {
				m.UpdatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)