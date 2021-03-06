// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/config_change.proto

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Classifies set of possible modifications to an object in the service
// configuration.
type ChangeType int32

const (
	// No value was provided.
	ChangeType_CHANGE_TYPE_UNSPECIFIED ChangeType = 0
	// The changed object exists in the 'new' service configuration, but not
	// in the 'old' service configuration.
	ChangeType_ADDED ChangeType = 1
	// The changed object exists in the 'old' service configuration, but not
	// in the 'new' service configuration.
	ChangeType_REMOVED ChangeType = 2
	// The changed object exists in both service configurations, but its value
	// is different.
	ChangeType_MODIFIED ChangeType = 3
)

var ChangeType_name = map[int32]string{
	0: "CHANGE_TYPE_UNSPECIFIED",
	1: "ADDED",
	2: "REMOVED",
	3: "MODIFIED",
}
var ChangeType_value = map[string]int32{
	"CHANGE_TYPE_UNSPECIFIED": 0,
	"ADDED":                   1,
	"REMOVED":                 2,
	"MODIFIED":                3,
}

func (x ChangeType) String() string {
	return proto.EnumName(ChangeType_name, int32(x))
}
func (ChangeType) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfigChange, []int{0} }

// Output generated from semantically comparing two versions of a service
// configuration.
//
// Includes detailed information about a field that have changed with
// applicable advice about potential consequences for the change, such as
// backwards-incompatibility.
type ConfigChange struct {
	// Object hierarchy path to the change, with levels separated by a '.'
	// character. For repeated fields, an applicable unique identifier field is
	// used for the index (usually selector, name, or id). For maps, the term
	// 'key' is used. If the field has no unique identifier, the numeric index
	// is used.
	// Examples:
	// - visibility.rules[selector=="google.LibraryService.CreateBook"].restriction
	// - quota.metric_rules[selector=="google"].metric_costs[key=="reads"].value
	// - logging.producer_destinations[0]
	Element string `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"`
	// Value of the changed object in the old Service configuration,
	// in JSON format. This field will not be populated if ChangeType == ADDED.
	OldValue string `protobuf:"bytes,2,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	// Value of the changed object in the new Service configuration,
	// in JSON format. This field will not be populated if ChangeType == REMOVED.
	NewValue string `protobuf:"bytes,3,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	// The type for this change, either ADDED, REMOVED, or MODIFIED.
	ChangeType ChangeType `protobuf:"varint,4,opt,name=change_type,json=changeType,proto3,enum=google.api.ChangeType" json:"change_type,omitempty"`
	// Collection of advice provided for this change, useful for determining the
	// possible impact of this change.
	Advices []*Advice `protobuf:"bytes,5,rep,name=advices" json:"advices,omitempty"`
}

func (m *ConfigChange) Reset()                    { *m = ConfigChange{} }
func (m *ConfigChange) String() string            { return proto.CompactTextString(m) }
func (*ConfigChange) ProtoMessage()               {}
func (*ConfigChange) Descriptor() ([]byte, []int) { return fileDescriptorConfigChange, []int{0} }

func (m *ConfigChange) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

func (m *ConfigChange) GetOldValue() string {
	if m != nil {
		return m.OldValue
	}
	return ""
}

func (m *ConfigChange) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

func (m *ConfigChange) GetChangeType() ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return ChangeType_CHANGE_TYPE_UNSPECIFIED
}

func (m *ConfigChange) GetAdvices() []*Advice {
	if m != nil {
		return m.Advices
	}
	return nil
}

// Generated advice about this change, used for providing more
// information about how a change will affect the existing service.
type Advice struct {
	// Useful description for why this advice was applied and what actions should
	// be taken to mitigate any implied risks.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Advice) Reset()                    { *m = Advice{} }
func (m *Advice) String() string            { return proto.CompactTextString(m) }
func (*Advice) ProtoMessage()               {}
func (*Advice) Descriptor() ([]byte, []int) { return fileDescriptorConfigChange, []int{1} }

func (m *Advice) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigChange)(nil), "google.api.ConfigChange")
	proto.RegisterType((*Advice)(nil), "google.api.Advice")
	proto.RegisterEnum("google.api.ChangeType", ChangeType_name, ChangeType_value)
}

func init() { proto.RegisterFile("google/api/config_change.proto", fileDescriptorConfigChange) }

var fileDescriptorConfigChange = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x6a, 0xc2, 0x40,
	0x10, 0x87, 0x1b, 0xff, 0x3b, 0x11, 0xb1, 0x7b, 0x68, 0x03, 0x42, 0x09, 0x9e, 0x44, 0x24, 0x82,
	0x3d, 0xf4, 0x1c, 0x93, 0xad, 0xf5, 0xa0, 0x86, 0xd4, 0x0a, 0x3d, 0x85, 0x34, 0x99, 0xa6, 0x0b,
	0x71, 0x77, 0xd1, 0x54, 0xf1, 0x75, 0xfa, 0x36, 0x7d, 0xab, 0x62, 0x56, 0x6b, 0x6e, 0x3b, 0xf3,
	0xfd, 0x96, 0xf9, 0x98, 0x81, 0x87, 0x44, 0x88, 0x24, 0xc5, 0x51, 0x28, 0xd9, 0x28, 0x12, 0xfc,
	0x93, 0x25, 0x41, 0xf4, 0x15, 0xf2, 0x04, 0x2d, 0xb9, 0x15, 0x99, 0x20, 0xa0, 0xb8, 0x15, 0x4a,
	0xd6, 0xfb, 0xd5, 0xa0, 0xe5, 0xe4, 0x19, 0x27, 0x8f, 0x10, 0x03, 0xea, 0x98, 0xe2, 0x06, 0x79,
	0x66, 0x68, 0xa6, 0xd6, 0x6f, 0xfa, 0x97, 0x92, 0x74, 0xa1, 0x29, 0xd2, 0x38, 0xd8, 0x87, 0xe9,
	0x37, 0x1a, 0xa5, 0x9c, 0x35, 0x44, 0x1a, 0xaf, 0x4f, 0xf5, 0x09, 0x72, 0x3c, 0x9c, 0x61, 0x59,
	0x41, 0x8e, 0x07, 0x05, 0x9f, 0x40, 0x57, 0x02, 0x41, 0x76, 0x94, 0x68, 0x54, 0x4c, 0xad, 0xdf,
	0x1e, 0xdf, 0x59, 0x57, 0x0d, 0x4b, 0x0d, 0x5f, 0x1d, 0x25, 0xfa, 0x10, 0xfd, 0xbf, 0xc9, 0x10,
	0xea, 0x61, 0xbc, 0x67, 0x11, 0xee, 0x8c, 0xaa, 0x59, 0xee, 0xeb, 0x63, 0x52, 0xfc, 0x64, 0xe7,
	0xc8, 0xbf, 0x44, 0x7a, 0x03, 0xa8, 0xa9, 0x16, 0x31, 0x41, 0x8f, 0x71, 0x17, 0x6d, 0x99, 0xcc,
	0x98, 0xe0, 0x67, 0xd9, 0x62, 0x6b, 0xb0, 0x04, 0xb8, 0xce, 0x24, 0x5d, 0xb8, 0x77, 0x5e, 0xec,
	0xc5, 0x94, 0x06, 0xab, 0x77, 0x8f, 0x06, 0x6f, 0x8b, 0x57, 0x8f, 0x3a, 0xb3, 0xe7, 0x19, 0x75,
	0x3b, 0x37, 0xa4, 0x09, 0x55, 0xdb, 0x75, 0xa9, 0xdb, 0xd1, 0x88, 0x0e, 0x75, 0x9f, 0xce, 0x97,
	0x6b, 0xea, 0x76, 0x4a, 0xa4, 0x05, 0x8d, 0xf9, 0xd2, 0x55, 0xa9, 0xf2, 0x64, 0x08, 0xed, 0x48,
	0x6c, 0x0a, 0x7a, 0x93, 0xdb, 0xe2, 0x5e, 0xbd, 0xd3, 0xe6, 0x3d, 0xed, 0xa7, 0x54, 0x99, 0xda,
	0xde, 0xec, 0xa3, 0x96, 0x5f, 0xe2, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xa3, 0xa2, 0xe2,
	0xab, 0x01, 0x00, 0x00,
}
