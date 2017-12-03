// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/duration.proto

package google_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A Duration represents a signed, fixed-length span of time represented
// as a count of seconds and fractions of seconds at nanosecond
// resolution. It is independent of any calendar and concepts like "day"
// or "month". It is related to Timestamp in that the difference between
// two Timestamp values is a Duration and it can be added or subtracted
// from a Timestamp. Range is approximately +-10,000 years.
//
// Example 1: Compute Duration from two Timestamps in pseudo code.
//
//     Timestamp start = ...;
//     Timestamp end = ...;
//     Duration duration = ...;
//
//     duration.seconds = end.seconds - start.seconds;
//     duration.nanos = end.nanos - start.nanos;
//
//     if (duration.seconds < 0 && duration.nanos > 0) {
//       duration.seconds += 1;
//       duration.nanos -= 1000000000;
//     } else if (durations.seconds > 0 && duration.nanos < 0) {
//       duration.seconds -= 1;
//       duration.nanos += 1000000000;
//     }
//
// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
//
//     Timestamp start = ...;
//     Duration duration = ...;
//     Timestamp end = ...;
//
//     end.seconds = start.seconds + duration.seconds;
//     end.nanos = start.nanos + duration.nanos;
//
//     if (end.nanos < 0) {
//       end.seconds -= 1;
//       end.nanos += 1000000000;
//     } else if (end.nanos >= 1000000000) {
//       end.seconds += 1;
//       end.nanos -= 1000000000;
//     }
//
// Example 3: Compute Duration from datetime.timedelta in Python.
//
//     td = datetime.timedelta(days=3, minutes=10)
//     duration = Duration()
//     duration.FromTimedelta(td)
//
//
type Duration struct {
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Duration) Reset()                    { *m = Duration{} }
func (m *Duration) String() string            { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }
func (*Duration) XXX_WellKnownType() string   { return "Duration" }

func (m *Duration) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Duration) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*Duration)(nil), "google.protobuf.Duration")
}

func init() { proto.RegisterFile("google/protobuf/duration.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
	0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
	0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0x14, 0xc0, 0x25, 0x9c, 0x9c,
	0x9f, 0xab, 0x87, 0x66, 0xa4, 0x13, 0x2f, 0xcc, 0xc0, 0x00, 0x90, 0x48, 0x00, 0xe3, 0x0f, 0x46,
	0xc6, 0x45, 0x4c, 0xcc, 0xee, 0x01, 0x4e, 0xab, 0x98, 0xe4, 0xdc, 0x21, 0x6a, 0x03, 0xa0, 0x6a,
	0xf5, 0xc2, 0x53, 0x73, 0x72, 0xbc, 0xf3, 0xf2, 0xcb, 0xf3, 0x42, 0x2a, 0x0b, 0x52, 0x8b, 0x93,
	0xd8, 0xc0, 0x86, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x08, 0x35, 0x91, 0xc7, 0x00,
	0x00, 0x00,
}
