// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/v2/logging_config.proto

package google_logging_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf6 "go.pedge.io/pb/go/google/protobuf"
import google_protobuf4 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Available log entry formats. Log entries can be written to Stackdriver
// Logging in either format and can be exported in either format.
// Version 2 is the preferred format.
type LogSink_VersionFormat int32

const (
	// An unspecified format version that will default to V2.
	LogSink_VERSION_FORMAT_UNSPECIFIED LogSink_VersionFormat = 0
	// `LogEntry` version 2 format.
	LogSink_V2 LogSink_VersionFormat = 1
	// `LogEntry` version 1 format.
	LogSink_V1 LogSink_VersionFormat = 2
)

var LogSink_VersionFormat_name = map[int32]string{
	0: "VERSION_FORMAT_UNSPECIFIED",
	1: "V2",
	2: "V1",
}
var LogSink_VersionFormat_value = map[string]int32{
	"VERSION_FORMAT_UNSPECIFIED": 0,
	"V2": 1,
	"V1": 2,
}

func (x LogSink_VersionFormat) String() string {
	return proto.EnumName(LogSink_VersionFormat_name, int32(x))
}
func (LogSink_VersionFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

// Describes a sink used to export log entries to one of the following
// destinations in any project: a Cloud Storage bucket, a BigQuery dataset, or a
// Cloud Pub/Sub topic.  A logs filter controls which log entries are
// exported. The sink must be created within a project, organization, billing
// account, or folder.
type LogSink struct {
	// Required. The client-assigned sink identifier, unique within the
	// project. Example: `"my-syslog-errors-to-pubsub"`.  Sink identifiers are
	// limited to 100 characters and can include only the following characters:
	// upper and lower-case alphanumeric characters, underscores, hyphens, and
	// periods.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Required. The export destination:
	//
	//     "storage.googleapis.com/[GCS_BUCKET]"
	//     "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]"
	//     "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]"
	//
	// The sink's `writer_identity`, set when the sink is created, must
	// have permission to write to the destination or else the log
	// entries are not exported.  For more information, see
	// [Exporting Logs With Sinks](/logging/docs/api/tasks/exporting-logs).
	Destination string `protobuf:"bytes,3,opt,name=destination" json:"destination,omitempty"`
	// Optional.
	// An [advanced logs filter](/logging/docs/view/advanced_filters).  The only
	// exported log entries are those that are in the resource owning the sink and
	// that match the filter.  For example:
	//
	//     logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR
	Filter string `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	// Deprecated. The log entry format to use for this sink's exported log
	// entries.  The v2 format is used by default and cannot be changed.
	OutputVersionFormat LogSink_VersionFormat `protobuf:"varint,6,opt,name=output_version_format,json=outputVersionFormat,enum=google.logging.v2.LogSink_VersionFormat" json:"output_version_format,omitempty"`
	// Output only. An IAM identity&mdash;a service account or group&mdash;under
	// which Stackdriver Logging writes the exported log entries to the sink's
	// destination.  This field is set by
	// [sinks.create](/logging/docs/api/reference/rest/v2/projects.sinks/create)
	// and
	// [sinks.update](/logging/docs/api/reference/rest/v2/projects.sinks/update),
	// based on the setting of `unique_writer_identity` in those methods.
	//
	// Until you grant this identity write-access to the destination, log entry
	// exports from this sink will fail. For more information,
	// see [Granting access for a
	// resource](/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource).
	// Consult the destination service's documentation to determine the
	// appropriate IAM roles to assign to the identity.
	WriterIdentity string `protobuf:"bytes,8,opt,name=writer_identity,json=writerIdentity" json:"writer_identity,omitempty"`
	// Optional. This field applies only to sinks owned by organizations and
	// folders. If the field is false, the default, only the logs owned by the
	// sink's parent resource are available for export. If the field is true, then
	// logs from all the projects, folders, and billing accounts contained in the
	// sink's parent resource are also available for export. Whether a particular
	// log entry from the children is exported depends on the sink's filter
	// expression. For example, if this field is true, then the filter
	// `resource.type=gce_instance` would export all Compute Engine VM instance
	// log entries from all projects in the sink's parent. To only export entries
	// from certain child projects, filter on the project part of the log name:
	//
	//     logName:("projects/test-project1/" OR "projects/test-project2/") AND
	//     resource.type=gce_instance
	IncludeChildren bool `protobuf:"varint,9,opt,name=include_children,json=includeChildren" json:"include_children,omitempty"`
	// Deprecated. This field is ignored when creating or updating sinks.
	StartTime *google_protobuf4.Timestamp `protobuf:"bytes,10,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Deprecated. This field is ignored when creating or updating sinks.
	EndTime *google_protobuf4.Timestamp `protobuf:"bytes,11,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *LogSink) Reset()                    { *m = LogSink{} }
func (m *LogSink) String() string            { return proto.CompactTextString(m) }
func (*LogSink) ProtoMessage()               {}
func (*LogSink) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *LogSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogSink) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *LogSink) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *LogSink) GetOutputVersionFormat() LogSink_VersionFormat {
	if m != nil {
		return m.OutputVersionFormat
	}
	return LogSink_VERSION_FORMAT_UNSPECIFIED
}

func (m *LogSink) GetWriterIdentity() string {
	if m != nil {
		return m.WriterIdentity
	}
	return ""
}

func (m *LogSink) GetIncludeChildren() bool {
	if m != nil {
		return m.IncludeChildren
	}
	return false
}

func (m *LogSink) GetStartTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *LogSink) GetEndTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// The parameters to `ListSinks`.
type ListSinksRequest struct {
	// Required. The parent resource whose sinks are to be listed:
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListSinksRequest) Reset()                    { *m = ListSinksRequest{} }
func (m *ListSinksRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSinksRequest) ProtoMessage()               {}
func (*ListSinksRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ListSinksRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListSinksRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListSinksRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Result returned from `ListSinks`.
type ListSinksResponse struct {
	// A list of sinks.
	Sinks []*LogSink `protobuf:"bytes,1,rep,name=sinks" json:"sinks,omitempty"`
	// If there might be more results than appear in this response, then
	// `nextPageToken` is included.  To get the next set of results, call the same
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListSinksResponse) Reset()                    { *m = ListSinksResponse{} }
func (m *ListSinksResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSinksResponse) ProtoMessage()               {}
func (*ListSinksResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ListSinksResponse) GetSinks() []*LogSink {
	if m != nil {
		return m.Sinks
	}
	return nil
}

func (m *ListSinksResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The parameters to `GetSink`.
type GetSinkRequest struct {
	// Required. The resource name of the sink:
	//
	//     "projects/[PROJECT_ID]/sinks/[SINK_ID]"
	//     "organizations/[ORGANIZATION_ID]/sinks/[SINK_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/sinks/[SINK_ID]"
	//     "folders/[FOLDER_ID]/sinks/[SINK_ID]"
	//
	// Example: `"projects/my-project-id/sinks/my-sink-id"`.
	SinkName string `protobuf:"bytes,1,opt,name=sink_name,json=sinkName" json:"sink_name,omitempty"`
}

func (m *GetSinkRequest) Reset()                    { *m = GetSinkRequest{} }
func (m *GetSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSinkRequest) ProtoMessage()               {}
func (*GetSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *GetSinkRequest) GetSinkName() string {
	if m != nil {
		return m.SinkName
	}
	return ""
}

// The parameters to `CreateSink`.
type CreateSinkRequest struct {
	// Required. The resource in which to create the sink:
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	//
	// Examples: `"projects/my-logging-project"`, `"organizations/123456789"`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The new sink, whose `name` parameter is a sink identifier that
	// is not already in use.
	Sink *LogSink `protobuf:"bytes,2,opt,name=sink" json:"sink,omitempty"`
	// Optional. Determines the kind of IAM identity returned as `writer_identity`
	// in the new sink.  If this value is omitted or set to false, and if the
	// sink's parent is a project, then the value returned as `writer_identity` is
	// the same group or service account used by Stackdriver Logging before the
	// addition of writer identities to this API. The sink's destination must be
	// in the same project as the sink itself.
	//
	// If this field is set to true, or if the sink is owned by a non-project
	// resource such as an organization, then the value of `writer_identity` will
	// be a unique service account used only for exports from the new sink.  For
	// more information, see `writer_identity` in [LogSink][google.logging.v2.LogSink].
	UniqueWriterIdentity bool `protobuf:"varint,3,opt,name=unique_writer_identity,json=uniqueWriterIdentity" json:"unique_writer_identity,omitempty"`
}

func (m *CreateSinkRequest) Reset()                    { *m = CreateSinkRequest{} }
func (m *CreateSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSinkRequest) ProtoMessage()               {}
func (*CreateSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CreateSinkRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateSinkRequest) GetSink() *LogSink {
	if m != nil {
		return m.Sink
	}
	return nil
}

func (m *CreateSinkRequest) GetUniqueWriterIdentity() bool {
	if m != nil {
		return m.UniqueWriterIdentity
	}
	return false
}

// The parameters to `UpdateSink`.
type UpdateSinkRequest struct {
	// Required. The full resource name of the sink to update, including the
	// parent resource and the sink identifier:
	//
	//     "projects/[PROJECT_ID]/sinks/[SINK_ID]"
	//     "organizations/[ORGANIZATION_ID]/sinks/[SINK_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/sinks/[SINK_ID]"
	//     "folders/[FOLDER_ID]/sinks/[SINK_ID]"
	//
	// Example: `"projects/my-project-id/sinks/my-sink-id"`.
	SinkName string `protobuf:"bytes,1,opt,name=sink_name,json=sinkName" json:"sink_name,omitempty"`
	// Required. The updated sink, whose name is the same identifier that appears
	// as part of `sink_name`.
	Sink *LogSink `protobuf:"bytes,2,opt,name=sink" json:"sink,omitempty"`
	// Optional. See
	// [sinks.create](/logging/docs/api/reference/rest/v2/projects.sinks/create)
	// for a description of this field.  When updating a sink, the effect of this
	// field on the value of `writer_identity` in the updated sink depends on both
	// the old and new values of this field:
	//
	// +   If the old and new values of this field are both false or both true,
	//     then there is no change to the sink's `writer_identity`.
	// +   If the old value is false and the new value is true, then
	//     `writer_identity` is changed to a unique service account.
	// +   It is an error if the old value is true and the new value is
	//     set to false or defaulted to false.
	UniqueWriterIdentity bool `protobuf:"varint,3,opt,name=unique_writer_identity,json=uniqueWriterIdentity" json:"unique_writer_identity,omitempty"`
	// Optional. Field mask that specifies the fields in `sink` that need
	// an update. A sink field will be overwritten if, and only if, it is
	// in the update mask.  `name` and output only fields cannot be updated.
	//
	// An empty updateMask is temporarily treated as using the following mask
	// for backwards compatibility purposes:
	//   destination,filter,includeChildren
	// At some point in the future, behavior will be removed and specifying an
	// empty updateMask will be an error.
	//
	// For a detailed `FieldMask` definition, see
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	//
	// Example: `updateMask=filter`.
	UpdateMask *google_protobuf6.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateSinkRequest) Reset()                    { *m = UpdateSinkRequest{} }
func (m *UpdateSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSinkRequest) ProtoMessage()               {}
func (*UpdateSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *UpdateSinkRequest) GetSinkName() string {
	if m != nil {
		return m.SinkName
	}
	return ""
}

func (m *UpdateSinkRequest) GetSink() *LogSink {
	if m != nil {
		return m.Sink
	}
	return nil
}

func (m *UpdateSinkRequest) GetUniqueWriterIdentity() bool {
	if m != nil {
		return m.UniqueWriterIdentity
	}
	return false
}

func (m *UpdateSinkRequest) GetUpdateMask() *google_protobuf6.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The parameters to `DeleteSink`.
type DeleteSinkRequest struct {
	// Required. The full resource name of the sink to delete, including the
	// parent resource and the sink identifier:
	//
	//     "projects/[PROJECT_ID]/sinks/[SINK_ID]"
	//     "organizations/[ORGANIZATION_ID]/sinks/[SINK_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/sinks/[SINK_ID]"
	//     "folders/[FOLDER_ID]/sinks/[SINK_ID]"
	//
	// Example: `"projects/my-project-id/sinks/my-sink-id"`.
	SinkName string `protobuf:"bytes,1,opt,name=sink_name,json=sinkName" json:"sink_name,omitempty"`
}

func (m *DeleteSinkRequest) Reset()                    { *m = DeleteSinkRequest{} }
func (m *DeleteSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSinkRequest) ProtoMessage()               {}
func (*DeleteSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *DeleteSinkRequest) GetSinkName() string {
	if m != nil {
		return m.SinkName
	}
	return ""
}

// Specifies a set of log entries that are not to be stored in Stackdriver
// Logging. If your project receives a large volume of logs, you might be able
// to use exclusions to reduce your chargeable logs. Exclusions are processed
// after log sinks, so you can export log entries before they are excluded.
// Audit log entries and log entries from Amazon Web Services are never
// excluded.
type LogExclusion struct {
	// Required. A client-assigned identifier, such as
	// `"load-balancer-exclusion"`. Identifiers are limited to 100 characters and
	// can include only letters, digits, underscores, hyphens, and periods.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optional. A description of this exclusion.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Required.
	// An [advanced logs filter](/logging/docs/view/advanced_filters)
	// that matches the log entries to be excluded. By using the
	// [sample function](/logging/docs/view/advanced_filters#sample),
	// you can exclude less than 100% of the matching log entries.
	// For example, the following filter matches 99% of low-severity log
	// entries from load balancers:
	//
	//     "resource.type=http_load_balancer severity<ERROR sample(insertId, 0.99)"
	Filter string `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
	// Optional. If set to True, then this exclusion is disabled and it does not
	// exclude any log entries. You can use
	// [exclusions.patch](/logging/docs/alpha-exclusion/docs/reference/v2/rest/v2/projects.exclusions/patch)
	// to change the value of this field.
	Disabled bool `protobuf:"varint,4,opt,name=disabled" json:"disabled,omitempty"`
}

func (m *LogExclusion) Reset()                    { *m = LogExclusion{} }
func (m *LogExclusion) String() string            { return proto.CompactTextString(m) }
func (*LogExclusion) ProtoMessage()               {}
func (*LogExclusion) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *LogExclusion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogExclusion) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LogExclusion) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *LogExclusion) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

// The parameters to `ListExclusions`.
type ListExclusionsRequest struct {
	// Required. The parent resource whose exclusions are to be listed.
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListExclusionsRequest) Reset()                    { *m = ListExclusionsRequest{} }
func (m *ListExclusionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListExclusionsRequest) ProtoMessage()               {}
func (*ListExclusionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ListExclusionsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListExclusionsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListExclusionsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Result returned from `ListExclusions`.
type ListExclusionsResponse struct {
	// A list of exclusions.
	Exclusions []*LogExclusion `protobuf:"bytes,1,rep,name=exclusions" json:"exclusions,omitempty"`
	// If there might be more results than appear in this response, then
	// `nextPageToken` is included.  To get the next set of results, call the same
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListExclusionsResponse) Reset()                    { *m = ListExclusionsResponse{} }
func (m *ListExclusionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListExclusionsResponse) ProtoMessage()               {}
func (*ListExclusionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *ListExclusionsResponse) GetExclusions() []*LogExclusion {
	if m != nil {
		return m.Exclusions
	}
	return nil
}

func (m *ListExclusionsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The parameters to `GetExclusion`.
type GetExclusionRequest struct {
	// Required. The resource name of an existing exclusion:
	//
	//     "projects/[PROJECT_ID]/exclusions/[EXCLUSION_ID]"
	//     "organizations/[ORGANIZATION_ID]/exclusions/[EXCLUSION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/exclusions/[EXCLUSION_ID]"
	//     "folders/[FOLDER_ID]/exclusions/[EXCLUSION_ID]"
	//
	// Example: `"projects/my-project-id/exclusions/my-exclusion-id"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetExclusionRequest) Reset()                    { *m = GetExclusionRequest{} }
func (m *GetExclusionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetExclusionRequest) ProtoMessage()               {}
func (*GetExclusionRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *GetExclusionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The parameters to `CreateExclusion`.
type CreateExclusionRequest struct {
	// Required. The parent resource in which to create the exclusion:
	//
	//     "projects/[PROJECT_ID]"
	//     "organizations/[ORGANIZATION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]"
	//     "folders/[FOLDER_ID]"
	//
	// Examples: `"projects/my-logging-project"`, `"organizations/123456789"`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The new exclusion, whose `name` parameter is an exclusion name
	// that is not already used in the parent resource.
	Exclusion *LogExclusion `protobuf:"bytes,2,opt,name=exclusion" json:"exclusion,omitempty"`
}

func (m *CreateExclusionRequest) Reset()                    { *m = CreateExclusionRequest{} }
func (m *CreateExclusionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateExclusionRequest) ProtoMessage()               {}
func (*CreateExclusionRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *CreateExclusionRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateExclusionRequest) GetExclusion() *LogExclusion {
	if m != nil {
		return m.Exclusion
	}
	return nil
}

// The parameters to `UpdateExclusion`.
type UpdateExclusionRequest struct {
	// Required. The resource name of the exclusion to update:
	//
	//     "projects/[PROJECT_ID]/exclusions/[EXCLUSION_ID]"
	//     "organizations/[ORGANIZATION_ID]/exclusions/[EXCLUSION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/exclusions/[EXCLUSION_ID]"
	//     "folders/[FOLDER_ID]/exclusions/[EXCLUSION_ID]"
	//
	// Example: `"projects/my-project-id/exclusions/my-exclusion-id"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Required. New values for the existing exclusion. Only the fields specified
	// in `update_mask` are relevant.
	Exclusion *LogExclusion `protobuf:"bytes,2,opt,name=exclusion" json:"exclusion,omitempty"`
	// Required. A nonempty list of fields to change in the existing exclusion.
	// New values for the fields are taken from the corresponding fields in the
	// [LogExclusion][google.logging.v2.LogExclusion] included in this request. Fields not mentioned in
	// `update_mask` are not changed and are ignored in the request.
	//
	// For example, to change the filter and description of an exclusion,
	// specify an `update_mask` of `"filter,description"`.
	UpdateMask *google_protobuf6.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateExclusionRequest) Reset()                    { *m = UpdateExclusionRequest{} }
func (m *UpdateExclusionRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateExclusionRequest) ProtoMessage()               {}
func (*UpdateExclusionRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *UpdateExclusionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateExclusionRequest) GetExclusion() *LogExclusion {
	if m != nil {
		return m.Exclusion
	}
	return nil
}

func (m *UpdateExclusionRequest) GetUpdateMask() *google_protobuf6.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The parameters to `DeleteExclusion`.
type DeleteExclusionRequest struct {
	// Required. The resource name of an existing exclusion to delete:
	//
	//     "projects/[PROJECT_ID]/exclusions/[EXCLUSION_ID]"
	//     "organizations/[ORGANIZATION_ID]/exclusions/[EXCLUSION_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/exclusions/[EXCLUSION_ID]"
	//     "folders/[FOLDER_ID]/exclusions/[EXCLUSION_ID]"
	//
	// Example: `"projects/my-project-id/exclusions/my-exclusion-id"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteExclusionRequest) Reset()                    { *m = DeleteExclusionRequest{} }
func (m *DeleteExclusionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteExclusionRequest) ProtoMessage()               {}
func (*DeleteExclusionRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *DeleteExclusionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*LogSink)(nil), "google.logging.v2.LogSink")
	proto.RegisterType((*ListSinksRequest)(nil), "google.logging.v2.ListSinksRequest")
	proto.RegisterType((*ListSinksResponse)(nil), "google.logging.v2.ListSinksResponse")
	proto.RegisterType((*GetSinkRequest)(nil), "google.logging.v2.GetSinkRequest")
	proto.RegisterType((*CreateSinkRequest)(nil), "google.logging.v2.CreateSinkRequest")
	proto.RegisterType((*UpdateSinkRequest)(nil), "google.logging.v2.UpdateSinkRequest")
	proto.RegisterType((*DeleteSinkRequest)(nil), "google.logging.v2.DeleteSinkRequest")
	proto.RegisterType((*LogExclusion)(nil), "google.logging.v2.LogExclusion")
	proto.RegisterType((*ListExclusionsRequest)(nil), "google.logging.v2.ListExclusionsRequest")
	proto.RegisterType((*ListExclusionsResponse)(nil), "google.logging.v2.ListExclusionsResponse")
	proto.RegisterType((*GetExclusionRequest)(nil), "google.logging.v2.GetExclusionRequest")
	proto.RegisterType((*CreateExclusionRequest)(nil), "google.logging.v2.CreateExclusionRequest")
	proto.RegisterType((*UpdateExclusionRequest)(nil), "google.logging.v2.UpdateExclusionRequest")
	proto.RegisterType((*DeleteExclusionRequest)(nil), "google.logging.v2.DeleteExclusionRequest")
	proto.RegisterEnum("google.logging.v2.LogSink_VersionFormat", LogSink_VersionFormat_name, LogSink_VersionFormat_value)
}

func init() { proto.RegisterFile("google/logging/v2/logging_config.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1086 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5d, 0x6f, 0xdb, 0xd4,
	0x1b, 0xff, 0x3b, 0xe9, 0x4b, 0xf2, 0x74, 0x6b, 0xda, 0xb3, 0x7f, 0x83, 0xe5, 0xb2, 0x2e, 0xb8,
	0x53, 0x49, 0x2b, 0x70, 0x46, 0x80, 0x0b, 0x86, 0xa6, 0x89, 0x75, 0x6d, 0x55, 0xa9, 0xeb, 0x2a,
	0xb7, 0x0b, 0x37, 0x95, 0x2c, 0x37, 0x3e, 0x09, 0x47, 0x4d, 0xec, 0xcc, 0x3e, 0x0e, 0xed, 0x60,
	0x12, 0x70, 0x89, 0xc4, 0x05, 0xe2, 0x4b, 0x20, 0x3e, 0x03, 0x9f, 0x80, 0xdb, 0x7d, 0x05, 0x3e,
	0x04, 0x97, 0xe8, 0xbc, 0x38, 0x71, 0x6c, 0xc7, 0xcd, 0x84, 0xe0, 0xaa, 0x3e, 0xcf, 0xcb, 0xf9,
	0xfd, 0x9e, 0xe7, 0x3c, 0xcf, 0xaf, 0x81, 0xad, 0xae, 0xe7, 0x75, 0x7b, 0xb8, 0xd1, 0xf3, 0xba,
	0x5d, 0xe2, 0x76, 0x1b, 0xc3, 0x66, 0xf4, 0x69, 0xb5, 0x3d, 0xb7, 0x43, 0xba, 0xc6, 0xc0, 0xf7,
	0xa8, 0x87, 0x56, 0x45, 0x9c, 0x21, 0x9d, 0xc6, 0xb0, 0xa9, 0xbd, 0x2b, 0x53, 0xed, 0x01, 0x69,
	0xd8, 0xae, 0xeb, 0x51, 0x9b, 0x12, 0xcf, 0x0d, 0x44, 0x82, 0xb6, 0x2e, 0xbd, 0xfc, 0x74, 0x11,
	0x76, 0x1a, 0xb8, 0x3f, 0xa0, 0xd7, 0xd2, 0x59, 0x4b, 0x3a, 0x3b, 0x04, 0xf7, 0x1c, 0xab, 0x6f,
	0x07, 0x97, 0x32, 0xe2, 0x5e, 0x32, 0x82, 0x92, 0x3e, 0x0e, 0xa8, 0xdd, 0x1f, 0x88, 0x00, 0xfd,
	0xf7, 0x22, 0x2c, 0x1e, 0x79, 0xdd, 0x53, 0xe2, 0x5e, 0x22, 0x04, 0x73, 0xae, 0xdd, 0xc7, 0xaa,
	0x52, 0x53, 0xea, 0x65, 0x93, 0x7f, 0xa3, 0x1a, 0x2c, 0x39, 0x38, 0xa0, 0xc4, 0xe5, 0xac, 0xd4,
	0x22, 0x77, 0xc5, 0x4d, 0xa8, 0x0a, 0x0b, 0x1d, 0xd2, 0xa3, 0xd8, 0x57, 0xe7, 0xb9, 0x53, 0x9e,
	0xd0, 0x39, 0xac, 0x79, 0x21, 0x1d, 0x84, 0xd4, 0x1a, 0x62, 0x3f, 0x20, 0x9e, 0x6b, 0x75, 0x3c,
	0xbf, 0x6f, 0x53, 0x75, 0xa1, 0xa6, 0xd4, 0x97, 0x9b, 0x75, 0x23, 0xd5, 0x0a, 0x43, 0x12, 0x31,
	0x5a, 0x22, 0x61, 0x9f, 0xc7, 0x9b, 0x77, 0xc4, 0x35, 0x13, 0x46, 0xf4, 0x3e, 0x54, 0xbe, 0xf6,
	0x09, 0xc5, 0xbe, 0x45, 0x1c, 0xec, 0x52, 0x42, 0xaf, 0xd5, 0x12, 0x87, 0x5f, 0x16, 0xe6, 0x43,
	0x69, 0x45, 0xdb, 0xb0, 0x42, 0xdc, 0x76, 0x2f, 0x74, 0xb0, 0xd5, 0xfe, 0x8a, 0xf4, 0x1c, 0x1f,
	0xbb, 0x6a, 0xb9, 0xa6, 0xd4, 0x4b, 0x66, 0x45, 0xda, 0x77, 0xa5, 0x19, 0x7d, 0x06, 0x10, 0x50,
	0xdb, 0xa7, 0x16, 0x6b, 0x92, 0x0a, 0x35, 0xa5, 0xbe, 0xd4, 0xd4, 0x22, 0x9a, 0x51, 0x07, 0x8d,
	0xb3, 0xa8, 0x83, 0x66, 0x99, 0x47, 0xb3, 0x33, 0xfa, 0x14, 0x4a, 0xd8, 0x75, 0x44, 0xe2, 0xd2,
	0x8d, 0x89, 0x8b, 0xd8, 0x75, 0xd8, 0x49, 0x7f, 0x0c, 0xb7, 0x27, 0xcb, 0xda, 0x00, 0xad, 0xb5,
	0x67, 0x9e, 0x1e, 0x3e, 0x3f, 0xb6, 0xf6, 0x9f, 0x9b, 0xcf, 0xbe, 0x38, 0xb3, 0x5e, 0x1c, 0x9f,
	0x9e, 0xec, 0xed, 0x1e, 0xee, 0x1f, 0xee, 0x3d, 0x5d, 0xf9, 0x1f, 0x5a, 0x80, 0x42, 0xab, 0xb9,
	0xa2, 0xf0, 0xbf, 0x1f, 0xad, 0x14, 0xf4, 0x0e, 0xac, 0x1c, 0x91, 0x80, 0xb2, 0xae, 0x05, 0x26,
	0x7e, 0x19, 0xe2, 0x80, 0xb2, 0x07, 0x19, 0xd8, 0x3e, 0x76, 0xa9, 0x7c, 0x48, 0x79, 0x42, 0x77,
	0x01, 0x06, 0x76, 0x17, 0x5b, 0xd4, 0xbb, 0xc4, 0xae, 0x5a, 0xe0, 0xbe, 0x32, 0xb3, 0x9c, 0x31,
	0x03, 0x5a, 0x07, 0x7e, 0xb0, 0x02, 0xf2, 0x0a, 0xf3, 0x77, 0x9e, 0x37, 0x4b, 0xcc, 0x70, 0x4a,
	0x5e, 0x61, 0xbd, 0x0f, 0xab, 0x31, 0x9c, 0x60, 0xe0, 0xb9, 0x01, 0x46, 0x0f, 0x60, 0x3e, 0x60,
	0x06, 0x55, 0xa9, 0x15, 0xe3, 0x15, 0xa7, 0x5f, 0xd4, 0x14, 0x81, 0x68, 0x0b, 0x2a, 0x2e, 0xbe,
	0xa2, 0x56, 0x8a, 0xc7, 0x6d, 0x66, 0x3e, 0x89, 0xb8, 0xe8, 0x1f, 0xc2, 0xf2, 0x01, 0xe6, 0x68,
	0x51, 0x51, 0xeb, 0x50, 0x66, 0x57, 0x58, 0xb1, 0x01, 0x2d, 0x31, 0xc3, 0xb1, 0xdd, 0xc7, 0xfa,
	0xcf, 0x0a, 0xac, 0xee, 0xfa, 0xd8, 0xa6, 0x38, 0x9e, 0x32, 0xad, 0x0f, 0x06, 0xcc, 0xb1, 0x4c,
	0x8e, 0x9c, 0xcf, 0x9a, 0xc7, 0xa1, 0x4f, 0xa0, 0x1a, 0xba, 0xe4, 0x65, 0x88, 0xad, 0xe4, 0xc4,
	0x15, 0xf9, 0x1c, 0xfd, 0x5f, 0x78, 0xbf, 0x9c, 0x98, 0x3b, 0xfd, 0x8d, 0x02, 0xab, 0x2f, 0x06,
	0x4e, 0x82, 0x53, 0x5e, 0x19, 0xff, 0x0d, 0x31, 0xf4, 0x39, 0x2c, 0x85, 0x9c, 0x17, 0xd7, 0x09,
	0x75, 0x6e, 0xca, 0xb4, 0xee, 0x33, 0x29, 0x79, 0x66, 0x07, 0x97, 0x26, 0x88, 0x70, 0xf6, 0xad,
	0x3f, 0x80, 0xd5, 0xa7, 0xb8, 0x87, 0x67, 0x2f, 0x4a, 0xbf, 0x82, 0x5b, 0x47, 0x5e, 0x77, 0xef,
	0xaa, 0xdd, 0x0b, 0xd9, 0x9c, 0xe7, 0x88, 0x4c, 0xdb, 0x27, 0x03, 0x2e, 0x32, 0x85, 0x91, 0xc8,
	0x44, 0xa6, 0x98, 0xc8, 0x14, 0x27, 0x44, 0x46, 0x83, 0x92, 0x43, 0x02, 0xfb, 0xa2, 0x87, 0x1d,
	0x5e, 0x49, 0xc9, 0x1c, 0x9d, 0xf5, 0x4b, 0x58, 0x63, 0x33, 0x3b, 0x82, 0xfe, 0x57, 0x17, 0xe4,
	0x7b, 0x05, 0xaa, 0x49, 0x34, 0xb9, 0x26, 0x8f, 0x01, 0xf0, 0xc8, 0x2a, 0x77, 0xe5, 0x5e, 0xf6,
	0xe3, 0x8e, 0xb2, 0xcd, 0x58, 0xca, 0xcc, 0x5b, 0xb3, 0x0d, 0x77, 0x0e, 0xf0, 0x98, 0x41, 0x54,
	0x6e, 0x46, 0xc7, 0x75, 0x0f, 0xaa, 0x62, 0x61, 0x52, 0xd1, 0xd3, 0x9a, 0xf3, 0x08, 0xca, 0x23,
	0x4a, 0x72, 0x42, 0x6f, 0x2c, 0x62, 0x9c, 0xa1, 0xff, 0xaa, 0x40, 0x55, 0xac, 0xc3, 0x2c, 0xfc,
	0xfe, 0x21, 0x5a, 0x72, 0xc6, 0x8b, 0x6f, 0x35, 0xe3, 0x1f, 0x40, 0x55, 0xcc, 0xf8, 0x2c, 0x4c,
	0x9b, 0xdf, 0x01, 0x54, 0x76, 0xf9, 0xbf, 0xf8, 0x53, 0xec, 0x0f, 0x49, 0x1b, 0xb7, 0x9a, 0x68,
	0x08, 0xe5, 0x91, 0x5a, 0xa2, 0xcd, 0x2c, 0xde, 0x09, 0xcd, 0xd6, 0xee, 0xe7, 0x07, 0x89, 0x49,
	0xd2, 0x37, 0x7e, 0x78, 0xf3, 0xe7, 0x2f, 0x05, 0x15, 0x55, 0xd9, 0xef, 0x8b, 0x6f, 0xc4, 0xc3,
	0x3c, 0xda, 0x69, 0xec, 0xbc, 0x6e, 0x08, 0x79, 0xed, 0xc3, 0xa2, 0x94, 0x4d, 0xf4, 0x5e, 0xc6,
	0x85, 0x93, 0x92, 0xaa, 0xe5, 0x08, 0x8c, 0xbe, 0xc9, 0x91, 0xee, 0xa2, 0x75, 0x8e, 0x34, 0xda,
	0x6e, 0x06, 0x26, 0xb0, 0x1a, 0x3b, 0xaf, 0x11, 0x05, 0x18, 0xab, 0x2e, 0xca, 0x2a, 0x21, 0x25,
	0xca, 0xb9, 0xa0, 0xf7, 0x39, 0xe8, 0x86, 0x3e, 0xa5, 0xbc, 0x87, 0x42, 0xf5, 0xae, 0x01, 0xc6,
	0xba, 0x9a, 0x89, 0x9a, 0x92, 0xdd, 0x5c, 0xd4, 0x6d, 0x8e, 0xba, 0xa9, 0xe5, 0x95, 0x2a, 0xa1,
	0x3d, 0x80, 0xb1, 0xfa, 0x65, 0x42, 0xa7, 0xc4, 0x51, 0xab, 0xa6, 0xa6, 0x6e, 0x8f, 0xfd, 0x82,
	0x8b, 0x3a, 0xbc, 0x93, 0xdb, 0xe1, 0x1f, 0x15, 0x58, 0x9e, 0x54, 0x15, 0x54, 0x9f, 0x32, 0x29,
	0x29, 0x99, 0xd3, 0xb6, 0x67, 0x88, 0x94, 0x83, 0x35, 0xf9, 0xdc, 0xf1, 0xce, 0xc7, 0x64, 0xe8,
	0x5b, 0xb8, 0x15, 0x97, 0x17, 0xb4, 0x95, 0x3d, 0x62, 0xc9, 0xad, 0xd1, 0x6e, 0x5a, 0xdc, 0x04,
	0xfa, 0xa8, 0x0b, 0x63, 0x68, 0xd6, 0x8a, 0x9f, 0x14, 0xa8, 0x24, 0x24, 0x0b, 0x6d, 0x4f, 0x1d,
	0xb9, 0xb7, 0x27, 0x61, 0x70, 0x12, 0x75, 0x3d, 0xaf, 0x05, 0x0f, 0x63, 0x12, 0xc3, 0xf8, 0x24,
	0x04, 0x2d, 0x93, 0x4f, 0xb6, 0xe8, 0xcd, 0xcc, 0xa7, 0x99, 0xd7, 0x94, 0x38, 0x9f, 0x6b, 0xa8,
	0x24, 0x54, 0x2b, 0x93, 0x4e, 0xb6, 0xb2, 0xcd, 0x38, 0xa5, 0xd9, 0x2c, 0x9e, 0x38, 0xb0, 0xd6,
	0xf6, 0xfa, 0x69, 0xb0, 0x27, 0xe8, 0x48, 0x7c, 0x0b, 0x7d, 0x3c, 0x61, 0x57, 0x9f, 0x28, 0x7f,
	0x29, 0xca, 0x6f, 0x85, 0x77, 0x0e, 0x44, 0xf0, 0x6e, 0xcf, 0x0b, 0x1d, 0x43, 0x86, 0x19, 0xad,
	0xe6, 0x1f, 0x91, 0xe7, 0x9c, 0x7b, 0xce, 0xa5, 0xe7, 0xbc, 0xd5, 0xbc, 0x58, 0xe0, 0xd4, 0x3e,
	0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xf4, 0xe6, 0x55, 0x6b, 0x0d, 0x00, 0x00,
}
