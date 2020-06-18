// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatService.proto

package pero

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ChatMessageResponse_MessageType int32

const (
	ChatMessageResponse_SYSTEM_JOIN_NEW_USER ChatMessageResponse_MessageType = 0
	ChatMessageResponse_SYSTEM_LEAVE_USER    ChatMessageResponse_MessageType = 1
	ChatMessageResponse_COMMON_MESSAGE       ChatMessageResponse_MessageType = 2
)

var ChatMessageResponse_MessageType_name = map[int32]string{
	0: "SYSTEM_JOIN_NEW_USER",
	1: "SYSTEM_LEAVE_USER",
	2: "COMMON_MESSAGE",
}

var ChatMessageResponse_MessageType_value = map[string]int32{
	"SYSTEM_JOIN_NEW_USER": 0,
	"SYSTEM_LEAVE_USER":    1,
	"COMMON_MESSAGE":       2,
}

func (x ChatMessageResponse_MessageType) String() string {
	return proto.EnumName(ChatMessageResponse_MessageType_name, int32(x))
}

func (ChatMessageResponse_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{5, 0}
}

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type Room struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{1}
}

func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Room) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Room) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Room) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type EntryRequest struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryRequest) Reset()         { *m = EntryRequest{} }
func (m *EntryRequest) String() string { return proto.CompactTextString(m) }
func (*EntryRequest) ProtoMessage()    {}
func (*EntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{2}
}

func (m *EntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryRequest.Unmarshal(m, b)
}
func (m *EntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryRequest.Marshal(b, m, deterministic)
}
func (m *EntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryRequest.Merge(m, src)
}
func (m *EntryRequest) XXX_Size() int {
	return xxx_messageInfo_EntryRequest.Size(m)
}
func (m *EntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntryRequest proto.InternalMessageInfo

func (m *EntryRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type ChatMessageRequest struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessageRequest) Reset()         { *m = ChatMessageRequest{} }
func (m *ChatMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ChatMessageRequest) ProtoMessage()    {}
func (*ChatMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{3}
}

func (m *ChatMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessageRequest.Unmarshal(m, b)
}
func (m *ChatMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessageRequest.Marshal(b, m, deterministic)
}
func (m *ChatMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessageRequest.Merge(m, src)
}
func (m *ChatMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ChatMessageRequest.Size(m)
}
func (m *ChatMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessageRequest proto.InternalMessageInfo

func (m *ChatMessageRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *ChatMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type BroadcastResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Room                 *Room    `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastResponse) Reset()         { *m = BroadcastResponse{} }
func (m *BroadcastResponse) String() string { return proto.CompactTextString(m) }
func (*BroadcastResponse) ProtoMessage()    {}
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{4}
}

func (m *BroadcastResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastResponse.Unmarshal(m, b)
}
func (m *BroadcastResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastResponse.Marshal(b, m, deterministic)
}
func (m *BroadcastResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastResponse.Merge(m, src)
}
func (m *BroadcastResponse) XXX_Size() int {
	return xxx_messageInfo_BroadcastResponse.Size(m)
}
func (m *BroadcastResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastResponse proto.InternalMessageInfo

func (m *BroadcastResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BroadcastResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

type ChatMessageResponse struct {
	MessageType ChatMessageResponse_MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=pero.ChatMessageResponse_MessageType" json:"message_type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ChatMessageResponse_SystemJoinNewUser
	//	*ChatMessageResponse_SystemLeaveUser
	//	*ChatMessageResponse_CommonMessage
	Payload              isChatMessageResponse_Payload `protobuf_oneof:"Payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ChatMessageResponse) Reset()         { *m = ChatMessageResponse{} }
func (m *ChatMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ChatMessageResponse) ProtoMessage()    {}
func (*ChatMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{5}
}

func (m *ChatMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessageResponse.Unmarshal(m, b)
}
func (m *ChatMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessageResponse.Marshal(b, m, deterministic)
}
func (m *ChatMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessageResponse.Merge(m, src)
}
func (m *ChatMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ChatMessageResponse.Size(m)
}
func (m *ChatMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessageResponse proto.InternalMessageInfo

func (m *ChatMessageResponse) GetMessageType() ChatMessageResponse_MessageType {
	if m != nil {
		return m.MessageType
	}
	return ChatMessageResponse_SYSTEM_JOIN_NEW_USER
}

type isChatMessageResponse_Payload interface {
	isChatMessageResponse_Payload()
}

type ChatMessageResponse_SystemJoinNewUser struct {
	SystemJoinNewUser *System_JoinNewUser `protobuf:"bytes,2,opt,name=system_join_new_user,json=systemJoinNewUser,proto3,oneof"`
}

type ChatMessageResponse_SystemLeaveUser struct {
	SystemLeaveUser *System_LeaveUser `protobuf:"bytes,3,opt,name=system_leave_user,json=systemLeaveUser,proto3,oneof"`
}

type ChatMessageResponse_CommonMessage struct {
	CommonMessage *CommonMessage `protobuf:"bytes,4,opt,name=common_message,json=commonMessage,proto3,oneof"`
}

func (*ChatMessageResponse_SystemJoinNewUser) isChatMessageResponse_Payload() {}

func (*ChatMessageResponse_SystemLeaveUser) isChatMessageResponse_Payload() {}

func (*ChatMessageResponse_CommonMessage) isChatMessageResponse_Payload() {}

func (m *ChatMessageResponse) GetPayload() isChatMessageResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChatMessageResponse) GetSystemJoinNewUser() *System_JoinNewUser {
	if x, ok := m.GetPayload().(*ChatMessageResponse_SystemJoinNewUser); ok {
		return x.SystemJoinNewUser
	}
	return nil
}

func (m *ChatMessageResponse) GetSystemLeaveUser() *System_LeaveUser {
	if x, ok := m.GetPayload().(*ChatMessageResponse_SystemLeaveUser); ok {
		return x.SystemLeaveUser
	}
	return nil
}

func (m *ChatMessageResponse) GetCommonMessage() *CommonMessage {
	if x, ok := m.GetPayload().(*ChatMessageResponse_CommonMessage); ok {
		return x.CommonMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ChatMessageResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ChatMessageResponse_SystemJoinNewUser)(nil),
		(*ChatMessageResponse_SystemLeaveUser)(nil),
		(*ChatMessageResponse_CommonMessage)(nil),
	}
}

type System_JoinNewUser struct {
	User                 *User                `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Room                 *Room                `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *System_JoinNewUser) Reset()         { *m = System_JoinNewUser{} }
func (m *System_JoinNewUser) String() string { return proto.CompactTextString(m) }
func (*System_JoinNewUser) ProtoMessage()    {}
func (*System_JoinNewUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{6}
}

func (m *System_JoinNewUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_System_JoinNewUser.Unmarshal(m, b)
}
func (m *System_JoinNewUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_System_JoinNewUser.Marshal(b, m, deterministic)
}
func (m *System_JoinNewUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_System_JoinNewUser.Merge(m, src)
}
func (m *System_JoinNewUser) XXX_Size() int {
	return xxx_messageInfo_System_JoinNewUser.Size(m)
}
func (m *System_JoinNewUser) XXX_DiscardUnknown() {
	xxx_messageInfo_System_JoinNewUser.DiscardUnknown(m)
}

var xxx_messageInfo_System_JoinNewUser proto.InternalMessageInfo

func (m *System_JoinNewUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *System_JoinNewUser) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *System_JoinNewUser) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *System_JoinNewUser) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type System_LeaveUser struct {
	User                 *User                `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Room                 *Room                `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *System_LeaveUser) Reset()         { *m = System_LeaveUser{} }
func (m *System_LeaveUser) String() string { return proto.CompactTextString(m) }
func (*System_LeaveUser) ProtoMessage()    {}
func (*System_LeaveUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{7}
}

func (m *System_LeaveUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_System_LeaveUser.Unmarshal(m, b)
}
func (m *System_LeaveUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_System_LeaveUser.Marshal(b, m, deterministic)
}
func (m *System_LeaveUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_System_LeaveUser.Merge(m, src)
}
func (m *System_LeaveUser) XXX_Size() int {
	return xxx_messageInfo_System_LeaveUser.Size(m)
}
func (m *System_LeaveUser) XXX_DiscardUnknown() {
	xxx_messageInfo_System_LeaveUser.DiscardUnknown(m)
}

var xxx_messageInfo_System_LeaveUser proto.InternalMessageInfo

func (m *System_LeaveUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *System_LeaveUser) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *System_LeaveUser) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *System_LeaveUser) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type CommonMessage struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *User                `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommonMessage) Reset()         { *m = CommonMessage{} }
func (m *CommonMessage) String() string { return proto.CompactTextString(m) }
func (*CommonMessage) ProtoMessage()    {}
func (*CommonMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7fde549b56f4d64, []int{8}
}

func (m *CommonMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonMessage.Unmarshal(m, b)
}
func (m *CommonMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonMessage.Marshal(b, m, deterministic)
}
func (m *CommonMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonMessage.Merge(m, src)
}
func (m *CommonMessage) XXX_Size() int {
	return xxx_messageInfo_CommonMessage.Size(m)
}
func (m *CommonMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CommonMessage proto.InternalMessageInfo

func (m *CommonMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CommonMessage) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CommonMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CommonMessage) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("pero.ChatMessageResponse_MessageType", ChatMessageResponse_MessageType_name, ChatMessageResponse_MessageType_value)
	proto.RegisterType((*User)(nil), "pero.User")
	proto.RegisterType((*Room)(nil), "pero.Room")
	proto.RegisterType((*EntryRequest)(nil), "pero.EntryRequest")
	proto.RegisterType((*ChatMessageRequest)(nil), "pero.ChatMessageRequest")
	proto.RegisterType((*BroadcastResponse)(nil), "pero.BroadcastResponse")
	proto.RegisterType((*ChatMessageResponse)(nil), "pero.ChatMessageResponse")
	proto.RegisterType((*System_JoinNewUser)(nil), "pero.System_JoinNewUser")
	proto.RegisterType((*System_LeaveUser)(nil), "pero.System_LeaveUser")
	proto.RegisterType((*CommonMessage)(nil), "pero.CommonMessage")
}

func init() {
	proto.RegisterFile("chatService.proto", fileDescriptor_a7fde549b56f4d64)
}

var fileDescriptor_a7fde549b56f4d64 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcf, 0x4f, 0xd4, 0x40,
	0x14, 0x66, 0x96, 0x02, 0xd9, 0xb7, 0xb0, 0xb2, 0x03, 0x4a, 0xe5, 0xa0, 0xa4, 0x89, 0x91, 0x53,
	0x31, 0xeb, 0x45, 0x13, 0x63, 0x02, 0xd8, 0xf0, 0x43, 0x76, 0x21, 0x53, 0xd0, 0x78, 0x9a, 0x0c,
	0xed, 0x13, 0x6a, 0xb6, 0x9d, 0xda, 0x19, 0x20, 0x7b, 0xf3, 0xea, 0xc9, 0x8b, 0xff, 0x83, 0xf1,
	0xbf, 0x34, 0xed, 0xb4, 0x6b, 0x97, 0x85, 0x10, 0x63, 0x4c, 0xbc, 0xf5, 0x7d, 0xef, 0xcd, 0xf7,
	0xbe, 0xf7, 0xa3, 0x0f, 0x3a, 0xc1, 0xb9, 0xd0, 0x3e, 0x66, 0x97, 0x51, 0x80, 0x6e, 0x9a, 0x49,
	0x2d, 0xa9, 0x95, 0x62, 0x26, 0x57, 0x1f, 0x9f, 0x49, 0x79, 0x36, 0xc0, 0x8d, 0x02, 0x3b, 0xbd,
	0xf8, 0xb8, 0xa1, 0xa3, 0x18, 0x95, 0x16, 0x71, 0x6a, 0xc2, 0x1c, 0x04, 0xeb, 0x44, 0x61, 0x46,
	0xdb, 0xd0, 0x88, 0x42, 0x9b, 0xac, 0x91, 0xf5, 0x26, 0x6b, 0x44, 0x21, 0xa5, 0x60, 0x25, 0x22,
	0x46, 0xbb, 0x51, 0x20, 0xc5, 0x37, 0x7d, 0x09, 0x10, 0x64, 0x28, 0x34, 0x86, 0x5c, 0x68, 0x7b,
	0x7a, 0x8d, 0xac, 0xb7, 0xba, 0xab, 0xae, 0xc9, 0xe0, 0x56, 0x19, 0xdc, 0xe3, 0x2a, 0x03, 0x6b,
	0x96, 0xd1, 0x9b, 0xda, 0xf9, 0x4a, 0xc0, 0x62, 0x52, 0xc6, 0x13, 0x79, 0x96, 0x61, 0x46, 0x47,
	0x7a, 0x50, 0x25, 0x32, 0x06, 0x5d, 0x83, 0x56, 0x88, 0x2a, 0xc8, 0xa2, 0x54, 0x47, 0x32, 0x29,
	0x52, 0x35, 0x59, 0x1d, 0xba, 0xa6, 0xc5, 0xfa, 0x13, 0x2d, 0x4f, 0x61, 0xde, 0x4b, 0x74, 0x36,
	0x64, 0xf8, 0xf9, 0x02, 0x95, 0xa6, 0x2b, 0x30, 0x97, 0x49, 0x19, 0xf3, 0x91, 0xae, 0xd9, 0xdc,
	0xdc, 0x0b, 0x9d, 0x1d, 0xa0, 0xdb, 0xe7, 0x42, 0xf7, 0x50, 0x29, 0x71, 0x86, 0x77, 0x85, 0x53,
	0x1b, 0xe6, 0x62, 0x13, 0x5a, 0x16, 0x53, 0x99, 0x4e, 0x0f, 0x3a, 0x5b, 0x99, 0x14, 0x61, 0x20,
	0x94, 0x66, 0xa8, 0x52, 0x99, 0x28, 0xac, 0x87, 0x93, 0xb1, 0x70, 0xfa, 0x08, 0xac, 0x9c, 0xb2,
	0x60, 0x69, 0x75, 0xc1, 0xcd, 0x27, 0xe9, 0xe6, 0xdd, 0x63, 0x05, 0xee, 0x7c, 0x9b, 0x86, 0xa5,
	0x31, 0x61, 0x25, 0xe3, 0x2e, 0xcc, 0x97, 0x14, 0x5c, 0x0f, 0x53, 0x43, 0xdb, 0xee, 0x3e, 0x31,
	0xef, 0x6f, 0x78, 0xe0, 0x96, 0xf6, 0xf1, 0x30, 0x45, 0xd6, 0x8a, 0x7f, 0x1b, 0xf4, 0x2d, 0x2c,
	0xab, 0xa1, 0xd2, 0x18, 0xf3, 0x4f, 0x32, 0x4a, 0x78, 0x82, 0x57, 0xfc, 0x42, 0x61, 0x56, 0x2a,
	0xb2, 0x0d, 0xa3, 0x6f, 0x22, 0xf6, 0x65, 0x94, 0xf4, 0xf1, 0x2a, 0xdf, 0xa2, 0xdd, 0x29, 0xd6,
	0x31, 0xef, 0x6a, 0x20, 0x7d, 0x03, 0x25, 0xc8, 0x07, 0x28, 0x2e, 0xd1, 0x30, 0x99, 0xed, 0x79,
	0x30, 0xc6, 0x74, 0x90, 0xbb, 0x4b, 0x9e, 0x7b, 0xe6, 0xc9, 0x08, 0xa2, 0xaf, 0xa0, 0x1d, 0xc8,
	0x38, 0x96, 0x09, 0xaf, 0xba, 0x66, 0x86, 0xbe, 0x54, 0x96, 0x57, 0xf8, 0xca, 0x82, 0x76, 0xa7,
	0xd8, 0x42, 0x50, 0x07, 0x1c, 0x06, 0xad, 0x5a, 0xb1, 0xd4, 0x86, 0x65, 0xff, 0x83, 0x7f, 0xec,
	0xf5, 0xf8, 0xfe, 0xe1, 0x5e, 0x9f, 0xf7, 0xbd, 0xf7, 0xfc, 0xc4, 0xf7, 0xd8, 0xe2, 0x14, 0xbd,
	0x0f, 0x9d, 0xd2, 0x73, 0xe0, 0x6d, 0xbe, 0xf3, 0x0c, 0x4c, 0x28, 0x85, 0xf6, 0xf6, 0x61, 0xaf,
	0x77, 0xd8, 0xe7, 0x3d, 0xcf, 0xf7, 0x37, 0x77, 0xbc, 0xc5, 0xc6, 0x56, 0x13, 0xe6, 0x8e, 0xc4,
	0x70, 0x20, 0x45, 0xe8, 0xfc, 0x24, 0x40, 0x27, 0xdb, 0x91, 0x0f, 0xb2, 0x28, 0x96, 0xd4, 0x07,
	0x99, 0x7b, 0x58, 0x81, 0xdf, 0x35, 0xe8, 0xfa, 0x8a, 0x4c, 0x8f, 0xaf, 0xc8, 0x5f, 0xac, 0xff,
	0x0f, 0x02, 0x8b, 0xd7, 0x1b, 0xfe, 0x7f, 0x2a, 0xfd, 0x4e, 0x60, 0x61, 0x6c, 0xae, 0x13, 0xd7,
	0xa3, 0x92, 0xdd, 0xb8, 0x45, 0xf6, 0xbf, 0x90, 0xd5, 0xfd, 0x42, 0xc0, 0x3a, 0xc2, 0x4c, 0xd2,
	0x17, 0x30, 0x53, 0x1c, 0x12, 0x4a, 0x4d, 0xe2, 0xfa, 0x55, 0x59, 0x7d, 0x78, 0xeb, 0x6f, 0xf7,
	0x8c, 0xd0, 0xd7, 0xd0, 0x1c, 0x1d, 0x04, 0x6a, 0xdf, 0x10, 0x69, 0x38, 0x56, 0x8c, 0x67, 0xe2,
	0x76, 0x9c, 0xce, 0x16, 0x0a, 0x9f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x69, 0x74, 0xd9, 0x06,
	0xf8, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PeroClient is the client API for Pero service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeroClient interface {
	Entry(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (Pero_EntryClient, error)
	Broadcast(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
}

type peroClient struct {
	cc grpc.ClientConnInterface
}

func NewPeroClient(cc grpc.ClientConnInterface) PeroClient {
	return &peroClient{cc}
}

func (c *peroClient) Entry(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (Pero_EntryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Pero_serviceDesc.Streams[0], "/pero.Pero/Entry", opts...)
	if err != nil {
		return nil, err
	}
	x := &peroEntryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pero_EntryClient interface {
	Recv() (*ChatMessageResponse, error)
	grpc.ClientStream
}

type peroEntryClient struct {
	grpc.ClientStream
}

func (x *peroEntryClient) Recv() (*ChatMessageResponse, error) {
	m := new(ChatMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peroClient) Broadcast(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/pero.Pero/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeroServer is the server API for Pero service.
type PeroServer interface {
	Entry(*EntryRequest, Pero_EntryServer) error
	Broadcast(context.Context, *ChatMessageRequest) (*BroadcastResponse, error)
}

// UnimplementedPeroServer can be embedded to have forward compatible implementations.
type UnimplementedPeroServer struct {
}

func (*UnimplementedPeroServer) Entry(req *EntryRequest, srv Pero_EntryServer) error {
	return status.Errorf(codes.Unimplemented, "method Entry not implemented")
}
func (*UnimplementedPeroServer) Broadcast(ctx context.Context, req *ChatMessageRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func RegisterPeroServer(s *grpc.Server, srv PeroServer) {
	s.RegisterService(&_Pero_serviceDesc, srv)
}

func _Pero_Entry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EntryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeroServer).Entry(m, &peroEntryServer{stream})
}

type Pero_EntryServer interface {
	Send(*ChatMessageResponse) error
	grpc.ServerStream
}

type peroEntryServer struct {
	grpc.ServerStream
}

func (x *peroEntryServer) Send(m *ChatMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Pero_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeroServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pero.Pero/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeroServer).Broadcast(ctx, req.(*ChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pero_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pero.Pero",
	HandlerType: (*PeroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _Pero_Broadcast_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Entry",
			Handler:       _Pero_Entry_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chatService.proto",
}
