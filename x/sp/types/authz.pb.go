// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/sp/authz.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

// DepositAuthorization defines authorization for sp deposit.
type DepositAuthorization struct {
	// max_deposit specifies the maximum amount of tokens can be deposit to a storage provider. If it is
	// empty, there is no spend limit and any amount of coins can be deposit.
	MaxDeposit *types.Coin `protobuf:"bytes,1,opt,name=max_deposit,json=maxDeposit,proto3" json:"max_deposit,omitempty"`
	// sp_address is the operator address of storage provider.
	SpAddress string `protobuf:"bytes,2,opt,name=sp_address,json=spAddress,proto3" json:"sp_address,omitempty"`
}

func (m *DepositAuthorization) Reset()         { *m = DepositAuthorization{} }
func (m *DepositAuthorization) String() string { return proto.CompactTextString(m) }
func (*DepositAuthorization) ProtoMessage()    {}
func (*DepositAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_eab7bba09f125b26, []int{0}
}
func (m *DepositAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositAuthorization.Merge(m, src)
}
func (m *DepositAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *DepositAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_DepositAuthorization proto.InternalMessageInfo

func (m *DepositAuthorization) GetMaxDeposit() *types.Coin {
	if m != nil {
		return m.MaxDeposit
	}
	return nil
}

func (m *DepositAuthorization) GetSpAddress() string {
	if m != nil {
		return m.SpAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*DepositAuthorization)(nil), "bnbchain.greenfield.sp.DepositAuthorization")
}

func init() { proto.RegisterFile("greenfield/sp/authz.proto", fileDescriptor_eab7bba09f125b26) }

var fileDescriptor_eab7bba09f125b26 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x6b, 0x06, 0xa4, 0xba, 0x5b, 0x15, 0xa1, 0xa4, 0x83, 0x55, 0x31, 0x55, 0x48, 0xb5,
	0x55, 0x18, 0x90, 0xd8, 0x5a, 0x7a, 0x82, 0xb2, 0xb1, 0x44, 0x76, 0x62, 0x12, 0x4b, 0xc4, 0xcf,
	0x8a, 0x1d, 0x14, 0x7a, 0x05, 0x16, 0x0e, 0xc3, 0x21, 0x18, 0x2b, 0x26, 0x46, 0x94, 0x5c, 0x04,
	0x25, 0xb6, 0x04, 0x9b, 0xdf, 0xff, 0xf9, 0xfd, 0xbf, 0xde, 0x8f, 0x93, 0xa2, 0x96, 0x52, 0x3f,
	0x29, 0xf9, 0x9c, 0x33, 0x6b, 0x18, 0x6f, 0x5c, 0x79, 0xa4, 0xa6, 0x06, 0x07, 0xf3, 0x0b, 0xa1,
	0x45, 0x56, 0x72, 0xa5, 0xe9, 0xdf, 0x1f, 0x6a, 0xcd, 0x82, 0x64, 0x60, 0x2b, 0xb0, 0x4c, 0x70,
	0x2b, 0xd9, 0xcb, 0x46, 0x48, 0xc7, 0x37, 0x2c, 0x03, 0xa5, 0xfd, 0xde, 0x22, 0xf1, 0x3c, 0x1d,
	0x27, 0xe6, 0x87, 0x80, 0xa2, 0x02, 0x0a, 0xf0, 0xfa, 0xf0, 0xf2, 0xea, 0xe5, 0x1b, 0xc2, 0xd1,
	0x5e, 0x1a, 0xb0, 0xca, 0x6d, 0x1b, 0x57, 0x42, 0xad, 0x8e, 0xdc, 0x29, 0xd0, 0xf3, 0x3b, 0x3c,
	0xab, 0x78, 0x9b, 0xe6, 0x9e, 0xc5, 0x68, 0x89, 0x56, 0xb3, 0xeb, 0x84, 0x06, 0xcb, 0x21, 0x9f,
	0x86, 0x7c, 0x7a, 0x0f, 0x4a, 0x1f, 0x70, 0xc5, 0xdb, 0x60, 0x34, 0xbf, 0xc5, 0xd8, 0x9a, 0x94,
	0xe7, 0x79, 0x2d, 0xad, 0x8d, 0xcf, 0x96, 0x68, 0x35, 0xdd, 0xc5, 0x5f, 0x1f, 0xeb, 0x28, 0x6c,
	0x6f, 0x3d, 0x79, 0x70, 0xb5, 0xd2, 0xc5, 0x61, 0x6a, 0x4d, 0x10, 0x76, 0xfb, 0xcf, 0x8e, 0xa0,
	0x53, 0x47, 0xd0, 0x4f, 0x47, 0xd0, 0x7b, 0x4f, 0x26, 0xa7, 0x9e, 0x4c, 0xbe, 0x7b, 0x32, 0x79,
	0xbc, 0x2a, 0x94, 0x2b, 0x1b, 0x41, 0x33, 0xa8, 0x98, 0xd0, 0x62, 0x3d, 0x96, 0xc3, 0xfe, 0x15,
	0xd8, 0x0e, 0x15, 0xba, 0x57, 0x23, 0xad, 0x38, 0x1f, 0x4f, 0xbb, 0xf9, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x74, 0x68, 0x28, 0x27, 0x60, 0x01, 0x00, 0x00,
}

func (m *DepositAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpAddress) > 0 {
		i -= len(m.SpAddress)
		copy(dAtA[i:], m.SpAddress)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.SpAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.MaxDeposit != nil {
		{
			size, err := m.MaxDeposit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxDeposit != nil {
		l = m.MaxDeposit.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = len(m.SpAddress)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: DepositAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxDeposit == nil {
				m.MaxDeposit = &types.Coin{}
			}
			if err := m.MaxDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
