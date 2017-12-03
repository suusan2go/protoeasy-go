// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/iam/v1/iam_policy.proto

package google_iam_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// REQUIRED: The resource for which the policy is being specified.
	// `resource` is usually specified as a path. For example, a Project
	// resource is specified as `projects/{project}`.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// REQUIRED: The complete policy to be applied to the `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as Projects)
	// might reject them.
	Policy *Policy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
}

func (m *SetIamPolicyRequest) Reset()                    { *m = SetIamPolicyRequest{} }
func (m *SetIamPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*SetIamPolicyRequest) ProtoMessage()               {}
func (*SetIamPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptorIamPolicy, []int{0} }

func (m *SetIamPolicyRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *SetIamPolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
	// REQUIRED: The resource for which the policy is being requested.
	// `resource` is usually specified as a path. For example, a Project
	// resource is specified as `projects/{project}`.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (m *GetIamPolicyRequest) Reset()                    { *m = GetIamPolicyRequest{} }
func (m *GetIamPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIamPolicyRequest) ProtoMessage()               {}
func (*GetIamPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptorIamPolicy, []int{1} }

func (m *GetIamPolicyRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

// Request message for `TestIamPermissions` method.
type TestIamPermissionsRequest struct {
	// REQUIRED: The resource for which the policy detail is being requested.
	// `resource` is usually specified as a path. For example, a Project
	// resource is specified as `projects/{project}`.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The set of permissions to check for the `resource`. Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For more
	// information see
	// [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `protobuf:"bytes,2,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *TestIamPermissionsRequest) Reset()         { *m = TestIamPermissionsRequest{} }
func (m *TestIamPermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*TestIamPermissionsRequest) ProtoMessage()    {}
func (*TestIamPermissionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorIamPolicy, []int{2}
}

func (m *TestIamPermissionsRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *TestIamPermissionsRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// Response message for `TestIamPermissions` method.
type TestIamPermissionsResponse struct {
	// A subset of `TestPermissionsRequest.permissions` that the caller is
	// allowed.
	Permissions []string `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *TestIamPermissionsResponse) Reset()         { *m = TestIamPermissionsResponse{} }
func (m *TestIamPermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*TestIamPermissionsResponse) ProtoMessage()    {}
func (*TestIamPermissionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorIamPolicy, []int{3}
}

func (m *TestIamPermissionsResponse) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*SetIamPolicyRequest)(nil), "google.iam.v1.SetIamPolicyRequest")
	proto.RegisterType((*GetIamPolicyRequest)(nil), "google.iam.v1.GetIamPolicyRequest")
	proto.RegisterType((*TestIamPermissionsRequest)(nil), "google.iam.v1.TestIamPermissionsRequest")
	proto.RegisterType((*TestIamPermissionsResponse)(nil), "google.iam.v1.TestIamPermissionsResponse")
}

func init() { proto.RegisterFile("google/iam/v1/iam_policy.proto", fileDescriptorIamPolicy) }

var fileDescriptorIamPolicy = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4c, 0xcc, 0xd5, 0x2f, 0x33, 0x04, 0x51, 0xf1, 0x05, 0xf9, 0x39, 0x99,
	0xc9, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x79, 0xbd, 0xcc, 0xc4, 0x5c,
	0xbd, 0x32, 0x43, 0x29, 0x19, 0xa8, 0xf2, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92,
	0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x62, 0x29, 0x29, 0x54, 0xc3, 0x90, 0x0d, 0x52, 0x4a,
	0xe0, 0x12, 0x0e, 0x4e, 0x2d, 0xf1, 0x4c, 0xcc, 0x0d, 0x00, 0x8b, 0x06, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x14, 0xa5, 0x16, 0xe7, 0x97, 0x16, 0x25, 0xa7, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0xba, 0x5c, 0x6c, 0x10, 0x23, 0x24, 0x98, 0x14,
	0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf5, 0x50, 0x1c, 0xa3, 0x07, 0x35, 0x09, 0xaa, 0x48, 0xc9, 0x90,
	0x4b, 0xd8, 0x9d, 0x34, 0x1b, 0x94, 0x22, 0xb9, 0x24, 0x43, 0x52, 0x8b, 0xc1, 0x7a, 0x52, 0x8b,
	0x72, 0x33, 0x8b, 0x8b, 0x41, 0x9e, 0x21, 0xc6, 0x69, 0x0a, 0x5c, 0xdc, 0x05, 0x08, 0x1d, 0x12,
	0x4c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x4a, 0x76, 0x5c, 0x52, 0xd8, 0x8c, 0x2e, 0x2e,
	0xc8, 0xcf, 0x2b, 0xc6, 0xd0, 0xcf, 0x88, 0xa1, 0xdf, 0x68, 0x0a, 0x33, 0x17, 0xa7, 0xa7, 0xa3,
	0x2f, 0xc4, 0x2f, 0x42, 0x25, 0x5c, 0x3c, 0xc8, 0xa1, 0x27, 0xa4, 0x84, 0x16, 0x14, 0x58, 0x82,
	0x56, 0x0a, 0x7b, 0x70, 0x29, 0x69, 0x36, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x59, 0x49, 0x0e, 0x14,
	0x45, 0xd5, 0x30, 0x1f, 0xd9, 0x6a, 0x69, 0xd5, 0x5a, 0x15, 0x23, 0x99, 0x62, 0xc5, 0xa8, 0x05,
	0xb2, 0xd5, 0x1d, 0x9f, 0xad, 0xee, 0x54, 0xb1, 0x35, 0x1d, 0xcd, 0xd6, 0x59, 0x8c, 0x5c, 0x42,
	0x98, 0x41, 0x27, 0xa4, 0x81, 0x66, 0x30, 0xce, 0x88, 0x93, 0xd2, 0x24, 0x42, 0x25, 0x24, 0x1e,
	0x94, 0xf4, 0xc1, 0xce, 0xd2, 0x54, 0x52, 0xc1, 0x74, 0x56, 0x09, 0x86, 0x2e, 0x2b, 0x46, 0x2d,
	0x27, 0x3b, 0x2e, 0xc1, 0xe4, 0xfc, 0x5c, 0x54, 0x0b, 0x9c, 0xf8, 0xe0, 0xee, 0x0f, 0x00, 0xa5,
	0xf5, 0x00, 0xc6, 0x1f, 0x8c, 0x8c, 0xab, 0x98, 0x84, 0xdd, 0x21, 0x8a, 0x9c, 0x73, 0xf2, 0x4b,
	0x53, 0xf4, 0x3c, 0x13, 0x73, 0xf5, 0xc2, 0x0c, 0x93, 0xd8, 0xc0, 0xb9, 0xc1, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x85, 0x39, 0xf6, 0xb7, 0x78, 0x03, 0x00, 0x00,
}
