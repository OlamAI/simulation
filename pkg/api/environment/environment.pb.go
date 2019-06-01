// Code generated by protoc-gen-go. DO NOT EDIT.
// source: environment.proto

package endpoints_terrariumai_environment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

// Taks we have to do
type Entity struct {
	// Unique integer identifier of the agent
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Entity stats
	Class  string `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	X      int32  `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32  `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
	Energy int32  `protobuf:"varint,5,opt,name=energy,proto3" json:"energy,omitempty"`
	Health int32  `protobuf:"varint,6,opt,name=health,proto3" json:"health,omitempty"`
	// Owner details
	OwnerUID             string   `protobuf:"bytes,7,opt,name=ownerUID,proto3" json:"ownerUID,omitempty"`
	ModelName            string   `protobuf:"bytes,8,opt,name=modelName,proto3" json:"modelName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entity) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *Entity) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Entity) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Entity) GetEnergy() int32 {
	if m != nil {
		return m.Energy
	}
	return 0
}

func (m *Entity) GetHealth() int32 {
	if m != nil {
		return m.Health
	}
	return 0
}

func (m *Entity) GetOwnerUID() string {
	if m != nil {
		return m.OwnerUID
	}
	return ""
}

func (m *Entity) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

// Request data to create new agent
type CreateEntityRequest struct {
	// agent
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntityRequest) Reset()         { *m = CreateEntityRequest{} }
func (m *CreateEntityRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntityRequest) ProtoMessage()    {}
func (*CreateEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{1}
}

func (m *CreateEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityRequest.Unmarshal(m, b)
}
func (m *CreateEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityRequest.Marshal(b, m, deterministic)
}
func (m *CreateEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityRequest.Merge(m, src)
}
func (m *CreateEntityRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntityRequest.Size(m)
}
func (m *CreateEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityRequest proto.InternalMessageInfo

func (m *CreateEntityRequest) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// Contains data of created agent
type CreateEntityResponse struct {
	// ID of created agent
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntityResponse) Reset()         { *m = CreateEntityResponse{} }
func (m *CreateEntityResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEntityResponse) ProtoMessage()    {}
func (*CreateEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{2}
}

func (m *CreateEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityResponse.Unmarshal(m, b)
}
func (m *CreateEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityResponse.Marshal(b, m, deterministic)
}
func (m *CreateEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityResponse.Merge(m, src)
}
func (m *CreateEntityResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEntityResponse.Size(m)
}
func (m *CreateEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityResponse proto.InternalMessageInfo

func (m *CreateEntityResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Request data to read entity
type GetEntityRequest struct {
	// Unique integer identifier of the agent
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntityRequest) Reset()         { *m = GetEntityRequest{} }
func (m *GetEntityRequest) String() string { return proto.CompactTextString(m) }
func (*GetEntityRequest) ProtoMessage()    {}
func (*GetEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{3}
}

func (m *GetEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityRequest.Unmarshal(m, b)
}
func (m *GetEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityRequest.Marshal(b, m, deterministic)
}
func (m *GetEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityRequest.Merge(m, src)
}
func (m *GetEntityRequest) XXX_Size() int {
	return xxx_messageInfo_GetEntityRequest.Size(m)
}
func (m *GetEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntityRequest proto.InternalMessageInfo

func (m *GetEntityRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Contains entity data specified in by ID request
type GetEntityResponse struct {
	// Task entity read by ID
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntityResponse) Reset()         { *m = GetEntityResponse{} }
func (m *GetEntityResponse) String() string { return proto.CompactTextString(m) }
func (*GetEntityResponse) ProtoMessage()    {}
func (*GetEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{4}
}

func (m *GetEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntityResponse.Unmarshal(m, b)
}
func (m *GetEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntityResponse.Marshal(b, m, deterministic)
}
func (m *GetEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntityResponse.Merge(m, src)
}
func (m *GetEntityResponse) XXX_Size() int {
	return xxx_messageInfo_GetEntityResponse.Size(m)
}
func (m *GetEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntityResponse proto.InternalMessageInfo

func (m *GetEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// Request data to delete agent
type DeleteEntityRequest struct {
	// Unique integer identifier of the agent to delete
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntityRequest) Reset()         { *m = DeleteEntityRequest{} }
func (m *DeleteEntityRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEntityRequest) ProtoMessage()    {}
func (*DeleteEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{5}
}

func (m *DeleteEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntityRequest.Unmarshal(m, b)
}
func (m *DeleteEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntityRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntityRequest.Merge(m, src)
}
func (m *DeleteEntityRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEntityRequest.Size(m)
}
func (m *DeleteEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntityRequest proto.InternalMessageInfo

func (m *DeleteEntityRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Contains status of delete operation
type DeleteEntityResponse struct {
	// Contains number of entities have beed deleted
	// Equals 1 in case of successful delete
	Deleted              int64    `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntityResponse) Reset()         { *m = DeleteEntityResponse{} }
func (m *DeleteEntityResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEntityResponse) ProtoMessage()    {}
func (*DeleteEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{6}
}

func (m *DeleteEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntityResponse.Unmarshal(m, b)
}
func (m *DeleteEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntityResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntityResponse.Merge(m, src)
}
func (m *DeleteEntityResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEntityResponse.Size(m)
}
func (m *DeleteEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntityResponse proto.InternalMessageInfo

func (m *DeleteEntityResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type ExecuteAgentActionRequest struct {
	// Id for the agent that should perform the action
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// identifier for the action
	// 0: MOVE
	// 1: EAT
	Action uint32 `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
	// direction to perform the action in
	// 0: UP
	// 1: DOWN
	// 2: LEFT
	// 3: RIGHT
	Direction            uint32   `protobuf:"varint,3,opt,name=direction,proto3" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteAgentActionRequest) Reset()         { *m = ExecuteAgentActionRequest{} }
func (m *ExecuteAgentActionRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteAgentActionRequest) ProtoMessage()    {}
func (*ExecuteAgentActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{7}
}

func (m *ExecuteAgentActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteAgentActionRequest.Unmarshal(m, b)
}
func (m *ExecuteAgentActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteAgentActionRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteAgentActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteAgentActionRequest.Merge(m, src)
}
func (m *ExecuteAgentActionRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteAgentActionRequest.Size(m)
}
func (m *ExecuteAgentActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteAgentActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteAgentActionRequest proto.InternalMessageInfo

func (m *ExecuteAgentActionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExecuteAgentActionRequest) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *ExecuteAgentActionRequest) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ExecuteAgentActionResponse struct {
	// Was the agent able to perform the action
	WasSuccessful        bool     `protobuf:"varint,3,opt,name=wasSuccessful,proto3" json:"wasSuccessful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteAgentActionResponse) Reset()         { *m = ExecuteAgentActionResponse{} }
func (m *ExecuteAgentActionResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteAgentActionResponse) ProtoMessage()    {}
func (*ExecuteAgentActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{8}
}

func (m *ExecuteAgentActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteAgentActionResponse.Unmarshal(m, b)
}
func (m *ExecuteAgentActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteAgentActionResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteAgentActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteAgentActionResponse.Merge(m, src)
}
func (m *ExecuteAgentActionResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteAgentActionResponse.Size(m)
}
func (m *ExecuteAgentActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteAgentActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteAgentActionResponse proto.InternalMessageInfo

func (m *ExecuteAgentActionResponse) GetWasSuccessful() bool {
	if m != nil {
		return m.WasSuccessful
	}
	return false
}

func init() {
	proto.RegisterType((*Entity)(nil), "endpoints.terrariumai.environment.Entity")
	proto.RegisterType((*CreateEntityRequest)(nil), "endpoints.terrariumai.environment.CreateEntityRequest")
	proto.RegisterType((*CreateEntityResponse)(nil), "endpoints.terrariumai.environment.CreateEntityResponse")
	proto.RegisterType((*GetEntityRequest)(nil), "endpoints.terrariumai.environment.GetEntityRequest")
	proto.RegisterType((*GetEntityResponse)(nil), "endpoints.terrariumai.environment.GetEntityResponse")
	proto.RegisterType((*DeleteEntityRequest)(nil), "endpoints.terrariumai.environment.DeleteEntityRequest")
	proto.RegisterType((*DeleteEntityResponse)(nil), "endpoints.terrariumai.environment.DeleteEntityResponse")
	proto.RegisterType((*ExecuteAgentActionRequest)(nil), "endpoints.terrariumai.environment.ExecuteAgentActionRequest")
	proto.RegisterType((*ExecuteAgentActionResponse)(nil), "endpoints.terrariumai.environment.ExecuteAgentActionResponse")
}

func init() { proto.RegisterFile("environment.proto", fileDescriptor_64e647b85623514a) }

var fileDescriptor_64e647b85623514a = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0x93, 0x26, 0x4d, 0x6e, 0x5b, 0x44, 0xa7, 0x51, 0x65, 0x0c, 0x8b, 0x32, 0x02, 0x54,
	0x36, 0x2e, 0x6a, 0x11, 0x6c, 0x00, 0x29, 0xd0, 0x08, 0xb1, 0x61, 0x61, 0xc4, 0x63, 0x3b, 0xb5,
	0x6f, 0xd3, 0x91, 0xec, 0x99, 0x30, 0x33, 0xa6, 0xf1, 0x92, 0x6f, 0xe0, 0x57, 0xf8, 0x25, 0xfe,
	0x03, 0x79, 0xc6, 0xcd, 0xc3, 0x24, 0xc2, 0x45, 0x2c, 0xcf, 0x99, 0xfb, 0x38, 0xf7, 0xde, 0xa3,
	0x81, 0x3d, 0x14, 0xdf, 0xb8, 0x92, 0x22, 0x43, 0x61, 0xc2, 0x89, 0x92, 0x46, 0x92, 0xfb, 0x28,
	0x92, 0x89, 0xe4, 0xc2, 0xe8, 0xd0, 0xa0, 0x52, 0x4c, 0xf1, 0x3c, 0x63, 0x3c, 0x5c, 0x08, 0x0c,
	0xee, 0x8e, 0xa5, 0x1c, 0xa7, 0x78, 0x6c, 0x13, 0xce, 0xf3, 0x8b, 0x63, 0xcc, 0x26, 0xa6, 0x70,
	0xf9, 0xf4, 0xa7, 0x07, 0xdd, 0x91, 0x30, 0xdc, 0x14, 0xe4, 0x16, 0xb4, 0x78, 0xe2, 0x7b, 0x87,
	0xde, 0x51, 0x3f, 0x6a, 0xf1, 0x84, 0x0c, 0xa0, 0x13, 0xa7, 0x4c, 0x6b, 0xbf, 0x65, 0x29, 0x07,
	0xc8, 0x0e, 0x78, 0x53, 0xbf, 0x7d, 0xe8, 0x1d, 0x75, 0x22, 0x6f, 0x5a, 0xa2, 0xc2, 0xdf, 0x74,
	0xa8, 0x20, 0x07, 0xd0, 0x45, 0x81, 0x6a, 0x5c, 0xf8, 0x1d, 0x4b, 0x55, 0xa8, 0xe4, 0x2f, 0x91,
	0xa5, 0xe6, 0xd2, 0xef, 0x3a, 0xde, 0x21, 0x12, 0x40, 0x4f, 0x5e, 0x09, 0x54, 0x1f, 0xdf, 0x9d,
	0xf9, 0x5b, 0xb6, 0xc9, 0x0c, 0x93, 0x7b, 0xd0, 0xcf, 0x64, 0x82, 0xe9, 0x7b, 0x96, 0xa1, 0xdf,
	0xb3, 0x8f, 0x73, 0x82, 0x7e, 0x81, 0xfd, 0x37, 0x0a, 0x99, 0x41, 0xa7, 0x3d, 0xc2, 0xaf, 0x39,
	0x6a, 0x43, 0x86, 0xa5, 0x80, 0x92, 0xb0, 0x63, 0x6c, 0x9f, 0x3c, 0x0e, 0xff, 0xba, 0x9e, 0xb0,
	0xaa, 0x50, 0x25, 0xd2, 0x47, 0x30, 0x58, 0xae, 0xac, 0x27, 0x52, 0x68, 0xac, 0x6f, 0x87, 0x52,
	0xb8, 0xfd, 0x16, 0xcd, 0x72, 0xfb, 0x7a, 0xcc, 0x27, 0xd8, 0x5b, 0x88, 0xa9, 0x0a, 0xfd, 0x07,
	0x8d, 0x0f, 0x61, 0xff, 0x0c, 0x53, 0xac, 0x4f, 0x5f, 0x6f, 0xff, 0x04, 0x06, 0xcb, 0x61, 0x95,
	0x02, 0x1f, 0xb6, 0x12, 0xcb, 0xbb, 0xe0, 0x76, 0x74, 0x0d, 0x29, 0x83, 0x3b, 0xa3, 0x29, 0xc6,
	0xb9, 0xc1, 0xe1, 0x18, 0x85, 0x19, 0xc6, 0x86, 0x4b, 0xb1, 0xa6, 0x7c, 0x79, 0x55, 0x66, 0x03,
	0xac, 0x41, 0x76, 0xa3, 0x0a, 0x95, 0x97, 0x4b, 0xb8, 0x42, 0xf7, 0xd4, 0xb6, 0x4f, 0x73, 0x82,
	0xbe, 0x86, 0x60, 0x55, 0x8b, 0x4a, 0xda, 0x03, 0xd8, 0xbd, 0x62, 0xfa, 0x43, 0x1e, 0xc7, 0xa8,
	0xf5, 0x45, 0x9e, 0xda, 0xfc, 0x5e, 0xb4, 0x4c, 0x9e, 0xfc, 0xda, 0x84, 0xed, 0xd1, 0x7c, 0x3d,
	0xe4, 0xbb, 0x07, 0x3b, 0x8b, 0x47, 0x23, 0xcf, 0x1a, 0xec, 0x74, 0x85, 0x7f, 0x82, 0xe7, 0x37,
	0xce, 0x73, 0xba, 0xe9, 0x06, 0x99, 0x42, 0x7f, 0x76, 0x6b, 0x72, 0xda, 0xa0, 0x4e, 0xdd, 0x3d,
	0xc1, 0xd3, 0x9b, 0x25, 0xcd, 0x3a, 0x97, 0xd3, 0x2f, 0xde, 0xb9, 0xd1, 0xf4, 0x2b, 0xfc, 0xd3,
	0x68, 0xfa, 0x55, 0x86, 0xa2, 0x1b, 0xe4, 0x87, 0x07, 0xe4, 0xcf, 0xb3, 0x92, 0x17, 0x4d, 0xbc,
	0xbd, 0xce, 0x70, 0xc1, 0xcb, 0x7f, 0xcc, 0x9e, 0xa9, 0x7a, 0x05, 0x10, 0xa1, 0x46, 0xf3, 0x59,
	0xaa, 0x34, 0x21, 0x07, 0xa1, 0xfb, 0x08, 0xc3, 0xeb, 0x8f, 0x30, 0x1c, 0x95, 0x1f, 0x61, 0xb0,
	0x86, 0xa7, 0x1b, 0xe7, 0x5d, 0xcb, 0x9c, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x92, 0x1e,
	0x6c, 0x78, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EnvironmentClient is the client API for Environment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnvironmentClient interface {
	// Create new agent
	CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error)
	// Get data for an entity
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error)
	// Delete an agent
	DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error)
	// Perform an action for an agent
	ExecuteAgentAction(ctx context.Context, in *ExecuteAgentActionRequest, opts ...grpc.CallOption) (*ExecuteAgentActionResponse, error)
	// Reset the world
	ResetWorld(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type environmentClient struct {
	cc *grpc.ClientConn
}

func NewEnvironmentClient(cc *grpc.ClientConn) EnvironmentClient {
	return &environmentClient{cc}
}

func (c *environmentClient) CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error) {
	out := new(CreateEntityResponse)
	err := c.cc.Invoke(ctx, "/endpoints.terrariumai.environment.Environment/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error) {
	out := new(GetEntityResponse)
	err := c.cc.Invoke(ctx, "/endpoints.terrariumai.environment.Environment/GetEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentClient) DeleteEntity(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error) {
	out := new(DeleteEntityResponse)
	err := c.cc.Invoke(ctx, "/endpoints.terrariumai.environment.Environment/DeleteEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentClient) ExecuteAgentAction(ctx context.Context, in *ExecuteAgentActionRequest, opts ...grpc.CallOption) (*ExecuteAgentActionResponse, error) {
	out := new(ExecuteAgentActionResponse)
	err := c.cc.Invoke(ctx, "/endpoints.terrariumai.environment.Environment/ExecuteAgentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentClient) ResetWorld(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/endpoints.terrariumai.environment.Environment/ResetWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentServer is the server API for Environment service.
type EnvironmentServer interface {
	// Create new agent
	CreateEntity(context.Context, *CreateEntityRequest) (*CreateEntityResponse, error)
	// Get data for an entity
	GetEntity(context.Context, *GetEntityRequest) (*GetEntityResponse, error)
	// Delete an agent
	DeleteEntity(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error)
	// Perform an action for an agent
	ExecuteAgentAction(context.Context, *ExecuteAgentActionRequest) (*ExecuteAgentActionResponse, error)
	// Reset the world
	ResetWorld(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterEnvironmentServer(s *grpc.Server, srv EnvironmentServer) {
	s.RegisterService(&_Environment_serviceDesc, srv)
}

func _Environment_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.terrariumai.environment.Environment/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServer).CreateEntity(ctx, req.(*CreateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environment_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.terrariumai.environment.Environment/GetEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environment_DeleteEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServer).DeleteEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.terrariumai.environment.Environment/DeleteEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServer).DeleteEntity(ctx, req.(*DeleteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environment_ExecuteAgentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteAgentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServer).ExecuteAgentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.terrariumai.environment.Environment/ExecuteAgentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServer).ExecuteAgentAction(ctx, req.(*ExecuteAgentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environment_ResetWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServer).ResetWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.terrariumai.environment.Environment/ResetWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServer).ResetWorld(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Environment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "endpoints.terrariumai.environment.Environment",
	HandlerType: (*EnvironmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntity",
			Handler:    _Environment_CreateEntity_Handler,
		},
		{
			MethodName: "GetEntity",
			Handler:    _Environment_GetEntity_Handler,
		},
		{
			MethodName: "DeleteEntity",
			Handler:    _Environment_DeleteEntity_Handler,
		},
		{
			MethodName: "ExecuteAgentAction",
			Handler:    _Environment_ExecuteAgentAction_Handler,
		},
		{
			MethodName: "ResetWorld",
			Handler:    _Environment_ResetWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "environment.proto",
}
