// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: managers.proto

package user_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_managers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_managers_proto_rawDescGZIP(), []int{0}
}

type ManagerPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ManagerPrimaryKey) Reset() {
	*x = ManagerPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerPrimaryKey) ProtoMessage() {}

func (x *ManagerPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_managers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerPrimaryKey.ProtoReflect.Descriptor instead.
func (*ManagerPrimaryKey) Descriptor() ([]byte, []int) {
	return file_managers_proto_rawDescGZIP(), []int{1}
}

func (x *ManagerPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName     string  `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone        string  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password     string  `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary       float64 `protobuf:"fixed64,4,opt,name=salary,proto3" json:"salary,omitempty"`
	SuperAdminId string  `protobuf:"bytes,5,opt,name=super_admin_id,json=superAdminId,proto3" json:"super_admin_id,omitempty"`
	BranchId     string  `protobuf:"bytes,6,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *CreateManager) Reset() {
	*x = CreateManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManager) ProtoMessage() {}

func (x *CreateManager) ProtoReflect() protoreflect.Message {
	mi := &file_managers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManager.ProtoReflect.Descriptor instead.
func (*CreateManager) Descriptor() ([]byte, []int) {
	return file_managers_proto_rawDescGZIP(), []int{2}
}

func (x *CreateManager) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateManager) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateManager) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateManager) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *CreateManager) GetSuperAdminId() string {
	if x != nil {
		return x.SuperAdminId
	}
	return ""
}

func (x *CreateManager) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type Manager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExtraId      string  `protobuf:"bytes,2,opt,name=extra_id,json=extraId,proto3" json:"extra_id,omitempty"`
	FullName     string  `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone        string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Password     string  `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Salary       float64 `protobuf:"fixed64,6,opt,name=salary,proto3" json:"salary,omitempty"`
	SuperAdminId string  `protobuf:"bytes,7,opt,name=super_admin_id,json=superAdminId,proto3" json:"super_admin_id,omitempty"`
	BranchId     string  `protobuf:"bytes,8,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	IsActive     int64   `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt    string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Manager) Reset() {
	*x = Manager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manager) ProtoMessage() {}

func (x *Manager) ProtoReflect() protoreflect.Message {
	mi := &file_managers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manager.ProtoReflect.Descriptor instead.
func (*Manager) Descriptor() ([]byte, []int) {
	return file_managers_proto_rawDescGZIP(), []int{3}
}

func (x *Manager) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Manager) GetExtraId() string {
	if x != nil {
		return x.ExtraId
	}
	return ""
}

func (x *Manager) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Manager) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Manager) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Manager) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *Manager) GetSuperAdminId() string {
	if x != nil {
		return x.SuperAdminId
	}
	return ""
}

func (x *Manager) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Manager) GetIsActive() int64 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

func (x *Manager) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UpdateManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName     string  `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone        string  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password     string  `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary       float64 `protobuf:"fixed64,4,opt,name=salary,proto3" json:"salary,omitempty"`
	SuperAdminId string  `protobuf:"bytes,5,opt,name=super_admin_id,json=superAdminId,proto3" json:"super_admin_id,omitempty"`
	BranchId     string  `protobuf:"bytes,6,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *UpdateManager) Reset() {
	*x = UpdateManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManager) ProtoMessage() {}

func (x *UpdateManager) ProtoReflect() protoreflect.Message {
	mi := &file_managers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManager.ProtoReflect.Descriptor instead.
func (*UpdateManager) Descriptor() ([]byte, []int) {
	return file_managers_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateManager) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateManager) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateManager) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateManager) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *UpdateManager) GetSuperAdminId() string {
	if x != nil {
		return x.SuperAdminId
	}
	return ""
}

func (x *UpdateManager) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type GetListManagersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListManagersRequest) Reset() {
	*x = GetListManagersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListManagersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListManagersRequest) ProtoMessage() {}

func (x *GetListManagersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListManagersRequest.ProtoReflect.Descriptor instead.
func (*GetListManagersRequest) Descriptor() ([]byte, []int) {
	return file_managers_proto_rawDescGZIP(), []int{5}
}

func (x *GetListManagersRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListManagersRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListManagersRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListManagersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64      `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Managers []*Manager `protobuf:"bytes,2,rep,name=Managers,proto3" json:"Managers,omitempty"`
}

func (x *GetListManagersResponse) Reset() {
	*x = GetListManagersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListManagersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListManagersResponse) ProtoMessage() {}

func (x *GetListManagersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListManagersResponse.ProtoReflect.Descriptor instead.
func (*GetListManagersResponse) Descriptor() ([]byte, []int) {
	return file_managers_proto_rawDescGZIP(), []int{6}
}

func (x *GetListManagersResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListManagersResponse) GetManagers() []*Manager {
	if x != nil {
		return x.Managers
	}
	return nil
}

var File_managers_proto protoreflect.FileDescriptor

var file_managers_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x49, 0x64, 0x22, 0x9a, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x5e, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5e, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x08, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x32, 0xc9, 0x02,
	0x0a, 0x0f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x20, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managers_proto_rawDescOnce sync.Once
	file_managers_proto_rawDescData = file_managers_proto_rawDesc
)

func file_managers_proto_rawDescGZIP() []byte {
	file_managers_proto_rawDescOnce.Do(func() {
		file_managers_proto_rawDescData = protoimpl.X.CompressGZIP(file_managers_proto_rawDescData)
	})
	return file_managers_proto_rawDescData
}

var file_managers_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_managers_proto_goTypes = []interface{}{
	(*Empty)(nil),                   // 0: managers.Empty
	(*ManagerPrimaryKey)(nil),       // 1: managers.ManagerPrimaryKey
	(*CreateManager)(nil),           // 2: managers.CreateManager
	(*Manager)(nil),                 // 3: managers.Manager
	(*UpdateManager)(nil),           // 4: managers.UpdateManager
	(*GetListManagersRequest)(nil),  // 5: managers.GetListManagersRequest
	(*GetListManagersResponse)(nil), // 6: managers.GetListManagersResponse
}
var file_managers_proto_depIdxs = []int32{
	3, // 0: managers.GetListManagersResponse.Managers:type_name -> managers.Manager
	2, // 1: managers.ManagersService.Create:input_type -> managers.CreateManager
	1, // 2: managers.ManagersService.GetById:input_type -> managers.ManagerPrimaryKey
	5, // 3: managers.ManagersService.GetAll:input_type -> managers.GetListManagersRequest
	4, // 4: managers.ManagersService.Update:input_type -> managers.UpdateManager
	1, // 5: managers.ManagersService.Delete:input_type -> managers.ManagerPrimaryKey
	3, // 6: managers.ManagersService.Create:output_type -> managers.Manager
	3, // 7: managers.ManagersService.GetById:output_type -> managers.Manager
	6, // 8: managers.ManagersService.GetAll:output_type -> managers.GetListManagersResponse
	3, // 9: managers.ManagersService.Update:output_type -> managers.Manager
	0, // 10: managers.ManagersService.Delete:output_type -> managers.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_managers_proto_init() }
func file_managers_proto_init() {
	if File_managers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerPrimaryKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManager); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manager); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManager); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListManagersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListManagersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managers_proto_goTypes,
		DependencyIndexes: file_managers_proto_depIdxs,
		MessageInfos:      file_managers_proto_msgTypes,
	}.Build()
	File_managers_proto = out.File
	file_managers_proto_rawDesc = nil
	file_managers_proto_goTypes = nil
	file_managers_proto_depIdxs = nil
}
