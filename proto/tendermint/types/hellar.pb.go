// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/types/hellar.proto

package types

import (
	bytes "bytes"
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

type VoteExtensionType int32

const (
	// Unsupported
	VoteExtensionType_DEFAULT VoteExtensionType = 0
	// Sign canonical form of vote extension and threshold-recover signatures.
	//
	// Deterministic vote extension - each validator in a quorum must provide the same vote extension data.
	VoteExtensionType_THRESHOLD_RECOVER VoteExtensionType = 1
	// Sign raw form of vote extension and threshold-recover signatures.
	//
	// Deterministic vote extension - each validator in a quorum must provide the same vote extension data.
	// Use with caution - it can have severe security consequences, like replay attacks.
	//
	// THRESHOLD_RECOVER_RAW alows overriding sign request ID with `sign_request_id` field
	// of ExtendVoteExtension.sign_request_id. If sign_request_id is provided, SHA256(sign_request_id) will be used as
	// a sign request ID.
	//
	// It also changes how threshold-recover signatures are generated. Instead of signing canonical form of
	// threshold-recover signatures, it signs SHA256 of raw form of the vote extension (`ExtendVoteExtension.extension`).
	VoteExtensionType_THRESHOLD_RECOVER_RAW VoteExtensionType = 2
)

var VoteExtensionType_name = map[int32]string{
	0: "DEFAULT",
	1: "THRESHOLD_RECOVER",
	2: "THRESHOLD_RECOVER_RAW",
}

var VoteExtensionType_value = map[string]int32{
	"DEFAULT":               0,
	"THRESHOLD_RECOVER":     1,
	"THRESHOLD_RECOVER_RAW": 2,
}

func (x VoteExtensionType) String() string {
	return proto.EnumName(VoteExtensionType_name, int32(x))
}

func (VoteExtensionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_098b09a14a95d15e, []int{0}
}

// CoreChainLock represents a core chain lock for synchronization between state data and core chain
type CoreChainLock struct {
	CoreBlockHeight uint32 `protobuf:"varint,1,opt,name=core_block_height,json=coreBlockHeight,proto3" json:"core_block_height,omitempty"`
	CoreBlockHash   []byte `protobuf:"bytes,2,opt,name=core_block_hash,json=coreBlockHash,proto3" json:"core_block_hash,omitempty"`
	Signature       []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *CoreChainLock) Reset()         { *m = CoreChainLock{} }
func (m *CoreChainLock) String() string { return proto.CompactTextString(m) }
func (*CoreChainLock) ProtoMessage()    {}
func (*CoreChainLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_098b09a14a95d15e, []int{0}
}
func (m *CoreChainLock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoreChainLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoreChainLock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoreChainLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreChainLock.Merge(m, src)
}
func (m *CoreChainLock) XXX_Size() int {
	return m.Size()
}
func (m *CoreChainLock) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreChainLock.DiscardUnknown(m)
}

var xxx_messageInfo_CoreChainLock proto.InternalMessageInfo

func (m *CoreChainLock) GetCoreBlockHeight() uint32 {
	if m != nil {
		return m.CoreBlockHeight
	}
	return 0
}

func (m *CoreChainLock) GetCoreBlockHash() []byte {
	if m != nil {
		return m.CoreBlockHash
	}
	return nil
}

func (m *CoreChainLock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type VoteExtension struct {
	Type      VoteExtensionType `protobuf:"varint,1,opt,name=type,proto3,enum=tendermint.types.VoteExtensionType" json:"type,omitempty"`
	Extension []byte            `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	Signature []byte            `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are valid to be assigned to XSignRequestId:
	//	*VoteExtension_SignRequestId
	XSignRequestId isVoteExtension_XSignRequestId `protobuf_oneof:"_sign_request_id"`
}

func (m *VoteExtension) Reset()         { *m = VoteExtension{} }
func (m *VoteExtension) String() string { return proto.CompactTextString(m) }
func (*VoteExtension) ProtoMessage()    {}
func (*VoteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_098b09a14a95d15e, []int{1}
}
func (m *VoteExtension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteExtension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteExtension.Merge(m, src)
}
func (m *VoteExtension) XXX_Size() int {
	return m.Size()
}
func (m *VoteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_VoteExtension proto.InternalMessageInfo

type isVoteExtension_XSignRequestId interface {
	isVoteExtension_XSignRequestId()
	MarshalTo([]byte) (int, error)
	Size() int
}

type VoteExtension_SignRequestId struct {
	SignRequestId []byte `protobuf:"bytes,4,opt,name=sign_request_id,json=signRequestId,proto3,oneof" json:"sign_request_id,omitempty"`
}

func (*VoteExtension_SignRequestId) isVoteExtension_XSignRequestId() {}

func (m *VoteExtension) GetXSignRequestId() isVoteExtension_XSignRequestId {
	if m != nil {
		return m.XSignRequestId
	}
	return nil
}

func (m *VoteExtension) GetType() VoteExtensionType {
	if m != nil {
		return m.Type
	}
	return VoteExtensionType_DEFAULT
}

func (m *VoteExtension) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *VoteExtension) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *VoteExtension) GetSignRequestId() []byte {
	if x, ok := m.GetXSignRequestId().(*VoteExtension_SignRequestId); ok {
		return x.SignRequestId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VoteExtension) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VoteExtension_SignRequestId)(nil),
	}
}

func init() {
	proto.RegisterEnum("tendermint.types.VoteExtensionType", VoteExtensionType_name, VoteExtensionType_value)
	proto.RegisterType((*CoreChainLock)(nil), "tendermint.types.CoreChainLock")
	proto.RegisterType((*VoteExtension)(nil), "tendermint.types.VoteExtension")
}

func init() { proto.RegisterFile("tendermint/types/hellar.proto", fileDescriptor_098b09a14a95d15e) }

var fileDescriptor_098b09a14a95d15e = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xe3, 0x40,
	0x1c, 0xc7, 0x33, 0xdd, 0xb2, 0xcb, 0xce, 0x6e, 0xb6, 0xe9, 0xb0, 0x85, 0xf8, 0x87, 0x58, 0x2a,
	0x48, 0xa9, 0x90, 0x80, 0x0a, 0x82, 0xb7, 0xfe, 0x89, 0x44, 0x28, 0x16, 0xc6, 0x5a, 0xc1, 0x4b,
	0x48, 0x93, 0x21, 0x09, 0xb5, 0x99, 0x98, 0x4c, 0xc1, 0x3e, 0x81, 0x1e, 0x7d, 0x04, 0x5f, 0x46,
	0xf0, 0xd8, 0xa3, 0x47, 0x69, 0x2f, 0x3e, 0x86, 0xcc, 0x44, 0x89, 0xb6, 0xe0, 0x2d, 0xf3, 0xf9,
	0x7e, 0xc2, 0xef, 0x9b, 0xfc, 0x06, 0x6e, 0x30, 0x12, 0x79, 0x24, 0x19, 0x87, 0x11, 0x33, 0xd8,
	0x34, 0x26, 0xa9, 0xe1, 0x39, 0x69, 0xa0, 0xc7, 0x09, 0x65, 0x14, 0x29, 0x79, 0xa8, 0x8b, 0x70,
	0xfd, 0xbf, 0x4f, 0x7d, 0x2a, 0x42, 0x83, 0x3f, 0x65, 0x5e, 0xed, 0x16, 0x40, 0xb9, 0x4d, 0x13,
	0xd2, 0x0e, 0x9c, 0x30, 0xea, 0x52, 0x77, 0x84, 0x1a, 0xb0, 0xec, 0xd2, 0x84, 0xd8, 0xc3, 0x2b,
	0xea, 0x8e, 0xec, 0x80, 0x84, 0x7e, 0xc0, 0x54, 0x50, 0x05, 0x75, 0x19, 0x97, 0x78, 0xd0, 0xe2,
	0xdc, 0x12, 0x18, 0xed, 0xc0, 0xd2, 0x67, 0xd7, 0x49, 0x03, 0xb5, 0x50, 0x05, 0xf5, 0xbf, 0x58,
	0xce, 0x4d, 0x27, 0x0d, 0xd0, 0x26, 0xfc, 0x9d, 0x86, 0x7e, 0xe4, 0xb0, 0x49, 0x42, 0xd4, 0x1f,
	0xc2, 0xc8, 0xc1, 0x51, 0xf1, 0xf5, 0x61, 0x0b, 0xd4, 0x1e, 0x01, 0x94, 0x07, 0x94, 0x11, 0xf3,
	0x86, 0x91, 0x28, 0x0d, 0x69, 0x84, 0x0e, 0x61, 0x91, 0x57, 0x17, 0xc3, 0xff, 0xed, 0x6d, 0xeb,
	0xcb, 0x9f, 0xa4, 0x7f, 0xd1, 0xfb, 0xd3, 0x98, 0x60, 0xf1, 0x02, 0x1f, 0x47, 0x3e, 0xf0, 0x7b,
	0xa1, 0x1c, 0x7c, 0x5f, 0x06, 0xed, 0xc2, 0x12, 0x3f, 0xd8, 0x09, 0xb9, 0x9e, 0x90, 0x94, 0xd9,
	0xa1, 0xa7, 0x16, 0xb9, 0x63, 0x49, 0x58, 0xe6, 0x01, 0xce, 0xf8, 0x89, 0x77, 0x07, 0x40, 0x0b,
	0x41, 0xc5, 0x5e, 0xb2, 0x1b, 0x18, 0x96, 0x57, 0x7a, 0xa1, 0x3f, 0xf0, 0x57, 0xc7, 0x3c, 0x6e,
	0x9e, 0x77, 0xfb, 0x8a, 0x84, 0x2a, 0xb0, 0xdc, 0xb7, 0xb0, 0x79, 0x66, 0xf5, 0xba, 0x1d, 0x1b,
	0x9b, 0xed, 0xde, 0xc0, 0xc4, 0x0a, 0x40, 0x6b, 0xb0, 0xb2, 0x82, 0x6d, 0xdc, 0xbc, 0x50, 0x0a,
	0xad, 0xd3, 0xa7, 0xb9, 0x06, 0x66, 0x73, 0x0d, 0xbc, 0xcc, 0x35, 0x70, 0xbf, 0xd0, 0xa4, 0xd9,
	0x42, 0x93, 0x9e, 0x17, 0x9a, 0x74, 0x79, 0xe0, 0x87, 0x2c, 0x98, 0x0c, 0x75, 0x97, 0x8e, 0xc5,
	0xfa, 0x63, 0x67, 0x6a, 0x64, 0xff, 0x89, 0x9f, 0x8c, 0x6c, 0xdf, 0xcb, 0x17, 0x65, 0xf8, 0x53,
	0xf0, 0xfd, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x70, 0x1d, 0xe4, 0x43, 0x02, 0x00, 0x00,
}

func (this *CoreChainLock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreChainLock)
	if !ok {
		that2, ok := that.(CoreChainLock)
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
	if this.CoreBlockHeight != that1.CoreBlockHeight {
		return false
	}
	if !bytes.Equal(this.CoreBlockHash, that1.CoreBlockHash) {
		return false
	}
	if !bytes.Equal(this.Signature, that1.Signature) {
		return false
	}
	return true
}
func (m *CoreChainLock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoreChainLock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoreChainLock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintHellar(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CoreBlockHash) > 0 {
		i -= len(m.CoreBlockHash)
		copy(dAtA[i:], m.CoreBlockHash)
		i = encodeVarintHellar(dAtA, i, uint64(len(m.CoreBlockHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.CoreBlockHeight != 0 {
		i = encodeVarintHellar(dAtA, i, uint64(m.CoreBlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VoteExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteExtension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteExtension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XSignRequestId != nil {
		{
			size := m.XSignRequestId.Size()
			i -= size
			if _, err := m.XSignRequestId.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintHellar(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Extension) > 0 {
		i -= len(m.Extension)
		copy(dAtA[i:], m.Extension)
		i = encodeVarintHellar(dAtA, i, uint64(len(m.Extension)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintHellar(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VoteExtension_SignRequestId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteExtension_SignRequestId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignRequestId != nil {
		i -= len(m.SignRequestId)
		copy(dAtA[i:], m.SignRequestId)
		i = encodeVarintHellar(dAtA, i, uint64(len(m.SignRequestId)))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func encodeVarintHellar(dAtA []byte, offset int, v uint64) int {
	offset -= sovHellar(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CoreChainLock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoreBlockHeight != 0 {
		n += 1 + sovHellar(uint64(m.CoreBlockHeight))
	}
	l = len(m.CoreBlockHash)
	if l > 0 {
		n += 1 + l + sovHellar(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovHellar(uint64(l))
	}
	return n
}

func (m *VoteExtension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovHellar(uint64(m.Type))
	}
	l = len(m.Extension)
	if l > 0 {
		n += 1 + l + sovHellar(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovHellar(uint64(l))
	}
	if m.XSignRequestId != nil {
		n += m.XSignRequestId.Size()
	}
	return n
}

func (m *VoteExtension_SignRequestId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignRequestId != nil {
		l = len(m.SignRequestId)
		n += 1 + l + sovHellar(uint64(l))
	}
	return n
}

func sovHellar(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHellar(x uint64) (n int) {
	return sovHellar(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoreChainLock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellar
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
			return fmt.Errorf("proto: CoreChainLock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoreChainLock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreBlockHeight", wireType)
			}
			m.CoreBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoreBlockHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellar
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
				return ErrInvalidLengthHellar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHellar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoreBlockHash = append(m.CoreBlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CoreBlockHash == nil {
				m.CoreBlockHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellar
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
				return ErrInvalidLengthHellar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHellar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHellar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHellar
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
func (m *VoteExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHellar
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
			return fmt.Errorf("proto: VoteExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= VoteExtensionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extension", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellar
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
				return ErrInvalidLengthHellar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHellar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extension = append(m.Extension[:0], dAtA[iNdEx:postIndex]...)
			if m.Extension == nil {
				m.Extension = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellar
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
				return ErrInvalidLengthHellar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHellar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignRequestId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHellar
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
				return ErrInvalidLengthHellar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHellar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.XSignRequestId = &VoteExtension_SignRequestId{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHellar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHellar
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
func skipHellar(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHellar
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
					return 0, ErrIntOverflowHellar
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
					return 0, ErrIntOverflowHellar
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
				return 0, ErrInvalidLengthHellar
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHellar
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHellar
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHellar        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHellar          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHellar = fmt.Errorf("proto: unexpected end of group")
)
