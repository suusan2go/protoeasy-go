// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/source_context.proto

package google_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `SourceContext` represents information about the source of a
// protobuf element, like the file in which it is defined.
type SourceContext struct {
	// The path-qualified name of the .proto file that contained the associated
	// protobuf element.  For example: `"google/protobuf/source_context.proto"`.
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
}

func (m *SourceContext) Reset()                    { *m = SourceContext{} }
func (m *SourceContext) String() string            { return proto.CompactTextString(m) }
func (*SourceContext) ProtoMessage()               {}
func (*SourceContext) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *SourceContext) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func init() {
	proto.RegisterType((*SourceContext)(nil), "google.protobuf.SourceContext")
}

func init() { proto.RegisterFile("google/protobuf/source_context.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xce, 0x2f, 0x2d,
	0x4a, 0x4e, 0x8d, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xad, 0x28, 0xd1, 0x03, 0x8b, 0x0b, 0xf1, 0x43,
	0x54, 0xe9, 0xc1, 0x54, 0x29, 0xe9, 0x70, 0xf1, 0x06, 0x83, 0x15, 0x3a, 0x43, 0xd4, 0x09, 0x49,
	0x73, 0x71, 0xa6, 0x65, 0xe6, 0xa4, 0xc6, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x71, 0x80, 0x04, 0xfc, 0x12, 0x73, 0x53, 0x9d, 0x82, 0xb8, 0x84, 0x93, 0xf3, 0x73,
	0xf5, 0xd0, 0x0c, 0x71, 0x12, 0x42, 0x31, 0x22, 0x00, 0x24, 0x1c, 0xc0, 0xb8, 0x88, 0x89, 0xd9,
	0x3d, 0xc0, 0x69, 0x15, 0x93, 0x9c, 0x3b, 0x44, 0x75, 0x00, 0x54, 0xb5, 0x5e, 0x78, 0x6a, 0x4e,
	0x8e, 0x77, 0x5e, 0x7e, 0x79, 0x5e, 0x48, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x18, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xb1, 0x33, 0xaa, 0xc1, 0x00, 0x00, 0x00,
}
