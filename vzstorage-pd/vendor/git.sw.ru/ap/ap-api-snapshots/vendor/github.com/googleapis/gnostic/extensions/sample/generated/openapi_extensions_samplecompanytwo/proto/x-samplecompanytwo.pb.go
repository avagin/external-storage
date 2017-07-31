// Code generated by protoc-gen-go.
// source: x-samplecompanytwo.proto
// DO NOT EDIT!

/*
Package samplecompanytwo is a generated protocol buffer package.

It is generated from these files:
	x-samplecompanytwo.proto

It has these top-level messages:
	SampleCompanyTwoBook
	SampleCompanyTwoShelve
*/
package samplecompanytwo

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

type SampleCompanyTwoBook struct {
	Code    int64 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message int64 `protobuf:"varint,2,opt,name=message" json:"message,omitempty"`
}

func (m *SampleCompanyTwoBook) Reset()                    { *m = SampleCompanyTwoBook{} }
func (m *SampleCompanyTwoBook) String() string            { return proto.CompactTextString(m) }
func (*SampleCompanyTwoBook) ProtoMessage()               {}
func (*SampleCompanyTwoBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SampleCompanyTwoBook) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SampleCompanyTwoBook) GetMessage() int64 {
	if m != nil {
		return m.Message
	}
	return 0
}

type SampleCompanyTwoShelve struct {
	Foo1 int64 `protobuf:"varint,1,opt,name=foo1" json:"foo1,omitempty"`
	Bar  int64 `protobuf:"varint,2,opt,name=bar" json:"bar,omitempty"`
}

func (m *SampleCompanyTwoShelve) Reset()                    { *m = SampleCompanyTwoShelve{} }
func (m *SampleCompanyTwoShelve) String() string            { return proto.CompactTextString(m) }
func (*SampleCompanyTwoShelve) ProtoMessage()               {}
func (*SampleCompanyTwoShelve) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SampleCompanyTwoShelve) GetFoo1() int64 {
	if m != nil {
		return m.Foo1
	}
	return 0
}

func (m *SampleCompanyTwoShelve) GetBar() int64 {
	if m != nil {
		return m.Bar
	}
	return 0
}

func init() {
	proto.RegisterType((*SampleCompanyTwoBook)(nil), "samplecompanytwo.SampleCompanyTwoBook")
	proto.RegisterType((*SampleCompanyTwoShelve)(nil), "samplecompanytwo.SampleCompanyTwoShelve")
}

func init() { proto.RegisterFile("x-samplecompanytwo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8f, 0x41, 0xeb, 0x82, 0x30,
	0x18, 0x87, 0x51, 0xff, 0xfc, 0x83, 0x9d, 0x64, 0x48, 0xec, 0x18, 0x1e, 0xa2, 0x4b, 0x83, 0xe8,
	0xde, 0xc1, 0xea, 0x2e, 0x19, 0xdd, 0xa7, 0xbe, 0x99, 0xa4, 0x7b, 0xc7, 0x26, 0x69, 0x5f, 0xa7,
	0x4f, 0x1a, 0x5b, 0x79, 0xb1, 0xdb, 0x6f, 0x0f, 0xe3, 0xe1, 0x7d, 0x08, 0x1b, 0xd6, 0x46, 0xb4,
	0xaa, 0x81, 0x02, 0x5b, 0x25, 0xe4, 0xb3, 0xeb, 0x91, 0x2b, 0x8d, 0x1d, 0xd2, 0x70, 0xca, 0xe3,
	0x03, 0x89, 0x32, 0xc7, 0xf6, 0x1f, 0x76, 0xee, 0x31, 0x41, 0xbc, 0x53, 0x4a, 0xfe, 0x0a, 0x2c,
	0x81, 0x79, 0x0b, 0x6f, 0x15, 0x9c, 0xdc, 0xa6, 0x8c, 0xcc, 0x5a, 0x30, 0x46, 0x54, 0xc0, 0x7c,
	0x87, 0xc7, 0x67, 0xbc, 0x23, 0xf3, 0xa9, 0x25, 0xbb, 0x41, 0xf3, 0x00, 0xeb, 0xb9, 0x22, 0x6e,
	0x46, 0x8f, 0xdd, 0x34, 0x24, 0x41, 0x2e, 0xf4, 0xd7, 0x61, 0x67, 0x92, 0x91, 0x25, 0xea, 0x8a,
	0xa3, 0x02, 0x29, 0x54, 0xcd, 0x61, 0xe8, 0x40, 0x9a, 0x1a, 0x25, 0x9f, 0xde, 0x9b, 0x44, 0x17,
	0x90, 0x25, 0xea, 0xe3, 0xf8, 0x23, 0xb5, 0x5d, 0xa9, 0xf7, 0xf2, 0x7f, 0xd2, 0xf2, 0x7f, 0xd7,
	0xbc, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x45, 0x2d, 0x48, 0x0f, 0x01, 0x00, 0x00,
}
