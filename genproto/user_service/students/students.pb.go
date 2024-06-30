// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: students.proto

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
		mi := &file_students_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_students_proto_msgTypes[0]
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
	return file_students_proto_rawDescGZIP(), []int{0}
}

type StudentPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StudentPrimaryKey) Reset() {
	*x = StudentPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_students_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentPrimaryKey) ProtoMessage() {}

func (x *StudentPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_students_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentPrimaryKey.ProtoReflect.Descriptor instead.
func (*StudentPrimaryKey) Descriptor() ([]byte, []int) {
	return file_students_proto_rawDescGZIP(), []int{1}
}

func (x *StudentPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateStudent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName    string  `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone       string  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password    string  `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PaidSum     float64 `protobuf:"fixed64,4,opt,name=paid_sum,json=paidSum,proto3" json:"paid_sum,omitempty"`
	CourseCount int32   `protobuf:"varint,5,opt,name=course_count,json=courseCount,proto3" json:"course_count,omitempty"`
	TotalSum    float64 `protobuf:"fixed64,6,opt,name=total_sum,json=totalSum,proto3" json:"total_sum,omitempty"`
	BranchId    string  `protobuf:"bytes,7,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	GroupId     string  `protobuf:"bytes,8,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *CreateStudent) Reset() {
	*x = CreateStudent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_students_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudent) ProtoMessage() {}

func (x *CreateStudent) ProtoReflect() protoreflect.Message {
	mi := &file_students_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudent.ProtoReflect.Descriptor instead.
func (*CreateStudent) Descriptor() ([]byte, []int) {
	return file_students_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStudent) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateStudent) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateStudent) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateStudent) GetPaidSum() float64 {
	if x != nil {
		return x.PaidSum
	}
	return 0
}

func (x *CreateStudent) GetCourseCount() int32 {
	if x != nil {
		return x.CourseCount
	}
	return 0
}

func (x *CreateStudent) GetTotalSum() float64 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

func (x *CreateStudent) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *CreateStudent) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExtraId     string  `protobuf:"bytes,2,opt,name=extra_id,json=extraId,proto3" json:"extra_id,omitempty"`
	FullName    string  `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone       string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Password    string  `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	PaidSum     float64 `protobuf:"fixed64,6,opt,name=paid_sum,json=paidSum,proto3" json:"paid_sum,omitempty"`
	CourseCount int32   `protobuf:"varint,7,opt,name=course_count,json=courseCount,proto3" json:"course_count,omitempty"`
	TotalSum    float64 `protobuf:"fixed64,8,opt,name=total_sum,json=totalSum,proto3" json:"total_sum,omitempty"`
	BranchId    string  `protobuf:"bytes,9,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	GroupId     string  `protobuf:"bytes,10,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	IsActive    int64   `protobuf:"varint,11,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt   string  `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_students_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_students_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_students_proto_rawDescGZIP(), []int{3}
}

func (x *Student) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Student) GetExtraId() string {
	if x != nil {
		return x.ExtraId
	}
	return ""
}

func (x *Student) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Student) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Student) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Student) GetPaidSum() float64 {
	if x != nil {
		return x.PaidSum
	}
	return 0
}

func (x *Student) GetCourseCount() int32 {
	if x != nil {
		return x.CourseCount
	}
	return 0
}

func (x *Student) GetTotalSum() float64 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

func (x *Student) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Student) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Student) GetIsActive() int64 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

func (x *Student) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UpdateStudent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName    string  `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone       string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password    string  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	PaidSum     float64 `protobuf:"fixed64,5,opt,name=paid_sum,json=paidSum,proto3" json:"paid_sum,omitempty"`
	CourseCount int32   `protobuf:"varint,6,opt,name=course_count,json=courseCount,proto3" json:"course_count,omitempty"`
	TotalSum    float64 `protobuf:"fixed64,7,opt,name=total_sum,json=totalSum,proto3" json:"total_sum,omitempty"`
	BranchId    string  `protobuf:"bytes,8,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	GroupId     string  `protobuf:"bytes,9,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *UpdateStudent) Reset() {
	*x = UpdateStudent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_students_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudent) ProtoMessage() {}

func (x *UpdateStudent) ProtoReflect() protoreflect.Message {
	mi := &file_students_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudent.ProtoReflect.Descriptor instead.
func (*UpdateStudent) Descriptor() ([]byte, []int) {
	return file_students_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStudent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStudent) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateStudent) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateStudent) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateStudent) GetPaidSum() float64 {
	if x != nil {
		return x.PaidSum
	}
	return 0
}

func (x *UpdateStudent) GetCourseCount() int32 {
	if x != nil {
		return x.CourseCount
	}
	return 0
}

func (x *UpdateStudent) GetTotalSum() float64 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

func (x *UpdateStudent) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UpdateStudent) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetListStudentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListStudentsRequest) Reset() {
	*x = GetListStudentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_students_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListStudentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListStudentsRequest) ProtoMessage() {}

func (x *GetListStudentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_students_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListStudentsRequest.ProtoReflect.Descriptor instead.
func (*GetListStudentsRequest) Descriptor() ([]byte, []int) {
	return file_students_proto_rawDescGZIP(), []int{5}
}

func (x *GetListStudentsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListStudentsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListStudentsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListStudentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64      `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Students []*Student `protobuf:"bytes,2,rep,name=Students,proto3" json:"Students,omitempty"`
}

func (x *GetListStudentsResponse) Reset() {
	*x = GetListStudentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_students_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListStudentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListStudentsResponse) ProtoMessage() {}

func (x *GetListStudentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_students_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListStudentsResponse.ProtoReflect.Descriptor instead.
func (*GetListStudentsResponse) Descriptor() ([]byte, []int) {
	return file_students_proto_rawDescGZIP(), []int{6}
}

func (x *GetListStudentsResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListStudentsResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

var File_students_proto protoreflect.FileDescriptor

var file_students_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf1, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x69,
	0x64, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x61, 0x69,
	0x64, 0x53, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xd2, 0x02, 0x0a,
	0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x61, 0x69, 0x64, 0x53, 0x75, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x81, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x61, 0x69, 0x64, 0x53, 0x75, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xc8, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x1a, 0x11, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x0f, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_students_proto_rawDescOnce sync.Once
	file_students_proto_rawDescData = file_students_proto_rawDesc
)

func file_students_proto_rawDescGZIP() []byte {
	file_students_proto_rawDescOnce.Do(func() {
		file_students_proto_rawDescData = protoimpl.X.CompressGZIP(file_students_proto_rawDescData)
	})
	return file_students_proto_rawDescData
}

var file_students_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_students_proto_goTypes = []interface{}{
	(*Empty)(nil),                   // 0: students.Empty
	(*StudentPrimaryKey)(nil),       // 1: students.StudentPrimaryKey
	(*CreateStudent)(nil),           // 2: students.CreateStudent
	(*Student)(nil),                 // 3: students.Student
	(*UpdateStudent)(nil),           // 4: students.UpdateStudent
	(*GetListStudentsRequest)(nil),  // 5: students.GetListStudentsRequest
	(*GetListStudentsResponse)(nil), // 6: students.GetListStudentsResponse
}
var file_students_proto_depIdxs = []int32{
	3, // 0: students.GetListStudentsResponse.Students:type_name -> students.Student
	2, // 1: students.StudentService.Create:input_type -> students.CreateStudent
	1, // 2: students.StudentService.GetById:input_type -> students.StudentPrimaryKey
	5, // 3: students.StudentService.GetAll:input_type -> students.GetListStudentsRequest
	4, // 4: students.StudentService.Update:input_type -> students.UpdateStudent
	1, // 5: students.StudentService.Delete:input_type -> students.StudentPrimaryKey
	3, // 6: students.StudentService.Create:output_type -> students.Student
	3, // 7: students.StudentService.GetById:output_type -> students.Student
	6, // 8: students.StudentService.GetAll:output_type -> students.GetListStudentsResponse
	3, // 9: students.StudentService.Update:output_type -> students.Student
	0, // 10: students.StudentService.Delete:output_type -> students.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_students_proto_init() }
func file_students_proto_init() {
	if File_students_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_students_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_students_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentPrimaryKey); i {
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
		file_students_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudent); i {
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
		file_students_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_students_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudent); i {
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
		file_students_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListStudentsRequest); i {
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
		file_students_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListStudentsResponse); i {
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
			RawDescriptor: file_students_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_students_proto_goTypes,
		DependencyIndexes: file_students_proto_depIdxs,
		MessageInfos:      file_students_proto_msgTypes,
	}.Build()
	File_students_proto = out.File
	file_students_proto_rawDesc = nil
	file_students_proto_goTypes = nil
	file_students_proto_depIdxs = nil
}
