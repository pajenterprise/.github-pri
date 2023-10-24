// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: HttpTransactionProto.proto

package forwarder

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TransactionPriorityProto int32

const (
	TransactionPriorityProto_NORMAL TransactionPriorityProto = 0
	TransactionPriorityProto_HIGH   TransactionPriorityProto = 1
)

var TransactionPriorityProto_name = map[int32]string{
	0: "NORMAL",
	1: "HIGH",
}

var TransactionPriorityProto_value = map[string]int32{
	"NORMAL": 0,
	"HIGH":   1,
}

func (x TransactionPriorityProto) String() string {
	return proto.EnumName(TransactionPriorityProto_name, int32(x))
}

func (TransactionPriorityProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6b31e85fc4b6d913, []int{0}
}

type HeaderValuesProto struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderValuesProto) Reset()         { *m = HeaderValuesProto{} }
func (m *HeaderValuesProto) String() string { return proto.CompactTextString(m) }
func (*HeaderValuesProto) ProtoMessage()    {}
func (*HeaderValuesProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b31e85fc4b6d913, []int{0}
}
func (m *HeaderValuesProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeaderValuesProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeaderValuesProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeaderValuesProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValuesProto.Merge(m, src)
}
func (m *HeaderValuesProto) XXX_Size() int {
	return m.Size()
}
func (m *HeaderValuesProto) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValuesProto.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValuesProto proto.InternalMessageInfo

func (m *HeaderValuesProto) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type EndpointProto struct {
	Route                string   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointProto) Reset()         { *m = EndpointProto{} }
func (m *EndpointProto) String() string { return proto.CompactTextString(m) }
func (*EndpointProto) ProtoMessage()    {}
func (*EndpointProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b31e85fc4b6d913, []int{1}
}
func (m *EndpointProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EndpointProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EndpointProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndpointProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointProto.Merge(m, src)
}
func (m *EndpointProto) XXX_Size() int {
	return m.Size()
}
func (m *EndpointProto) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointProto.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointProto proto.InternalMessageInfo

func (m *EndpointProto) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *EndpointProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HttpTransactionProto struct {
	Domain               string                        `protobuf:"bytes,1,opt,name=Domain,proto3" json:"Domain,omitempty"`
	Endpoint             *EndpointProto                `protobuf:"bytes,2,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	Headers              map[string]*HeaderValuesProto `protobuf:"bytes,3,rep,name=Headers,proto3" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              []byte                        `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	ErrorCount           int64                         `protobuf:"varint,5,opt,name=ErrorCount,proto3" json:"ErrorCount,omitempty"`
	CreatedAt            int64                         `protobuf:"varint,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Retryable            bool                          `protobuf:"varint,7,opt,name=Retryable,proto3" json:"Retryable,omitempty"`
	Priority             TransactionPriorityProto      `protobuf:"varint,8,opt,name=priority,proto3,enum=forwarder.TransactionPriorityProto" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HttpTransactionProto) Reset()         { *m = HttpTransactionProto{} }
func (m *HttpTransactionProto) String() string { return proto.CompactTextString(m) }
func (*HttpTransactionProto) ProtoMessage()    {}
func (*HttpTransactionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b31e85fc4b6d913, []int{2}
}
func (m *HttpTransactionProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpTransactionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpTransactionProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpTransactionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpTransactionProto.Merge(m, src)
}
func (m *HttpTransactionProto) XXX_Size() int {
	return m.Size()
}
func (m *HttpTransactionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpTransactionProto.DiscardUnknown(m)
}

var xxx_messageInfo_HttpTransactionProto proto.InternalMessageInfo

func (m *HttpTransactionProto) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *HttpTransactionProto) GetEndpoint() *EndpointProto {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *HttpTransactionProto) GetHeaders() map[string]*HeaderValuesProto {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpTransactionProto) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HttpTransactionProto) GetErrorCount() int64 {
	if m != nil {
		return m.ErrorCount
	}
	return 0
}

func (m *HttpTransactionProto) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *HttpTransactionProto) GetRetryable() bool {
	if m != nil {
		return m.Retryable
	}
	return false
}

func (m *HttpTransactionProto) GetPriority() TransactionPriorityProto {
	if m != nil {
		return m.Priority
	}
	return TransactionPriorityProto_NORMAL
}

type HttpTransactionProtoCollection struct {
	Version              int32                   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Values               []*HttpTransactionProto `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HttpTransactionProtoCollection) Reset()         { *m = HttpTransactionProtoCollection{} }
func (m *HttpTransactionProtoCollection) String() string { return proto.CompactTextString(m) }
func (*HttpTransactionProtoCollection) ProtoMessage()    {}
func (*HttpTransactionProtoCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b31e85fc4b6d913, []int{3}
}
func (m *HttpTransactionProtoCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpTransactionProtoCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpTransactionProtoCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpTransactionProtoCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpTransactionProtoCollection.Merge(m, src)
}
func (m *HttpTransactionProtoCollection) XXX_Size() int {
	return m.Size()
}
func (m *HttpTransactionProtoCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpTransactionProtoCollection.DiscardUnknown(m)
}

var xxx_messageInfo_HttpTransactionProtoCollection proto.InternalMessageInfo

func (m *HttpTransactionProtoCollection) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HttpTransactionProtoCollection) GetValues() []*HttpTransactionProto {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("forwarder.TransactionPriorityProto", TransactionPriorityProto_name, TransactionPriorityProto_value)
	proto.RegisterType((*HeaderValuesProto)(nil), "forwarder.HeaderValuesProto")
	proto.RegisterType((*EndpointProto)(nil), "forwarder.EndpointProto")
	proto.RegisterType((*HttpTransactionProto)(nil), "forwarder.HttpTransactionProto")
	proto.RegisterMapType((map[string]*HeaderValuesProto)(nil), "forwarder.HttpTransactionProto.HeadersEntry")
	proto.RegisterType((*HttpTransactionProtoCollection)(nil), "forwarder.HttpTransactionProtoCollection")
}

func init() { proto.RegisterFile("HttpTransactionProto.proto", fileDescriptor_6b31e85fc4b6d913) }

var fileDescriptor_6b31e85fc4b6d913 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdb, 0x8a, 0xd4, 0x40,
	0x10, 0xdd, 0x9e, 0xcc, 0x25, 0xa9, 0x5d, 0x25, 0x36, 0x8b, 0x34, 0xcb, 0x12, 0x43, 0x7c, 0x09,
	0x2a, 0x83, 0x44, 0xc1, 0xcb, 0x8b, 0xac, 0xe3, 0x68, 0x04, 0x2f, 0x4b, 0x23, 0xe2, 0x6b, 0xaf,
	0x69, 0x21, 0x98, 0xed, 0x0e, 0x95, 0x9e, 0x95, 0xfc, 0x89, 0x3f, 0xe3, 0xbb, 0x8f, 0x7e, 0x82,
	0x8c, 0x3f, 0x22, 0xe9, 0x24, 0xb3, 0x19, 0x19, 0xd9, 0x97, 0xa6, 0x4f, 0xf5, 0xa9, 0xae, 0x53,
	0xa7, 0x0a, 0x8e, 0x52, 0x63, 0xca, 0x0f, 0x28, 0x54, 0x25, 0x3e, 0x9b, 0x5c, 0xab, 0x53, 0xd4,
	0x46, 0xcf, 0xcb, 0xe6, 0xa4, 0xde, 0x17, 0x8d, 0xdf, 0x04, 0x66, 0x12, 0xa3, 0xbb, 0x70, 0x23,
	0x95, 0x22, 0x93, 0xf8, 0x51, 0x14, 0x2b, 0x59, 0x59, 0x16, 0xbd, 0x09, 0xd3, 0x0b, 0x0b, 0x19,
	0x09, 0x9d, 0xd8, 0xe3, 0x1d, 0x8a, 0x9e, 0xc0, 0xb5, 0xa5, 0xca, 0x4a, 0x9d, 0x2b, 0xd3, 0x12,
	0x0f, 0x61, 0x82, 0x7a, 0x65, 0x24, 0x23, 0x21, 0x89, 0x3d, 0xde, 0x02, 0x4a, 0x61, 0xac, 0xc4,
	0xb9, 0x64, 0x23, 0x1b, 0xb4, 0xf7, 0xe8, 0x87, 0x03, 0x87, 0xbb, 0x14, 0x35, 0xb5, 0x5e, 0xe8,
	0x73, 0x91, 0xab, 0xee, 0x8f, 0x0e, 0xd1, 0x87, 0xe0, 0xf6, 0xb5, 0xec, 0x47, 0xfb, 0x09, 0x9b,
	0x6f, 0x64, 0xcf, 0xb7, 0x64, 0xf0, 0x0d, 0x93, 0xbe, 0x84, 0x59, 0xdb, 0x4e, 0xc5, 0x9c, 0xd0,
	0x89, 0xf7, 0x93, 0x7b, 0x83, 0xa4, 0x9d, 0x8e, 0x74, 0xf4, 0xa5, 0x32, 0x58, 0xf3, 0x3e, 0x99,
	0x32, 0x98, 0x9d, 0x8a, 0xba, 0xd0, 0x22, 0x63, 0xe3, 0x90, 0xc4, 0x07, 0xbc, 0x87, 0x34, 0x00,
	0x58, 0x22, 0x6a, 0x5c, 0xe8, 0x95, 0x32, 0x6c, 0x12, 0x92, 0xd8, 0xe1, 0x83, 0x08, 0x3d, 0x06,
	0x6f, 0x81, 0x52, 0x18, 0x99, 0x9d, 0x18, 0x36, 0xb5, 0xcf, 0x97, 0x81, 0xe6, 0x95, 0x4b, 0x83,
	0xb5, 0x38, 0x2b, 0x24, 0x9b, 0x85, 0x24, 0x76, 0xf9, 0x65, 0x80, 0x3e, 0x03, 0xb7, 0xc4, 0x5c,
	0x63, 0x6e, 0x6a, 0xe6, 0x86, 0x24, 0xbe, 0x9e, 0xdc, 0x1e, 0xc8, 0xdf, 0x92, 0xde, 0xb2, 0xba,
	0xf6, 0xfb, 0xa4, 0xa3, 0x4f, 0x70, 0x30, 0xec, 0x87, 0xfa, 0xe0, 0x7c, 0x95, 0x75, 0xe7, 0x6c,
	0x73, 0xa5, 0x09, 0x4c, 0xec, 0x30, 0x3b, 0x4f, 0x8f, 0x87, 0xf6, 0xfc, 0xbb, 0x07, 0xbc, 0xa5,
	0x3e, 0x1d, 0x3d, 0x26, 0x51, 0x05, 0xc1, 0x2e, 0xfb, 0x16, 0xba, 0x28, 0xa4, 0x85, 0x8d, 0x65,
	0x17, 0x12, 0xab, 0x5c, 0xb7, 0x93, 0x9c, 0xf0, 0x1e, 0xd2, 0x47, 0x9b, 0x75, 0x1a, 0xd9, 0x99,
	0xdc, 0xba, 0x62, 0x26, 0xfd, 0xbe, 0xdd, 0xb9, 0x0f, 0xec, 0x7f, 0x4d, 0x53, 0x80, 0xe9, 0xbb,
	0xf7, 0xfc, 0xed, 0xc9, 0x1b, 0x7f, 0x8f, 0xba, 0x30, 0x4e, 0x5f, 0xbf, 0x4a, 0x7d, 0xf2, 0xdc,
	0xff, 0xb9, 0x0e, 0xc8, 0xaf, 0x75, 0x40, 0x7e, 0xaf, 0x03, 0xf2, 0xfd, 0x4f, 0xb0, 0x77, 0x36,
	0xb5, 0x2b, 0xff, 0xe0, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x16, 0xcd, 0xc7, 0x10, 0x03,
	0x00, 0x00,
}

func (m *HeaderValuesProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeaderValuesProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeaderValuesProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarintHttpTransactionProto(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EndpointProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndpointProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EndpointProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Route) > 0 {
		i -= len(m.Route)
		copy(dAtA[i:], m.Route)
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(len(m.Route)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HttpTransactionProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpTransactionProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpTransactionProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Priority != 0 {
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x40
	}
	if m.Retryable {
		i--
		if m.Retryable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CreatedAt != 0 {
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x30
	}
	if m.ErrorCount != 0 {
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(m.ErrorCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Headers) > 0 {
		for k := range m.Headers {
			v := m.Headers[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintHttpTransactionProto(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintHttpTransactionProto(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintHttpTransactionProto(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Endpoint != nil {
		{
			size, err := m.Endpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttpTransactionProto(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HttpTransactionProtoCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpTransactionProtoCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpTransactionProtoCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Values[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHttpTransactionProto(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Version != 0 {
		i = encodeVarintHttpTransactionProto(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHttpTransactionProto(dAtA []byte, offset int, v uint64) int {
	offset -= sovHttpTransactionProto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeaderValuesProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovHttpTransactionProto(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EndpointProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Route)
	if l > 0 {
		n += 1 + l + sovHttpTransactionProto(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHttpTransactionProto(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpTransactionProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovHttpTransactionProto(uint64(l))
	}
	if m.Endpoint != nil {
		l = m.Endpoint.Size()
		n += 1 + l + sovHttpTransactionProto(uint64(l))
	}
	if len(m.Headers) > 0 {
		for k, v := range m.Headers {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovHttpTransactionProto(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovHttpTransactionProto(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovHttpTransactionProto(uint64(mapEntrySize))
		}
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovHttpTransactionProto(uint64(l))
	}
	if m.ErrorCount != 0 {
		n += 1 + sovHttpTransactionProto(uint64(m.ErrorCount))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovHttpTransactionProto(uint64(m.CreatedAt))
	}
	if m.Retryable {
		n += 2
	}
	if m.Priority != 0 {
		n += 1 + sovHttpTransactionProto(uint64(m.Priority))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpTransactionProtoCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovHttpTransactionProto(uint64(m.Version))
	}
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovHttpTransactionProto(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHttpTransactionProto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHttpTransactionProto(x uint64) (n int) {
	return sovHttpTransactionProto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeaderValuesProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpTransactionProto
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
			return fmt.Errorf("proto: HeaderValuesProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderValuesProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpTransactionProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpTransactionProto
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
func (m *EndpointProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpTransactionProto
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
			return fmt.Errorf("proto: EndpointProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndpointProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Route = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpTransactionProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpTransactionProto
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
func (m *HttpTransactionProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpTransactionProto
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
			return fmt.Errorf("proto: HttpTransactionProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpTransactionProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Endpoint == nil {
				m.Endpoint = &EndpointProto{}
			}
			if err := m.Endpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Headers == nil {
				m.Headers = make(map[string]*HeaderValuesProto)
			}
			var mapkey string
			var mapvalue *HeaderValuesProto
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHttpTransactionProto
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHttpTransactionProto
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthHttpTransactionProto
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthHttpTransactionProto
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHttpTransactionProto
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthHttpTransactionProto
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthHttpTransactionProto
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &HeaderValuesProto{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipHttpTransactionProto(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthHttpTransactionProto
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Headers[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCount", wireType)
			}
			m.ErrorCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Retryable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
			m.Retryable = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= TransactionPriorityProto(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHttpTransactionProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpTransactionProto
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
func (m *HttpTransactionProtoCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpTransactionProto
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
			return fmt.Errorf("proto: HttpTransactionProtoCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpTransactionProtoCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpTransactionProto
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
				return ErrInvalidLengthHttpTransactionProto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &HttpTransactionProto{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpTransactionProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpTransactionProto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpTransactionProto
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
func skipHttpTransactionProto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpTransactionProto
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
					return 0, ErrIntOverflowHttpTransactionProto
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
					return 0, ErrIntOverflowHttpTransactionProto
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
				return 0, ErrInvalidLengthHttpTransactionProto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHttpTransactionProto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHttpTransactionProto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHttpTransactionProto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpTransactionProto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHttpTransactionProto = fmt.Errorf("proto: unexpected end of group")
)
