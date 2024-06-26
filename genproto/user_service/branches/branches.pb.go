// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: branches.proto

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[1]
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
	return file_branches_proto_rawDescGZIP(), []int{1}
}

type BranchePrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BranchePrimaryKey) Reset() {
	*x = BranchePrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchePrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchePrimaryKey) ProtoMessage() {}

func (x *BranchePrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchePrimaryKey.ProtoReflect.Descriptor instead.
func (*BranchePrimaryKey) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{2}
}

func (x *BranchePrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location     *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	SuperAdminId string    `protobuf:"bytes,3,opt,name=super_admin_id,json=superAdminId,proto3" json:"super_admin_id,omitempty"`
}

func (x *CreateBranch) Reset() {
	*x = CreateBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBranch) ProtoMessage() {}

func (x *CreateBranch) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBranch.ProtoReflect.Descriptor instead.
func (*CreateBranch) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBranch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBranch) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *CreateBranch) GetSuperAdminId() string {
	if x != nil {
		return x.SuperAdminId
	}
	return ""
}

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location     *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	SuperAdminId string    `protobuf:"bytes,4,opt,name=super_admin_id,json=superAdminId,proto3" json:"super_admin_id,omitempty"`
	IsActive     int64     `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt    string    `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{4}
}

func (x *Branch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Branch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Branch) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Branch) GetSuperAdminId() string {
	if x != nil {
		return x.SuperAdminId
	}
	return ""
}

func (x *Branch) GetIsActive() int64 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

func (x *Branch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UpdateBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location     *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	SuperAdminId string    `protobuf:"bytes,4,opt,name=super_admin_id,json=superAdminId,proto3" json:"super_admin_id,omitempty"`
	IsActive     int64     `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *UpdateBranch) Reset() {
	*x = UpdateBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranch) ProtoMessage() {}

func (x *UpdateBranch) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranch.ProtoReflect.Descriptor instead.
func (*UpdateBranch) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBranch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBranch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBranch) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *UpdateBranch) GetSuperAdminId() string {
	if x != nil {
		return x.SuperAdminId
	}
	return ""
}

func (x *UpdateBranch) GetIsActive() int64 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

type GetListBranchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListBranchesRequest) Reset() {
	*x = GetListBranchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchesRequest) ProtoMessage() {}

func (x *GetListBranchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchesRequest.ProtoReflect.Descriptor instead.
func (*GetListBranchesRequest) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{6}
}

func (x *GetListBranchesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListBranchesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListBranchesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListBranchesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64     `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Branches []*Branch `protobuf:"bytes,2,rep,name=Branches,proto3" json:"Branches,omitempty"`
}

func (x *GetListBranchesResponse) Reset() {
	*x = GetListBranchesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchesResponse) ProtoMessage() {}

func (x *GetListBranchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchesResponse.ProtoReflect.Descriptor instead.
func (*GetListBranchesResponse) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{7}
}

func (x *GetListBranchesResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListBranchesResponse) GetBranches() []*Branch {
	if x != nil {
		return x.Branches
	}
	return nil
}

var File_branches_proto protoreflect.FileDescriptor

var file_branches_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x78,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xbe, 0x01, 0x0a, 0x06, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x22, 0x5d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73,
	0x32, 0xc4, 0x02, 0x0a, 0x0f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x20, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x1a, 0x0f, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branches_proto_rawDescOnce sync.Once
	file_branches_proto_rawDescData = file_branches_proto_rawDesc
)

func file_branches_proto_rawDescGZIP() []byte {
	file_branches_proto_rawDescOnce.Do(func() {
		file_branches_proto_rawDescData = protoimpl.X.CompressGZIP(file_branches_proto_rawDescData)
	})
	return file_branches_proto_rawDescData
}

var file_branches_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_branches_proto_goTypes = []interface{}{
	(*Location)(nil),                // 0: branches.Location
	(*Empty)(nil),                   // 1: branches.Empty
	(*BranchePrimaryKey)(nil),       // 2: branches.BranchePrimaryKey
	(*CreateBranch)(nil),            // 3: branches.CreateBranch
	(*Branch)(nil),                  // 4: branches.Branch
	(*UpdateBranch)(nil),            // 5: branches.UpdateBranch
	(*GetListBranchesRequest)(nil),  // 6: branches.GetListBranchesRequest
	(*GetListBranchesResponse)(nil), // 7: branches.GetListBranchesResponse
}
var file_branches_proto_depIdxs = []int32{
	0, // 0: branches.CreateBranch.location:type_name -> branches.Location
	0, // 1: branches.Branch.location:type_name -> branches.Location
	0, // 2: branches.UpdateBranch.location:type_name -> branches.Location
	4, // 3: branches.GetListBranchesResponse.Branches:type_name -> branches.Branch
	3, // 4: branches.BranchesService.Create:input_type -> branches.CreateBranch
	2, // 5: branches.BranchesService.GetById:input_type -> branches.BranchePrimaryKey
	6, // 6: branches.BranchesService.GetAll:input_type -> branches.GetListBranchesRequest
	5, // 7: branches.BranchesService.Update:input_type -> branches.UpdateBranch
	2, // 8: branches.BranchesService.Delete:input_type -> branches.BranchePrimaryKey
	4, // 9: branches.BranchesService.Create:output_type -> branches.Branch
	4, // 10: branches.BranchesService.GetById:output_type -> branches.Branch
	7, // 11: branches.BranchesService.GetAll:output_type -> branches.GetListBranchesResponse
	4, // 12: branches.BranchesService.Update:output_type -> branches.Branch
	1, // 13: branches.BranchesService.Delete:output_type -> branches.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_branches_proto_init() }
func file_branches_proto_init() {
	if File_branches_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_branches_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_branches_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_branches_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchePrimaryKey); i {
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
		file_branches_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBranch); i {
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
		file_branches_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
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
		file_branches_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBranch); i {
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
		file_branches_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchesRequest); i {
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
		file_branches_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchesResponse); i {
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
			RawDescriptor: file_branches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_branches_proto_goTypes,
		DependencyIndexes: file_branches_proto_depIdxs,
		MessageInfos:      file_branches_proto_msgTypes,
	}.Build()
	File_branches_proto = out.File
	file_branches_proto_rawDesc = nil
	file_branches_proto_goTypes = nil
	file_branches_proto_depIdxs = nil
}
