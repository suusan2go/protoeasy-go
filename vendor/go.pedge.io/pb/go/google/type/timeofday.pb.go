// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/timeofday.proto

package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a time of day. The date and time zone are either not significant
// or are specified elsewhere. An API may chose to allow leap seconds. Related
// types are [google.type.Date][google.type.Date] and `google.protobuf.Timestamp`.
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	// to allow the value "24:00:00" for scenarios like business closing time.
	Hours int32 `protobuf:"varint,1,opt,name=hours" json:"hours,omitempty"`
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,2,opt,name=minutes" json:"minutes,omitempty"`
	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	// allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,3,opt,name=seconds" json:"seconds,omitempty"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos int32 `protobuf:"varint,4,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *TimeOfDay) Reset()                    { *m = TimeOfDay{} }
func (m *TimeOfDay) String() string            { return proto.CompactTextString(m) }
func (*TimeOfDay) ProtoMessage()               {}
func (*TimeOfDay) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *TimeOfDay) GetHours() int32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *TimeOfDay) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *TimeOfDay) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TimeOfDay) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*TimeOfDay)(nil), "google.type.TimeOfDay")
}

func init() { proto.RegisterFile("google/type/timeofday.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0xc9, 0xcc, 0x4d, 0xcd, 0x4f, 0x4b, 0x49,
	0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xea, 0x81, 0x24, 0x95, 0xb2,
	0xb9, 0x38, 0x43, 0x32, 0x73, 0x53, 0xfd, 0xd3, 0x5c, 0x12, 0x2b, 0x85, 0x44, 0xb8, 0x58, 0x33,
	0xf2, 0x4b, 0x8b, 0x8a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x21, 0x09, 0x2e,
	0xf6, 0xdc, 0xcc, 0xbc, 0xd2, 0x92, 0xd4, 0x62, 0x09, 0x26, 0xb0, 0x38, 0x8c, 0x0b, 0x92, 0x29,
	0x4e, 0x4d, 0xce, 0xcf, 0x4b, 0x29, 0x96, 0x60, 0x86, 0xc8, 0x40, 0xb9, 0x20, 0x93, 0xf2, 0x12,
	0xf3, 0xf2, 0x8b, 0x25, 0x58, 0x20, 0x26, 0x81, 0x39, 0x4e, 0x9a, 0x5c, 0xfc, 0xc9, 0xf9, 0xb9,
	0x7a, 0x48, 0xf6, 0x3b, 0xf1, 0xc1, 0x6d, 0x0f, 0x00, 0x39, 0x2e, 0x80, 0x71, 0x11, 0x13, 0xb3,
	0x7b, 0x48, 0x40, 0x12, 0x1b, 0xd8, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x46,
	0x4a, 0x01, 0xca, 0x00, 0x00, 0x00,
}
