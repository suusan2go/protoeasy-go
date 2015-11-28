// Code generated by protoc-gen-go.
// source: exec/pkgexec.proto
// DO NOT EDIT!

package pkgexec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RunningCommand struct {
	Args string `protobuf:"bytes,1,opt,name=args" json:"args,omitempty"`
}

func (m *RunningCommand) Reset()                    { *m = RunningCommand{} }
func (m *RunningCommand) String() string            { return proto.CompactTextString(m) }
func (*RunningCommand) ProtoMessage()               {}
func (*RunningCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*RunningCommand)(nil), "pkgexec.RunningCommand")
}

var fileDescriptor0 = []byte{
	// 82 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xad, 0x48, 0x4d,
	0xd6, 0x2f, 0xc8, 0x4e, 0x07, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x1c, 0x17, 0x5f, 0x50, 0x69, 0x5e, 0x5e, 0x66, 0x5e, 0xba, 0x73, 0x7e, 0x6e, 0x6e, 0x62,
	0x5e, 0x8a, 0x10, 0x0f, 0x17, 0x4b, 0x62, 0x51, 0x7a, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x12, 0x1b, 0x58, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x66, 0x99, 0xbe, 0xf8, 0x45, 0x00,
	0x00, 0x00,
}
