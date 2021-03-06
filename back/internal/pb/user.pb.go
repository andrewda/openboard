// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

// openboard
//
// user semantics:
//
//  Add{T}(s)    (Add{T}(s)Req)    returns {T}(s)Resp // POST
//  Ovr{T}(s)    (Ovr{T}(s)Req)    returns {T}(s)Resp // PUT
//  Mod{T}(s)    (Mod{T}(s)Req)    returns {T}(s)Resp // PATCH
//  Get{T}       (Get{T}Req)       returns {T}Resp    // GET
//  Fnd{T}s      (Fnd{T}sReq)      returns {T}sResp   // GET
//  Rmv{T}(s)    (Rmv{T}(s)Req)    returns EmptyResp  // DELETE
//  Unr{T}       (Unr{T}Req)       returns {T}Resp    // PATCH

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RoleResp struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleResp) Reset()         { *m = RoleResp{} }
func (m *RoleResp) String() string { return proto.CompactTextString(m) }
func (*RoleResp) ProtoMessage()    {}
func (*RoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *RoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleResp.Unmarshal(m, b)
}
func (m *RoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleResp.Marshal(b, m, deterministic)
}
func (m *RoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleResp.Merge(m, src)
}
func (m *RoleResp) XXX_Size() int {
	return xxx_messageInfo_RoleResp.Size(m)
}
func (m *RoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_RoleResp proto.InternalMessageInfo

func (m *RoleResp) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RoleResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddRoleReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRoleReq) Reset()         { *m = AddRoleReq{} }
func (m *AddRoleReq) String() string { return proto.CompactTextString(m) }
func (*AddRoleReq) ProtoMessage()    {}
func (*AddRoleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *AddRoleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleReq.Unmarshal(m, b)
}
func (m *AddRoleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleReq.Marshal(b, m, deterministic)
}
func (m *AddRoleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleReq.Merge(m, src)
}
func (m *AddRoleReq) XXX_Size() int {
	return xxx_messageInfo_AddRoleReq.Size(m)
}
func (m *AddRoleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleReq proto.InternalMessageInfo

func (m *AddRoleReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RolesResp struct {
	Items                []*RoleResp `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total                uint32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RolesResp) Reset()         { *m = RolesResp{} }
func (m *RolesResp) String() string { return proto.CompactTextString(m) }
func (*RolesResp) ProtoMessage()    {}
func (*RolesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *RolesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesResp.Unmarshal(m, b)
}
func (m *RolesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesResp.Marshal(b, m, deterministic)
}
func (m *RolesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesResp.Merge(m, src)
}
func (m *RolesResp) XXX_Size() int {
	return xxx_messageInfo_RolesResp.Size(m)
}
func (m *RolesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesResp.DiscardUnknown(m)
}

var xxx_messageInfo_RolesResp proto.InternalMessageInfo

func (m *RolesResp) GetItems() []*RoleResp {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *RolesResp) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type FndRolesReq struct {
	RoleIds              []uint32 `protobuf:"varint,1,rep,packed,name=roleIds,proto3" json:"roleIds,omitempty"`
	RoleNames            []string `protobuf:"bytes,2,rep,name=roleNames,proto3" json:"roleNames,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Lapse                uint32   `protobuf:"varint,4,opt,name=lapse,proto3" json:"lapse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FndRolesReq) Reset()         { *m = FndRolesReq{} }
func (m *FndRolesReq) String() string { return proto.CompactTextString(m) }
func (*FndRolesReq) ProtoMessage()    {}
func (*FndRolesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *FndRolesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FndRolesReq.Unmarshal(m, b)
}
func (m *FndRolesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FndRolesReq.Marshal(b, m, deterministic)
}
func (m *FndRolesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FndRolesReq.Merge(m, src)
}
func (m *FndRolesReq) XXX_Size() int {
	return xxx_messageInfo_FndRolesReq.Size(m)
}
func (m *FndRolesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FndRolesReq.DiscardUnknown(m)
}

var xxx_messageInfo_FndRolesReq proto.InternalMessageInfo

func (m *FndRolesReq) GetRoleIds() []uint32 {
	if m != nil {
		return m.RoleIds
	}
	return nil
}

func (m *FndRolesReq) GetRoleNames() []string {
	if m != nil {
		return m.RoleNames
	}
	return nil
}

func (m *FndRolesReq) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FndRolesReq) GetLapse() uint32 {
	if m != nil {
		return m.Lapse
	}
	return 0
}

type UserResp struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	EmailHold            bool                 `protobuf:"varint,4,opt,name=emailHold,proto3" json:"emailHold,omitempty"`
	Altmail              string               `protobuf:"bytes,5,opt,name=altmail,proto3" json:"altmail,omitempty"`
	AltmailHold          bool                 `protobuf:"varint,6,opt,name=altmailHold,proto3" json:"altmailHold,omitempty"`
	FirstName            string               `protobuf:"bytes,7,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string               `protobuf:"bytes,8,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Avatar               string               `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Roles                []*RoleResp          `protobuf:"bytes,10,rep,name=roles,proto3" json:"roles,omitempty"`
	LastLogin            *timestamp.Timestamp `protobuf:"bytes,11,opt,name=lastLogin,proto3" json:"lastLogin,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,12,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,13,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted              *timestamp.Timestamp `protobuf:"bytes,14,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Blocked              *timestamp.Timestamp `protobuf:"bytes,15,opt,name=blocked,proto3" json:"blocked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserResp) Reset()         { *m = UserResp{} }
func (m *UserResp) String() string { return proto.CompactTextString(m) }
func (*UserResp) ProtoMessage()    {}
func (*UserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResp.Unmarshal(m, b)
}
func (m *UserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResp.Marshal(b, m, deterministic)
}
func (m *UserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResp.Merge(m, src)
}
func (m *UserResp) XXX_Size() int {
	return xxx_messageInfo_UserResp.Size(m)
}
func (m *UserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserResp proto.InternalMessageInfo

func (m *UserResp) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResp) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResp) GetEmailHold() bool {
	if m != nil {
		return m.EmailHold
	}
	return false
}

func (m *UserResp) GetAltmail() string {
	if m != nil {
		return m.Altmail
	}
	return ""
}

func (m *UserResp) GetAltmailHold() bool {
	if m != nil {
		return m.AltmailHold
	}
	return false
}

func (m *UserResp) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserResp) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserResp) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserResp) GetRoles() []*RoleResp {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *UserResp) GetLastLogin() *timestamp.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

func (m *UserResp) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *UserResp) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *UserResp) GetDeleted() *timestamp.Timestamp {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *UserResp) GetBlocked() *timestamp.Timestamp {
	if m != nil {
		return m.Blocked
	}
	return nil
}

type AddUserReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	EmailHold            bool     `protobuf:"varint,3,opt,name=emailHold,proto3" json:"emailHold,omitempty"`
	Altmail              string   `protobuf:"bytes,4,opt,name=altmail,proto3" json:"altmail,omitempty"`
	AltmailHold          bool     `protobuf:"varint,5,opt,name=altmailHold,proto3" json:"altmailHold,omitempty"`
	FirstName            string   `protobuf:"bytes,6,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,7,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Avatar               string   `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Password             string   `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	RoleIds              []uint32 `protobuf:"varint,10,rep,packed,name=roleIds,proto3" json:"roleIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserReq) Reset()         { *m = AddUserReq{} }
func (m *AddUserReq) String() string { return proto.CompactTextString(m) }
func (*AddUserReq) ProtoMessage()    {}
func (*AddUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *AddUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserReq.Unmarshal(m, b)
}
func (m *AddUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserReq.Marshal(b, m, deterministic)
}
func (m *AddUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserReq.Merge(m, src)
}
func (m *AddUserReq) XXX_Size() int {
	return xxx_messageInfo_AddUserReq.Size(m)
}
func (m *AddUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserReq proto.InternalMessageInfo

func (m *AddUserReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddUserReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddUserReq) GetEmailHold() bool {
	if m != nil {
		return m.EmailHold
	}
	return false
}

func (m *AddUserReq) GetAltmail() string {
	if m != nil {
		return m.Altmail
	}
	return ""
}

func (m *AddUserReq) GetAltmailHold() bool {
	if m != nil {
		return m.AltmailHold
	}
	return false
}

func (m *AddUserReq) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *AddUserReq) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *AddUserReq) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *AddUserReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddUserReq) GetRoleIds() []uint32 {
	if m != nil {
		return m.RoleIds
	}
	return nil
}

type OvrUserReq struct {
	Id                   uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Req                  *AddUserReq `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OvrUserReq) Reset()         { *m = OvrUserReq{} }
func (m *OvrUserReq) String() string { return proto.CompactTextString(m) }
func (*OvrUserReq) ProtoMessage()    {}
func (*OvrUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *OvrUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OvrUserReq.Unmarshal(m, b)
}
func (m *OvrUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OvrUserReq.Marshal(b, m, deterministic)
}
func (m *OvrUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OvrUserReq.Merge(m, src)
}
func (m *OvrUserReq) XXX_Size() int {
	return xxx_messageInfo_OvrUserReq.Size(m)
}
func (m *OvrUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OvrUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_OvrUserReq proto.InternalMessageInfo

func (m *OvrUserReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OvrUserReq) GetReq() *AddUserReq {
	if m != nil {
		return m.Req
	}
	return nil
}

type UsersResp struct {
	Items                []*UserResp `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total                uint32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UsersResp) Reset()         { *m = UsersResp{} }
func (m *UsersResp) String() string { return proto.CompactTextString(m) }
func (*UsersResp) ProtoMessage()    {}
func (*UsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResp.Unmarshal(m, b)
}
func (m *UsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResp.Marshal(b, m, deterministic)
}
func (m *UsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResp.Merge(m, src)
}
func (m *UsersResp) XXX_Size() int {
	return xxx_messageInfo_UsersResp.Size(m)
}
func (m *UsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResp proto.InternalMessageInfo

func (m *UsersResp) GetItems() []*UserResp {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *UsersResp) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type FndUsersReq struct {
	RoleIds              []uint32             `protobuf:"varint,1,rep,packed,name=roleIds,proto3" json:"roleIds,omitempty"`
	Emails               []string             `protobuf:"bytes,2,rep,name=emails,proto3" json:"emails,omitempty"`
	EmailHold            bool                 `protobuf:"varint,3,opt,name=emailHold,proto3" json:"emailHold,omitempty"`
	Altmails             []string             `protobuf:"bytes,4,rep,name=altmails,proto3" json:"altmails,omitempty"`
	AltmailHold          bool                 `protobuf:"varint,5,opt,name=altmailHold,proto3" json:"altmailHold,omitempty"`
	CreatedFirst         *timestamp.Timestamp `protobuf:"bytes,6,opt,name=createdFirst,proto3" json:"createdFirst,omitempty"`
	CreatedFinal         *timestamp.Timestamp `protobuf:"bytes,7,opt,name=createdFinal,proto3" json:"createdFinal,omitempty"`
	CreatedDesc          bool                 `protobuf:"varint,8,opt,name=createdDesc,proto3" json:"createdDesc,omitempty"`
	Limit                uint32               `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	Lapse                uint32               `protobuf:"varint,10,opt,name=lapse,proto3" json:"lapse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FndUsersReq) Reset()         { *m = FndUsersReq{} }
func (m *FndUsersReq) String() string { return proto.CompactTextString(m) }
func (*FndUsersReq) ProtoMessage()    {}
func (*FndUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *FndUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FndUsersReq.Unmarshal(m, b)
}
func (m *FndUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FndUsersReq.Marshal(b, m, deterministic)
}
func (m *FndUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FndUsersReq.Merge(m, src)
}
func (m *FndUsersReq) XXX_Size() int {
	return xxx_messageInfo_FndUsersReq.Size(m)
}
func (m *FndUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FndUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_FndUsersReq proto.InternalMessageInfo

func (m *FndUsersReq) GetRoleIds() []uint32 {
	if m != nil {
		return m.RoleIds
	}
	return nil
}

func (m *FndUsersReq) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *FndUsersReq) GetEmailHold() bool {
	if m != nil {
		return m.EmailHold
	}
	return false
}

func (m *FndUsersReq) GetAltmails() []string {
	if m != nil {
		return m.Altmails
	}
	return nil
}

func (m *FndUsersReq) GetAltmailHold() bool {
	if m != nil {
		return m.AltmailHold
	}
	return false
}

func (m *FndUsersReq) GetCreatedFirst() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedFirst
	}
	return nil
}

func (m *FndUsersReq) GetCreatedFinal() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedFinal
	}
	return nil
}

func (m *FndUsersReq) GetCreatedDesc() bool {
	if m != nil {
		return m.CreatedDesc
	}
	return false
}

func (m *FndUsersReq) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FndUsersReq) GetLapse() uint32 {
	if m != nil {
		return m.Lapse
	}
	return 0
}

type RmvUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmvUserResp) Reset()         { *m = RmvUserResp{} }
func (m *RmvUserResp) String() string { return proto.CompactTextString(m) }
func (*RmvUserResp) ProtoMessage()    {}
func (*RmvUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *RmvUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmvUserResp.Unmarshal(m, b)
}
func (m *RmvUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmvUserResp.Marshal(b, m, deterministic)
}
func (m *RmvUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmvUserResp.Merge(m, src)
}
func (m *RmvUserResp) XXX_Size() int {
	return xxx_messageInfo_RmvUserResp.Size(m)
}
func (m *RmvUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RmvUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_RmvUserResp proto.InternalMessageInfo

type RmvUserReq struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmvUserReq) Reset()         { *m = RmvUserReq{} }
func (m *RmvUserReq) String() string { return proto.CompactTextString(m) }
func (*RmvUserReq) ProtoMessage()    {}
func (*RmvUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *RmvUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmvUserReq.Unmarshal(m, b)
}
func (m *RmvUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmvUserReq.Marshal(b, m, deterministic)
}
func (m *RmvUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmvUserReq.Merge(m, src)
}
func (m *RmvUserReq) XXX_Size() int {
	return xxx_messageInfo_RmvUserReq.Size(m)
}
func (m *RmvUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RmvUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_RmvUserReq proto.InternalMessageInfo

func (m *RmvUserReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*RoleResp)(nil), "pb.RoleResp")
	proto.RegisterType((*AddRoleReq)(nil), "pb.AddRoleReq")
	proto.RegisterType((*RolesResp)(nil), "pb.RolesResp")
	proto.RegisterType((*FndRolesReq)(nil), "pb.FndRolesReq")
	proto.RegisterType((*UserResp)(nil), "pb.UserResp")
	proto.RegisterType((*AddUserReq)(nil), "pb.AddUserReq")
	proto.RegisterType((*OvrUserReq)(nil), "pb.OvrUserReq")
	proto.RegisterType((*UsersResp)(nil), "pb.UsersResp")
	proto.RegisterType((*FndUsersReq)(nil), "pb.FndUsersReq")
	proto.RegisterType((*RmvUserResp)(nil), "pb.RmvUserResp")
	proto.RegisterType((*RmvUserReq)(nil), "pb.RmvUserReq")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x8e, 0xd3, 0x4a,
	0x10, 0x7e, 0x76, 0xfe, 0xec, 0xca, 0xdf, 0x53, 0xeb, 0xbd, 0x91, 0x65, 0x8d, 0x84, 0xe5, 0x55,
	0x98, 0x45, 0x22, 0x05, 0x16, 0x30, 0x8b, 0x48, 0x48, 0x30, 0x02, 0x09, 0x81, 0xd4, 0x82, 0x03,
	0x74, 0xc6, 0x3d, 0x23, 0x0b, 0xff, 0xc5, 0xed, 0x0c, 0x0b, 0xc4, 0x86, 0x2b, 0x70, 0x01, 0x4e,
	0xc0, 0x65, 0xb8, 0xc2, 0x6c, 0xb8, 0x05, 0xea, 0xea, 0xee, 0xd8, 0x1e, 0x26, 0xc9, 0xec, 0xfc,
	0x75, 0xd5, 0x57, 0x55, 0xae, 0xfa, 0xba, 0x1a, 0x60, 0x2b, 0x78, 0x39, 0x2f, 0xca, 0xbc, 0xca,
	0x89, 0x5d, 0xac, 0xfd, 0xd3, 0xeb, 0x3c, 0xbf, 0x4e, 0xf8, 0x82, 0x15, 0xf1, 0x82, 0x65, 0x59,
	0x5e, 0xb1, 0x2a, 0xce, 0x33, 0xa1, 0x3c, 0xfc, 0x47, 0xda, 0x8a, 0x68, 0xbd, 0xbd, 0x5a, 0x54,
	0x71, 0xca, 0x45, 0xc5, 0xd2, 0x42, 0x39, 0x84, 0x73, 0x70, 0x68, 0x9e, 0x70, 0xca, 0x45, 0x41,
	0x26, 0x60, 0xc7, 0x91, 0x67, 0x05, 0xd6, 0x6c, 0x4c, 0xed, 0x38, 0x22, 0x04, 0xba, 0x19, 0x4b,
	0xb9, 0x67, 0x07, 0xd6, 0xcc, 0xa5, 0xf8, 0x1d, 0x06, 0x00, 0x2f, 0xa2, 0x48, 0x51, 0x36, 0x3b,
	0x0f, 0xab, 0xe1, 0xf1, 0x0a, 0x5c, 0x69, 0x16, 0x18, 0x32, 0x84, 0x5e, 0x5c, 0xf1, 0x54, 0x78,
	0x56, 0xd0, 0x99, 0x0d, 0x97, 0xa3, 0x79, 0xb1, 0x9e, 0x9b, 0x7c, 0x54, 0x99, 0xc8, 0x7f, 0xd0,
	0xab, 0xf2, 0x8a, 0x25, 0x98, 0x67, 0x4c, 0x15, 0x08, 0x37, 0x30, 0xbc, 0xc8, 0x22, 0x1d, 0x69,
	0x43, 0x3c, 0x18, 0x94, 0x79, 0xc2, 0xdf, 0x44, 0x2a, 0xd4, 0x98, 0x1a, 0x48, 0x4e, 0xc1, 0x95,
	0x9f, 0xef, 0x58, 0xca, 0x85, 0x67, 0x07, 0x9d, 0x99, 0x4b, 0xeb, 0x03, 0x19, 0x3c, 0x89, 0xd3,
	0xb8, 0xf2, 0x3a, 0x2a, 0x38, 0x02, 0x3c, 0x65, 0x85, 0xe0, 0x5e, 0x57, 0x9f, 0x4a, 0x10, 0xfe,
	0xec, 0x82, 0xf3, 0x51, 0xf0, 0xf2, 0xde, 0x66, 0xf8, 0xe0, 0xc8, 0xce, 0x37, 0x1a, 0xb2, 0xc3,
	0x32, 0x1c, 0x4f, 0x59, 0x9c, 0x60, 0x12, 0x97, 0x2a, 0x20, 0x0b, 0xc3, 0x8f, 0xd7, 0x79, 0x12,
	0x61, 0x22, 0x87, 0xd6, 0x07, 0xf2, 0x87, 0x58, 0x52, 0x21, 0xab, 0x87, 0x2c, 0x03, 0x49, 0x00,
	0x43, 0xfd, 0x89, 0xcc, 0x3e, 0x32, 0x9b, 0x47, 0x32, 0xf2, 0x55, 0x5c, 0x8a, 0x4a, 0xfe, 0xa2,
	0x37, 0x40, 0x76, 0x7d, 0x20, 0x2b, 0x4d, 0x98, 0x36, 0x3a, 0xaa, 0x52, 0x83, 0xc9, 0x09, 0xf4,
	0xd9, 0x0d, 0xab, 0x58, 0xe9, 0xb9, 0x68, 0xd1, 0x48, 0xce, 0x49, 0xf6, 0x4c, 0x78, 0x70, 0xdf,
	0x9c, 0xd0, 0x44, 0x9e, 0x81, 0x2b, 0xe3, 0xbc, 0xcd, 0xaf, 0xe3, 0xcc, 0x1b, 0x06, 0xd6, 0x6c,
	0xb8, 0xf4, 0xe7, 0x4a, 0x5f, 0x73, 0xa3, 0xaf, 0xf9, 0x07, 0xa3, 0x2f, 0x5a, 0x3b, 0x93, 0xa7,
	0x30, 0xb8, 0x2c, 0x39, 0xab, 0x78, 0xe4, 0x8d, 0x8e, 0xf2, 0x8c, 0xab, 0x64, 0x6d, 0x8b, 0x08,
	0x59, 0xe3, 0xe3, 0x2c, 0xed, 0x2a, 0x59, 0x11, 0x4f, 0xb8, 0x64, 0x4d, 0x8e, 0xb3, 0xb4, 0xab,
	0x64, 0xad, 0x93, 0xfc, 0xf2, 0x13, 0x8f, 0xbc, 0xe9, 0x71, 0x96, 0x76, 0x0d, 0x7f, 0xd8, 0x78,
	0x1b, 0x94, 0x66, 0x36, 0x2d, 0x89, 0x58, 0xfb, 0x24, 0x62, 0xef, 0x95, 0x48, 0xe7, 0x80, 0x44,
	0xba, 0x07, 0x25, 0xd2, 0x3b, 0x22, 0x91, 0xfe, 0x21, 0x89, 0x0c, 0xf6, 0x4a, 0xc4, 0x69, 0x49,
	0xc4, 0x07, 0xa7, 0x60, 0x42, 0x7c, 0xce, 0xcb, 0x48, 0x8b, 0x67, 0x87, 0x9b, 0xb7, 0x13, 0x5a,
	0xb7, 0x33, 0x5c, 0x01, 0xbc, 0xbf, 0x29, 0x4d, 0x87, 0xee, 0x5e, 0xaa, 0x00, 0x3a, 0x25, 0xdf,
	0x60, 0x4f, 0x86, 0xcb, 0x89, 0x14, 0x5d, 0xdd, 0x4e, 0x2a, 0x4d, 0x72, 0x9b, 0x48, 0xbc, 0x7f,
	0x9b, 0x98, 0x0b, 0x7b, 0x78, 0x9b, 0xfc, 0xb6, 0x71, 0x9d, 0xe8, 0x50, 0x87, 0xd6, 0xc9, 0x09,
	0xf4, 0x71, 0x02, 0x66, 0x97, 0x68, 0x74, 0x64, 0x54, 0x3e, 0x38, 0xba, 0xfb, 0xc2, 0xeb, 0x22,
	0x6f, 0x87, 0x1f, 0x30, 0xac, 0x15, 0x8c, 0xb4, 0xe8, 0x2f, 0xe4, 0x88, 0x70, 0x5e, 0x87, 0x25,
	0xd8, 0xf2, 0x6f, 0xf1, 0x33, 0x96, 0xe0, 0x48, 0x1f, 0xca, 0xcf, 0x18, 0xca, 0x49, 0xe3, 0x97,
	0x5c, 0x5c, 0xe2, 0xdc, 0x1d, 0xda, 0x3c, 0xaa, 0xd7, 0xa8, 0x7b, 0xef, 0x1a, 0x85, 0xe6, 0x1a,
	0x1d, 0xc3, 0x90, 0xa6, 0x37, 0x66, 0x2e, 0xe1, 0x29, 0xc0, 0x0e, 0xfe, 0xa5, 0x80, 0xe5, 0xad,
	0x0d, 0x5d, 0x69, 0x23, 0x8f, 0x61, 0xa0, 0x1f, 0x16, 0x62, 0x84, 0xa0, 0x5f, 0x19, 0xbf, 0xb5,
	0x8d, 0xc2, 0x7f, 0xc8, 0x39, 0x38, 0xe6, 0x69, 0x20, 0x53, 0x69, 0x6b, 0x3c, 0x14, 0xfe, 0xd8,
	0x38, 0xa3, 0x64, 0xc2, 0xc9, 0xb7, 0x5f, 0xb7, 0xdf, 0x6d, 0x87, 0xf4, 0x17, 0x6a, 0x89, 0x3d,
	0xc7, 0x34, 0x98, 0xf1, 0x8e, 0xde, 0xfc, 0x96, 0x9c, 0xc2, 0x7f, 0x91, 0x08, 0x61, 0x6f, 0x21,
	0xef, 0xf0, 0xb9, 0x75, 0x46, 0x56, 0x30, 0xd0, 0x52, 0x56, 0xd4, 0x5a, 0xd7, 0x77, 0xa8, 0xff,
	0x23, 0x75, 0xea, 0x03, 0x52, 0x17, 0x5f, 0xe2, 0xe8, 0xab, 0xe4, 0xab, 0xb2, 0x51, 0x82, 0xbb,
	0xb2, 0x8d, 0x20, 0x55, 0xd9, 0x3b, 0xa5, 0x37, 0xca, 0xde, 0xa2, 0xff, 0x0a, 0x06, 0xba, 0x89,
	0x2a, 0x77, 0xdd, 0x51, 0x7f, 0xda, 0xc2, 0xa2, 0x08, 0x09, 0x72, 0x47, 0x67, 0x8d, 0xf4, 0xeb,
	0x3e, 0x6a, 0xe0, 0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x76, 0x86, 0xa6, 0x3e, 0x08,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*RoleResp, error)
	FndRoles(ctx context.Context, in *FndRolesReq, opts ...grpc.CallOption) (*RolesResp, error)
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*UserResp, error)
	OvrUser(ctx context.Context, in *OvrUserReq, opts ...grpc.CallOption) (*UserResp, error)
	FndUsers(ctx context.Context, in *FndUsersReq, opts ...grpc.CallOption) (*UsersResp, error)
	RmvUser(ctx context.Context, in *RmvUserReq, opts ...grpc.CallOption) (*RmvUserResp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*RoleResp, error) {
	out := new(RoleResp)
	err := c.cc.Invoke(ctx, "/pb.User/AddRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FndRoles(ctx context.Context, in *FndRolesReq, opts ...grpc.CallOption) (*RolesResp, error) {
	out := new(RolesResp)
	err := c.cc.Invoke(ctx, "/pb.User/FndRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*UserResp, error) {
	out := new(UserResp)
	err := c.cc.Invoke(ctx, "/pb.User/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) OvrUser(ctx context.Context, in *OvrUserReq, opts ...grpc.CallOption) (*UserResp, error) {
	out := new(UserResp)
	err := c.cc.Invoke(ctx, "/pb.User/OvrUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FndUsers(ctx context.Context, in *FndUsersReq, opts ...grpc.CallOption) (*UsersResp, error) {
	out := new(UsersResp)
	err := c.cc.Invoke(ctx, "/pb.User/FndUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RmvUser(ctx context.Context, in *RmvUserReq, opts ...grpc.CallOption) (*RmvUserResp, error) {
	out := new(RmvUserResp)
	err := c.cc.Invoke(ctx, "/pb.User/RmvUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	AddRole(context.Context, *AddRoleReq) (*RoleResp, error)
	FndRoles(context.Context, *FndRolesReq) (*RolesResp, error)
	AddUser(context.Context, *AddUserReq) (*UserResp, error)
	OvrUser(context.Context, *OvrUserReq) (*UserResp, error)
	FndUsers(context.Context, *FndUsersReq) (*UsersResp, error)
	RmvUser(context.Context, *RmvUserReq) (*RmvUserResp, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/AddRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddRole(ctx, req.(*AddRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FndRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FndRolesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FndRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/FndRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FndRoles(ctx, req.(*FndRolesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_OvrUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OvrUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).OvrUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/OvrUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).OvrUser(ctx, req.(*OvrUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FndUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FndUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FndUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/FndUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FndUsers(ctx, req.(*FndUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RmvUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RmvUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RmvUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/RmvUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RmvUser(ctx, req.(*RmvUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRole",
			Handler:    _User_AddRole_Handler,
		},
		{
			MethodName: "FndRoles",
			Handler:    _User_FndRoles_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _User_AddUser_Handler,
		},
		{
			MethodName: "OvrUser",
			Handler:    _User_OvrUser_Handler,
		},
		{
			MethodName: "FndUsers",
			Handler:    _User_FndUsers_Handler,
		},
		{
			MethodName: "RmvUser",
			Handler:    _User_RmvUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
