// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/longrunning/operations.proto

package google_longrunning

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import google_protobuf1 "go.pedge.io/pb/gogo/google/protobuf"
import _ "go.pedge.io/pb/gogo/google/protobuf"
import google_rpc "go.pedge.io/pb/gogo/google/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// The server-assigned name, which is only unique within the same service that
	// originally returns it. If you use the default HTTP mapping, the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *google_protobuf1.Any `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If true, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// The operation result, which can be either an `error` or a valid `response`.
	// If `done` == `false`, neither `error` nor `response` is set.
	// If `done` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result isOperation_Result `protobuf_oneof:"result"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{0} }

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *google_rpc.Status `protobuf:"bytes,4,opt,name=error,oneof"`
}
type Operation_Response struct {
	Response *google_protobuf1.Any `protobuf:"bytes,5,opt,name=response,oneof"`
}

func (*Operation_Error) isOperation_Result()    {}
func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operation) GetMetadata() *google_protobuf1.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Operation) GetError() *google_rpc.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *google_protobuf1.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Operation_Response:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Result has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 4: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_rpc.Status)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Error{msg}
		return true, err
	case 5: // result.response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Response{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Response:
		s := proto.Size(x.Response)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request message for [Operations.GetOperation][google.longrunning.Operations.GetOperation].
type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetOperationRequest) Reset()                    { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()               {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{1} }

func (m *GetOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsRequest struct {
	// The name of the operation collection.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The standard list filter.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The standard list page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The standard list page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListOperationsRequest) Reset()                    { *m = ListOperationsRequest{} }
func (m *ListOperationsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsRequest) ProtoMessage()               {}
func (*ListOperationsRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{2} }

func (m *ListOperationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListOperationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListOperationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListOperationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsResponse struct {
	// A list of operations that matches the specified filter in the request.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations" json:"operations,omitempty"`
	// The standard List next-page token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListOperationsResponse) Reset()                    { *m = ListOperationsResponse{} }
func (m *ListOperationsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsResponse) ProtoMessage()               {}
func (*ListOperationsResponse) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{3} }

func (m *ListOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListOperationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Operations.CancelOperation][google.longrunning.Operations.CancelOperation].
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *CancelOperationRequest) Reset()                    { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()               {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{4} }

func (m *CancelOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.DeleteOperation][google.longrunning.Operations.DeleteOperation].
type DeleteOperationRequest struct {
	// The name of the operation resource to be deleted.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteOperationRequest) Reset()                    { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()               {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{5} }

func (m *DeleteOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Operation)(nil), "google.longrunning.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "google.longrunning.GetOperationRequest")
	proto.RegisterType((*ListOperationsRequest)(nil), "google.longrunning.ListOperationsRequest")
	proto.RegisterType((*ListOperationsResponse)(nil), "google.longrunning.ListOperationsResponse")
	proto.RegisterType((*CancelOperationRequest)(nil), "google.longrunning.CancelOperationRequest")
	proto.RegisterType((*DeleteOperationRequest)(nil), "google.longrunning.DeleteOperationRequest")
}

func init() { proto.RegisterFile("google/longrunning/operations.proto", fileDescriptorOperations) }

var fileDescriptorOperations = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xd3, 0x24, 0x4a, 0xa6, 0x40, 0xa4, 0x81, 0xba, 0xc6, 0x25, 0x22, 0x32, 0x08, 0x52,
	0x53, 0xd9, 0x10, 0x6e, 0x95, 0x7a, 0x20, 0x80, 0xe8, 0x01, 0x89, 0xc8, 0xe5, 0x5e, 0xb9, 0xe9,
	0x34, 0xb2, 0x70, 0x76, 0xcd, 0x7a, 0x03, 0x6d, 0x51, 0x55, 0xc1, 0x81, 0x13, 0x37, 0x3e, 0x8c,
	0x03, 0xbf, 0xc0, 0x87, 0x20, 0x6f, 0x9c, 0xd8, 0xa4, 0x1b, 0x94, 0xdb, 0xee, 0xce, 0xdb, 0x37,
	0x6f, 0xde, 0xdb, 0x85, 0x07, 0x23, 0xce, 0x47, 0x31, 0xf9, 0x31, 0x67, 0x23, 0x31, 0x61, 0x2c,
	0x62, 0x23, 0x9f, 0x27, 0x24, 0x42, 0x19, 0x71, 0x96, 0x7a, 0x89, 0xe0, 0x92, 0x23, 0x4e, 0x41,
	0x5e, 0x09, 0x64, 0xdf, 0xcb, 0x2f, 0x86, 0x49, 0xe4, 0x87, 0x8c, 0x71, 0x59, 0xbe, 0x61, 0xdf,
	0xcd, 0xab, 0x6a, 0x77, 0x3c, 0x39, 0xf5, 0x43, 0x76, 0x9e, 0x97, 0xb6, 0x17, 0x4b, 0x34, 0x4e,
	0xe4, 0xac, 0xb8, 0x95, 0x17, 0x45, 0x32, 0xf4, 0x53, 0x19, 0xca, 0x49, 0x4e, 0xe8, 0xfc, 0x32,
	0xa0, 0xf9, 0x6e, 0xa6, 0x0b, 0x11, 0xaa, 0x2c, 0x1c, 0x93, 0x65, 0x74, 0x8c, 0x6e, 0x33, 0x50,
	0x6b, 0x7c, 0x0a, 0x8d, 0x31, 0xc9, 0xf0, 0x24, 0x94, 0xa1, 0x55, 0xe9, 0x18, 0xdd, 0x8d, 0xde,
	0x1d, 0x2f, 0xd7, 0x3d, 0x6b, 0xe5, 0xbd, 0x60, 0xe7, 0xc1, 0x1c, 0x95, 0xb1, 0x9c, 0x70, 0x46,
	0xd6, 0x7a, 0xc7, 0xe8, 0x36, 0x02, 0xb5, 0x46, 0x17, 0x6a, 0x24, 0x04, 0x17, 0x56, 0x55, 0x51,
	0xe0, 0x8c, 0x42, 0x24, 0x43, 0xef, 0x50, 0x09, 0x3a, 0x58, 0x0b, 0xa6, 0x10, 0xec, 0x41, 0x43,
	0x50, 0x9a, 0x70, 0x96, 0x92, 0x55, 0x5b, 0xde, 0xf1, 0x60, 0x2d, 0x98, 0xe3, 0xfa, 0x0d, 0xa8,
	0x0b, 0x4a, 0x27, 0xb1, 0x74, 0x76, 0xe0, 0xf6, 0x1b, 0x92, 0xf3, 0x99, 0x02, 0xfa, 0x38, 0xa1,
	0x54, 0xea, 0x46, 0x73, 0xae, 0x60, 0xf3, 0x6d, 0x94, 0x16, 0xd8, 0x74, 0x11, 0x5c, 0x2d, 0xf9,
	0x60, 0x42, 0xfd, 0x34, 0x8a, 0x25, 0x89, 0x9c, 0x22, 0xdf, 0xe1, 0x36, 0x34, 0x93, 0x70, 0x44,
	0x47, 0x69, 0x74, 0x41, 0xca, 0xa0, 0x5a, 0xd0, 0xc8, 0x0e, 0x0e, 0xa3, 0x0b, 0xc2, 0x36, 0x80,
	0x2a, 0x4a, 0xfe, 0x81, 0x98, 0x32, 0xa4, 0x19, 0x28, 0xf8, 0xfb, 0xec, 0xc0, 0xb9, 0x02, 0x73,
	0x51, 0xc0, 0x74, 0x1e, 0xdc, 0x07, 0x28, 0x9e, 0x8b, 0x65, 0x74, 0xd6, 0xbb, 0x1b, 0xbd, 0xb6,
	0x77, 0xfd, 0xbd, 0x78, 0xc5, 0xa0, 0xa5, 0x0b, 0xf8, 0x08, 0x5a, 0x8c, 0xce, 0xe4, 0x51, 0xa9,
	0x79, 0x45, 0x35, 0xbf, 0x99, 0x1d, 0x0f, 0xe6, 0x02, 0x76, 0xc1, 0x7c, 0x19, 0xb2, 0x21, 0xc5,
	0x2b, 0xf9, 0xb5, 0x0b, 0xe6, 0x2b, 0x8a, 0x49, 0xd2, 0x2a, 0xe8, 0xde, 0x8f, 0x2a, 0x40, 0x31,
	0x19, 0x9e, 0xc1, 0x8d, 0x72, 0x2e, 0xf8, 0x58, 0x37, 0x8d, 0x26, 0x39, 0xfb, 0xff, 0x63, 0x3b,
	0x9d, 0x6f, 0xbf, 0xff, 0xfc, 0xac, 0xd8, 0x68, 0xf9, 0x9f, 0x9e, 0xf9, 0x5f, 0xb2, 0xce, 0xfb,
	0x85, 0x11, 0xbe, 0xeb, 0x5e, 0xe2, 0x77, 0x03, 0x6e, 0xfd, 0x6b, 0x33, 0xee, 0xe8, 0x38, 0xb5,
	0x6f, 0xc1, 0x76, 0x57, 0x81, 0x4e, 0x53, 0x73, 0xda, 0x4a, 0xcb, 0x16, 0x6e, 0xea, 0xb4, 0x5c,
	0xe2, 0x57, 0x03, 0x5a, 0x0b, 0x76, 0xa3, 0x96, 0x5e, 0x9f, 0x89, 0x6d, 0x5e, 0xfb, 0x06, 0xaf,
	0xb3, 0x3f, 0xee, 0xb8, 0xaa, 0xed, 0x43, 0xe7, 0xfe, 0x32, 0x0b, 0xf6, 0x86, 0x8a, 0x70, 0xcf,
	0x70, 0xf1, 0x33, 0xb4, 0x16, 0x32, 0xd4, 0x4b, 0xd0, 0x07, 0xbd, 0x54, 0x42, 0x9e, 0x82, 0xbb,
	0x34, 0x85, 0xfe, 0x13, 0x30, 0x87, 0x7c, 0xac, 0x69, 0xd5, 0x6f, 0x15, 0x4e, 0x0e, 0x32, 0xd6,
	0x81, 0x71, 0x5c, 0x57, 0xf4, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x2c, 0x30, 0xa8,
	0x47, 0x05, 0x00, 0x00,
}
