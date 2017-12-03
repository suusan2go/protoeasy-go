// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/service.proto

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api2 "go.pedge.io/pb/gogo/google/api/experimental"
import _ "github.com/gogo/protobuf/types"
import google_protobuf5 "go.pedge.io/pb/gogo/google/protobuf"
import google_protobuf4 "go.pedge.io/pb/gogo/google/protobuf"
import google_protobuf6 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Service` is the root object of Google service configuration schema. It
// describes basic information about a service, such as the name and the
// title, and delegates other aspects to sub-sections. Each sub-section is
// either a proto message or a repeated proto message that configures a
// specific aspect, such as auth. See each proto message definition for details.
//
// Example:
//
//     type: google.api.Service
//     config_version: 3
//     name: calendar.googleapis.com
//     title: Google Calendar API
//     apis:
//     - name: google.calendar.v3.Calendar
//     authentication:
//       providers:
//       - id: google_calendar_auth
//         jwks_uri: https://www.googleapis.com/oauth2/v1/certs
//         issuer: https://securetoken.google.com
//       rules:
//       - selector: "*"
//         requirements:
//           provider_id: google_calendar_auth
type Service struct {
	// The semantic version of the service configuration. The config version
	// affects the interpretation of the service configuration. For example,
	// certain features are enabled by default for certain config versions.
	// The latest config version is `3`.
	ConfigVersion *google_protobuf6.UInt32Value `protobuf:"bytes,20,opt,name=config_version,json=configVersion" json:"config_version,omitempty"`
	// The DNS address at which this service is available,
	// e.g. `calendar.googleapis.com`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique ID for a specific instance of this message, typically assigned
	// by the client for tracking purpose. If empty, the server may choose to
	// generate one instead.
	Id string `protobuf:"bytes,33,opt,name=id,proto3" json:"id,omitempty"`
	// The product title for this service.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// The Google project that owns this service.
	ProducerProjectId string `protobuf:"bytes,22,opt,name=producer_project_id,json=producerProjectId,proto3" json:"producer_project_id,omitempty"`
	// A list of API interfaces exported by this service. Only the `name` field
	// of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration
	// author, as the remaining fields will be derived from the IDL during the
	// normalization process. It is an error to specify an API interface here
	// which cannot be resolved against the associated IDL files.
	Apis []*google_protobuf5.Api `protobuf:"bytes,3,rep,name=apis" json:"apis,omitempty"`
	// A list of all proto message types included in this API service.
	// Types referenced directly or indirectly by the `apis` are
	// automatically included.  Messages which are not referenced but
	// shall be included, such as types used by the `google.protobuf.Any` type,
	// should be listed here by name. Example:
	//
	//     types:
	//     - name: google.protobuf.Int32
	Types []*google_protobuf4.Type `protobuf:"bytes,4,rep,name=types" json:"types,omitempty"`
	// A list of all enum types included in this API service.  Enums
	// referenced directly or indirectly by the `apis` are automatically
	// included.  Enums which are not referenced but shall be included
	// should be listed here by name. Example:
	//
	//     enums:
	//     - name: google.someapi.v1.SomeEnum
	Enums []*google_protobuf4.Enum `protobuf:"bytes,5,rep,name=enums" json:"enums,omitempty"`
	// Additional API documentation.
	Documentation *Documentation `protobuf:"bytes,6,opt,name=documentation" json:"documentation,omitempty"`
	// API backend configuration.
	Backend *Backend `protobuf:"bytes,8,opt,name=backend" json:"backend,omitempty"`
	// HTTP configuration.
	Http *Http `protobuf:"bytes,9,opt,name=http" json:"http,omitempty"`
	// Quota configuration.
	Quota *Quota `protobuf:"bytes,10,opt,name=quota" json:"quota,omitempty"`
	// Auth configuration.
	Authentication *Authentication `protobuf:"bytes,11,opt,name=authentication" json:"authentication,omitempty"`
	// Context configuration.
	Context *Context `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
	// Configuration controlling usage of this service.
	Usage *Usage `protobuf:"bytes,15,opt,name=usage" json:"usage,omitempty"`
	// Configuration for network endpoints.  If this is empty, then an endpoint
	// with the same name as the service is automatically generated to service all
	// defined APIs.
	Endpoints []*Endpoint `protobuf:"bytes,18,rep,name=endpoints" json:"endpoints,omitempty"`
	// Configuration for the service control plane.
	Control *Control `protobuf:"bytes,21,opt,name=control" json:"control,omitempty"`
	// Defines the logs used by this service.
	Logs []*LogDescriptor `protobuf:"bytes,23,rep,name=logs" json:"logs,omitempty"`
	// Defines the metrics used by this service.
	Metrics []*MetricDescriptor `protobuf:"bytes,24,rep,name=metrics" json:"metrics,omitempty"`
	// Defines the monitored resources used by this service. This is required
	// by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations.
	MonitoredResources []*MonitoredResourceDescriptor `protobuf:"bytes,25,rep,name=monitored_resources,json=monitoredResources" json:"monitored_resources,omitempty"`
	// Billing configuration.
	Billing *Billing `protobuf:"bytes,26,opt,name=billing" json:"billing,omitempty"`
	// Logging configuration.
	Logging *Logging `protobuf:"bytes,27,opt,name=logging" json:"logging,omitempty"`
	// Monitoring configuration.
	Monitoring *Monitoring `protobuf:"bytes,28,opt,name=monitoring" json:"monitoring,omitempty"`
	// System parameter configuration.
	SystemParameters *SystemParameters `protobuf:"bytes,29,opt,name=system_parameters,json=systemParameters" json:"system_parameters,omitempty"`
	// Output only. The source information for this configuration if available.
	SourceInfo *SourceInfo `protobuf:"bytes,37,opt,name=source_info,json=sourceInfo" json:"source_info,omitempty"`
	// Experimental configuration.
	Experimental *google_api2.Experimental `protobuf:"bytes,101,opt,name=experimental" json:"experimental,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{0} }

func (m *Service) GetConfigVersion() *google_protobuf6.UInt32Value {
	if m != nil {
		return m.ConfigVersion
	}
	return nil
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Service) GetProducerProjectId() string {
	if m != nil {
		return m.ProducerProjectId
	}
	return ""
}

func (m *Service) GetApis() []*google_protobuf5.Api {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *Service) GetTypes() []*google_protobuf4.Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Service) GetEnums() []*google_protobuf4.Enum {
	if m != nil {
		return m.Enums
	}
	return nil
}

func (m *Service) GetDocumentation() *Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func (m *Service) GetBackend() *Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Service) GetHttp() *Http {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Service) GetQuota() *Quota {
	if m != nil {
		return m.Quota
	}
	return nil
}

func (m *Service) GetAuthentication() *Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Service) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Service) GetUsage() *Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetControl() *Control {
	if m != nil {
		return m.Control
	}
	return nil
}

func (m *Service) GetLogs() []*LogDescriptor {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Service) GetMetrics() []*MetricDescriptor {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Service) GetMonitoredResources() []*MonitoredResourceDescriptor {
	if m != nil {
		return m.MonitoredResources
	}
	return nil
}

func (m *Service) GetBilling() *Billing {
	if m != nil {
		return m.Billing
	}
	return nil
}

func (m *Service) GetLogging() *Logging {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (m *Service) GetMonitoring() *Monitoring {
	if m != nil {
		return m.Monitoring
	}
	return nil
}

func (m *Service) GetSystemParameters() *SystemParameters {
	if m != nil {
		return m.SystemParameters
	}
	return nil
}

func (m *Service) GetSourceInfo() *SourceInfo {
	if m != nil {
		return m.SourceInfo
	}
	return nil
}

func (m *Service) GetExperimental() *google_api2.Experimental {
	if m != nil {
		return m.Experimental
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "google.api.Service")
}

func init() { proto.RegisterFile("google/api/service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0xdf, 0x4e, 0xf3, 0x36,
	0x18, 0x87, 0xd5, 0x7e, 0xed, 0xc7, 0x70, 0xf9, 0x18, 0x98, 0x02, 0xa6, 0x74, 0x13, 0xec, 0x8f,
	0xa8, 0x36, 0x51, 0xa4, 0x22, 0xb1, 0x93, 0x49, 0x13, 0x05, 0xb4, 0x55, 0x62, 0x52, 0x67, 0x06,
	0xda, 0x59, 0x95, 0x26, 0x6e, 0xf0, 0x96, 0xd8, 0x9e, 0xed, 0x30, 0x7a, 0x1d, 0xbb, 0x83, 0x5d,
	0xe9, 0x14, 0xdb, 0xa1, 0x4e, 0x13, 0xce, 0x1a, 0x3f, 0xcf, 0xfb, 0xea, 0x75, 0xec, 0xfc, 0x0a,
	0x50, 0xcc, 0x79, 0x9c, 0x90, 0x8b, 0x40, 0xd0, 0x0b, 0x45, 0xe4, 0x0b, 0x0d, 0xc9, 0x50, 0x48,
	0xae, 0x39, 0x04, 0x96, 0x0c, 0x03, 0x41, 0x7b, 0x7d, 0xcf, 0x0a, 0x18, 0xe3, 0x3a, 0xd0, 0x94,
	0x33, 0x65, 0xcd, 0xde, 0xbe, 0x4f, 0x33, 0xfd, 0xec, 0x96, 0xfd, 0xd6, 0xf3, 0x20, 0xfc, 0x8b,
	0xb0, 0xa8, 0x8e, 0xd0, 0x24, 0xa1, 0x2c, 0xae, 0x21, 0x21, 0x67, 0x9a, 0xbc, 0xea, 0x77, 0x88,
	0xe4, 0x89, 0x23, 0x5f, 0x7a, 0x24, 0xe2, 0x61, 0x96, 0x12, 0x66, 0xe7, 0x73, 0xfc, 0xc8, 0xe3,
	0x84, 0x45, 0x82, 0x53, 0x56, 0x34, 0xfd, 0xce, 0x47, 0xaf, 0x82, 0x48, 0x6a, 0x8a, 0x93, 0xd2,
	0x43, 0xcd, 0x2e, 0x9f, 0xb5, 0x16, 0x6e, 0xf9, 0xc0, 0x5b, 0x4e, 0x82, 0x39, 0x29, 0xf4, 0xae,
	0xbf, 0xce, 0xeb, 0xf6, 0x97, 0xf0, 0x38, 0x5e, 0xed, 0xfc, 0xd0, 0x23, 0x29, 0xd1, 0x92, 0x86,
	0x0e, 0x7c, 0xed, 0x03, 0xce, 0xa8, 0xe6, 0x92, 0x44, 0x33, 0x49, 0x14, 0xcf, 0x64, 0x71, 0x58,
	0xbd, 0xe3, 0xaa, 0xb4, 0x6a, 0xed, 0x8f, 0xf8, 0x77, 0xc6, 0x75, 0xe0, 0xd6, 0xfd, 0x53, 0xb5,
	0xdd, 0x66, 0x94, 0x2d, 0xb8, 0xa3, 0xa7, 0x3e, 0x5d, 0x2a, 0x4d, 0xd2, 0x99, 0x08, 0x64, 0x90,
	0x12, 0x4d, 0x64, 0x4d, 0xe3, 0x4c, 0x05, 0x31, 0x59, 0x7b, 0xe3, 0xe6, 0x69, 0x9e, 0x2d, 0x2e,
	0x02, 0xb6, 0x7c, 0x17, 0x09, 0xea, 0x50, 0x6f, 0x1d, 0xe9, 0xa5, 0x20, 0x6b, 0x67, 0xfc, 0xc6,
	0xfe, 0x91, 0x81, 0x10, 0x44, 0xba, 0x2b, 0xf8, 0xd5, 0xbf, 0x00, 0x6c, 0x3c, 0xd8, 0xeb, 0x0b,
	0x6f, 0xc0, 0x76, 0xc8, 0xd9, 0x82, 0xc6, 0xb3, 0x17, 0x22, 0x15, 0xe5, 0x0c, 0x75, 0x4f, 0x1a,
	0x83, 0xce, 0xa8, 0x3f, 0x74, 0x37, 0xba, 0x68, 0x32, 0x7c, 0x9c, 0x30, 0x7d, 0x39, 0x7a, 0x0a,
	0x92, 0x8c, 0xe0, 0x4f, 0xb6, 0xe6, 0xc9, 0x96, 0x40, 0x08, 0x5a, 0x2c, 0x48, 0x09, 0x6a, 0x9c,
	0x34, 0x06, 0x9b, 0xd8, 0xfc, 0x86, 0xdb, 0xa0, 0x49, 0x23, 0x74, 0x6a, 0x56, 0x9a, 0x34, 0x82,
	0x5d, 0xd0, 0xd6, 0x54, 0x27, 0x04, 0x35, 0xcd, 0x92, 0x7d, 0x80, 0x43, 0xb0, 0x27, 0x24, 0x8f,
	0xb2, 0x90, 0xc8, 0x99, 0x90, 0xfc, 0x4f, 0x12, 0xea, 0x19, 0x8d, 0xd0, 0x81, 0x71, 0x76, 0x0b,
	0x34, 0xb5, 0x64, 0x12, 0xc1, 0x01, 0x68, 0x05, 0x82, 0x2a, 0xf4, 0xe1, 0xe4, 0xc3, 0xa0, 0x33,
	0xea, 0x56, 0x86, 0xbc, 0x16, 0x14, 0x1b, 0x03, 0x7e, 0x0f, 0xda, 0xf9, 0x2b, 0x51, 0xa8, 0x65,
	0xd4, 0xfd, 0x8a, 0xfa, 0xfb, 0x52, 0x10, 0x6c, 0x9d, 0x5c, 0x26, 0x2c, 0x4b, 0x15, 0x6a, 0xbf,
	0x23, 0xdf, 0xb1, 0x2c, 0xc5, 0xd6, 0x81, 0x3f, 0x81, 0x4f, 0xa5, 0x2f, 0x07, 0x7d, 0x34, 0x6f,
	0xec, 0x68, 0xb8, 0xca, 0x80, 0xe1, 0xad, 0x2f, 0xe0, 0xb2, 0x0f, 0xcf, 0xc1, 0x86, 0xfb, 0xc4,
	0xd1, 0x67, 0xa6, 0x74, 0xcf, 0x2f, 0x1d, 0x5b, 0x84, 0x0b, 0x07, 0x7e, 0x03, 0x5a, 0xf9, 0x27,
	0x84, 0x36, 0x8d, 0xbb, 0xe3, 0xbb, 0xbf, 0x68, 0x2d, 0xb0, 0xa1, 0xf0, 0x0c, 0xb4, 0xcd, 0x75,
	0x45, 0xc0, 0x68, 0xbb, 0xbe, 0xf6, 0x5b, 0x0e, 0xb0, 0xe5, 0x70, 0x0c, 0xb6, 0xf3, 0xdc, 0x21,
	0x4c, 0xd3, 0xd0, 0xce, 0xdf, 0x31, 0x15, 0x3d, 0xbf, 0xe2, 0xba, 0x64, 0xe0, 0xb5, 0x8a, 0x7c,
	0x07, 0x2e, 0x70, 0xd0, 0x56, 0x75, 0x07, 0x37, 0x16, 0xe1, 0xc2, 0xc9, 0x67, 0x33, 0x37, 0x1e,
	0x7d, 0x5e, 0x9d, 0xed, 0x31, 0x07, 0xd8, 0x72, 0x38, 0x02, 0x9b, 0x45, 0xe8, 0x28, 0x04, 0xcb,
	0x67, 0x9c, 0xcb, 0x77, 0x0e, 0xe2, 0x95, 0x56, 0xcc, 0x22, 0x79, 0x82, 0xf6, 0xeb, 0x67, 0x91,
	0x3c, 0xc1, 0x85, 0x03, 0xcf, 0x41, 0x2b, 0xe1, 0xb1, 0x42, 0x87, 0xa6, 0x7b, 0xe9, 0xd0, 0xee,
	0x79, 0x7c, 0x4b, 0x54, 0x28, 0xa9, 0xd0, 0x5c, 0x62, 0xa3, 0xc1, 0x2b, 0xb0, 0x61, 0x03, 0x46,
	0x21, 0x64, 0x2a, 0xfa, 0x7e, 0xc5, 0xaf, 0x06, 0x79, 0x45, 0x85, 0x0c, 0xff, 0x00, 0x7b, 0xd5,
	0xfc, 0x51, 0xe8, 0xc8, 0xf4, 0x38, 0x2b, 0xf5, 0x28, 0x34, 0xec, 0x2c, 0xaf, 0x1d, 0x4c, 0xd7,
	0xa1, 0xd9, 0xaf, 0xfb, 0x1b, 0x40, 0xbd, 0x9a, 0xdb, 0x63, 0x11, 0x2e, 0x9c, 0x5c, 0x77, 0xd9,
	0x89, 0x8e, 0xab, 0xfa, 0xbd, 0x45, 0xb8, 0x70, 0xe0, 0x15, 0x00, 0xab, 0x48, 0x44, 0x7d, 0x53,
	0x71, 0x50, 0x33, 0x6e, 0x5e, 0xe4, 0x99, 0x70, 0x02, 0x76, 0xd7, 0x73, 0x4f, 0xa1, 0x2f, 0xca,
	0x51, 0x92, 0x97, 0x3f, 0x18, 0x69, 0xfa, 0xe6, 0xe0, 0x1d, 0xb5, 0xb6, 0x02, 0x7f, 0x00, 0x1d,
	0x2f, 0x60, 0xd1, 0xb7, 0xd5, 0x19, 0x1e, 0x0c, 0x9e, 0xb0, 0x05, 0xc7, 0x40, 0xbd, 0xfd, 0x86,
	0x3f, 0x82, 0x2d, 0xff, 0xaf, 0x08, 0x11, 0x53, 0x89, 0x4a, 0x17, 0xc8, 0xe3, 0xb8, 0x64, 0x8f,
	0xcf, 0xf2, 0x24, 0x4c, 0x3d, 0x79, 0xbc, 0xe5, 0x42, 0x72, 0x9a, 0xa7, 0xc0, 0xb4, 0xf1, 0x5f,
	0xb3, 0xf5, 0xf3, 0xf5, 0x74, 0x32, 0xff, 0x68, 0x52, 0xe1, 0xf2, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf2, 0xb5, 0x30, 0x24, 0x0e, 0x08, 0x00, 0x00,
}
