// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: huawei-grpc-dialin.proto

package dialin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`    //Subscribed sensor-path.
	Depth uint32 `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"` //Sampling depth of the subscribed sensor-path. The value 1 indicates the current level, and the value 2 indicates the current level and its sub-level, and so on.
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_dialin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_dialin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_dialin_proto_rawDescGZIP(), []int{0}
}

func (x *Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Path) GetDepth() uint32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

type SubsArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId         uint64  `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`                         //Request ID, which is transferred by the invoker.
	Encoding          uint32  `protobuf:"varint,2,opt,name=encoding,proto3" json:"encoding,omitempty"`                                            //Encoding type. Currently, the value can only be 0,indicating GPB encoding. 1 : json
	Path              []*Path `protobuf:"bytes,5,rep,name=path,proto3" json:"path,omitempty"`                                                     //Subscribed path structure.
	SampleInterval    uint64  `protobuf:"varint,6,opt,name=sample_interval,json=sampleInterval,proto3" json:"sample_interval,omitempty"`          //Sampling interval.
	HeartbeatInterval uint64  `protobuf:"varint,7,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"` //Redundancy suppression interval. This parameter is valid only when suppress_redundant is set to 1.
	// Types that are assignable to Suppress:
	//	*SubsArgs_SuppressRedundant
	//	*SubsArgs_DelayTime
	Suppress isSubsArgs_Suppress `protobuf_oneof:"Suppress"`
}

func (x *SubsArgs) Reset() {
	*x = SubsArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_dialin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubsArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsArgs) ProtoMessage() {}

func (x *SubsArgs) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_dialin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsArgs.ProtoReflect.Descriptor instead.
func (*SubsArgs) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_dialin_proto_rawDescGZIP(), []int{1}
}

func (x *SubsArgs) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *SubsArgs) GetEncoding() uint32 {
	if x != nil {
		return x.Encoding
	}
	return 0
}

func (x *SubsArgs) GetPath() []*Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *SubsArgs) GetSampleInterval() uint64 {
	if x != nil {
		return x.SampleInterval
	}
	return 0
}

func (x *SubsArgs) GetHeartbeatInterval() uint64 {
	if x != nil {
		return x.HeartbeatInterval
	}
	return 0
}

func (m *SubsArgs) GetSuppress() isSubsArgs_Suppress {
	if m != nil {
		return m.Suppress
	}
	return nil
}

func (x *SubsArgs) GetSuppressRedundant() bool {
	if x, ok := x.GetSuppress().(*SubsArgs_SuppressRedundant); ok {
		return x.SuppressRedundant
	}
	return false
}

func (x *SubsArgs) GetDelayTime() uint64 {
	if x, ok := x.GetSuppress().(*SubsArgs_DelayTime); ok {
		return x.DelayTime
	}
	return 0
}

type isSubsArgs_Suppress interface {
	isSubsArgs_Suppress()
}

type SubsArgs_SuppressRedundant struct {
	SuppressRedundant bool `protobuf:"varint,8,opt,name=suppress_redundant,json=suppressRedundant,proto3,oneof"` //Redundancy suppression. The sampled data is not reported if the data content remains unchanged. 0: disabled, 1: enabled.
}

type SubsArgs_DelayTime struct {
	DelayTime uint64 `protobuf:"varint,9,opt,name=delay_time,json=delayTime,proto3,oneof"` //Delay time, which ranges from 100 to 60000, in ms. After data changes, the new data is reported after a specified delay. The new changes during the delay are accumulated. This parameter is used to prevent a large amount of data from being reported during service flapping.
}

func (*SubsArgs_SuppressRedundant) isSubsArgs_Suppress() {}

func (*SubsArgs_DelayTime) isSubsArgs_Suppress() {}

type SubsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId uint32 `protobuf:"varint,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"` //If the subscribing is successful, the subscription ID is returned. If the subscribing fails, 0 is returned.
	RequestId      uint64 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`                //ID of the subscribing request.
	ResponseCode   string `protobuf:"bytes,3,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`        //Return code. The value 200 indicates a success.
	// Types that are assignable to MessageData:
	//	*SubsReply_Message
	//	*SubsReply_MessageJson
	MessageData isSubsReply_MessageData `protobuf_oneof:"MessageData"`
}

func (x *SubsReply) Reset() {
	*x = SubsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_dialin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsReply) ProtoMessage() {}

func (x *SubsReply) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_dialin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsReply.ProtoReflect.Descriptor instead.
func (*SubsReply) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_dialin_proto_rawDescGZIP(), []int{2}
}

func (x *SubsReply) GetSubscriptionId() uint32 {
	if x != nil {
		return x.SubscriptionId
	}
	return 0
}

func (x *SubsReply) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *SubsReply) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

func (m *SubsReply) GetMessageData() isSubsReply_MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (x *SubsReply) GetMessage() []byte {
	if x, ok := x.GetMessageData().(*SubsReply_Message); ok {
		return x.Message
	}
	return nil
}

func (x *SubsReply) GetMessageJson() string {
	if x, ok := x.GetMessageData().(*SubsReply_MessageJson); ok {
		return x.MessageJson
	}
	return ""
}

type isSubsReply_MessageData interface {
	isSubsReply_MessageData()
}

type SubsReply_Message struct {
	Message []byte `protobuf:"bytes,4,opt,name=message,proto3,oneof"` //Sampled data in GPB encoding format is carried.
}

type SubsReply_MessageJson struct {
	MessageJson string `protobuf:"bytes,5,opt,name=message_json,json=messageJson,proto3,oneof"` //Sampled data in JSON encoding format is carried.
}

func (*SubsReply_Message) isSubsReply_MessageData() {}

func (*SubsReply_MessageJson) isSubsReply_MessageData() {}

type CancelArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId      uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`                //Request ID, which is transferred by the invoker.
	SubscriptionId uint32 `protobuf:"varint,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"` //ID of the subscription to be canceled.
}

func (x *CancelArgs) Reset() {
	*x = CancelArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_dialin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelArgs) ProtoMessage() {}

func (x *CancelArgs) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_dialin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelArgs.ProtoReflect.Descriptor instead.
func (*CancelArgs) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_dialin_proto_rawDescGZIP(), []int{3}
}

func (x *CancelArgs) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *CancelArgs) GetSubscriptionId() uint32 {
	if x != nil {
		return x.SubscriptionId
	}
	return 0
}

type CancelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId    uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`         //Request ID, which is transferred by the invoker.
	ResponseCode string `protobuf:"bytes,2,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"` //Return code. The value 200 indicates a success.
	Message      string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`                               //Error description.
}

func (x *CancelReply) Reset() {
	*x = CancelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_dialin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelReply) ProtoMessage() {}

func (x *CancelReply) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_dialin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelReply.ProtoReflect.Descriptor instead.
func (*CancelReply) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_dialin_proto_rawDescGZIP(), []int{4}
}

func (x *CancelReply) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *CancelReply) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

func (x *CancelReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_huawei_grpc_dialin_proto protoreflect.FileDescriptor

var file_huawei_grpc_dialin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x69,
	0x61, 0x6c, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x75, 0x61, 0x77,
	0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x69, 0x6e, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x22, 0xa4, 0x02, 0x0a, 0x08,
	0x53, 0x75, 0x62, 0x73, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x69,
	0x6e, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x11, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x11, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x64, 0x75,
	0x6e, 0x64, 0x61, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x42, 0x0d,
	0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a,
	0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x97, 0x01, 0x0a, 0x0e, 0x67, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f,
	0x70, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x17, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x69, 0x6e,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x18, 0x2e, 0x68, 0x75, 0x61, 0x77,
	0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x12, 0x19, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x69,
	0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1a, 0x2e, 0x68,
	0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x68, 0x75,
	0x61, 0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x69, 0x6e, 0x2f, 0x64, 0x69, 0x61, 0x6c,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_huawei_grpc_dialin_proto_rawDescOnce sync.Once
	file_huawei_grpc_dialin_proto_rawDescData = file_huawei_grpc_dialin_proto_rawDesc
)

func file_huawei_grpc_dialin_proto_rawDescGZIP() []byte {
	file_huawei_grpc_dialin_proto_rawDescOnce.Do(func() {
		file_huawei_grpc_dialin_proto_rawDescData = protoimpl.X.CompressGZIP(file_huawei_grpc_dialin_proto_rawDescData)
	})
	return file_huawei_grpc_dialin_proto_rawDescData
}

var file_huawei_grpc_dialin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_huawei_grpc_dialin_proto_goTypes = []interface{}{
	(*Path)(nil),        // 0: huawei_dialin.Path
	(*SubsArgs)(nil),    // 1: huawei_dialin.SubsArgs
	(*SubsReply)(nil),   // 2: huawei_dialin.SubsReply
	(*CancelArgs)(nil),  // 3: huawei_dialin.CancelArgs
	(*CancelReply)(nil), // 4: huawei_dialin.CancelReply
}
var file_huawei_grpc_dialin_proto_depIdxs = []int32{
	0, // 0: huawei_dialin.SubsArgs.path:type_name -> huawei_dialin.Path
	1, // 1: huawei_dialin.gRPCConfigOper.Subscribe:input_type -> huawei_dialin.SubsArgs
	3, // 2: huawei_dialin.gRPCConfigOper.Cancel:input_type -> huawei_dialin.CancelArgs
	2, // 3: huawei_dialin.gRPCConfigOper.Subscribe:output_type -> huawei_dialin.SubsReply
	4, // 4: huawei_dialin.gRPCConfigOper.Cancel:output_type -> huawei_dialin.CancelReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_huawei_grpc_dialin_proto_init() }
func file_huawei_grpc_dialin_proto_init() {
	if File_huawei_grpc_dialin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_huawei_grpc_dialin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_huawei_grpc_dialin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubsArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_huawei_grpc_dialin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubsReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_huawei_grpc_dialin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_huawei_grpc_dialin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_huawei_grpc_dialin_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SubsArgs_SuppressRedundant)(nil),
		(*SubsArgs_DelayTime)(nil),
	}
	file_huawei_grpc_dialin_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SubsReply_Message)(nil),
		(*SubsReply_MessageJson)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_huawei_grpc_dialin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_huawei_grpc_dialin_proto_goTypes,
		DependencyIndexes: file_huawei_grpc_dialin_proto_depIdxs,
		MessageInfos:      file_huawei_grpc_dialin_proto_msgTypes,
	}.Build()
	File_huawei_grpc_dialin_proto = out.File
	file_huawei_grpc_dialin_proto_rawDesc = nil
	file_huawei_grpc_dialin_proto_goTypes = nil
	file_huawei_grpc_dialin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GRPCConfigOperClient is the client API for GRPCConfigOper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCConfigOperClient interface {
	Subscribe(ctx context.Context, in *SubsArgs, opts ...grpc.CallOption) (GRPCConfigOper_SubscribeClient, error)
	Cancel(ctx context.Context, in *CancelArgs, opts ...grpc.CallOption) (*CancelReply, error)
}

type gRPCConfigOperClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCConfigOperClient(cc grpc.ClientConnInterface) GRPCConfigOperClient {
	return &gRPCConfigOperClient{cc}
}

func (c *gRPCConfigOperClient) Subscribe(ctx context.Context, in *SubsArgs, opts ...grpc.CallOption) (GRPCConfigOper_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GRPCConfigOper_serviceDesc.Streams[0], "/huawei_dialin.gRPCConfigOper/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCConfigOperSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GRPCConfigOper_SubscribeClient interface {
	Recv() (*SubsReply, error)
	grpc.ClientStream
}

type gRPCConfigOperSubscribeClient struct {
	grpc.ClientStream
}

func (x *gRPCConfigOperSubscribeClient) Recv() (*SubsReply, error) {
	m := new(SubsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCConfigOperClient) Cancel(ctx context.Context, in *CancelArgs, opts ...grpc.CallOption) (*CancelReply, error) {
	out := new(CancelReply)
	err := c.cc.Invoke(ctx, "/huawei_dialin.gRPCConfigOper/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCConfigOperServer is the server API for GRPCConfigOper service.
type GRPCConfigOperServer interface {
	Subscribe(*SubsArgs, GRPCConfigOper_SubscribeServer) error
	Cancel(context.Context, *CancelArgs) (*CancelReply, error)
}

// UnimplementedGRPCConfigOperServer can be embedded to have forward compatible implementations.
type UnimplementedGRPCConfigOperServer struct {
}

func (*UnimplementedGRPCConfigOperServer) Subscribe(*SubsArgs, GRPCConfigOper_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedGRPCConfigOperServer) Cancel(context.Context, *CancelArgs) (*CancelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterGRPCConfigOperServer(s *grpc.Server, srv GRPCConfigOperServer) {
	s.RegisterService(&_GRPCConfigOper_serviceDesc, srv)
}

func _GRPCConfigOper_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubsArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GRPCConfigOperServer).Subscribe(m, &gRPCConfigOperSubscribeServer{stream})
}

type GRPCConfigOper_SubscribeServer interface {
	Send(*SubsReply) error
	grpc.ServerStream
}

type gRPCConfigOperSubscribeServer struct {
	grpc.ServerStream
}

func (x *gRPCConfigOperSubscribeServer) Send(m *SubsReply) error {
	return x.ServerStream.SendMsg(m)
}

func _GRPCConfigOper_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCConfigOperServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huawei_dialin.gRPCConfigOper/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCConfigOperServer).Cancel(ctx, req.(*CancelArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCConfigOper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "huawei_dialin.gRPCConfigOper",
	HandlerType: (*GRPCConfigOperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Cancel",
			Handler:    _GRPCConfigOper_Cancel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _GRPCConfigOper_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "huawei-grpc-dialin.proto",
}
