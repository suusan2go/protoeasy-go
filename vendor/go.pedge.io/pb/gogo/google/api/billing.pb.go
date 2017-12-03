// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/billing.proto

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Billing related configuration of the service.
//
// The following example shows how to configure metrics for billing:
//
//     metrics:
//     - name: library.googleapis.com/read_calls
//       metric_kind: DELTA
//       value_type: INT64
//     - name: library.googleapis.com/write_calls
//       metric_kind: DELTA
//       value_type: INT64
//     billing:
//       metrics:
//       - library.googleapis.com/read_calls
//       - library.googleapis.com/write_calls
//
// The next example shows how to enable billing status check and customize the
// check behavior. It makes sure billing status check is included in the `Check`
// method of [Service Control API](https://cloud.google.com/service-control/).
// In the example, "google.storage.Get" method can be served when the billing
// status is either `current` or `delinquent`, while "google.storage.Write"
// method can only be served when the billing status is `current`:
//
//     billing:
//       rules:
//       - selector: google.storage.Get
//         allowed_statuses:
//         - current
//         - delinquent
//       - selector: google.storage.Write
//         allowed_statuses: current
//
// Mostly services should only allow `current` status when serving requests.
// In addition, services can choose to allow both `current` and `delinquent`
// statuses when serving read-only requests to resources. If there's no
// matching selector for operation, no billing status check will be performed.
//
type Billing struct {
	// Names of the metrics to report to billing. Each name must
	// be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics []string `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// A list of billing status rules for configuring billing status check.
	Rules []*BillingStatusRule `protobuf:"bytes,5,rep,name=rules" json:"rules,omitempty"`
}

func (m *Billing) Reset()                    { *m = Billing{} }
func (m *Billing) String() string            { return proto.CompactTextString(m) }
func (*Billing) ProtoMessage()               {}
func (*Billing) Descriptor() ([]byte, []int) { return fileDescriptorBilling, []int{0} }

func (m *Billing) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Billing) GetRules() []*BillingStatusRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Defines the billing status requirements for operations.
//
// When used with
// [Service Control API](https://cloud.google.com/service-control/), the
// following statuses are supported:
//
// - **current**: the associated billing account is up to date and capable of
//                paying for resource usages.
// - **delinquent**: the associated billing account has a correctable problem,
//                   such as late payment.
//
// Mostly services should only allow `current` status when serving requests.
// In addition, services can choose to allow both `current` and `delinquent`
// statuses when serving read-only requests to resources. If the list of
// allowed_statuses is empty, it means no billing requirement.
//
type BillingStatusRule struct {
	// Selects the operation names to which this rule applies.
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Allowed billing statuses. The billing status check passes if the actual
	// billing status matches any of the provided values here.
	AllowedStatuses []string `protobuf:"bytes,2,rep,name=allowed_statuses,json=allowedStatuses" json:"allowed_statuses,omitempty"`
}

func (m *BillingStatusRule) Reset()                    { *m = BillingStatusRule{} }
func (m *BillingStatusRule) String() string            { return proto.CompactTextString(m) }
func (*BillingStatusRule) ProtoMessage()               {}
func (*BillingStatusRule) Descriptor() ([]byte, []int) { return fileDescriptorBilling, []int{1} }

func (m *BillingStatusRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *BillingStatusRule) GetAllowedStatuses() []string {
	if m != nil {
		return m.AllowedStatuses
	}
	return nil
}

func init() {
	proto.RegisterType((*Billing)(nil), "google.api.Billing")
	proto.RegisterType((*BillingStatusRule)(nil), "google.api.BillingStatusRule")
}

func init() { proto.RegisterFile("google/api/billing.proto", fileDescriptorBilling) }

var fileDescriptorBilling = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x59, 0x4b, 0xad, 0x1d, 0xc5, 0x3f, 0xb9, 0x18, 0x16, 0x85, 0xd0, 0xd3, 0x7a, 0x49,
	0xc1, 0x7e, 0x83, 0xfd, 0x04, 0x25, 0xbd, 0x88, 0x17, 0x49, 0xd7, 0x61, 0x09, 0x4c, 0x33, 0x4b,
	0x92, 0xc5, 0xaf, 0x2f, 0x6e, 0xa2, 0x2e, 0xf4, 0x38, 0xf3, 0x7b, 0x33, 0xef, 0x3d, 0x90, 0x3d,
	0x73, 0x4f, 0xb8, 0xb5, 0x83, 0xdb, 0x1e, 0x1d, 0x91, 0xf3, 0xbd, 0x1e, 0x02, 0x27, 0x16, 0x90,
	0x89, 0xb6, 0x83, 0xab, 0x9f, 0x66, 0x2a, 0xeb, 0x3d, 0x27, 0x9b, 0x1c, 0xfb, 0x98, 0x95, 0xf5,
	0xe3, 0x8c, 0x9e, 0x30, 0x05, 0xd7, 0x65, 0xb0, 0x79, 0x83, 0x55, 0x9b, 0x7f, 0x0a, 0x09, 0xab,
	0x8c, 0xa2, 0xac, 0xd4, 0xa2, 0x59, 0x9b, 0xdf, 0x51, 0xec, 0x60, 0x19, 0x46, 0xc2, 0x28, 0x97,
	0x6a, 0xd1, 0x5c, 0xbf, 0x3e, 0xeb, 0x7f, 0x5f, 0x5d, 0xae, 0x0f, 0xc9, 0xa6, 0x31, 0x9a, 0x91,
	0xd0, 0x64, 0xed, 0xe6, 0x1d, 0x1e, 0xce, 0x98, 0xa8, 0xe1, 0x2a, 0x22, 0x61, 0x97, 0x38, 0xc8,
	0x4a, 0x55, 0xcd, 0xda, 0xfc, 0xcd, 0xe2, 0x05, 0xee, 0x2d, 0x11, 0x7f, 0xe1, 0xe7, 0x47, 0x9c,
	0x2e, 0x30, 0xca, 0x8b, 0x29, 0xc8, 0x5d, 0xd9, 0x1f, 0xca, 0xba, 0x55, 0x70, 0xdb, 0xf1, 0x69,
	0x16, 0xa3, 0xbd, 0x29, 0x5e, 0xfb, 0x9f, 0x56, 0xfb, 0xea, 0x78, 0x39, 0xd5, 0xdb, 0x7d, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x80, 0xdf, 0xe4, 0x3d, 0x01, 0x00, 0x00,
}
