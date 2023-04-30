// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: benchmark/common2/benchmark.proto

package benchmark

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

type ResumeProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email        string                   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone        int64                    `protobuf:"varint,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Address      *ResumeProtoAddress      `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
	PersonalInfo *ResumeProtoPersonalInfo `protobuf:"bytes,5,opt,name=PersonalInfo,proto3" json:"PersonalInfo,omitempty"`
}

func (x *ResumeProto) Reset() {
	*x = ResumeProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_common2_benchmark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeProto) ProtoMessage() {}

func (x *ResumeProto) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_common2_benchmark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeProto.ProtoReflect.Descriptor instead.
func (*ResumeProto) Descriptor() ([]byte, []int) {
	return file_benchmark_common2_benchmark_proto_rawDescGZIP(), []int{0}
}

func (x *ResumeProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResumeProto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResumeProto) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *ResumeProto) GetAddress() *ResumeProtoAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ResumeProto) GetPersonalInfo() *ResumeProtoPersonalInfo {
	if x != nil {
		return x.PersonalInfo
	}
	return nil
}

type ResumeProtoAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street string `protobuf:"bytes,1,opt,name=Street,proto3" json:"Street,omitempty"`
	City   string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	State  string `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	Zip    int32  `protobuf:"varint,4,opt,name=Zip,proto3" json:"Zip,omitempty"`
}

func (x *ResumeProtoAddress) Reset() {
	*x = ResumeProtoAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_common2_benchmark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeProtoAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeProtoAddress) ProtoMessage() {}

func (x *ResumeProtoAddress) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_common2_benchmark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeProtoAddress.ProtoReflect.Descriptor instead.
func (*ResumeProtoAddress) Descriptor() ([]byte, []int) {
	return file_benchmark_common2_benchmark_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResumeProtoAddress) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *ResumeProtoAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ResumeProtoAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ResumeProtoAddress) GetZip() int32 {
	if x != nil {
		return x.Zip
	}
	return 0
}

type ResumeProtoPersonalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Education  *ResumeProtoPersonalInfoEducation  `protobuf:"bytes,1,opt,name=Education,proto3" json:"Education,omitempty"`
	Experience *ResumeProtoPersonalInfoExperience `protobuf:"bytes,2,opt,name=Experience,proto3" json:"Experience,omitempty"`
}

func (x *ResumeProtoPersonalInfo) Reset() {
	*x = ResumeProtoPersonalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_common2_benchmark_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeProtoPersonalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeProtoPersonalInfo) ProtoMessage() {}

func (x *ResumeProtoPersonalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_common2_benchmark_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeProtoPersonalInfo.ProtoReflect.Descriptor instead.
func (*ResumeProtoPersonalInfo) Descriptor() ([]byte, []int) {
	return file_benchmark_common2_benchmark_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ResumeProtoPersonalInfo) GetEducation() *ResumeProtoPersonalInfoEducation {
	if x != nil {
		return x.Education
	}
	return nil
}

func (x *ResumeProtoPersonalInfo) GetExperience() *ResumeProtoPersonalInfoExperience {
	if x != nil {
		return x.Experience
	}
	return nil
}

type ResumeProtoPersonalInfoEducation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Degree   string                       `protobuf:"bytes,1,opt,name=Degree,proto3" json:"Degree,omitempty"`
	Major    string                       `protobuf:"bytes,2,opt,name=Major,proto3" json:"Major,omitempty"`
	School   string                       `protobuf:"bytes,3,opt,name=School,proto3" json:"School,omitempty"`
	Graduate *ResumeProtoPersonalInfoDate `protobuf:"bytes,4,opt,name=Graduate,proto3" json:"Graduate,omitempty"`
}

func (x *ResumeProtoPersonalInfoEducation) Reset() {
	*x = ResumeProtoPersonalInfoEducation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_common2_benchmark_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeProtoPersonalInfoEducation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeProtoPersonalInfoEducation) ProtoMessage() {}

func (x *ResumeProtoPersonalInfoEducation) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_common2_benchmark_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeProtoPersonalInfoEducation.ProtoReflect.Descriptor instead.
func (*ResumeProtoPersonalInfoEducation) Descriptor() ([]byte, []int) {
	return file_benchmark_common2_benchmark_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *ResumeProtoPersonalInfoEducation) GetDegree() string {
	if x != nil {
		return x.Degree
	}
	return ""
}

func (x *ResumeProtoPersonalInfoEducation) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *ResumeProtoPersonalInfoEducation) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *ResumeProtoPersonalInfoEducation) GetGraduate() *ResumeProtoPersonalInfoDate {
	if x != nil {
		return x.Graduate
	}
	return nil
}

type ResumeProtoPersonalInfoExperience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job []*ResumeProtoPersonalInfoExperienceJob `protobuf:"bytes,1,rep,name=Job,proto3" json:"Job,omitempty"`
}

func (x *ResumeProtoPersonalInfoExperience) Reset() {
	*x = ResumeProtoPersonalInfoExperience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_common2_benchmark_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeProtoPersonalInfoExperience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeProtoPersonalInfoExperience) ProtoMessage() {}

func (x *ResumeProtoPersonalInfoExperience) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_common2_benchmark_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeProtoPersonalInfoExperience.ProtoReflect.Descriptor instead.
func (*ResumeProtoPersonalInfoExperience) Descriptor() ([]byte, []int) {
	return file_benchmark_common2_benchmark_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *ResumeProtoPersonalInfoExperience) GetJob() []*ResumeProtoPersonalInfoExperienceJob {
	if x != nil {
		return x.Job
	}
	return nil
}

type ResumeProtoPersonalInfoDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *ResumeProtoPersonalInfoDate) Reset() {
	*x = ResumeProtoPersonalInfoDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_common2_benchmark_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeProtoPersonalInfoDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeProtoPersonalInfoDate) ProtoMessage() {}

func (x *ResumeProtoPersonalInfoDate) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_common2_benchmark_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeProtoPersonalInfoDate.ProtoReflect.Descriptor instead.
func (*ResumeProtoPersonalInfoDate) Descriptor() ([]byte, []int) {
	return file_benchmark_common2_benchmark_proto_rawDescGZIP(), []int{0, 1, 2}
}

func (x *ResumeProtoPersonalInfoDate) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *ResumeProtoPersonalInfoDate) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *ResumeProtoPersonalInfoDate) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type ResumeProtoPersonalInfoExperienceJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string                       `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Company   string                       `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	StartDate *ResumeProtoPersonalInfoDate `protobuf:"bytes,3,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate   *ResumeProtoPersonalInfoDate `protobuf:"bytes,4,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
}

func (x *ResumeProtoPersonalInfoExperienceJob) Reset() {
	*x = ResumeProtoPersonalInfoExperienceJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_common2_benchmark_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeProtoPersonalInfoExperienceJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeProtoPersonalInfoExperienceJob) ProtoMessage() {}

func (x *ResumeProtoPersonalInfoExperienceJob) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_common2_benchmark_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeProtoPersonalInfoExperienceJob.ProtoReflect.Descriptor instead.
func (*ResumeProtoPersonalInfoExperienceJob) Descriptor() ([]byte, []int) {
	return file_benchmark_common2_benchmark_proto_rawDescGZIP(), []int{0, 1, 1, 0}
}

func (x *ResumeProtoPersonalInfoExperienceJob) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResumeProtoPersonalInfoExperienceJob) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *ResumeProtoPersonalInfoExperienceJob) GetStartDate() *ResumeProtoPersonalInfoDate {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *ResumeProtoPersonalInfoExperienceJob) GetEndDate() *ResumeProtoPersonalInfoDate {
	if x != nil {
		return x.EndDate
	}
	return nil
}

var File_benchmark_common2_benchmark_proto protoreflect.FileDescriptor

var file_benchmark_common2_benchmark_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x32, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d, 0x4c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xd4,
	0x08, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x48,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d, 0x4c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d, 0x4c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x5d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x5a, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x5a, 0x69, 0x70,
	0x1a, 0x82, 0x06, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x5b, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d,
	0x4c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5e,
	0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d, 0x4c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0xa7,
	0x01, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d,
	0x4c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x52, 0x08,
	0x47, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x1a, 0xc6, 0x02, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d,
	0x4c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x52, 0x03, 0x4a, 0x6f, 0x62, 0x1a, 0xe1, 0x01,
	0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x56, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x58, 0x4d, 0x4c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x52, 0x0a,
	0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x58, 0x4d, 0x4c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x1a, 0x42, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x64, 0x61, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benchmark_common2_benchmark_proto_rawDescOnce sync.Once
	file_benchmark_common2_benchmark_proto_rawDescData = file_benchmark_common2_benchmark_proto_rawDesc
)

func file_benchmark_common2_benchmark_proto_rawDescGZIP() []byte {
	file_benchmark_common2_benchmark_proto_rawDescOnce.Do(func() {
		file_benchmark_common2_benchmark_proto_rawDescData = protoimpl.X.CompressGZIP(file_benchmark_common2_benchmark_proto_rawDescData)
	})
	return file_benchmark_common2_benchmark_proto_rawDescData
}

var file_benchmark_common2_benchmark_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_benchmark_common2_benchmark_proto_goTypes = []interface{}{
	(*ResumeProto)(nil),                          // 0: dynamicXMLproto.benchmark.resumeProto
	(*ResumeProtoAddress)(nil),                   // 1: dynamicXMLproto.benchmark.resumeProto.address
	(*ResumeProtoPersonalInfo)(nil),              // 2: dynamicXMLproto.benchmark.resumeProto.personalInfo
	(*ResumeProtoPersonalInfoEducation)(nil),     // 3: dynamicXMLproto.benchmark.resumeProto.personalInfo.education
	(*ResumeProtoPersonalInfoExperience)(nil),    // 4: dynamicXMLproto.benchmark.resumeProto.personalInfo.experience
	(*ResumeProtoPersonalInfoDate)(nil),          // 5: dynamicXMLproto.benchmark.resumeProto.personalInfo.date
	(*ResumeProtoPersonalInfoExperienceJob)(nil), // 6: dynamicXMLproto.benchmark.resumeProto.personalInfo.experience.job
}
var file_benchmark_common2_benchmark_proto_depIdxs = []int32{
	1, // 0: dynamicXMLproto.benchmark.resumeProto.Address:type_name -> dynamicXMLproto.benchmark.resumeProto.address
	2, // 1: dynamicXMLproto.benchmark.resumeProto.PersonalInfo:type_name -> dynamicXMLproto.benchmark.resumeProto.personalInfo
	3, // 2: dynamicXMLproto.benchmark.resumeProto.personalInfo.Education:type_name -> dynamicXMLproto.benchmark.resumeProto.personalInfo.education
	4, // 3: dynamicXMLproto.benchmark.resumeProto.personalInfo.Experience:type_name -> dynamicXMLproto.benchmark.resumeProto.personalInfo.experience
	5, // 4: dynamicXMLproto.benchmark.resumeProto.personalInfo.education.Graduate:type_name -> dynamicXMLproto.benchmark.resumeProto.personalInfo.date
	6, // 5: dynamicXMLproto.benchmark.resumeProto.personalInfo.experience.Job:type_name -> dynamicXMLproto.benchmark.resumeProto.personalInfo.experience.job
	5, // 6: dynamicXMLproto.benchmark.resumeProto.personalInfo.experience.job.StartDate:type_name -> dynamicXMLproto.benchmark.resumeProto.personalInfo.date
	5, // 7: dynamicXMLproto.benchmark.resumeProto.personalInfo.experience.job.EndDate:type_name -> dynamicXMLproto.benchmark.resumeProto.personalInfo.date
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_benchmark_common2_benchmark_proto_init() }
func file_benchmark_common2_benchmark_proto_init() {
	if File_benchmark_common2_benchmark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benchmark_common2_benchmark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeProto); i {
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
		file_benchmark_common2_benchmark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeProtoAddress); i {
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
		file_benchmark_common2_benchmark_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeProtoPersonalInfo); i {
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
		file_benchmark_common2_benchmark_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeProtoPersonalInfoEducation); i {
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
		file_benchmark_common2_benchmark_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeProtoPersonalInfoExperience); i {
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
		file_benchmark_common2_benchmark_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeProtoPersonalInfoDate); i {
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
		file_benchmark_common2_benchmark_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeProtoPersonalInfoExperienceJob); i {
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
			RawDescriptor: file_benchmark_common2_benchmark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_benchmark_common2_benchmark_proto_goTypes,
		DependencyIndexes: file_benchmark_common2_benchmark_proto_depIdxs,
		MessageInfos:      file_benchmark_common2_benchmark_proto_msgTypes,
	}.Build()
	File_benchmark_common2_benchmark_proto = out.File
	file_benchmark_common2_benchmark_proto_rawDesc = nil
	file_benchmark_common2_benchmark_proto_goTypes = nil
	file_benchmark_common2_benchmark_proto_depIdxs = nil
}
