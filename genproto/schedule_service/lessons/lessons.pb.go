// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: lessons.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_proto_msgTypes[0]
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
	return file_lessons_proto_rawDescGZIP(), []int{0}
}

type LessonPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LessonPrimaryKey) Reset() {
	*x = LessonPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LessonPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LessonPrimaryKey) ProtoMessage() {}

func (x *LessonPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LessonPrimaryKey.ProtoReflect.Descriptor instead.
func (*LessonPrimaryKey) Descriptor() ([]byte, []int) {
	return file_lessons_proto_rawDescGZIP(), []int{1}
}

func (x *LessonPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateLesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId string `protobuf:"bytes,2,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *CreateLesson) Reset() {
	*x = CreateLesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLesson) ProtoMessage() {}

func (x *CreateLesson) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLesson.ProtoReflect.Descriptor instead.
func (*CreateLesson) Descriptor() ([]byte, []int) {
	return file_lessons_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLesson) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

type Lesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ScheduleId string `protobuf:"bytes,2,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_lessons_proto_rawDescGZIP(), []int{3}
}

func (x *Lesson) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lesson) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *Lesson) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UpdateLesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ScheduleId string `protobuf:"bytes,2,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *UpdateLesson) Reset() {
	*x = UpdateLesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLesson) ProtoMessage() {}

func (x *UpdateLesson) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLesson.ProtoReflect.Descriptor instead.
func (*UpdateLesson) Descriptor() ([]byte, []int) {
	return file_lessons_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateLesson) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateLesson) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

type GetListLessonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetListLessonRequest) Reset() {
	*x = GetListLessonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListLessonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListLessonRequest) ProtoMessage() {}

func (x *GetListLessonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListLessonRequest.ProtoReflect.Descriptor instead.
func (*GetListLessonRequest) Descriptor() ([]byte, []int) {
	return file_lessons_proto_rawDescGZIP(), []int{5}
}

func (x *GetListLessonRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListLessonRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetListLesssonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int64     `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Lessons []*Lesson `protobuf:"bytes,2,rep,name=Lessons,proto3" json:"Lessons,omitempty"`
}

func (x *GetListLesssonResponse) Reset() {
	*x = GetListLesssonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lessons_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListLesssonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListLesssonResponse) ProtoMessage() {}

func (x *GetListLesssonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lessons_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListLesssonResponse.ProtoReflect.Descriptor instead.
func (*GetListLesssonResponse) Descriptor() ([]byte, []int) {
	return file_lessons_proto_rawDescGZIP(), []int{6}
}

func (x *GetListLesssonResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListLesssonResponse) GetLessons() []*Lesson {
	if x != nil {
		return x.Lessons
	}
	return nil
}

var File_lessons_proto protoreflect.FileDescriptor

var file_lessons_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x22, 0x0a, 0x10, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x3f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x44, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x59, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x32, 0xfd, 0x01, 0x0a, 0x0e, 0x4c, 0x65, 0x73, 0x73, 0x73, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a,
	0x0f, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73,
	0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x1a,
	0x0f, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lessons_proto_rawDescOnce sync.Once
	file_lessons_proto_rawDescData = file_lessons_proto_rawDesc
)

func file_lessons_proto_rawDescGZIP() []byte {
	file_lessons_proto_rawDescOnce.Do(func() {
		file_lessons_proto_rawDescData = protoimpl.X.CompressGZIP(file_lessons_proto_rawDescData)
	})
	return file_lessons_proto_rawDescData
}

var file_lessons_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lessons_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: lessons.Empty
	(*LessonPrimaryKey)(nil),       // 1: lessons.LessonPrimaryKey
	(*CreateLesson)(nil),           // 2: lessons.CreateLesson
	(*Lesson)(nil),                 // 3: lessons.Lesson
	(*UpdateLesson)(nil),           // 4: lessons.UpdateLesson
	(*GetListLessonRequest)(nil),   // 5: lessons.GetListLessonRequest
	(*GetListLesssonResponse)(nil), // 6: lessons.GetListLesssonResponse
}
var file_lessons_proto_depIdxs = []int32{
	3, // 0: lessons.GetListLesssonResponse.Lessons:type_name -> lessons.Lesson
	2, // 1: lessons.LesssonService.Create:input_type -> lessons.CreateLesson
	1, // 2: lessons.LesssonService.GetById:input_type -> lessons.LessonPrimaryKey
	5, // 3: lessons.LesssonService.GetAll:input_type -> lessons.GetListLessonRequest
	4, // 4: lessons.LesssonService.Update:input_type -> lessons.UpdateLesson
	3, // 5: lessons.LesssonService.Create:output_type -> lessons.Lesson
	3, // 6: lessons.LesssonService.GetById:output_type -> lessons.Lesson
	6, // 7: lessons.LesssonService.GetAll:output_type -> lessons.GetListLesssonResponse
	3, // 8: lessons.LesssonService.Update:output_type -> lessons.Lesson
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lessons_proto_init() }
func file_lessons_proto_init() {
	if File_lessons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lessons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_lessons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LessonPrimaryKey); i {
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
		file_lessons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLesson); i {
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
		file_lessons_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
		file_lessons_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLesson); i {
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
		file_lessons_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListLessonRequest); i {
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
		file_lessons_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListLesssonResponse); i {
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
			RawDescriptor: file_lessons_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lessons_proto_goTypes,
		DependencyIndexes: file_lessons_proto_depIdxs,
		MessageInfos:      file_lessons_proto_msgTypes,
	}.Build()
	File_lessons_proto = out.File
	file_lessons_proto_rawDesc = nil
	file_lessons_proto_goTypes = nil
	file_lessons_proto_depIdxs = nil
}
