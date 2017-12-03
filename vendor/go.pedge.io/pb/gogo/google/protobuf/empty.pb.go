// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/protobuf/empty.proto

package google_protobuf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorEmpty, []int{0} }
func (*Empty) XXX_WellKnownType() string   { return "Empty" }

func init() {
	proto.RegisterType((*Empty)(nil), "google.protobuf.Empty")
}

func init() { proto.RegisterFile("google/protobuf/empty.proto", fileDescriptorEmpty) }

var fileDescriptorEmpty = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xcd, 0x2d, 0x28,
	0xa9, 0xd4, 0x03, 0x73, 0x85, 0xf8, 0x21, 0x92, 0x7a, 0x30, 0x49, 0x25, 0x76, 0x2e, 0x56, 0x57,
	0x90, 0xbc, 0x93, 0x2f, 0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xbc, 0x13, 0x17, 0x58, 0x36,
	0x00, 0xc4, 0x0d, 0x60, 0xfc, 0xc1, 0xc8, 0xb8, 0x88, 0x89, 0xd9, 0x3d, 0xc0, 0x69, 0x15, 0x93,
	0x9c, 0x3b, 0x44, 0x61, 0x00, 0x54, 0xa1, 0x5e, 0x78, 0x6a, 0x4e, 0x8e, 0x77, 0x5e, 0x7e, 0x79,
	0x5e, 0x48, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x04, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0x47, 0x3d, 0xdd, 0x8e, 0x00, 0x00, 0x00,
}
