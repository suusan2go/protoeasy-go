// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/struct.proto

package google_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}
var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return proto.EnumName(NullValue_name, int32(x))
}
func (NullValue) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }
func (NullValue) XXX_WellKnownType() string       { return "NullValue" }

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	// Unordered map of dynamically typed values.
	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Struct) Reset()                    { *m = Struct{} }
func (m *Struct) String() string            { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()               {}
func (*Struct) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }
func (*Struct) XXX_WellKnownType() string   { return "Struct" }

func (m *Struct) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of that
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
type Value struct {
	// The kind of value.
	//
	// Types that are valid to be assigned to Kind:
	//	*Value_NullValue
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_ListValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }
func (*Value) XXX_WellKnownType() string   { return "Value" }

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,enum=google.protobuf.NullValue,oneof"`
}
type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,2,opt,name=number_value,json=numberValue,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,oneof"`
}
type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,oneof"`
}
type Value_StructValue struct {
	StructValue *Struct `protobuf:"bytes,5,opt,name=struct_value,json=structValue,oneof"`
}
type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,6,opt,name=list_value,json=listValue,oneof"`
}

func (*Value_NullValue) isValue_Kind()   {}
func (*Value_NumberValue) isValue_Kind() {}
func (*Value_StringValue) isValue_Kind() {}
func (*Value_BoolValue) isValue_Kind()   {}
func (*Value_StructValue) isValue_Kind() {}
func (*Value_ListValue) isValue_Kind()   {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetNullValue() NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

func (m *Value) GetNumberValue() float64 {
	if x, ok := m.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetStructValue() *Struct {
	if x, ok := m.GetKind().(*Value_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.NullValue))
	case *Value_NumberValue:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.NumberValue))
	case *Value_StringValue:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_StructValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructValue); err != nil {
			return err
		}
	case *Value_ListValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.Kind has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // kind.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_NullValue{NullValue(x)}
		return true, err
	case 2: // kind.number_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Kind = &Value_NumberValue{math.Float64frombits(x)}
		return true, err
	case 3: // kind.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Value_StringValue{x}
		return true, err
	case 4: // kind.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_BoolValue{x != 0}
		return true, err
	case 5: // kind.struct_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Struct)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_StructValue{msg}
		return true, err
	case 6: // kind.list_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_ListValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Value_NumberValue:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *Value_StringValue:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BoolValue:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case *Value_StructValue:
		s := proto.Size(x.StructValue)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ListValue:
		s := proto.Size(x.ListValue)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is JSON array.
type ListValue struct {
	// Repeated field of dynamically typed values.
	Values []*Value `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ListValue) Reset()                    { *m = ListValue{} }
func (m *ListValue) String() string            { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()               {}
func (*ListValue) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }
func (*ListValue) XXX_WellKnownType() string   { return "ListValue" }

func (m *ListValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Struct)(nil), "google.protobuf.Struct")
	proto.RegisterType((*Value)(nil), "google.protobuf.Value")
	proto.RegisterType((*ListValue)(nil), "google.protobuf.ListValue")
	proto.RegisterEnum("google.protobuf.NullValue", NullValue_name, NullValue_value)
}

func init() { proto.RegisterFile("google/protobuf/struct.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0xef, 0xd2, 0x40,
	0x10, 0xc5, 0xd9, 0x16, 0x1a, 0x3b, 0x35, 0x48, 0xd6, 0x44, 0x09, 0x1a, 0x6d, 0xe0, 0xd2, 0x18,
	0x53, 0x12, 0xbc, 0x18, 0xf1, 0x62, 0x13, 0x84, 0xc4, 0x86, 0xd4, 0x2a, 0x78, 0x24, 0xb6, 0x14,
	0xd2, 0xb0, 0xec, 0x92, 0x76, 0x57, 0xc3, 0x37, 0xf1, 0xec, 0xd1, 0xa3, 0x9f, 0xce, 0xa3, 0xd9,
	0xdd, 0xb6, 0xff, 0x7f, 0x20, 0xdc, 0x3a, 0x6f, 0x7e, 0xf3, 0x66, 0xdf, 0x14, 0x9e, 0xef, 0x19,
	0xdb, 0x93, 0x6c, 0x7c, 0x2a, 0x18, 0x67, 0x89, 0xd8, 0x8d, 0x4b, 0x5e, 0x88, 0x94, 0xfb, 0xaa,
	0xc6, 0x8f, 0x74, 0xd7, 0xaf, 0xbb, 0xc3, 0x5f, 0x08, 0xac, 0x2f, 0x8a, 0xc0, 0x53, 0xb0, 0x76,
	0x79, 0x46, 0xb6, 0x65, 0x1f, 0xb9, 0xa6, 0xe7, 0x4c, 0x46, 0xfe, 0x05, 0xec, 0x6b, 0xd0, 0xff,
	0xa8, 0xa8, 0x19, 0xe5, 0xc5, 0x39, 0xae, 0x46, 0x06, 0x9f, 0xc1, 0xb9, 0x27, 0xe3, 0x1e, 0x98,
	0x87, 0xec, 0xdc, 0x47, 0x2e, 0xf2, 0xec, 0x58, 0x7e, 0xe2, 0xd7, 0xd0, 0xf9, 0xf1, 0x9d, 0x88,
	0xac, 0x6f, 0xb8, 0xc8, 0x73, 0x26, 0x4f, 0xae, 0xcc, 0xd7, 0xb2, 0x1b, 0x6b, 0xe8, 0x9d, 0xf1,
	0x16, 0x0d, 0xff, 0x1a, 0xd0, 0x51, 0x22, 0x9e, 0x02, 0x50, 0x41, 0xc8, 0x46, 0x1b, 0x48, 0xd3,
	0xee, 0x64, 0x70, 0x65, 0xb0, 0x14, 0x84, 0x28, 0x7e, 0xd1, 0x8a, 0x6d, 0x5a, 0x17, 0x78, 0x04,
	0x0f, 0xa9, 0x38, 0x26, 0x59, 0xb1, 0xb9, 0xdb, 0x8f, 0x16, 0xad, 0xd8, 0xd1, 0x6a, 0x03, 0x95,
	0xbc, 0xc8, 0xe9, 0xbe, 0x82, 0x4c, 0xf9, 0x70, 0x09, 0x69, 0x55, 0x43, 0x2f, 0x01, 0x12, 0xc6,
	0xea, 0x67, 0xb4, 0x5d, 0xe4, 0x3d, 0x90, 0xab, 0xa4, 0xa6, 0x81, 0xf7, 0xca, 0x45, 0xa4, 0xbc,
	0x42, 0x3a, 0x2a, 0xea, 0xd3, 0x1b, 0x77, 0xac, 0xec, 0x45, 0xca, 0x9b, 0x94, 0x24, 0x2f, 0xeb,
	0x59, 0x4b, 0xcd, 0x5e, 0xa7, 0x0c, 0xf3, 0x92, 0x37, 0x29, 0x49, 0x5d, 0x04, 0x16, 0xb4, 0x0f,
	0x39, 0xdd, 0x0e, 0xa7, 0x60, 0x37, 0x04, 0xf6, 0xc1, 0x52, 0x66, 0xf5, 0x1f, 0xbd, 0x75, 0xf4,
	0x8a, 0x7a, 0xf5, 0x0c, 0xec, 0xe6, 0x88, 0xb8, 0x0b, 0xb0, 0x5c, 0x85, 0xe1, 0x66, 0xfd, 0x21,
	0x5c, 0xcd, 0x7a, 0xad, 0x60, 0x09, 0x8f, 0x53, 0x76, 0xbc, 0x74, 0x08, 0x1c, 0x1d, 0x26, 0x92,
	0x75, 0x84, 0xfe, 0x21, 0xf4, 0xdb, 0x30, 0xe7, 0x51, 0xf0, 0xc7, 0x78, 0x31, 0xd7, 0x64, 0x54,
	0xef, 0xfa, 0x96, 0x11, 0xf2, 0x89, 0xb2, 0x9f, 0xf4, 0xeb, 0xf9, 0x94, 0x95, 0x89, 0xa5, 0x2c,
	0xde, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x83, 0x7f, 0xda, 0xb1, 0x02, 0x00, 0x00,
}
