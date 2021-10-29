// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_schema_validation.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
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

type Validation struct {
	XId                  string              `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	SchemaId             string              `protobuf:"bytes,2,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	Fields               []*Validation_Field `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Validation) Reset()         { *m = Validation{} }
func (m *Validation) String() string { return proto.CompactTextString(m) }
func (*Validation) ProtoMessage()    {}
func (*Validation) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{0}
}

func (m *Validation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validation.Unmarshal(m, b)
}
func (m *Validation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validation.Marshal(b, m, deterministic)
}
func (m *Validation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validation.Merge(m, src)
}
func (m *Validation) XXX_Size() int {
	return xxx_messageInfo_Validation.Size(m)
}
func (m *Validation) XXX_DiscardUnknown() {
	xxx_messageInfo_Validation.DiscardUnknown(m)
}

var xxx_messageInfo_Validation proto.InternalMessageInfo

func (m *Validation) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Validation) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *Validation) GetFields() []*Validation_Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Validation_Field struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	ValidationType       string   `protobuf:"bytes,2,opt,name=validation_type,json=validationType,proto3" json:"validation_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validation_Field) Reset()         { *m = Validation_Field{} }
func (m *Validation_Field) String() string { return proto.CompactTextString(m) }
func (*Validation_Field) ProtoMessage()    {}
func (*Validation_Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{0, 0}
}

func (m *Validation_Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validation_Field.Unmarshal(m, b)
}
func (m *Validation_Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validation_Field.Marshal(b, m, deterministic)
}
func (m *Validation_Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validation_Field.Merge(m, src)
}
func (m *Validation_Field) XXX_Size() int {
	return xxx_messageInfo_Validation_Field.Size(m)
}
func (m *Validation_Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Validation_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Validation_Field proto.InternalMessageInfo

func (m *Validation_Field) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Validation_Field) GetValidationType() string {
	if m != nil {
		return m.ValidationType
	}
	return ""
}

type GetValidationRequest struct {
	// Every gRPC request must have a valid auth config
	Auth *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	// Validation ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetValidationRequest) Reset()         { *m = GetValidationRequest{} }
func (m *GetValidationRequest) String() string { return proto.CompactTextString(m) }
func (*GetValidationRequest) ProtoMessage()    {}
func (*GetValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{1}
}

func (m *GetValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetValidationRequest.Unmarshal(m, b)
}
func (m *GetValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetValidationRequest.Marshal(b, m, deterministic)
}
func (m *GetValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValidationRequest.Merge(m, src)
}
func (m *GetValidationRequest) XXX_Size() int {
	return xxx_messageInfo_GetValidationRequest.Size(m)
}
func (m *GetValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetValidationRequest proto.InternalMessageInfo

func (m *GetValidationRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetValidationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetValidationResponse struct {
	Validation           *Validation `protobuf:"bytes,1,opt,name=validation,proto3" json:"validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetValidationResponse) Reset()         { *m = GetValidationResponse{} }
func (m *GetValidationResponse) String() string { return proto.CompactTextString(m) }
func (*GetValidationResponse) ProtoMessage()    {}
func (*GetValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{2}
}

func (m *GetValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetValidationResponse.Unmarshal(m, b)
}
func (m *GetValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetValidationResponse.Marshal(b, m, deterministic)
}
func (m *GetValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValidationResponse.Merge(m, src)
}
func (m *GetValidationResponse) XXX_Size() int {
	return xxx_messageInfo_GetValidationResponse.Size(m)
}
func (m *GetValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetValidationResponse proto.InternalMessageInfo

func (m *GetValidationResponse) GetValidation() *Validation {
	if m != nil {
		return m.Validation
	}
	return nil
}

type GetAllValidationsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	// Validation ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllValidationsRequest) Reset()         { *m = GetAllValidationsRequest{} }
func (m *GetAllValidationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllValidationsRequest) ProtoMessage()    {}
func (*GetAllValidationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{3}
}

func (m *GetAllValidationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllValidationsRequest.Unmarshal(m, b)
}
func (m *GetAllValidationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllValidationsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllValidationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllValidationsRequest.Merge(m, src)
}
func (m *GetAllValidationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllValidationsRequest.Size(m)
}
func (m *GetAllValidationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllValidationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllValidationsRequest proto.InternalMessageInfo

func (m *GetAllValidationsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetAllValidationsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAllValidationsResponse struct {
	Validations          []*Validation `protobuf:"bytes,1,rep,name=validations,proto3" json:"validations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllValidationsResponse) Reset()         { *m = GetAllValidationsResponse{} }
func (m *GetAllValidationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllValidationsResponse) ProtoMessage()    {}
func (*GetAllValidationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{4}
}

func (m *GetAllValidationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllValidationsResponse.Unmarshal(m, b)
}
func (m *GetAllValidationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllValidationsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllValidationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllValidationsResponse.Merge(m, src)
}
func (m *GetAllValidationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllValidationsResponse.Size(m)
}
func (m *GetAllValidationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllValidationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllValidationsResponse proto.InternalMessageInfo

func (m *GetAllValidationsResponse) GetValidations() []*Validation {
	if m != nil {
		return m.Validations
	}
	return nil
}

type CreateValidationRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Validation           *Validation  `protobuf:"bytes,1,opt,name=validation,proto3" json:"validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateValidationRequest) Reset()         { *m = CreateValidationRequest{} }
func (m *CreateValidationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateValidationRequest) ProtoMessage()    {}
func (*CreateValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{5}
}

func (m *CreateValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateValidationRequest.Unmarshal(m, b)
}
func (m *CreateValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateValidationRequest.Marshal(b, m, deterministic)
}
func (m *CreateValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateValidationRequest.Merge(m, src)
}
func (m *CreateValidationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateValidationRequest.Size(m)
}
func (m *CreateValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateValidationRequest proto.InternalMessageInfo

func (m *CreateValidationRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateValidationRequest) GetValidation() *Validation {
	if m != nil {
		return m.Validation
	}
	return nil
}

type CreateValidationResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	Validation           *Validation    `protobuf:"bytes,1,opt,name=validation,proto3" json:"validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateValidationResponse) Reset()         { *m = CreateValidationResponse{} }
func (m *CreateValidationResponse) String() string { return proto.CompactTextString(m) }
func (*CreateValidationResponse) ProtoMessage()    {}
func (*CreateValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{6}
}

func (m *CreateValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateValidationResponse.Unmarshal(m, b)
}
func (m *CreateValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateValidationResponse.Marshal(b, m, deterministic)
}
func (m *CreateValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateValidationResponse.Merge(m, src)
}
func (m *CreateValidationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateValidationResponse.Size(m)
}
func (m *CreateValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateValidationResponse proto.InternalMessageInfo

func (m *CreateValidationResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateValidationResponse) GetValidation() *Validation {
	if m != nil {
		return m.Validation
	}
	return nil
}

type UpdateValidationRequest struct {
	// Every gRPC request must have a valid auth config
	Auth *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	// Validation ID
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Validation           *Validation `protobuf:"bytes,2,opt,name=validation,proto3" json:"validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateValidationRequest) Reset()         { *m = UpdateValidationRequest{} }
func (m *UpdateValidationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateValidationRequest) ProtoMessage()    {}
func (*UpdateValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{7}
}

func (m *UpdateValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateValidationRequest.Unmarshal(m, b)
}
func (m *UpdateValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateValidationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateValidationRequest.Merge(m, src)
}
func (m *UpdateValidationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateValidationRequest.Size(m)
}
func (m *UpdateValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateValidationRequest proto.InternalMessageInfo

func (m *UpdateValidationRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UpdateValidationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateValidationRequest) GetValidation() *Validation {
	if m != nil {
		return m.Validation
	}
	return nil
}

type UpdateValidationResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	Validation           *Validation    `protobuf:"bytes,1,opt,name=validation,proto3" json:"validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateValidationResponse) Reset()         { *m = UpdateValidationResponse{} }
func (m *UpdateValidationResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateValidationResponse) ProtoMessage()    {}
func (*UpdateValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{8}
}

func (m *UpdateValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateValidationResponse.Unmarshal(m, b)
}
func (m *UpdateValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateValidationResponse.Marshal(b, m, deterministic)
}
func (m *UpdateValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateValidationResponse.Merge(m, src)
}
func (m *UpdateValidationResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateValidationResponse.Size(m)
}
func (m *UpdateValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateValidationResponse proto.InternalMessageInfo

func (m *UpdateValidationResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *UpdateValidationResponse) GetValidation() *Validation {
	if m != nil {
		return m.Validation
	}
	return nil
}

type DeleteValidationRequest struct {
	// Every gRPC request must have a valid auth config
	Auth *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	// Validation ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteValidationRequest) Reset()         { *m = DeleteValidationRequest{} }
func (m *DeleteValidationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteValidationRequest) ProtoMessage()    {}
func (*DeleteValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{9}
}

func (m *DeleteValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteValidationRequest.Unmarshal(m, b)
}
func (m *DeleteValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteValidationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteValidationRequest.Merge(m, src)
}
func (m *DeleteValidationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteValidationRequest.Size(m)
}
func (m *DeleteValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteValidationRequest proto.InternalMessageInfo

func (m *DeleteValidationRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DeleteValidationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteValidationResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteValidationResponse) Reset()         { *m = DeleteValidationResponse{} }
func (m *DeleteValidationResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteValidationResponse) ProtoMessage()    {}
func (*DeleteValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b42dfee40054c1, []int{10}
}

func (m *DeleteValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteValidationResponse.Unmarshal(m, b)
}
func (m *DeleteValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteValidationResponse.Marshal(b, m, deterministic)
}
func (m *DeleteValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteValidationResponse.Merge(m, src)
}
func (m *DeleteValidationResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteValidationResponse.Size(m)
}
func (m *DeleteValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteValidationResponse proto.InternalMessageInfo

func (m *DeleteValidationResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Validation)(nil), "protos.Validation")
	proto.RegisterType((*Validation_Field)(nil), "protos.Validation.Field")
	proto.RegisterType((*GetValidationRequest)(nil), "protos.GetValidationRequest")
	proto.RegisterType((*GetValidationResponse)(nil), "protos.GetValidationResponse")
	proto.RegisterType((*GetAllValidationsRequest)(nil), "protos.GetAllValidationsRequest")
	proto.RegisterType((*GetAllValidationsResponse)(nil), "protos.GetAllValidationsResponse")
	proto.RegisterType((*CreateValidationRequest)(nil), "protos.CreateValidationRequest")
	proto.RegisterType((*CreateValidationResponse)(nil), "protos.CreateValidationResponse")
	proto.RegisterType((*UpdateValidationRequest)(nil), "protos.UpdateValidationRequest")
	proto.RegisterType((*UpdateValidationResponse)(nil), "protos.UpdateValidationResponse")
	proto.RegisterType((*DeleteValidationRequest)(nil), "protos.DeleteValidationRequest")
	proto.RegisterType((*DeleteValidationResponse)(nil), "protos.DeleteValidationResponse")
}

func init() { proto.RegisterFile("ps_schema_validation.proto", fileDescriptor_56b42dfee40054c1) }

var fileDescriptor_56b42dfee40054c1 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcf, 0x6a, 0xd4, 0x40,
	0x18, 0x27, 0xd9, 0xba, 0xda, 0x2f, 0xd0, 0xc2, 0x68, 0xd9, 0x31, 0x45, 0x58, 0x72, 0x31, 0x17,
	0x13, 0x89, 0xc5, 0x7b, 0xb5, 0x58, 0xaa, 0x20, 0x9a, 0x56, 0x0f, 0x5e, 0xc2, 0x24, 0x33, 0x36,
	0x03, 0xc9, 0xce, 0x98, 0x99, 0x28, 0xbd, 0x78, 0xf5, 0xe8, 0xbb, 0xf8, 0x54, 0x3e, 0x86, 0xec,
	0x4c, 0x74, 0x4a, 0xd2, 0x83, 0xec, 0x8a, 0xa7, 0x1d, 0xbe, 0xef, 0xb7, 0xbf, 0x7f, 0x13, 0x06,
	0x42, 0xa9, 0x0a, 0x55, 0xd5, 0xac, 0x25, 0xc5, 0x67, 0xd2, 0x70, 0x4a, 0x34, 0x17, 0xab, 0x44,
	0x76, 0x42, 0x0b, 0x34, 0x37, 0x3f, 0x2a, 0x3c, 0xac, 0x44, 0xdb, 0x8a, 0x55, 0x2a, 0x55, 0x61,
	0x4f, 0x05, 0xe9, 0x75, 0x6d, 0x41, 0xe1, 0x83, 0xc9, 0x52, 0x69, 0xa2, 0x7b, 0x65, 0xd7, 0xd1,
	0x0f, 0x0f, 0xe0, 0xfd, 0x1f, 0x62, 0xb4, 0x0f, 0xb3, 0x82, 0x53, 0xec, 0x2d, 0xbd, 0x78, 0x37,
	0xf7, 0xcf, 0x28, 0x3a, 0x84, 0xdd, 0x41, 0x9e, 0x53, 0xec, 0x9b, 0xf1, 0x1d, 0x3b, 0x38, 0xa3,
	0xe8, 0x31, 0xcc, 0x3f, 0x72, 0xd6, 0x50, 0x85, 0x67, 0xcb, 0x59, 0x1c, 0x64, 0xd8, 0x92, 0xaa,
	0xc4, 0x31, 0x26, 0x2f, 0xd6, 0x80, 0x7c, 0xc0, 0x85, 0x27, 0x70, 0xcb, 0x0c, 0x10, 0x82, 0x1d,
	0x49, 0x74, 0x3d, 0x28, 0x99, 0x33, 0x7a, 0x08, 0xfb, 0x2e, 0x63, 0xa1, 0xaf, 0x24, 0x1b, 0x14,
	0xf7, 0xdc, 0xf8, 0xe2, 0x4a, 0xb2, 0xe8, 0x0d, 0xdc, 0x3b, 0x65, 0xda, 0x89, 0xe4, 0xec, 0x53,
	0xcf, 0x94, 0x46, 0x31, 0xec, 0xac, 0x93, 0xe3, 0xef, 0xaf, 0x97, 0x5e, 0x1c, 0x64, 0x77, 0x7f,
	0xdb, 0xb1, 0xc1, 0x93, 0xe3, 0x5e, 0xd7, 0xb9, 0x41, 0xa0, 0x3d, 0xf0, 0x5d, 0x4c, 0x4e, 0xa3,
	0x57, 0x70, 0x30, 0x62, 0x54, 0x52, 0xac, 0x14, 0x43, 0x19, 0x80, 0x13, 0x37, 0x7f, 0x08, 0x32,
	0x34, 0x8d, 0x99, 0x5f, 0x43, 0x45, 0x17, 0x80, 0x4f, 0x99, 0x3e, 0x6e, 0x1a, 0xb7, 0x57, 0xdb,
	0x5b, 0x7c, 0x0b, 0xf7, 0x6f, 0x60, 0x1d, 0x6c, 0x1e, 0x41, 0xe0, 0x0c, 0x28, 0xec, 0x99, 0xeb,
	0xb8, 0xc9, 0xe7, 0x75, 0x58, 0xf4, 0x05, 0x16, 0xcf, 0x3b, 0x46, 0x34, 0xdb, 0xa6, 0xca, 0x4d,
	0x1a, 0xfa, 0x0a, 0x78, 0x2a, 0x3c, 0x44, 0x49, 0x60, 0x6e, 0xbf, 0x50, 0xfc, 0xf3, 0xb6, 0x21,
	0x3b, 0x18, 0x69, 0x9f, 0x9b, 0x6d, 0x3e, 0xa0, 0x36, 0xd2, 0xff, 0xe6, 0xc1, 0xe2, 0x9d, 0xa4,
	0x5b, 0x26, 0x1f, 0xdd, 0xd0, 0xc8, 0x89, 0xff, 0xb7, 0x4d, 0x4c, 0x8d, 0xfc, 0xc7, 0x26, 0xce,
	0x61, 0x71, 0xc2, 0x1a, 0xf6, 0x4f, 0x8b, 0x88, 0x5e, 0x02, 0x9e, 0x92, 0x6e, 0x16, 0xea, 0xd9,
	0xd3, 0x0f, 0x47, 0x97, 0x5c, 0xd7, 0x7d, 0xb9, 0xde, 0xa7, 0x25, 0xd1, 0x55, 0x5d, 0x89, 0x4e,
	0xa6, 0xb2, 0xe9, 0xdb, 0x92, 0x75, 0x8f, 0xec, 0x63, 0xa4, 0xd2, 0xb2, 0xe7, 0x0d, 0x4d, 0x2f,
	0x45, 0x6a, 0xd9, 0x4a, 0xfb, 0x38, 0x3e, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x79, 0xf0, 0xd8,
	0x02, 0x41, 0x05, 0x00, 0x00,
}
