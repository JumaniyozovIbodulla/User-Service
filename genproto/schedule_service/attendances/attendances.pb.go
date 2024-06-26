// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: attendances.proto

package schedule_service

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

type AttendType int32

const (
	AttendType_attended AttendType = 0
	AttendType_absent   AttendType = 1
	AttendType_late     AttendType = 2
)

// Enum value maps for AttendType.
var (
	AttendType_name = map[int32]string{
		0: "attended",
		1: "absent",
		2: "late",
	}
	AttendType_value = map[string]int32{
		"attended": 0,
		"absent":   1,
		"late":     2,
	}
)

func (x AttendType) Enum() *AttendType {
	p := new(AttendType)
	*p = x
	return p
}

func (x AttendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttendType) Descriptor() protoreflect.EnumDescriptor {
	return file_attendances_proto_enumTypes[0].Descriptor()
}

func (AttendType) Type() protoreflect.EnumType {
	return &file_attendances_proto_enumTypes[0]
}

func (x AttendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttendType.Descriptor instead.
func (AttendType) EnumDescriptor() ([]byte, []int) {
	return file_attendances_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendances_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_attendances_proto_msgTypes[0]
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
	return file_attendances_proto_rawDescGZIP(), []int{0}
}

type AttendancePrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AttendancePrimaryKey) Reset() {
	*x = AttendancePrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendances_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttendancePrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttendancePrimaryKey) ProtoMessage() {}

func (x *AttendancePrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_attendances_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttendancePrimaryKey.ProtoReflect.Descriptor instead.
func (*AttendancePrimaryKey) Descriptor() ([]byte, []int) {
	return file_attendances_proto_rawDescGZIP(), []int{1}
}

func (x *AttendancePrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateAttendance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId   string     `protobuf:"bytes,2,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	StudentId  string     `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Status     AttendType `protobuf:"varint,4,opt,name=status,proto3,enum=attendances.AttendType" json:"status,omitempty"`
	LateMinute int32      `protobuf:"varint,5,opt,name=late_minute,json=lateMinute,proto3" json:"late_minute,omitempty"`
}

func (x *CreateAttendance) Reset() {
	*x = CreateAttendance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendances_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttendance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttendance) ProtoMessage() {}

func (x *CreateAttendance) ProtoReflect() protoreflect.Message {
	mi := &file_attendances_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttendance.ProtoReflect.Descriptor instead.
func (*CreateAttendance) Descriptor() ([]byte, []int) {
	return file_attendances_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAttendance) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *CreateAttendance) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateAttendance) GetStatus() AttendType {
	if x != nil {
		return x.Status
	}
	return AttendType_attended
}

func (x *CreateAttendance) GetLateMinute() int32 {
	if x != nil {
		return x.LateMinute
	}
	return 0
}

type Attendance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LessonId   string     `protobuf:"bytes,2,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	StudentId  string     `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Status     AttendType `protobuf:"varint,4,opt,name=status,proto3,enum=attendances.AttendType" json:"status,omitempty"`
	LateMinute int32      `protobuf:"varint,5,opt,name=late_minute,json=lateMinute,proto3" json:"late_minute,omitempty"`
}

func (x *Attendance) Reset() {
	*x = Attendance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendances_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attendance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attendance) ProtoMessage() {}

func (x *Attendance) ProtoReflect() protoreflect.Message {
	mi := &file_attendances_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attendance.ProtoReflect.Descriptor instead.
func (*Attendance) Descriptor() ([]byte, []int) {
	return file_attendances_proto_rawDescGZIP(), []int{3}
}

func (x *Attendance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attendance) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *Attendance) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *Attendance) GetStatus() AttendType {
	if x != nil {
		return x.Status
	}
	return AttendType_attended
}

func (x *Attendance) GetLateMinute() int32 {
	if x != nil {
		return x.LateMinute
	}
	return 0
}

type UpdateAttendance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LessonId   string     `protobuf:"bytes,2,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	StudentId  string     `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Status     AttendType `protobuf:"varint,4,opt,name=status,proto3,enum=attendances.AttendType" json:"status,omitempty"`
	LateMinute int32      `protobuf:"varint,5,opt,name=late_minute,json=lateMinute,proto3" json:"late_minute,omitempty"`
}

func (x *UpdateAttendance) Reset() {
	*x = UpdateAttendance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendances_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAttendance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAttendance) ProtoMessage() {}

func (x *UpdateAttendance) ProtoReflect() protoreflect.Message {
	mi := &file_attendances_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAttendance.ProtoReflect.Descriptor instead.
func (*UpdateAttendance) Descriptor() ([]byte, []int) {
	return file_attendances_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAttendance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAttendance) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *UpdateAttendance) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *UpdateAttendance) GetStatus() AttendType {
	if x != nil {
		return x.Status
	}
	return AttendType_attended
}

func (x *UpdateAttendance) GetLateMinute() int32 {
	if x != nil {
		return x.LateMinute
	}
	return 0
}

type GetListAttendanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetListAttendanceRequest) Reset() {
	*x = GetListAttendanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendances_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAttendanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAttendanceRequest) ProtoMessage() {}

func (x *GetListAttendanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attendances_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAttendanceRequest.ProtoReflect.Descriptor instead.
func (*GetListAttendanceRequest) Descriptor() ([]byte, []int) {
	return file_attendances_proto_rawDescGZIP(), []int{5}
}

func (x *GetListAttendanceRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListAttendanceRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetListAttendanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int64         `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Attendances []*Attendance `protobuf:"bytes,2,rep,name=Attendances,proto3" json:"Attendances,omitempty"`
}

func (x *GetListAttendanceResponse) Reset() {
	*x = GetListAttendanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attendances_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAttendanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAttendanceResponse) ProtoMessage() {}

func (x *GetListAttendanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attendances_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAttendanceResponse.ProtoReflect.Descriptor instead.
func (*GetListAttendanceResponse) Descriptor() ([]byte, []int) {
	return file_attendances_proto_rawDescGZIP(), []int{6}
}

func (x *GetListAttendanceResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListAttendanceResponse) GetAttendances() []*Attendance {
	if x != nil {
		return x.Attendances
	}
	return nil
}

var File_attendances_proto protoreflect.FileDescriptor

var file_attendances_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x41, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xa1, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x4d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x74, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x74,
	0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0x48, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x6c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2a,
	0x31, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0c,
	0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x61, 0x62, 0x73, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x6c, 0x61, 0x74, 0x65,
	0x10, 0x02, 0x32, 0x82, 0x03, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63,
	0x65, 0x1a, 0x17, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x17, 0x2e, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x25, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x17, 0x2e, 0x61, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x1a, 0x12, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attendances_proto_rawDescOnce sync.Once
	file_attendances_proto_rawDescData = file_attendances_proto_rawDesc
)

func file_attendances_proto_rawDescGZIP() []byte {
	file_attendances_proto_rawDescOnce.Do(func() {
		file_attendances_proto_rawDescData = protoimpl.X.CompressGZIP(file_attendances_proto_rawDescData)
	})
	return file_attendances_proto_rawDescData
}

var file_attendances_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_attendances_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_attendances_proto_goTypes = []interface{}{
	(AttendType)(0),                   // 0: attendances.Attend_type
	(*Empty)(nil),                     // 1: attendances.Empty
	(*AttendancePrimaryKey)(nil),      // 2: attendances.AttendancePrimaryKey
	(*CreateAttendance)(nil),          // 3: attendances.CreateAttendance
	(*Attendance)(nil),                // 4: attendances.Attendance
	(*UpdateAttendance)(nil),          // 5: attendances.UpdateAttendance
	(*GetListAttendanceRequest)(nil),  // 6: attendances.GetListAttendanceRequest
	(*GetListAttendanceResponse)(nil), // 7: attendances.GetListAttendanceResponse
}
var file_attendances_proto_depIdxs = []int32{
	0, // 0: attendances.CreateAttendance.status:type_name -> attendances.Attend_type
	0, // 1: attendances.Attendance.status:type_name -> attendances.Attend_type
	0, // 2: attendances.UpdateAttendance.status:type_name -> attendances.Attend_type
	4, // 3: attendances.GetListAttendanceResponse.Attendances:type_name -> attendances.Attendance
	3, // 4: attendances.AttendanceService.Create:input_type -> attendances.CreateAttendance
	2, // 5: attendances.AttendanceService.GetById:input_type -> attendances.AttendancePrimaryKey
	6, // 6: attendances.AttendanceService.GetAll:input_type -> attendances.GetListAttendanceRequest
	5, // 7: attendances.AttendanceService.Update:input_type -> attendances.UpdateAttendance
	2, // 8: attendances.AttendanceService.Delete:input_type -> attendances.AttendancePrimaryKey
	4, // 9: attendances.AttendanceService.Create:output_type -> attendances.Attendance
	4, // 10: attendances.AttendanceService.GetById:output_type -> attendances.Attendance
	7, // 11: attendances.AttendanceService.GetAll:output_type -> attendances.GetListAttendanceResponse
	4, // 12: attendances.AttendanceService.Update:output_type -> attendances.Attendance
	1, // 13: attendances.AttendanceService.Delete:output_type -> attendances.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_attendances_proto_init() }
func file_attendances_proto_init() {
	if File_attendances_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attendances_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_attendances_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttendancePrimaryKey); i {
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
		file_attendances_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttendance); i {
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
		file_attendances_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attendance); i {
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
		file_attendances_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAttendance); i {
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
		file_attendances_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAttendanceRequest); i {
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
		file_attendances_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAttendanceResponse); i {
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
			RawDescriptor: file_attendances_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attendances_proto_goTypes,
		DependencyIndexes: file_attendances_proto_depIdxs,
		EnumInfos:         file_attendances_proto_enumTypes,
		MessageInfos:      file_attendances_proto_msgTypes,
	}.Build()
	File_attendances_proto = out.File
	file_attendances_proto_rawDesc = nil
	file_attendances_proto_goTypes = nil
	file_attendances_proto_depIdxs = nil
}
