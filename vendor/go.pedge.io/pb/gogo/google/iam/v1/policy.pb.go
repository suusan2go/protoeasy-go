// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/iam/v1/policy.proto

package google_iam_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines an Identity and Access Management (IAM) policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list of
// `members` to a `role`, where the members can be user accounts, Google groups,
// Google domains, and service accounts. A `role` is a named list of permissions
// defined by IAM.
//
// **Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//             "serviceAccount:my-other-app@appspot.gserviceaccount.com",
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam).
type Policy struct {
	// Version of the `Policy`. The default version is 0.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Associates a list of `members` to a `role`.
	// Multiple `bindings` must not be specified for the same `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `protobuf:"bytes,4,rep,name=bindings" json:"bindings,omitempty"`
	// `etag` is used for optimistic concurrency control as a way to help
	// prevent simultaneous updates of a policy from overwriting each other.
	// It is strongly suggested that systems make use of the `etag` in the
	// read-modify-write cycle to perform policy updates in order to avoid race
	// conditions: An `etag` is returned in the response to `getIamPolicy`, and
	// systems are expected to put that etag in the request to `setIamPolicy` to
	// ensure that their change will be applied to the same version of the policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the existing
	// policy is overwritten blindly.
	Etag []byte `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptorPolicy, []int{0} }

func (m *Policy) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Policy) GetBindings() []*Binding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func (m *Policy) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

// Associates `members` with a `role`.
type Binding struct {
	// Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	// Required
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// Specifies the identities requesting access for a Cloud Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents anyone
	//    who is authenticated with a Google account or a service account.
	//
	// * `user:{emailid}`: An email address that represents a specific Google
	//    account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a service
	//    account. For example, `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google group.
	//    For example, `admins@example.com`.
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all the
	//    users of that domain. For example, `google.com` or `example.com`.
	//
	//
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *Binding) Reset()                    { *m = Binding{} }
func (m *Binding) String() string            { return proto.CompactTextString(m) }
func (*Binding) ProtoMessage()               {}
func (*Binding) Descriptor() ([]byte, []int) { return fileDescriptorPolicy, []int{1} }

func (m *Binding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Binding) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "google.iam.v1.Policy")
	proto.RegisterType((*Binding)(nil), "google.iam.v1.Binding")
}

func init() { proto.RegisterFile("google/iam/v1/policy.proto", fileDescriptorPolicy) }

var fileDescriptorPolicy = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0xe5, 0xa6, 0xb4, 0xd4, 0x85, 0x01, 0x23, 0x21, 0xab, 0x62, 0x88, 0x3a, 0x65, 0x72,
	0x94, 0x32, 0x30, 0xb0, 0x85, 0x01, 0xb1, 0x45, 0x1e, 0xd8, 0x9d, 0xd6, 0x8a, 0x1e, 0x8a, 0xfd,
	0x22, 0xdb, 0x44, 0xe2, 0x2f, 0xf1, 0x0b, 0x19, 0x51, 0xec, 0x16, 0x29, 0xdb, 0x3b, 0xdd, 0xf7,
	0x74, 0x77, 0x74, 0xd7, 0x21, 0x76, 0xbd, 0x2e, 0x41, 0x99, 0x72, 0xac, 0xca, 0x01, 0x7b, 0x38,
	0x7e, 0x8b, 0xc1, 0x61, 0x40, 0x76, 0x9b, 0x3c, 0x01, 0xca, 0x88, 0xb1, 0xda, 0x3d, 0x9e, 0x51,
	0x35, 0x40, 0xa9, 0xac, 0xc5, 0xa0, 0x02, 0xa0, 0xf5, 0x09, 0xde, 0x7f, 0xd2, 0x55, 0x13, 0x9f,
	0x19, 0xa7, 0xeb, 0x51, 0x3b, 0x0f, 0x68, 0x39, 0xc9, 0x49, 0x71, 0x25, 0x2f, 0x92, 0x1d, 0xe8,
	0x75, 0x0b, 0xf6, 0x04, 0xb6, 0xf3, 0x7c, 0x99, 0x67, 0xc5, 0xf6, 0xf0, 0x20, 0x66, 0x19, 0xa2,
	0x4e, 0xb6, 0xfc, 0xe7, 0x18, 0xa3, 0x4b, 0x1d, 0x54, 0xc7, 0xb3, 0x9c, 0x14, 0x37, 0x32, 0xde,
	0xfb, 0x67, 0xba, 0x3e, 0x83, 0x93, 0xed, 0xb0, 0xd7, 0x31, 0x69, 0x23, 0xe3, 0x3d, 0x15, 0x30,
	0xda, 0xb4, 0xda, 0x79, 0xbe, 0xc8, 0xb3, 0x62, 0x23, 0x2f, 0xb2, 0x7e, 0xa1, 0x77, 0x47, 0x34,
	0xf3, 0xcc, 0x7a, 0x9b, 0x7a, 0x37, 0xd3, 0x8c, 0x86, 0xfc, 0x12, 0xf2, 0xb3, 0xb8, 0x7f, 0x4b,
	0xc4, 0x6b, 0x8f, 0x5f, 0x27, 0xf1, 0xae, 0x8c, 0xf8, 0xa8, 0xda, 0x55, 0x1c, 0xfa, 0xf4, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x13, 0x85, 0xbd, 0xcb, 0x33, 0x01, 0x00, 0x00,
}
