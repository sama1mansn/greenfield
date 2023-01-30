// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/sp/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	// max_storage_providers is the maximum number of validators.
	MaxStorageProviders uint32 `protobuf:"varint,1,opt,name=max_storage_providers,json=maxStorageProviders,proto3" json:"max_storage_providers,omitempty"`
	// deposit_denom defines the staking coin denomination.
	DepositDenom string `protobuf:"bytes,2,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	// min_deposit_amount defines the minimum deposit amount for storage providers.
	MinDeposit github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=min_deposit,json=minDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_deposit"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5353d8e6e407d7e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxStorageProviders() uint32 {
	if m != nil {
		return m.MaxStorageProviders
	}
	return 0
}

func (m *Params) GetDepositDenom() string {
	if m != nil {
		return m.DepositDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "bnbchain.greenfield.sp.Params")
}

func init() { proto.RegisterFile("greenfield/sp/params.proto", fileDescriptor_a5353d8e6e407d7e) }

var fileDescriptor_a5353d8e6e407d7e = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2f, 0x4a, 0x4d,
	0xcd, 0x4b, 0xcb, 0x4c, 0xcd, 0x49, 0xd1, 0x2f, 0x2e, 0xd0, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4b, 0xca, 0x4b, 0x4a, 0xce, 0x48, 0xcc, 0xcc,
	0xd3, 0x43, 0x28, 0xd2, 0x2b, 0x2e, 0x90, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07,
	0xab, 0xd2, 0x87, 0x70, 0x20, 0x5a, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0x21, 0xe2, 0x20, 0x16,
	0x44, 0x54, 0xe9, 0x14, 0x23, 0x17, 0x5b, 0x00, 0xd8, 0x64, 0x21, 0x23, 0x2e, 0xd1, 0xdc, 0xc4,
	0x8a, 0xf8, 0xe2, 0x92, 0xfc, 0xa2, 0xc4, 0xf4, 0x54, 0x90, 0x11, 0x65, 0x99, 0x29, 0xa9, 0x45,
	0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0xc2, 0xb9, 0x89, 0x15, 0xc1, 0x10, 0xb9, 0x00,
	0x98, 0x94, 0x90, 0x32, 0x17, 0x6f, 0x4a, 0x6a, 0x41, 0x7e, 0x71, 0x66, 0x49, 0x7c, 0x4a, 0x6a,
	0x5e, 0x7e, 0xae, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x0f, 0x54, 0xd0, 0x05, 0x24, 0x26,
	0x14, 0xcb, 0xc5, 0x9d, 0x9b, 0x99, 0x17, 0x0f, 0x15, 0x93, 0x60, 0x06, 0x29, 0x71, 0xb2, 0x39,
	0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0xb5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0xa8, 0x7b, 0xa1, 0x94, 0x6e, 0x71, 0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0xb1,
	0x9e, 0x67, 0x5e, 0xc9, 0xa5, 0x2d, 0xba, 0x5c, 0x50, 0xef, 0x78, 0xe6, 0x95, 0x04, 0x71, 0xe5,
	0x66, 0xe6, 0xb9, 0x40, 0xcc, 0xb3, 0xe2, 0x98, 0xb1, 0x40, 0x9e, 0xe1, 0xc5, 0x02, 0x79, 0x46,
	0x27, 0x97, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x42, 0xb2, 0x25,
	0x29, 0x2f, 0x49, 0x17, 0x1c, 0x76, 0xfa, 0x48, 0x01, 0x5c, 0x01, 0x0a, 0x62, 0xb0, 0x6d, 0x49,
	0x6c, 0xe0, 0x90, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xaf, 0x8a, 0xca, 0x80, 0x01,
	0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxStorageProviders != that1.MaxStorageProviders {
		return false
	}
	if this.DepositDenom != that1.DepositDenom {
		return false
	}
	if !this.MinDeposit.Equal(that1.MinDeposit) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinDeposit.Size()
		i -= size
		if _, err := m.MinDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.DepositDenom) > 0 {
		i -= len(m.DepositDenom)
		copy(dAtA[i:], m.DepositDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DepositDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.MaxStorageProviders != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxStorageProviders))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxStorageProviders != 0 {
		n += 1 + sovParams(uint64(m.MaxStorageProviders))
	}
	l = len(m.DepositDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinDeposit.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStorageProviders", wireType)
			}
			m.MaxStorageProviders = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxStorageProviders |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
