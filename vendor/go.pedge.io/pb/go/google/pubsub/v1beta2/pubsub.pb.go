// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/pubsub/v1beta2/pubsub.proto

package google_pubsub_v1beta2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A topic resource.
type Topic struct {
	// Name of the topic.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A message data and its attributes.
type PubsubMessage struct {
	// The message payload. For JSON requests, the value of this field must be
	// base64-encoded.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Optional attributes for this message.
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ID of this message assigned by the server at publication time. Guaranteed
	// to be unique within the topic. This value may be read by a subscriber
	// that receives a PubsubMessage via a Pull call or a push delivery. It must
	// not be populated by a publisher in a Publish call.
	MessageId string `protobuf:"bytes,3,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
}

func (m *PubsubMessage) Reset()                    { *m = PubsubMessage{} }
func (m *PubsubMessage) String() string            { return proto.CompactTextString(m) }
func (*PubsubMessage) ProtoMessage()               {}
func (*PubsubMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PubsubMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PubsubMessage) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *PubsubMessage) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

// Request for the GetTopic method.
type GetTopicRequest struct {
	// The name of the topic to get.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *GetTopicRequest) Reset()                    { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()               {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

// Request for the Publish method.
type PublishRequest struct {
	// The messages in the request will be published on this topic.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	// The messages to publish.
	Messages []*PubsubMessage `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetMessages() []*PubsubMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Response for the Publish method.
type PublishResponse struct {
	// The server-assigned ID of each published message, in the same order as
	// the messages in the request. IDs are guaranteed to be unique within
	// the topic.
	MessageIds []string `protobuf:"bytes,1,rep,name=message_ids,json=messageIds" json:"message_ids,omitempty"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublishResponse) GetMessageIds() []string {
	if m != nil {
		return m.MessageIds
	}
	return nil
}

// Request for the ListTopics method.
type ListTopicsRequest struct {
	// The name of the cloud project that topics belong to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// Maximum number of topics to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The value returned by the last ListTopicsResponse; indicates that this is
	// a continuation of a prior ListTopics call, and that the system should
	// return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListTopicsRequest) Reset()                    { *m = ListTopicsRequest{} }
func (m *ListTopicsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTopicsRequest) ProtoMessage()               {}
func (*ListTopicsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListTopicsRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ListTopicsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTopicsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response for the ListTopics method.
type ListTopicsResponse struct {
	// The resulting topics.
	Topics []*Topic `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
	// If not empty, indicates that there may be more topics that match the
	// request; this value should be passed in a new ListTopicsRequest.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListTopicsResponse) Reset()                    { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()               {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListTopicsResponse) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *ListTopicsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request for the ListTopicSubscriptions method.
type ListTopicSubscriptionsRequest struct {
	// The name of the topic that subscriptions are attached to.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	// Maximum number of subscription names to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The value returned by the last ListTopicSubscriptionsResponse; indicates
	// that this is a continuation of a prior ListTopicSubscriptions call, and
	// that the system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListTopicSubscriptionsRequest) Reset()                    { *m = ListTopicSubscriptionsRequest{} }
func (m *ListTopicSubscriptionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTopicSubscriptionsRequest) ProtoMessage()               {}
func (*ListTopicSubscriptionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListTopicSubscriptionsRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ListTopicSubscriptionsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTopicSubscriptionsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response for the ListTopicSubscriptions method.
type ListTopicSubscriptionsResponse struct {
	// The names of the subscriptions that match the request.
	Subscriptions []string `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// If not empty, indicates that there may be more subscriptions that match
	// the request; this value should be passed in a new
	// ListTopicSubscriptionsRequest to get more subscriptions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListTopicSubscriptionsResponse) Reset()                    { *m = ListTopicSubscriptionsResponse{} }
func (m *ListTopicSubscriptionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTopicSubscriptionsResponse) ProtoMessage()               {}
func (*ListTopicSubscriptionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListTopicSubscriptionsResponse) GetSubscriptions() []string {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *ListTopicSubscriptionsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request for the DeleteTopic method.
type DeleteTopicRequest struct {
	// Name of the topic to delete.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *DeleteTopicRequest) Reset()                    { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()               {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

// A subscription resource.
type Subscription struct {
	// Name of the subscription.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name of the topic from which this subscription is receiving messages.
	// This will be present if and only if the subscription has not been detached
	// from its topic.
	Topic string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	// If push delivery is used with this subscription, this field is
	// used to configure it. An empty pushConfig signifies that the subscriber
	// will pull and ack messages using API methods.
	PushConfig *PushConfig `protobuf:"bytes,4,opt,name=push_config,json=pushConfig" json:"push_config,omitempty"`
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	//
	// For pull delivery this value
	// is used as the initial value for the ack deadline. It may be overridden
	// for a specific message by calling ModifyAckDeadline.
	//
	// For push delivery, this value is also used to set the request timeout for
	// the call to the push endpoint.
	//
	// If the subscriber never acknowledges the message, the Pub/Sub
	// system will eventually redeliver the message.
	AckDeadlineSeconds int32 `protobuf:"varint,5,opt,name=ack_deadline_seconds,json=ackDeadlineSeconds" json:"ack_deadline_seconds,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Subscription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subscription) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Subscription) GetPushConfig() *PushConfig {
	if m != nil {
		return m.PushConfig
	}
	return nil
}

func (m *Subscription) GetAckDeadlineSeconds() int32 {
	if m != nil {
		return m.AckDeadlineSeconds
	}
	return 0
}

// Configuration for a push delivery endpoint.
type PushConfig struct {
	// A URL locating the endpoint to which messages should be pushed.
	// For example, a Webhook endpoint might use "https://example.com/push".
	PushEndpoint string `protobuf:"bytes,1,opt,name=push_endpoint,json=pushEndpoint" json:"push_endpoint,omitempty"`
	// Endpoint configuration attributes.
	//
	// Every endpoint has a set of API supported attributes that can be used to
	// control different aspects of the message delivery.
	//
	// The currently supported attribute is `x-goog-version`, which you can
	// use to change the format of the push message. This attribute
	// indicates the version of the data expected by the endpoint. This
	// controls the shape of the envelope (i.e. its fields and metadata).
	// The endpoint version is based on the version of the Pub/Sub
	// API.
	//
	// If not present during the CreateSubscription call, it will default to
	// the version of the API used to make such call. If not present during a
	// ModifyPushConfig call, its value will not be changed. GetSubscription
	// calls will always return a valid version, even if the subscription was
	// created without this attribute.
	//
	// The possible values for this attribute are:
	//
	// * `v1beta1`: uses the push format defined in the v1beta1 Pub/Sub API.
	// * `v1beta2`: uses the push format defined in the v1beta2 Pub/Sub API.
	//
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PushConfig) Reset()                    { *m = PushConfig{} }
func (m *PushConfig) String() string            { return proto.CompactTextString(m) }
func (*PushConfig) ProtoMessage()               {}
func (*PushConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PushConfig) GetPushEndpoint() string {
	if m != nil {
		return m.PushEndpoint
	}
	return ""
}

func (m *PushConfig) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A message and its corresponding acknowledgment ID.
type ReceivedMessage struct {
	// This ID can be used to acknowledge the received message.
	AckId string `protobuf:"bytes,1,opt,name=ack_id,json=ackId" json:"ack_id,omitempty"`
	// The message.
	Message *PubsubMessage `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ReceivedMessage) Reset()                    { *m = ReceivedMessage{} }
func (m *ReceivedMessage) String() string            { return proto.CompactTextString(m) }
func (*ReceivedMessage) ProtoMessage()               {}
func (*ReceivedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReceivedMessage) GetAckId() string {
	if m != nil {
		return m.AckId
	}
	return ""
}

func (m *ReceivedMessage) GetMessage() *PubsubMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

// Request for the GetSubscription method.
type GetSubscriptionRequest struct {
	// The name of the subscription to get.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *GetSubscriptionRequest) Reset()                    { *m = GetSubscriptionRequest{} }
func (m *GetSubscriptionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSubscriptionRequest) ProtoMessage()               {}
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetSubscriptionRequest) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

// Request for the ListSubscriptions method.
type ListSubscriptionsRequest struct {
	// The name of the cloud project that subscriptions belong to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// Maximum number of subscriptions to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The value returned by the last ListSubscriptionsResponse; indicates that
	// this is a continuation of a prior ListSubscriptions call, and that the
	// system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListSubscriptionsRequest) Reset()                    { *m = ListSubscriptionsRequest{} }
func (m *ListSubscriptionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSubscriptionsRequest) ProtoMessage()               {}
func (*ListSubscriptionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ListSubscriptionsRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ListSubscriptionsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSubscriptionsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response for the ListSubscriptions method.
type ListSubscriptionsResponse struct {
	// The subscriptions that match the request.
	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// If not empty, indicates that there may be more subscriptions that match
	// the request; this value should be passed in a new ListSubscriptionsRequest
	// to get more subscriptions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListSubscriptionsResponse) Reset()                    { *m = ListSubscriptionsResponse{} }
func (m *ListSubscriptionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSubscriptionsResponse) ProtoMessage()               {}
func (*ListSubscriptionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ListSubscriptionsResponse) GetSubscriptions() []*Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *ListSubscriptionsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request for the DeleteSubscription method.
type DeleteSubscriptionRequest struct {
	// The subscription to delete.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *DeleteSubscriptionRequest) Reset()                    { *m = DeleteSubscriptionRequest{} }
func (m *DeleteSubscriptionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubscriptionRequest) ProtoMessage()               {}
func (*DeleteSubscriptionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *DeleteSubscriptionRequest) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

// Request for the ModifyPushConfig method.
type ModifyPushConfigRequest struct {
	// The name of the subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// The push configuration for future deliveries.
	//
	// An empty pushConfig indicates that the Pub/Sub system should
	// stop pushing messages from the given subscription and allow
	// messages to be pulled and acknowledged - effectively pausing
	// the subscription if Pull is not called.
	PushConfig *PushConfig `protobuf:"bytes,2,opt,name=push_config,json=pushConfig" json:"push_config,omitempty"`
}

func (m *ModifyPushConfigRequest) Reset()                    { *m = ModifyPushConfigRequest{} }
func (m *ModifyPushConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyPushConfigRequest) ProtoMessage()               {}
func (*ModifyPushConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ModifyPushConfigRequest) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *ModifyPushConfigRequest) GetPushConfig() *PushConfig {
	if m != nil {
		return m.PushConfig
	}
	return nil
}

// Request for the Pull method.
type PullRequest struct {
	// The subscription from which messages should be pulled.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// If this is specified as true the system will respond immediately even if
	// it is not able to return a message in the Pull response. Otherwise the
	// system is allowed to wait until at least one message is available rather
	// than returning no messages. The client may cancel the request if it does
	// not wish to wait any longer for the response.
	ReturnImmediately bool `protobuf:"varint,2,opt,name=return_immediately,json=returnImmediately" json:"return_immediately,omitempty"`
	// The maximum number of messages returned for this request. The Pub/Sub
	// system may return fewer than the number specified.
	MaxMessages int32 `protobuf:"varint,3,opt,name=max_messages,json=maxMessages" json:"max_messages,omitempty"`
}

func (m *PullRequest) Reset()                    { *m = PullRequest{} }
func (m *PullRequest) String() string            { return proto.CompactTextString(m) }
func (*PullRequest) ProtoMessage()               {}
func (*PullRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *PullRequest) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *PullRequest) GetReturnImmediately() bool {
	if m != nil {
		return m.ReturnImmediately
	}
	return false
}

func (m *PullRequest) GetMaxMessages() int32 {
	if m != nil {
		return m.MaxMessages
	}
	return 0
}

// Response for the Pull method.
type PullResponse struct {
	// Received Pub/Sub messages. The Pub/Sub system will return zero messages if
	// there are no more available in the backlog. The Pub/Sub system may return
	// fewer than the maxMessages requested even if there are more messages
	// available in the backlog.
	ReceivedMessages []*ReceivedMessage `protobuf:"bytes,1,rep,name=received_messages,json=receivedMessages" json:"received_messages,omitempty"`
}

func (m *PullResponse) Reset()                    { *m = PullResponse{} }
func (m *PullResponse) String() string            { return proto.CompactTextString(m) }
func (*PullResponse) ProtoMessage()               {}
func (*PullResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *PullResponse) GetReceivedMessages() []*ReceivedMessage {
	if m != nil {
		return m.ReceivedMessages
	}
	return nil
}

// Request for the ModifyAckDeadline method.
type ModifyAckDeadlineRequest struct {
	// The name of the subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// The acknowledgment ID.
	AckId string `protobuf:"bytes,2,opt,name=ack_id,json=ackId" json:"ack_id,omitempty"`
	// The new ack deadline with respect to the time this request was sent to the
	// Pub/Sub system. Must be >= 0. For example, if the value is 10, the new ack
	// deadline will expire 10 seconds after the ModifyAckDeadline call was made.
	// Specifying zero may immediately make the message available for another pull
	// request.
	AckDeadlineSeconds int32 `protobuf:"varint,3,opt,name=ack_deadline_seconds,json=ackDeadlineSeconds" json:"ack_deadline_seconds,omitempty"`
}

func (m *ModifyAckDeadlineRequest) Reset()                    { *m = ModifyAckDeadlineRequest{} }
func (m *ModifyAckDeadlineRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyAckDeadlineRequest) ProtoMessage()               {}
func (*ModifyAckDeadlineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *ModifyAckDeadlineRequest) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *ModifyAckDeadlineRequest) GetAckId() string {
	if m != nil {
		return m.AckId
	}
	return ""
}

func (m *ModifyAckDeadlineRequest) GetAckDeadlineSeconds() int32 {
	if m != nil {
		return m.AckDeadlineSeconds
	}
	return 0
}

// Request for the Acknowledge method.
type AcknowledgeRequest struct {
	// The subscription whose message is being acknowledged.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// The acknowledgment ID for the messages being acknowledged that was returned
	// by the Pub/Sub system in the Pull response. Must not be empty.
	AckIds []string `protobuf:"bytes,2,rep,name=ack_ids,json=ackIds" json:"ack_ids,omitempty"`
}

func (m *AcknowledgeRequest) Reset()                    { *m = AcknowledgeRequest{} }
func (m *AcknowledgeRequest) String() string            { return proto.CompactTextString(m) }
func (*AcknowledgeRequest) ProtoMessage()               {}
func (*AcknowledgeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *AcknowledgeRequest) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *AcknowledgeRequest) GetAckIds() []string {
	if m != nil {
		return m.AckIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Topic)(nil), "google.pubsub.v1beta2.Topic")
	proto.RegisterType((*PubsubMessage)(nil), "google.pubsub.v1beta2.PubsubMessage")
	proto.RegisterType((*GetTopicRequest)(nil), "google.pubsub.v1beta2.GetTopicRequest")
	proto.RegisterType((*PublishRequest)(nil), "google.pubsub.v1beta2.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "google.pubsub.v1beta2.PublishResponse")
	proto.RegisterType((*ListTopicsRequest)(nil), "google.pubsub.v1beta2.ListTopicsRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "google.pubsub.v1beta2.ListTopicsResponse")
	proto.RegisterType((*ListTopicSubscriptionsRequest)(nil), "google.pubsub.v1beta2.ListTopicSubscriptionsRequest")
	proto.RegisterType((*ListTopicSubscriptionsResponse)(nil), "google.pubsub.v1beta2.ListTopicSubscriptionsResponse")
	proto.RegisterType((*DeleteTopicRequest)(nil), "google.pubsub.v1beta2.DeleteTopicRequest")
	proto.RegisterType((*Subscription)(nil), "google.pubsub.v1beta2.Subscription")
	proto.RegisterType((*PushConfig)(nil), "google.pubsub.v1beta2.PushConfig")
	proto.RegisterType((*ReceivedMessage)(nil), "google.pubsub.v1beta2.ReceivedMessage")
	proto.RegisterType((*GetSubscriptionRequest)(nil), "google.pubsub.v1beta2.GetSubscriptionRequest")
	proto.RegisterType((*ListSubscriptionsRequest)(nil), "google.pubsub.v1beta2.ListSubscriptionsRequest")
	proto.RegisterType((*ListSubscriptionsResponse)(nil), "google.pubsub.v1beta2.ListSubscriptionsResponse")
	proto.RegisterType((*DeleteSubscriptionRequest)(nil), "google.pubsub.v1beta2.DeleteSubscriptionRequest")
	proto.RegisterType((*ModifyPushConfigRequest)(nil), "google.pubsub.v1beta2.ModifyPushConfigRequest")
	proto.RegisterType((*PullRequest)(nil), "google.pubsub.v1beta2.PullRequest")
	proto.RegisterType((*PullResponse)(nil), "google.pubsub.v1beta2.PullResponse")
	proto.RegisterType((*ModifyAckDeadlineRequest)(nil), "google.pubsub.v1beta2.ModifyAckDeadlineRequest")
	proto.RegisterType((*AcknowledgeRequest)(nil), "google.pubsub.v1beta2.AcknowledgeRequest")
}

func init() { proto.RegisterFile("google/pubsub/v1beta2/pubsub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1081 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x25, 0x5b, 0xb6, 0x86, 0x72, 0x15, 0x2f, 0x12, 0x87, 0x96, 0x9b, 0xd6, 0x61, 0x52,
	0x57, 0x09, 0x10, 0x39, 0x51, 0x5d, 0xa0, 0x28, 0xfa, 0x67, 0x27, 0x46, 0x20, 0xa0, 0x46, 0x15,
	0xda, 0x87, 0xa2, 0x28, 0x22, 0x50, 0xe4, 0x5a, 0x62, 0x45, 0x91, 0x0c, 0x77, 0xe9, 0x5a, 0xb9,
	0x15, 0x45, 0xd1, 0x9e, 0xfa, 0x28, 0x7d, 0x85, 0x1e, 0xfa, 0x10, 0x7d, 0x9d, 0x82, 0xbb, 0x4b,
	0x8a, 0xb4, 0xb9, 0x32, 0xed, 0xa0, 0x37, 0x72, 0x76, 0x7e, 0xbe, 0x99, 0x9d, 0x6f, 0x38, 0x04,
	0x7d, 0xe4, 0xfb, 0x23, 0x17, 0xef, 0x06, 0xd1, 0x90, 0x44, 0xc3, 0xdd, 0xb3, 0x67, 0x43, 0x4c,
	0xcd, 0xae, 0x78, 0xed, 0x04, 0xa1, 0x4f, 0x7d, 0x74, 0x87, 0xeb, 0x74, 0x84, 0x50, 0xe8, 0xb4,
	0xb6, 0x12, 0xd3, 0x58, 0x69, 0x18, 0x9d, 0xee, 0xe2, 0x69, 0x40, 0x67, 0xdc, 0x46, 0xdf, 0x82,
	0xe5, 0x13, 0x3f, 0x70, 0x2c, 0x84, 0x60, 0xc9, 0x33, 0xa7, 0x58, 0x53, 0xb6, 0x95, 0x76, 0xdd,
	0x60, 0xcf, 0xfa, 0xbf, 0x0a, 0xac, 0xf5, 0x99, 0xb3, 0x23, 0x4c, 0x88, 0x39, 0xc2, 0xb1, 0x96,
	0x6d, 0x52, 0x93, 0x69, 0x35, 0x0c, 0xf6, 0x8c, 0x4e, 0x00, 0x4c, 0x4a, 0x43, 0x67, 0x18, 0x51,
	0x4c, 0xb4, 0xca, 0x76, 0xb5, 0xad, 0x76, 0xf7, 0x3a, 0x85, 0x58, 0x3a, 0x39, 0x6f, 0x9d, 0xfd,
	0xd4, 0xec, 0xd0, 0xa3, 0xe1, 0xcc, 0xc8, 0xf8, 0x41, 0xf7, 0x00, 0xa6, 0x5c, 0x6d, 0xe0, 0xd8,
	0x5a, 0x95, 0xa1, 0xaa, 0x0b, 0x49, 0xcf, 0x6e, 0x7d, 0x09, 0xcd, 0x0b, 0xd6, 0xe8, 0x16, 0x54,
	0x27, 0x78, 0x26, 0x12, 0x88, 0x1f, 0xd1, 0x6d, 0x58, 0x3e, 0x33, 0xdd, 0x08, 0x6b, 0x15, 0x26,
	0xe3, 0x2f, 0x9f, 0x57, 0x3e, 0x53, 0xf4, 0x8f, 0xa1, 0xf9, 0x12, 0x53, 0x96, 0xb9, 0x81, 0xdf,
	0x44, 0x98, 0xd0, 0x58, 0x99, 0xc6, 0xef, 0xc2, 0x01, 0x7f, 0xd1, 0xc7, 0xf0, 0x5e, 0x3f, 0x1a,
	0xba, 0x0e, 0x19, 0x2f, 0xd4, 0x43, 0xdf, 0xc0, 0xaa, 0x00, 0x97, 0x94, 0xe0, 0x61, 0x99, 0x12,
	0x18, 0xa9, 0x95, 0xde, 0x85, 0x66, 0x1a, 0x89, 0x04, 0xbe, 0x47, 0x30, 0xfa, 0x10, 0xd4, 0x79,
	0x0d, 0x88, 0xa6, 0x6c, 0x57, 0xdb, 0x75, 0x03, 0xd2, 0x22, 0x10, 0xdd, 0x81, 0xf5, 0x6f, 0x1d,
	0xc2, 0xf3, 0x20, 0x09, 0x40, 0x0d, 0x56, 0x82, 0xd0, 0xff, 0x09, 0x5b, 0x54, 0x40, 0x4c, 0x5e,
	0xd1, 0x16, 0xd4, 0x83, 0xd8, 0x19, 0x71, 0xde, 0xf2, 0x9a, 0x2c, 0x1b, 0xab, 0xb1, 0xe0, 0xd8,
	0x79, 0x8b, 0xe3, 0x82, 0xb3, 0x43, 0xea, 0x4f, 0xb0, 0x97, 0x14, 0x3c, 0x96, 0x9c, 0xc4, 0x02,
	0x3d, 0x04, 0x94, 0x0d, 0x25, 0x10, 0xee, 0x41, 0x8d, 0xe5, 0xcf, 0xc1, 0xa9, 0xdd, 0xf7, 0x25,
	0x49, 0xf3, 0x4a, 0x0b, 0x5d, 0xb4, 0x03, 0x4d, 0x0f, 0x9f, 0xd3, 0x41, 0x26, 0x1e, 0xbf, 0xa1,
	0xb5, 0x58, 0xdc, 0x4f, 0x63, 0xbe, 0x81, 0x7b, 0x69, 0xcc, 0xe3, 0x68, 0x48, 0xac, 0xd0, 0x09,
	0xa8, 0xe3, 0x7b, 0x64, 0xf1, 0x5d, 0xbc, 0x4b, 0x9a, 0x1e, 0x7c, 0x20, 0x0b, 0x29, 0x52, 0x7e,
	0x08, 0x6b, 0x24, 0x7b, 0x20, 0xae, 0x25, 0x2f, 0x2c, 0x9d, 0xe2, 0x63, 0x40, 0x2f, 0xb0, 0x8b,
	0x29, 0x2e, 0xd1, 0x8b, 0x7f, 0x29, 0xd0, 0xc8, 0x62, 0x2a, 0xe2, 0xec, 0xdc, 0xb4, 0x92, 0x2d,
	0xc9, 0x01, 0xa8, 0x41, 0x44, 0xc6, 0x03, 0xcb, 0xf7, 0x4e, 0x9d, 0x91, 0xb6, 0xb4, 0xad, 0xb4,
	0xd5, 0xee, 0x7d, 0x69, 0x87, 0x92, 0xf1, 0x73, 0xa6, 0x68, 0x40, 0x90, 0x3e, 0xa3, 0xa7, 0x70,
	0xdb, 0xb4, 0x26, 0x03, 0x1b, 0x9b, 0xb6, 0xeb, 0x78, 0x78, 0x40, 0xb0, 0xe5, 0x7b, 0x36, 0xd1,
	0x96, 0x59, 0x85, 0x91, 0x69, 0x4d, 0x5e, 0x88, 0xa3, 0x63, 0x7e, 0xa2, 0xff, 0xa3, 0x00, 0xcc,
	0x9d, 0xa1, 0x07, 0xb0, 0xc6, 0x40, 0x60, 0xcf, 0x0e, 0x7c, 0xc7, 0x4b, 0xda, 0xb3, 0x11, 0x0b,
	0x0f, 0x85, 0x0c, 0xbd, 0x2a, 0x98, 0x26, 0xcf, 0xae, 0x04, 0xba, 0x68, 0x94, 0xbc, 0xeb, 0xac,
	0x18, 0x43, 0xd3, 0xc0, 0x16, 0x76, 0xce, 0xb0, 0x9d, 0x8c, 0xc1, 0x3b, 0x50, 0x8b, 0x4b, 0xe1,
	0xd8, 0xc9, 0x05, 0x99, 0xd6, 0xa4, 0x67, 0xa3, 0xaf, 0x60, 0x45, 0x90, 0x93, 0x79, 0x29, 0x3b,
	0x03, 0x12, 0x23, 0xfd, 0x0b, 0xd8, 0x78, 0x89, 0x69, 0xf6, 0x8a, 0x93, 0x86, 0xd0, 0xa1, 0x91,
	0xed, 0xaf, 0xa4, 0x72, 0x59, 0x99, 0x1e, 0x80, 0x16, 0xb7, 0x6e, 0x21, 0x51, 0xfe, 0x9f, 0x99,
	0xf0, 0xa7, 0x02, 0x9b, 0x05, 0x21, 0x05, 0x51, 0x7a, 0x45, 0x44, 0x51, 0xbb, 0x0f, 0x24, 0x35,
	0xc9, 0xa5, 0x7d, 0x43, 0x36, 0x7d, 0x0d, 0x9b, 0x9c, 0x4d, 0x37, 0xad, 0xe1, 0x2f, 0x0a, 0xdc,
	0x3d, 0xf2, 0x6d, 0xe7, 0x74, 0x96, 0x21, 0x41, 0x79, 0xfb, 0x8b, 0x3c, 0xab, 0xdc, 0x80, 0x67,
	0xfa, 0xaf, 0x0a, 0xa8, 0xfd, 0xc8, 0x75, 0xaf, 0x13, 0xf7, 0x09, 0xa0, 0x10, 0xd3, 0x28, 0xf4,
	0x06, 0xce, 0x74, 0x8a, 0x6d, 0xc7, 0xa4, 0xd8, 0x9d, 0xb1, 0xf0, 0xab, 0xc6, 0x3a, 0x3f, 0xe9,
	0xcd, 0x0f, 0xd0, 0x7d, 0x68, 0x4c, 0xcd, 0xf3, 0x41, 0xfa, 0xc5, 0xaa, 0xb2, 0x7b, 0x57, 0xa7,
	0xe6, 0xf9, 0x51, 0xf2, 0x39, 0xb2, 0xa0, 0xc1, 0x41, 0x88, 0xdb, 0x3c, 0x86, 0xf5, 0x50, 0xb0,
	0x60, 0x6e, 0xc7, 0x6f, 0x74, 0x47, 0x92, 0xdf, 0x05, 0xd6, 0x18, 0xb7, 0xc2, 0xbc, 0x80, 0xe8,
	0xbf, 0x2b, 0xa0, 0xf1, 0x72, 0xef, 0xcf, 0xa7, 0xc7, 0x75, 0xf2, 0x9e, 0x13, 0xb1, 0x92, 0x25,
	0xa2, 0x6c, 0x54, 0x55, 0xa5, 0xa3, 0xea, 0x15, 0xa0, 0x7d, 0x6b, 0xe2, 0xf9, 0x3f, 0xbb, 0xd8,
	0x1e, 0x5d, 0x0b, 0xc2, 0x5d, 0x58, 0xe1, 0x10, 0xf8, 0xb4, 0xaa, 0x1b, 0x35, 0x86, 0x81, 0x74,
	0xff, 0xa8, 0x01, 0x88, 0x3e, 0x1c, 0xe2, 0x10, 0xbd, 0x06, 0xf4, 0x3c, 0xc4, 0x66, 0xbe, 0x37,
	0x51, 0x19, 0x36, 0xb4, 0xca, 0x28, 0x21, 0xcc, 0x56, 0x9a, 0x9c, 0xe8, 0x89, 0xc4, 0xae, 0x78,
	0xc8, 0x94, 0x0b, 0x73, 0xc6, 0x57, 0x8e, 0x1c, 0xe5, 0xd1, 0xae, 0xc4, 0x52, 0x36, 0x8f, 0x5a,
	0x4f, 0xcb, 0x1b, 0x88, 0xfe, 0x7b, 0x9d, 0x7c, 0x28, 0x73, 0x68, 0x64, 0x7e, 0xa4, 0x53, 0xa0,
	0xb5, 0x91, 0x5a, 0x88, 0x75, 0xb8, 0x73, 0x18, 0xaf, 0xc3, 0xe8, 0x47, 0x58, 0xbf, 0xd4, 0x89,
	0xd2, 0xbc, 0x64, 0x3d, 0x2b, 0xf5, 0xde, 0x07, 0x35, 0xd3, 0x5e, 0xe8, 0x91, 0xc4, 0xef, 0xe5,
	0x16, 0x94, 0x7a, 0xfc, 0x0e, 0x96, 0x62, 0x7e, 0x22, 0x5d, 0x3a, 0x5c, 0xd2, 0x09, 0x22, 0xbd,
	0xd8, 0x1c, 0xc1, 0x7f, 0x80, 0x5b, 0x17, 0x27, 0x1f, 0xea, 0x2c, 0xcc, 0xff, 0xd2, 0x88, 0x94,
	0x81, 0xed, 0xfe, 0xbd, 0x04, 0x75, 0xb1, 0xdc, 0xe2, 0x10, 0xf5, 0x40, 0xe5, 0x4c, 0xe0, 0x7f,
	0x1e, 0x0b, 0x77, 0xc6, 0xd6, 0xc2, 0x53, 0xf4, 0x3d, 0xac, 0x08, 0xbf, 0xe8, 0x23, 0xf9, 0xb7,
	0x36, 0xb3, 0xbe, 0xb7, 0x76, 0xae, 0x52, 0x13, 0xe5, 0xe8, 0xc3, 0x6a, 0xf2, 0x87, 0x80, 0x76,
	0xe4, 0x3c, 0xca, 0xae, 0x6d, 0x57, 0x60, 0x35, 0x01, 0xe6, 0x1b, 0x34, 0x6a, 0x2f, 0x60, 0x40,
	0x6e, 0x9f, 0x6f, 0x3d, 0x2a, 0xa1, 0x29, 0x40, 0xff, 0xa6, 0xc0, 0x46, 0xf1, 0xfa, 0x8a, 0xf6,
	0xae, 0xf2, 0x52, 0xc8, 0xd3, 0x4f, 0xaf, 0x69, 0x95, 0x16, 0x4f, 0xcd, 0x6c, 0xb5, 0xd2, 0x76,
	0xbf, 0xbc, 0xf9, 0xca, 0x3a, 0xe8, 0xe0, 0x31, 0x6c, 0x5a, 0xfe, 0xb4, 0xd8, 0xcf, 0x81, 0xca,
	0xf7, 0xa9, 0x7e, 0x6c, 0xd2, 0x57, 0x86, 0x35, 0x66, 0xfb, 0xc9, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0xde, 0xcb, 0x49, 0x34, 0x0f, 0x00, 0x00,
}
