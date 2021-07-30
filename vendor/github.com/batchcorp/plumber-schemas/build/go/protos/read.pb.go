// Code generated by protoc-gen-go. DO NOT EDIT.
// source: read.proto

package protos

import (
	fmt "fmt"
	args "github.com/batchcorp/plumber-schemas/build/go/protos/args"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	encoding "github.com/batchcorp/plumber-schemas/build/go/protos/encoding"
	records "github.com/batchcorp/plumber-schemas/build/go/protos/records"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ReadOptions_Type int32

const (
	ReadOptions_ONE_TIME   ReadOptions_Type = 0
	ReadOptions_CONTINUOUS ReadOptions_Type = 1
)

var ReadOptions_Type_name = map[int32]string{
	0: "ONE_TIME",
	1: "CONTINUOUS",
}

var ReadOptions_Type_value = map[string]int32{
	"ONE_TIME":   0,
	"CONTINUOUS": 1,
}

func (x ReadOptions_Type) String() string {
	return proto.EnumName(ReadOptions_Type_name, int32(x))
}

func (ReadOptions_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{0, 0}
}

type SampleOptions_Interval int32

const (
	SampleOptions_SECOND SampleOptions_Interval = 0
	SampleOptions_MINUTE SampleOptions_Interval = 1
)

var SampleOptions_Interval_name = map[int32]string{
	0: "SECOND",
	1: "MINUTE",
}

var SampleOptions_Interval_value = map[string]int32{
	"SECOND": 0,
	"MINUTE": 1,
}

func (x SampleOptions_Interval) String() string {
	return proto.EnumName(SampleOptions_Interval_name, int32(x))
}

func (SampleOptions_Interval) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{1, 0}
}

type ReadOptions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOptions) Reset()         { *m = ReadOptions{} }
func (m *ReadOptions) String() string { return proto.CompactTextString(m) }
func (*ReadOptions) ProtoMessage()    {}
func (*ReadOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{0}
}

func (m *ReadOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOptions.Unmarshal(m, b)
}
func (m *ReadOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOptions.Marshal(b, m, deterministic)
}
func (m *ReadOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOptions.Merge(m, src)
}
func (m *ReadOptions) XXX_Size() int {
	return xxx_messageInfo_ReadOptions.Size(m)
}
func (m *ReadOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOptions proto.InternalMessageInfo

type SampleOptions struct {
	SampleRate           uint32                 `protobuf:"varint,1,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	SampleInterval       SampleOptions_Interval `protobuf:"varint,2,opt,name=sample_interval,json=sampleInterval,proto3,enum=protos.SampleOptions_Interval" json:"sample_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SampleOptions) Reset()         { *m = SampleOptions{} }
func (m *SampleOptions) String() string { return proto.CompactTextString(m) }
func (*SampleOptions) ProtoMessage()    {}
func (*SampleOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{1}
}

func (m *SampleOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleOptions.Unmarshal(m, b)
}
func (m *SampleOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleOptions.Marshal(b, m, deterministic)
}
func (m *SampleOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleOptions.Merge(m, src)
}
func (m *SampleOptions) XXX_Size() int {
	return xxx_messageInfo_SampleOptions.Size(m)
}
func (m *SampleOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SampleOptions proto.InternalMessageInfo

func (m *SampleOptions) GetSampleRate() uint32 {
	if m != nil {
		return m.SampleRate
	}
	return 0
}

func (m *SampleOptions) GetSampleInterval() SampleOptions_Interval {
	if m != nil {
		return m.SampleInterval
	}
	return SampleOptions_SECOND
}

type Read struct {
	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectionId  string            `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ReadOptions   *ReadOptions      `protobuf:"bytes,3,opt,name=read_options,json=readOptions,proto3" json:"read_options,omitempty"`
	SampleOptions *SampleOptions    `protobuf:"bytes,4,opt,name=sample_options,json=sampleOptions,proto3" json:"sample_options,omitempty"`
	DecodeOptions *encoding.Options `protobuf:"bytes,5,opt,name=decode_options,json=decodeOptions,proto3" json:"decode_options,omitempty"`
	// Types that are valid to be assigned to Args:
	//	*Read_Kafka
	Args                 isRead_Args `protobuf_oneof:"Args"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Read) Reset()         { *m = Read{} }
func (m *Read) String() string { return proto.CompactTextString(m) }
func (*Read) ProtoMessage()    {}
func (*Read) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{2}
}

func (m *Read) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Read.Unmarshal(m, b)
}
func (m *Read) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Read.Marshal(b, m, deterministic)
}
func (m *Read) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Read.Merge(m, src)
}
func (m *Read) XXX_Size() int {
	return xxx_messageInfo_Read.Size(m)
}
func (m *Read) XXX_DiscardUnknown() {
	xxx_messageInfo_Read.DiscardUnknown(m)
}

var xxx_messageInfo_Read proto.InternalMessageInfo

func (m *Read) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Read) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Read) GetReadOptions() *ReadOptions {
	if m != nil {
		return m.ReadOptions
	}
	return nil
}

func (m *Read) GetSampleOptions() *SampleOptions {
	if m != nil {
		return m.SampleOptions
	}
	return nil
}

func (m *Read) GetDecodeOptions() *encoding.Options {
	if m != nil {
		return m.DecodeOptions
	}
	return nil
}

type isRead_Args interface {
	isRead_Args()
}

type Read_Kafka struct {
	Kafka *args.Kafka `protobuf:"bytes,100,opt,name=kafka,proto3,oneof"`
}

func (*Read_Kafka) isRead_Args() {}

func (m *Read) GetArgs() isRead_Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Read) GetKafka() *args.Kafka {
	if x, ok := m.GetArgs().(*Read_Kafka); ok {
		return x.Kafka
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Read) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Read_Kafka)(nil),
	}
}

type StartReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Read                 *Read        `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StartReadRequest) Reset()         { *m = StartReadRequest{} }
func (m *StartReadRequest) String() string { return proto.CompactTextString(m) }
func (*StartReadRequest) ProtoMessage()    {}
func (*StartReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{3}
}

func (m *StartReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartReadRequest.Unmarshal(m, b)
}
func (m *StartReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartReadRequest.Marshal(b, m, deterministic)
}
func (m *StartReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartReadRequest.Merge(m, src)
}
func (m *StartReadRequest) XXX_Size() int {
	return xxx_messageInfo_StartReadRequest.Size(m)
}
func (m *StartReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartReadRequest proto.InternalMessageInfo

func (m *StartReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *StartReadRequest) GetRead() *Read {
	if m != nil {
		return m.Read
	}
	return nil
}

type StartReadResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// Assigned and returned by plumber-server to identify a successful read request
	ReadId               string   `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartReadResponse) Reset()         { *m = StartReadResponse{} }
func (m *StartReadResponse) String() string { return proto.CompactTextString(m) }
func (*StartReadResponse) ProtoMessage()    {}
func (*StartReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{4}
}

func (m *StartReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartReadResponse.Unmarshal(m, b)
}
func (m *StartReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartReadResponse.Marshal(b, m, deterministic)
}
func (m *StartReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartReadResponse.Merge(m, src)
}
func (m *StartReadResponse) XXX_Size() int {
	return xxx_messageInfo_StartReadResponse.Size(m)
}
func (m *StartReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartReadResponse proto.InternalMessageInfo

func (m *StartReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StartReadResponse) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type StopReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ReadId               string       `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StopReadRequest) Reset()         { *m = StopReadRequest{} }
func (m *StopReadRequest) String() string { return proto.CompactTextString(m) }
func (*StopReadRequest) ProtoMessage()    {}
func (*StopReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{5}
}

func (m *StopReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopReadRequest.Unmarshal(m, b)
}
func (m *StopReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopReadRequest.Marshal(b, m, deterministic)
}
func (m *StopReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopReadRequest.Merge(m, src)
}
func (m *StopReadRequest) XXX_Size() int {
	return xxx_messageInfo_StopReadRequest.Size(m)
}
func (m *StopReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopReadRequest proto.InternalMessageInfo

func (m *StopReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *StopReadRequest) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type StopReadResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StopReadResponse) Reset()         { *m = StopReadResponse{} }
func (m *StopReadResponse) String() string { return proto.CompactTextString(m) }
func (*StopReadResponse) ProtoMessage()    {}
func (*StopReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{6}
}

func (m *StopReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopReadResponse.Unmarshal(m, b)
}
func (m *StopReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopReadResponse.Marshal(b, m, deterministic)
}
func (m *StopReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopReadResponse.Merge(m, src)
}
func (m *StopReadResponse) XXX_Size() int {
	return xxx_messageInfo_StopReadResponse.Size(m)
}
func (m *StopReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopReadResponse proto.InternalMessageInfo

func (m *StopReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type StreamReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ReadId               string       `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamReadRequest) Reset()         { *m = StreamReadRequest{} }
func (m *StreamReadRequest) String() string { return proto.CompactTextString(m) }
func (*StreamReadRequest) ProtoMessage()    {}
func (*StreamReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{7}
}

func (m *StreamReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamReadRequest.Unmarshal(m, b)
}
func (m *StreamReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamReadRequest.Marshal(b, m, deterministic)
}
func (m *StreamReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamReadRequest.Merge(m, src)
}
func (m *StreamReadRequest) XXX_Size() int {
	return xxx_messageInfo_StreamReadRequest.Size(m)
}
func (m *StreamReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamReadRequest proto.InternalMessageInfo

func (m *StreamReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *StreamReadRequest) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type StreamReadResponse struct {
	Messages             []*records.Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Status               *common.Status     `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StreamReadResponse) Reset()         { *m = StreamReadResponse{} }
func (m *StreamReadResponse) String() string { return proto.CompactTextString(m) }
func (*StreamReadResponse) ProtoMessage()    {}
func (*StreamReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{8}
}

func (m *StreamReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamReadResponse.Unmarshal(m, b)
}
func (m *StreamReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamReadResponse.Marshal(b, m, deterministic)
}
func (m *StreamReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamReadResponse.Merge(m, src)
}
func (m *StreamReadResponse) XXX_Size() int {
	return xxx_messageInfo_StreamReadResponse.Size(m)
}
func (m *StreamReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamReadResponse proto.InternalMessageInfo

func (m *StreamReadResponse) GetMessages() []*records.Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *StreamReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetAllReadsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllReadsRequest) Reset()         { *m = GetAllReadsRequest{} }
func (m *GetAllReadsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllReadsRequest) ProtoMessage()    {}
func (*GetAllReadsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{9}
}

func (m *GetAllReadsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllReadsRequest.Unmarshal(m, b)
}
func (m *GetAllReadsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllReadsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllReadsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllReadsRequest.Merge(m, src)
}
func (m *GetAllReadsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllReadsRequest.Size(m)
}
func (m *GetAllReadsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllReadsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllReadsRequest proto.InternalMessageInfo

func (m *GetAllReadsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetAllReadsResponse struct {
	Read                 []*Read        `protobuf:"bytes,1,rep,name=read,proto3" json:"read,omitempty"`
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetAllReadsResponse) Reset()         { *m = GetAllReadsResponse{} }
func (m *GetAllReadsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllReadsResponse) ProtoMessage()    {}
func (*GetAllReadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{10}
}

func (m *GetAllReadsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllReadsResponse.Unmarshal(m, b)
}
func (m *GetAllReadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllReadsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllReadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllReadsResponse.Merge(m, src)
}
func (m *GetAllReadsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllReadsResponse.Size(m)
}
func (m *GetAllReadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllReadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllReadsResponse proto.InternalMessageInfo

func (m *GetAllReadsResponse) GetRead() []*Read {
	if m != nil {
		return m.Read
	}
	return nil
}

func (m *GetAllReadsResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.ReadOptions_Type", ReadOptions_Type_name, ReadOptions_Type_value)
	proto.RegisterEnum("protos.SampleOptions_Interval", SampleOptions_Interval_name, SampleOptions_Interval_value)
	proto.RegisterType((*ReadOptions)(nil), "protos.ReadOptions")
	proto.RegisterType((*SampleOptions)(nil), "protos.SampleOptions")
	proto.RegisterType((*Read)(nil), "protos.Read")
	proto.RegisterType((*StartReadRequest)(nil), "protos.StartReadRequest")
	proto.RegisterType((*StartReadResponse)(nil), "protos.StartReadResponse")
	proto.RegisterType((*StopReadRequest)(nil), "protos.StopReadRequest")
	proto.RegisterType((*StopReadResponse)(nil), "protos.StopReadResponse")
	proto.RegisterType((*StreamReadRequest)(nil), "protos.StreamReadRequest")
	proto.RegisterType((*StreamReadResponse)(nil), "protos.StreamReadResponse")
	proto.RegisterType((*GetAllReadsRequest)(nil), "protos.GetAllReadsRequest")
	proto.RegisterType((*GetAllReadsResponse)(nil), "protos.GetAllReadsResponse")
}

func init() { proto.RegisterFile("read.proto", fileDescriptor_7b10ec61df6818dd) }

var fileDescriptor_7b10ec61df6818dd = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xae, 0xd3, 0x34, 0x6d, 0x4f, 0x2e, 0x4d, 0xa7, 0xfa, 0xff, 0x5a, 0x5d, 0x40, 0x64, 0x58,
	0x44, 0x48, 0xd8, 0x52, 0x8a, 0xba, 0x42, 0xa0, 0xb6, 0x44, 0x25, 0x42, 0x4d, 0xa4, 0x49, 0xca,
	0x02, 0x21, 0xa2, 0x89, 0x3d, 0x38, 0x56, 0x63, 0x8f, 0x99, 0x19, 0x23, 0xf5, 0x29, 0x78, 0x02,
	0xb6, 0x3c, 0x1b, 0x8f, 0x81, 0xe6, 0xe2, 0x5c, 0xa4, 0xb2, 0x68, 0xc5, 0x2a, 0x33, 0xe7, 0x7c,
	0xdf, 0x37, 0xdf, 0xb9, 0xc4, 0x00, 0x9c, 0x92, 0xc8, 0xcf, 0x39, 0x93, 0x0c, 0xd5, 0xf4, 0x8f,
	0x38, 0x39, 0x0c, 0x59, 0x9a, 0xb2, 0x2c, 0x20, 0x85, 0x9c, 0x9b, 0xd4, 0xc9, 0x91, 0x0d, 0x09,
	0x49, 0x64, 0x21, 0x6c, 0xb0, 0x4d, 0x78, 0x2c, 0x82, 0x5b, 0xf2, 0xf5, 0x96, 0xd8, 0xc8, 0xff,
	0x34, 0x0b, 0x59, 0x94, 0x64, 0x71, 0xc0, 0x72, 0x99, 0xb0, 0xac, 0x44, 0x22, 0x4e, 0x43, 0xc6,
	0x23, 0x11, 0xcc, 0x88, 0xa0, 0x26, 0xe6, 0x9d, 0x42, 0x1d, 0x53, 0x12, 0x8d, 0x0c, 0xd0, 0x7b,
	0x0e, 0xd5, 0xc9, 0x5d, 0x4e, 0x51, 0x03, 0xf6, 0x46, 0xc3, 0xfe, 0x74, 0x32, 0xb8, 0xee, 0xb7,
	0xb7, 0x50, 0x0b, 0xe0, 0x72, 0x34, 0x9c, 0x0c, 0x86, 0x37, 0xa3, 0x9b, 0x71, 0xdb, 0xf1, 0x7e,
	0x3a, 0xd0, 0x1c, 0x93, 0x34, 0x5f, 0x50, 0xcb, 0x43, 0x4f, 0xa1, 0x2e, 0x74, 0x60, 0xca, 0x89,
	0xa4, 0xae, 0xd3, 0x71, 0xba, 0x4d, 0x0c, 0x26, 0x84, 0x89, 0xa4, 0xe8, 0x0a, 0x0e, 0x2c, 0x20,
	0xc9, 0x24, 0xe5, 0xdf, 0xc9, 0xc2, 0xad, 0x74, 0x9c, 0x6e, 0xab, 0xf7, 0xc4, 0x18, 0x11, 0xfe,
	0x86, 0xa0, 0x3f, 0xb0, 0x28, 0xdc, 0x32, 0xb4, 0xf2, 0xee, 0x79, 0xb0, 0x57, 0x9e, 0x11, 0x40,
	0x6d, 0xdc, 0xbf, 0x1c, 0x0d, 0xdf, 0xb5, 0xb7, 0xd4, 0xf9, 0x7a, 0x30, 0xbc, 0x99, 0xf4, 0xdb,
	0x8e, 0xf7, 0xab, 0x02, 0x55, 0x55, 0x15, 0x6a, 0x41, 0x25, 0x89, 0xb4, 0x9b, 0x7d, 0x5c, 0x49,
	0x22, 0xf4, 0x0c, 0x9a, 0x21, 0xcb, 0x32, 0x1a, 0xaa, 0x47, 0xa6, 0x49, 0xa4, 0x3d, 0xec, 0xe3,
	0xc6, 0x2a, 0x38, 0x88, 0xd0, 0x19, 0x34, 0xd4, 0x38, 0xa6, 0xb6, 0x79, 0xee, 0x76, 0xc7, 0xe9,
	0xd6, 0x7b, 0x47, 0xa5, 0xcf, 0xb5, 0x76, 0xe1, 0x3a, 0x5f, 0x5d, 0xd0, 0x6b, 0xb0, 0x5e, 0x97,
	0xcc, 0xaa, 0x66, 0xfe, 0x77, 0x6f, 0x85, 0xb8, 0x29, 0x36, 0x3a, 0xf8, 0x16, 0x5a, 0x11, 0x0d,
	0x59, 0xb4, 0x62, 0xef, 0x68, 0xb6, 0x5b, 0xb2, 0xcb, 0xa1, 0xfa, 0x4b, 0x01, 0x83, 0x2f, 0x05,
	0x5e, 0xc0, 0x8e, 0x5e, 0x02, 0x37, 0xd2, 0x3c, 0x54, 0xf2, 0xd4, 0x7a, 0xf8, 0x1f, 0x54, 0xe6,
	0xfd, 0x16, 0x36, 0x90, 0x8b, 0x1a, 0x54, 0xcf, 0x79, 0x2c, 0xbc, 0x2f, 0xd0, 0x1e, 0x4b, 0xc2,
	0xa5, 0xaa, 0x09, 0xd3, 0x6f, 0x05, 0x15, 0x12, 0x75, 0xa1, 0xaa, 0x56, 0xce, 0xfd, 0x31, 0xdc,
	0xac, 0xdb, 0xec, 0x9e, 0x7f, 0x5e, 0xc8, 0x39, 0xd6, 0x08, 0xd4, 0x81, 0xaa, 0xaa, 0x5f, 0xf7,
	0xb7, 0xde, 0x6b, 0xac, 0x37, 0x08, 0xeb, 0x8c, 0xf7, 0x19, 0x0e, 0xd7, 0xf4, 0x45, 0xce, 0x32,
	0x41, 0x91, 0x0f, 0x35, 0xb3, 0xc0, 0xee, 0xef, 0xdd, 0xcd, 0x06, 0xd9, 0x27, 0xc6, 0x3a, 0x8b,
	0x2d, 0x0a, 0x1d, 0xc3, 0xae, 0x9e, 0xc7, 0x72, 0x92, 0x35, 0x75, 0x1d, 0x44, 0xde, 0x04, 0x0e,
	0xc6, 0x92, 0xe5, 0x8f, 0x33, 0xff, 0x57, 0xd5, 0x0b, 0xd5, 0x93, 0x52, 0xf5, 0x71, 0x96, 0xbd,
	0x8f, 0xaa, 0x6e, 0x4e, 0x49, 0xfa, 0x8f, 0xbd, 0xdd, 0x01, 0x5a, 0xd7, 0xb5, 0xee, 0x4e, 0x61,
	0x2f, 0xa5, 0x42, 0x90, 0x98, 0x0a, 0xd7, 0xe9, 0x6c, 0x77, 0xeb, 0xbd, 0xe3, 0x52, 0xdb, 0xfe,
	0xe3, 0xfd, 0x6b, 0x93, 0xc7, 0x4b, 0xe0, 0x83, 0x4b, 0x7a, 0x03, 0xe8, 0x8a, 0xca, 0xf3, 0xc5,
	0x42, 0x3d, 0x2d, 0x1e, 0x5c, 0x93, 0x17, 0xc3, 0xd1, 0x06, 0xdf, 0x7a, 0x5f, 0xed, 0xd0, 0xf6,
	0xfd, 0x3b, 0xf4, 0x50, 0xa3, 0x17, 0x67, 0x9f, 0x5e, 0xc5, 0x89, 0x9c, 0x17, 0x33, 0x95, 0x0f,
	0x66, 0x44, 0x86, 0xf3, 0x90, 0xf1, 0x3c, 0xc8, 0x17, 0x45, 0x3a, 0xa3, 0xfc, 0xa5, 0x08, 0xe7,
	0x34, 0x25, 0x22, 0x98, 0x15, 0xc9, 0x22, 0x0a, 0x62, 0x16, 0x18, 0xb5, 0x99, 0xf9, 0xee, 0x9e,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xed, 0x24, 0x18, 0x07, 0x8c, 0x05, 0x00, 0x00,
}
