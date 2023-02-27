// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/storage/common.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type SourceType int32

const (
	SOURCE_TYPE_ORIGIN          SourceType = 0
	SOURCE_TYPE_BSC_CROSS_CHAIN SourceType = 1
)

var SourceType_name = map[int32]string{
	0: "SOURCE_TYPE_ORIGIN",
	1: "SOURCE_TYPE_BSC_CROSS_CHAIN",
}

var SourceType_value = map[string]int32{
	"SOURCE_TYPE_ORIGIN":          0,
	"SOURCE_TYPE_BSC_CROSS_CHAIN": 1,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}

func (SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{0}
}

// TODO: Need to confirm the read quota enum which determined by payment module
// TODO: Make this field be configured through governance
type ReadQuota int32

const (
	READ_QUOTA_FREE ReadQuota = 0
	READ_QUOTA_1G   ReadQuota = 1
	READ_QUOTA_10G  ReadQuota = 2
)

var ReadQuota_name = map[int32]string{
	0: "READ_QUOTA_FREE",
	1: "READ_QUOTA_1G",
	2: "READ_QUOTA_10G",
}

var ReadQuota_value = map[string]int32{
	"READ_QUOTA_FREE": 0,
	"READ_QUOTA_1G":   1,
	"READ_QUOTA_10G":  2,
}

func (x ReadQuota) String() string {
	return proto.EnumName(ReadQuota_name, int32(x))
}

func (ReadQuota) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{1}
}

type RedundancyType int32

const (
	REDUNDANCY_REPLICA_TYPE RedundancyType = 0
	REDUNDANCY_EC_TYPE      RedundancyType = 1
	REDUNDANCY_INLINE_TYPE  RedundancyType = 2
)

var RedundancyType_name = map[int32]string{
	0: "REDUNDANCY_REPLICA_TYPE",
	1: "REDUNDANCY_EC_TYPE",
	2: "REDUNDANCY_INLINE_TYPE",
}

var RedundancyType_value = map[string]int32{
	"REDUNDANCY_REPLICA_TYPE": 0,
	"REDUNDANCY_EC_TYPE":      1,
	"REDUNDANCY_INLINE_TYPE":  2,
}

func (x RedundancyType) String() string {
	return proto.EnumName(RedundancyType_name, int32(x))
}

func (RedundancyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{2}
}

type ObjectStatus int32

const (
	OBJECT_STATUS_INIT       ObjectStatus = 0
	OBJECT_STATUS_IN_SERVICE ObjectStatus = 1
)

var ObjectStatus_name = map[int32]string{
	0: "OBJECT_STATUS_INIT",
	1: "OBJECT_STATUS_IN_SERVICE",
}

var ObjectStatus_value = map[string]int32{
	"OBJECT_STATUS_INIT":       0,
	"OBJECT_STATUS_IN_SERVICE": 1,
}

func (x ObjectStatus) String() string {
	return proto.EnumName(ObjectStatus_name, int32(x))
}

func (ObjectStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{3}
}

type Approval struct {
	ExpiredHeight uint64 `protobuf:"varint,1,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
	Sig           []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *Approval) Reset()         { *m = Approval{} }
func (m *Approval) String() string { return proto.CompactTextString(m) }
func (*Approval) ProtoMessage()    {}
func (*Approval) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{0}
}
func (m *Approval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Approval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Approval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Approval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Approval.Merge(m, src)
}
func (m *Approval) XXX_Size() int {
	return m.Size()
}
func (m *Approval) XXX_DiscardUnknown() {
	xxx_messageInfo_Approval.DiscardUnknown(m)
}

var xxx_messageInfo_Approval proto.InternalMessageInfo

func (m *Approval) GetExpiredHeight() uint64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

func (m *Approval) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// SecondarySpSignDoc used to generate seal signature of secondary SP
type SecondarySpSignDoc struct {
	SpAddress string `protobuf:"bytes,1,opt,name=sp_address,json=spAddress,proto3" json:"sp_address,omitempty"`
	Checksum  []byte `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *SecondarySpSignDoc) Reset()         { *m = SecondarySpSignDoc{} }
func (m *SecondarySpSignDoc) String() string { return proto.CompactTextString(m) }
func (*SecondarySpSignDoc) ProtoMessage()    {}
func (*SecondarySpSignDoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{1}
}
func (m *SecondarySpSignDoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecondarySpSignDoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecondarySpSignDoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecondarySpSignDoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondarySpSignDoc.Merge(m, src)
}
func (m *SecondarySpSignDoc) XXX_Size() int {
	return m.Size()
}
func (m *SecondarySpSignDoc) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondarySpSignDoc.DiscardUnknown(m)
}

var xxx_messageInfo_SecondarySpSignDoc proto.InternalMessageInfo

func (m *SecondarySpSignDoc) GetSpAddress() string {
	if m != nil {
		return m.SpAddress
	}
	return ""
}

func (m *SecondarySpSignDoc) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func init() {
	proto.RegisterEnum("bnbchain.greenfield.storage.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("bnbchain.greenfield.storage.ReadQuota", ReadQuota_name, ReadQuota_value)
	proto.RegisterEnum("bnbchain.greenfield.storage.RedundancyType", RedundancyType_name, RedundancyType_value)
	proto.RegisterEnum("bnbchain.greenfield.storage.ObjectStatus", ObjectStatus_name, ObjectStatus_value)
	proto.RegisterType((*Approval)(nil), "bnbchain.greenfield.storage.Approval")
	proto.RegisterType((*SecondarySpSignDoc)(nil), "bnbchain.greenfield.storage.SecondarySpSignDoc")
}

func init() { proto.RegisterFile("greenfield/storage/common.proto", fileDescriptor_4eff6c0fa4aaf4c9) }

var fileDescriptor_4eff6c0fa4aaf4c9 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xed, 0x52, 0xa1, 0x76, 0xd4, 0x06, 0xb3, 0x54, 0x25, 0x24, 0xc8, 0xad, 0x2a, 0x21,
	0x55, 0x95, 0x1a, 0x83, 0x38, 0x70, 0x76, 0x1c, 0x93, 0xba, 0x14, 0xbb, 0xdd, 0x75, 0x90, 0xca,
	0x65, 0x65, 0xaf, 0x17, 0xc7, 0xd0, 0x78, 0x2d, 0xaf, 0x8d, 0x9a, 0x37, 0xe0, 0xc8, 0x3b, 0xf0,
	0x0a, 0x3c, 0x04, 0xc7, 0x8a, 0x13, 0x47, 0x94, 0xbc, 0x08, 0x8a, 0x6d, 0x82, 0xc5, 0x6d, 0xe6,
	0x9b, 0xf1, 0x3f, 0xbf, 0x57, 0x3f, 0x1c, 0xc4, 0x39, 0xe7, 0xe9, 0x87, 0x84, 0xdf, 0x44, 0x86,
	0x2c, 0x44, 0x1e, 0xc4, 0xdc, 0x60, 0x62, 0x36, 0x13, 0xe9, 0x20, 0xcb, 0x45, 0x21, 0x50, 0x3f,
	0x4c, 0x43, 0x36, 0x0d, 0x92, 0x74, 0xf0, 0x6f, 0x73, 0xd0, 0x6c, 0xf6, 0x9e, 0x30, 0x21, 0x67,
	0x42, 0xd2, 0x6a, 0xd5, 0xa8, 0x9b, 0xfa, 0xbb, 0xde, 0x5e, 0x2c, 0x62, 0x51, 0xf3, 0x55, 0x55,
	0xd3, 0x23, 0x0b, 0xb6, 0xcc, 0x2c, 0xcb, 0xc5, 0xe7, 0xe0, 0x06, 0x3d, 0x83, 0x0e, 0xbf, 0xcd,
	0x92, 0x9c, 0x47, 0x74, 0xca, 0x93, 0x78, 0x5a, 0x74, 0xd5, 0x43, 0xf5, 0x78, 0x13, 0xef, 0x36,
	0xf4, 0xac, 0x82, 0x48, 0x83, 0x7b, 0x32, 0x89, 0xbb, 0x1b, 0x87, 0xea, 0xf1, 0x0e, 0x5e, 0x95,
	0x47, 0x09, 0x20, 0xc2, 0x99, 0x48, 0xa3, 0x20, 0x9f, 0x93, 0x8c, 0x24, 0x71, 0x3a, 0x12, 0x0c,
	0xbd, 0x02, 0x90, 0x19, 0x0d, 0xa2, 0x28, 0xe7, 0x52, 0x56, 0x52, 0xdb, 0xc3, 0xee, 0xcf, 0xef,
	0xa7, 0x7b, 0x8d, 0x2d, 0xb3, 0x9e, 0x90, 0x22, 0x4f, 0xd2, 0x18, 0x6f, 0xcb, 0xac, 0x01, 0xa8,
	0x07, 0x5b, 0x6c, 0xca, 0xd9, 0x27, 0x59, 0xce, 0x9a, 0x2b, 0xeb, 0xfe, 0xe4, 0x0d, 0x00, 0x11,
	0x65, 0xce, 0xb8, 0x3f, 0xcf, 0x38, 0xda, 0x07, 0x44, 0xbc, 0x09, 0xb6, 0x6c, 0xea, 0x5f, 0x5f,
	0xda, 0xd4, 0xc3, 0xce, 0xd8, 0x71, 0x35, 0x05, 0x1d, 0x40, 0xbf, 0xcd, 0x87, 0xc4, 0xa2, 0x16,
	0xf6, 0x08, 0xa1, 0xd6, 0x99, 0xe9, 0xb8, 0x9a, 0xda, 0xdb, 0xfc, 0xf2, 0x4d, 0x57, 0x4e, 0xde,
	0xc2, 0x36, 0xe6, 0x41, 0x74, 0x55, 0x8a, 0x22, 0x40, 0x8f, 0xe0, 0x01, 0xb6, 0xcd, 0x11, 0xbd,
	0x9a, 0x78, 0xbe, 0x49, 0x5f, 0x63, 0xdb, 0xd6, 0x14, 0xf4, 0x10, 0x76, 0x5b, 0xf0, 0xc5, 0x58,
	0x53, 0x11, 0x82, 0x4e, 0x1b, 0x3d, 0x1f, 0x6b, 0x1b, 0x8d, 0x5c, 0x0c, 0x1d, 0xcc, 0xa3, 0x32,
	0x8d, 0x82, 0x94, 0xcd, 0x2b, 0x7f, 0x7d, 0x78, 0x8c, 0xed, 0xd1, 0xc4, 0x1d, 0x99, 0xae, 0x75,
	0x4d, 0xb1, 0x7d, 0x79, 0xe1, 0x58, 0x66, 0xe5, 0x49, 0x53, 0x56, 0xe6, 0x5b, 0x43, 0xdb, 0xaa,
	0xb9, 0x8a, 0x7a, 0xb0, 0xdf, 0xe2, 0x8e, 0x7b, 0xe1, 0xb8, 0xf5, 0x7f, 0xac, 0x0f, 0x9d, 0xc3,
	0x8e, 0x17, 0x7e, 0xe4, 0xac, 0x20, 0x45, 0x50, 0x94, 0x72, 0xa5, 0xe4, 0x0d, 0xcf, 0x6d, 0xcb,
	0xa7, 0xc4, 0x37, 0xfd, 0x09, 0xa1, 0x8e, 0xeb, 0xf8, 0x9a, 0x82, 0x9e, 0x42, 0xf7, 0x7f, 0x4e,
	0x89, 0x8d, 0xdf, 0x39, 0x96, 0xfd, 0xf7, 0x0d, 0x86, 0xce, 0x8f, 0x85, 0xae, 0xde, 0x2d, 0x74,
	0xf5, 0xf7, 0x42, 0x57, 0xbf, 0x2e, 0x75, 0xe5, 0x6e, 0xa9, 0x2b, 0xbf, 0x96, 0xba, 0xf2, 0xde,
	0x88, 0x93, 0x62, 0x5a, 0x86, 0x03, 0x26, 0x66, 0x46, 0x98, 0x86, 0xa7, 0x55, 0xe8, 0x8c, 0x56,
	0x3c, 0x6f, 0xd7, 0x01, 0x2d, 0xe6, 0x19, 0x97, 0xe1, 0xfd, 0x2a, 0x52, 0x2f, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xda, 0x13, 0xaa, 0xc5, 0xc3, 0x02, 0x00, 0x00,
}

func (m *Approval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Approval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Approval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x12
	}
	if m.ExpiredHeight != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SecondarySpSignDoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecondarySpSignDoc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecondarySpSignDoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpAddress) > 0 {
		i -= len(m.SpAddress)
		copy(dAtA[i:], m.SpAddress)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.SpAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Approval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpiredHeight != 0 {
		n += 1 + sovCommon(uint64(m.ExpiredHeight))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *SecondarySpSignDoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpAddress)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Approval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Approval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Approval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiredHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *SecondarySpSignDoc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: SecondarySpSignDoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecondarySpSignDoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
