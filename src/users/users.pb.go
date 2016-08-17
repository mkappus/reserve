// Code generated by protoc-gen-go.
// source: users.proto
// DO NOT EDIT!

/*
Package users is a generated protocol buffer package.

It is generated from these files:
	users.proto

It has these top-level messages:
	Teacher
	Student
	Course
*/
package users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Teacher struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
}

func (m *Teacher) Reset()                    { *m = Teacher{} }
func (m *Teacher) String() string            { return proto.CompactTextString(m) }
func (*Teacher) ProtoMessage()               {}
func (*Teacher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Student struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
}

func (m *Student) Reset()                    { *m = Student{} }
func (m *Student) String() string            { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()               {}
func (*Student) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Course struct {
	Id      uint64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Teacher *Teacher `protobuf:"bytes,2,opt,name=teacher" json:"teacher,omitempty"`
	// map<uint64, Student > students = 3;
	Students []*Student `protobuf:"bytes,3,rep,name=students" json:"students,omitempty"`
}

func (m *Course) Reset()                    { *m = Course{} }
func (m *Course) String() string            { return proto.CompactTextString(m) }
func (*Course) ProtoMessage()               {}
func (*Course) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Course) GetTeacher() *Teacher {
	if m != nil {
		return m.Teacher
	}
	return nil
}

func (m *Course) GetStudents() []*Student {
	if m != nil {
		return m.Students
	}
	return nil
}

func init() {
	proto.RegisterType((*Teacher)(nil), "users.Teacher")
	proto.RegisterType((*Student)(nil), "users.Student")
	proto.RegisterType((*Course)(nil), "users.Course")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xe2, 0xb9, 0xd8, 0x43,
	0x52, 0x13, 0x93, 0x33, 0x52, 0x8b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x82, 0x80, 0x2c, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xa0,
	0x08, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc2, 0xc5, 0x9a, 0x9a, 0x9b, 0x98, 0x99, 0x23, 0xc1, 0x0c,
	0x16, 0x84, 0x70, 0x84, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24,
	0x58, 0xc0, 0x12, 0x70, 0x3e, 0xc8, 0x82, 0xe0, 0x92, 0xd2, 0x94, 0xd4, 0xbc, 0x12, 0x1a, 0x59,
	0x90, 0xc7, 0xc5, 0xe6, 0x9c, 0x5f, 0x5a, 0x54, 0x9c, 0x8a, 0x61, 0xbe, 0x06, 0x17, 0x7b, 0x09,
	0xc4, 0x6f, 0x60, 0x2b, 0xb8, 0x8d, 0xf8, 0xf4, 0x20, 0x21, 0x00, 0xf5, 0x71, 0x10, 0x4c, 0x5a,
	0x48, 0x8b, 0x8b, 0xa3, 0x18, 0xe2, 0xc8, 0x62, 0xa0, 0xc5, 0xcc, 0x48, 0x4a, 0xa1, 0x6e, 0x0f,
	0x82, 0xcb, 0x27, 0xb1, 0x81, 0xc3, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x87, 0x8a,
	0xbe, 0x4e, 0x01, 0x00, 0x00,
}
