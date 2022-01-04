// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pair.proto

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

type Pair_Type int32

const (
	Pair_ORIGINAL Pair_Type = 0
	Pair_WRAPPED  Pair_Type = 1
)

var Pair_Type_name = map[int32]string{
	0: "ORIGINAL",
	1: "WRAPPED",
}

var Pair_Type_value = map[string]int32{
	"ORIGINAL": 0,
	"WRAPPED":  1,
}

func (x Pair_Type) String() string {
	return proto.EnumName(Pair_Type_name, int32(x))
}

func (Pair_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{0, 0}
}

type Pair struct {
	Inaddr               string     `protobuf:"bytes,1,opt,name=inaddr,proto3" json:"inaddr,omitempty"`
	Outaddr              string     `protobuf:"bytes,2,opt,name=outaddr,proto3" json:"outaddr,omitempty"`
	Intype               Pair_Type  `protobuf:"varint,3,opt,name=intype,proto3,enum=tak1827.evmbridge.cli.Pair_Type" json:"intype,omitempty"`
	Outtype              Pair_Type  `protobuf:"varint,4,opt,name=outtype,proto3,enum=tak1827.evmbridge.cli.Pair_Type" json:"outtype,omitempty"`
	UpdatedAt            *time.Time `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Pair) Reset()      { *m = Pair{} }
func (*Pair) ProtoMessage() {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6c646fab57af36d, []int{0}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return m.Size()
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetInaddr() string {
	if m != nil {
		return m.Inaddr
	}
	return ""
}

func (m *Pair) GetOutaddr() string {
	if m != nil {
		return m.Outaddr
	}
	return ""
}

func (m *Pair) GetIntype() Pair_Type {
	if m != nil {
		return m.Intype
	}
	return Pair_ORIGINAL
}

func (m *Pair) GetOuttype() Pair_Type {
	if m != nil {
		return m.Outtype
	}
	return Pair_ORIGINAL
}

func (m *Pair) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("tak1827.evmbridge.cli.Pair_Type", Pair_Type_name, Pair_Type_value)
	proto.RegisterType((*Pair)(nil), "tak1827.evmbridge.cli.Pair")
}

func init() { proto.RegisterFile("pair.proto", fileDescriptor_b6c646fab57af36d) }

var fileDescriptor_b6c646fab57af36d = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0xcc, 0x2c,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2d, 0x49, 0xcc, 0x36, 0xb4, 0x30, 0x32, 0xd7,
	0x4b, 0x2d, 0xcb, 0x4d, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0xd5, 0x4b, 0xce, 0xc9, 0x94, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd0, 0x07, 0xb1, 0x20, 0x8a, 0xa5, 0xe4, 0xd3, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0x92, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4,
	0xdc, 0x02, 0x88, 0x02, 0xa5, 0x49, 0x4c, 0x5c, 0x2c, 0x01, 0x89, 0x99, 0x45, 0x42, 0x62, 0x5c,
	0x6c, 0x99, 0x79, 0x89, 0x29, 0x29, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e,
	0x90, 0x04, 0x17, 0x7b, 0x7e, 0x69, 0x09, 0x58, 0x82, 0x09, 0x2c, 0x01, 0xe3, 0x0a, 0x59, 0x80,
	0x74, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x19, 0x29, 0xe8, 0x61, 0x75,
	0x99, 0x1e, 0xc8, 0x78, 0xbd, 0x90, 0xca, 0x82, 0xd4, 0x20, 0xa8, 0x7a, 0x21, 0x2b, 0xb0, 0x99,
	0x60, 0xad, 0x2c, 0x44, 0x6a, 0x85, 0x69, 0x10, 0xb2, 0xe7, 0xe2, 0x2a, 0x2d, 0x48, 0x49, 0x2c,
	0x49, 0x4d, 0x89, 0x4f, 0x2c, 0x91, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2, 0x83, 0x78,
	0x53, 0x0f, 0xe6, 0x4d, 0xbd, 0x10, 0x98, 0x37, 0x9d, 0x58, 0x26, 0xdc, 0x97, 0x67, 0x0c, 0xe2,
	0x84, 0xea, 0x71, 0x2c, 0x51, 0x52, 0xe4, 0x62, 0x01, 0x99, 0x28, 0xc4, 0xc3, 0xc5, 0xe1, 0x1f,
	0xe4, 0xe9, 0xee, 0xe9, 0xe7, 0xe8, 0x23, 0xc0, 0x20, 0xc4, 0xcd, 0xc5, 0x1e, 0x1e, 0xe4, 0x18,
	0x10, 0xe0, 0xea, 0x22, 0xc0, 0xe8, 0xe4, 0x76, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c,
	0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0x52, 0x49, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x3a, 0x5d, 0x3f, 0xb5, 0x2c, 0x57, 0x17, 0xe2, 0x76,
	0xfd, 0xe4, 0x9c, 0x4c, 0xfd, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x7b, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x24, 0x72, 0xa9, 0x1e, 0xbf, 0x01, 0x00, 0x00,
}

func (this *Pair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pair)
	if !ok {
		that2, ok := that.(Pair)
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
	if this.Inaddr != that1.Inaddr {
		return false
	}
	if this.Outaddr != that1.Outaddr {
		return false
	}
	if this.Intype != that1.Intype {
		return false
	}
	if this.Outtype != that1.Outtype {
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
func (this *Pair) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&pb.Pair{")
	s = append(s, "Inaddr: "+fmt.Sprintf("%#v", this.Inaddr)+",\n")
	s = append(s, "Outaddr: "+fmt.Sprintf("%#v", this.Outaddr)+",\n")
	s = append(s, "Intype: "+fmt.Sprintf("%#v", this.Intype)+",\n")
	s = append(s, "Outtype: "+fmt.Sprintf("%#v", this.Outtype)+",\n")
	s = append(s, "UpdatedAt: "+fmt.Sprintf("%#v", this.UpdatedAt)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPair(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Pair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintPair(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x2a
	}
	if m.Outtype != 0 {
		i = encodeVarintPair(dAtA, i, uint64(m.Outtype))
		i--
		dAtA[i] = 0x20
	}
	if m.Intype != 0 {
		i = encodeVarintPair(dAtA, i, uint64(m.Intype))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Outaddr) > 0 {
		i -= len(m.Outaddr)
		copy(dAtA[i:], m.Outaddr)
		i = encodeVarintPair(dAtA, i, uint64(len(m.Outaddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Inaddr) > 0 {
		i -= len(m.Inaddr)
		copy(dAtA[i:], m.Inaddr)
		i = encodeVarintPair(dAtA, i, uint64(len(m.Inaddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Inaddr)
	if l > 0 {
		n += 1 + l + sovPair(uint64(l))
	}
	l = len(m.Outaddr)
	if l > 0 {
		n += 1 + l + sovPair(uint64(l))
	}
	if m.Intype != 0 {
		n += 1 + sovPair(uint64(m.Intype))
	}
	if m.Outtype != 0 {
		n += 1 + sovPair(uint64(m.Outtype))
	}
	if m.UpdatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.UpdatedAt)
		n += 1 + l + sovPair(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPair(x uint64) (n int) {
	return sovPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Pair) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Pair{`,
		`Inaddr:` + fmt.Sprintf("%v", this.Inaddr) + `,`,
		`Outaddr:` + fmt.Sprintf("%v", this.Outaddr) + `,`,
		`Intype:` + fmt.Sprintf("%v", this.Intype) + `,`,
		`Outtype:` + fmt.Sprintf("%v", this.Outtype) + `,`,
		`UpdatedAt:` + strings.Replace(fmt.Sprintf("%v", this.UpdatedAt), "Timestamp", "timestamppb.Timestamp", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPair(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Pair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPair
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
			return fmt.Errorf("proto: Pair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inaddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inaddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outaddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Outaddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Intype", wireType)
			}
			m.Intype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Intype |= Pair_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outtype", wireType)
			}
			m.Outtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Outtype |= Pair_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPair
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
				return ErrInvalidLengthPair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPair
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
			skippy, err := skipPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPair
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
func skipPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPair
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
					return 0, ErrIntOverflowPair
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
					return 0, ErrIntOverflowPair
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
				return 0, ErrInvalidLengthPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPair = fmt.Errorf("proto: unexpected end of group")
)