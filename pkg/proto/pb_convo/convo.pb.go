// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pb_convo/convo.proto

package pb_convo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pb_enum "lark/pkg/proto/pb_enum"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConvoListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatIds string `protobuf:"bytes,1,opt,name=chat_ids,json=chatIds,proto3" json:"chat_ids,omitempty"` // chat_id 列表(base64(gzip(chat_id1,chat_id2)))
}

func (x *ConvoListReq) Reset() {
	*x = ConvoListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_convo_convo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvoListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvoListReq) ProtoMessage() {}

func (x *ConvoListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_convo_convo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvoListReq.ProtoReflect.Descriptor instead.
func (*ConvoListReq) Descriptor() ([]byte, []int) {
	return file_pb_convo_convo_proto_rawDescGZIP(), []int{0}
}

func (x *ConvoListReq) GetChatIds() string {
	if x != nil {
		return x.ChatIds
	}
	return ""
}

type ConvoListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	List []*ConvoMessage `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ConvoListResp) Reset() {
	*x = ConvoListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_convo_convo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvoListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvoListResp) ProtoMessage() {}

func (x *ConvoListResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_convo_convo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvoListResp.ProtoReflect.Descriptor instead.
func (*ConvoListResp) Descriptor() ([]byte, []int) {
	return file_pb_convo_convo_proto_rawDescGZIP(), []int{1}
}

func (x *ConvoListResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConvoListResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ConvoListResp) GetList() []*ConvoMessage {
	if x != nil {
		return x.List
	}
	return nil
}

type ConvoChatSeqListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	LastCid   int64  `protobuf:"varint,2,opt,name=last_cid,json=lastCid,proto3" json:"last_cid,omitempty"`
	LastTs    int64  `protobuf:"varint,3,opt,name=last_ts,json=lastTs,proto3" json:"last_ts,omitempty"`
	Limit     int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`           //后面的已弃用
	ChatIds   string `protobuf:"bytes,6,opt,name=chat_ids,json=chatIds,proto3" json:"chat_ids,omitempty"` // chat_id 列表(base64(gzip(chat_id1,chat_id2)))
}

func (x *ConvoChatSeqListReq) Reset() {
	*x = ConvoChatSeqListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_convo_convo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvoChatSeqListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvoChatSeqListReq) ProtoMessage() {}

func (x *ConvoChatSeqListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_convo_convo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvoChatSeqListReq.ProtoReflect.Descriptor instead.
func (*ConvoChatSeqListReq) Descriptor() ([]byte, []int) {
	return file_pb_convo_convo_proto_rawDescGZIP(), []int{2}
}

func (x *ConvoChatSeqListReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ConvoChatSeqListReq) GetLastCid() int64 {
	if x != nil {
		return x.LastCid
	}
	return 0
}

func (x *ConvoChatSeqListReq) GetLastTs() int64 {
	if x != nil {
		return x.LastTs
	}
	return 0
}

func (x *ConvoChatSeqListReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ConvoChatSeqListReq) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ConvoChatSeqListReq) GetChatIds() string {
	if x != nil {
		return x.ChatIds
	}
	return ""
}

type ConvoChatSeqListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	List []*ConvoChatSeq `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ConvoChatSeqListResp) Reset() {
	*x = ConvoChatSeqListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_convo_convo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvoChatSeqListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvoChatSeqListResp) ProtoMessage() {}

func (x *ConvoChatSeqListResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_convo_convo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvoChatSeqListResp.ProtoReflect.Descriptor instead.
func (*ConvoChatSeqListResp) Descriptor() ([]byte, []int) {
	return file_pb_convo_convo_proto_rawDescGZIP(), []int{3}
}

func (x *ConvoChatSeqListResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConvoChatSeqListResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ConvoChatSeqListResp) GetList() []*ConvoChatSeq {
	if x != nil {
		return x.List
	}
	return nil
}

type ConvoChatSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  int64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	SeqId   int64 `protobuf:"varint,2,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	SrvTs   int64 `protobuf:"varint,3,opt,name=srv_ts,json=srvTs,proto3" json:"srv_ts,omitempty"`
	ReadSeq int64 `protobuf:"varint,4,opt,name=read_seq,json=readSeq,proto3" json:"read_seq,omitempty"`
}

func (x *ConvoChatSeq) Reset() {
	*x = ConvoChatSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_convo_convo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvoChatSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvoChatSeq) ProtoMessage() {}

func (x *ConvoChatSeq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_convo_convo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvoChatSeq.ProtoReflect.Descriptor instead.
func (*ConvoChatSeq) Descriptor() ([]byte, []int) {
	return file_pb_convo_convo_proto_rawDescGZIP(), []int{4}
}

func (x *ConvoChatSeq) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *ConvoChatSeq) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *ConvoChatSeq) GetSrvTs() int64 {
	if x != nil {
		return x.SrvTs
	}
	return 0
}

func (x *ConvoChatSeq) GetReadSeq() int64 {
	if x != nil {
		return x.ReadSeq
	}
	return 0
}

type ConvoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId     int64             `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`                              // 会话ID
	SeqId      int64             `protobuf:"varint,2,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`                                 // 消息唯一ID
	SenderId   int64             `protobuf:"varint,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`                        // 发送者uid
	SenderName string            `protobuf:"bytes,4,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"`                   // 发送者姓名
	ChatType   pb_enum.CHAT_TYPE `protobuf:"varint,5,opt,name=chat_type,json=chatType,proto3,enum=pb_enum.CHAT_TYPE" json:"chat_type,omitempty"` // 会话类型
	MsgFrom    pb_enum.MSG_FROM  `protobuf:"varint,6,opt,name=msg_from,json=msgFrom,proto3,enum=pb_enum.MSG_FROM" json:"msg_from,omitempty"`     // 消息来源
	MsgType    pb_enum.MSG_TYPE  `protobuf:"varint,7,opt,name=msg_type,json=msgType,proto3,enum=pb_enum.MSG_TYPE" json:"msg_type,omitempty"`     // 消息类型
	Body       string            `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`                                                 // 消息本体
	Status     int32             `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`                                            // 消息状态
	SrvTs      int64             `protobuf:"varint,10,opt,name=srv_ts,json=srvTs,proto3" json:"srv_ts,omitempty"`                                // 服务端接收消息的时间
}

func (x *ConvoMessage) Reset() {
	*x = ConvoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_convo_convo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvoMessage) ProtoMessage() {}

func (x *ConvoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_convo_convo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvoMessage.ProtoReflect.Descriptor instead.
func (*ConvoMessage) Descriptor() ([]byte, []int) {
	return file_pb_convo_convo_proto_rawDescGZIP(), []int{5}
}

func (x *ConvoMessage) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *ConvoMessage) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *ConvoMessage) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *ConvoMessage) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

func (x *ConvoMessage) GetChatType() pb_enum.CHAT_TYPE {
	if x != nil {
		return x.ChatType
	}
	return pb_enum.CHAT_TYPE(0)
}

func (x *ConvoMessage) GetMsgFrom() pb_enum.MSG_FROM {
	if x != nil {
		return x.MsgFrom
	}
	return pb_enum.MSG_FROM(0)
}

func (x *ConvoMessage) GetMsgType() pb_enum.MSG_TYPE {
	if x != nil {
		return x.MsgType
	}
	return pb_enum.MSG_TYPE(0)
}

func (x *ConvoMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ConvoMessage) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ConvoMessage) GetSrvTs() int64 {
	if x != nil {
		return x.SrvTs
	}
	return 0
}

var File_pb_convo_convo_proto protoreflect.FileDescriptor

var file_pb_convo_convo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x6f,
	0x1a, 0x12, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x73, 0x22,
	0x61, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6c, 0x61, 0x73, 0x74, 0x43, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x73, 0x22,
	0x68, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x71, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2a, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62,
	0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x65, 0x71, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x76, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x72, 0x76,
	0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x72, 0x76, 0x54, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x53, 0x65, 0x71, 0x22, 0xcc, 0x02, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d,
	0x73, 0x67, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x53, 0x47, 0x5f, 0x46, 0x52, 0x4f, 0x4d,
	0x52, 0x07, 0x6d, 0x73, 0x67, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x07,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x72, 0x76, 0x5f, 0x74, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x72, 0x76, 0x54, 0x73, 0x32, 0x98, 0x01, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x76, 0x6f, 0x12, 0x3c, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x51, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x71, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x76,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x71, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x71, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x63, 0x6f,
	0x6e, 0x76, 0x6f, 0x3b, 0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_convo_convo_proto_rawDescOnce sync.Once
	file_pb_convo_convo_proto_rawDescData = file_pb_convo_convo_proto_rawDesc
)

func file_pb_convo_convo_proto_rawDescGZIP() []byte {
	file_pb_convo_convo_proto_rawDescOnce.Do(func() {
		file_pb_convo_convo_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_convo_convo_proto_rawDescData)
	})
	return file_pb_convo_convo_proto_rawDescData
}

var file_pb_convo_convo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_convo_convo_proto_goTypes = []interface{}{
	(*ConvoListReq)(nil),         // 0: pb_convo.ConvoListReq
	(*ConvoListResp)(nil),        // 1: pb_convo.ConvoListResp
	(*ConvoChatSeqListReq)(nil),  // 2: pb_convo.ConvoChatSeqListReq
	(*ConvoChatSeqListResp)(nil), // 3: pb_convo.ConvoChatSeqListResp
	(*ConvoChatSeq)(nil),         // 4: pb_convo.ConvoChatSeq
	(*ConvoMessage)(nil),         // 5: pb_convo.ConvoMessage
	(pb_enum.CHAT_TYPE)(0),       // 6: pb_enum.CHAT_TYPE
	(pb_enum.MSG_FROM)(0),        // 7: pb_enum.MSG_FROM
	(pb_enum.MSG_TYPE)(0),        // 8: pb_enum.MSG_TYPE
}
var file_pb_convo_convo_proto_depIdxs = []int32{
	5, // 0: pb_convo.ConvoListResp.list:type_name -> pb_convo.ConvoMessage
	4, // 1: pb_convo.ConvoChatSeqListResp.list:type_name -> pb_convo.ConvoChatSeq
	6, // 2: pb_convo.ConvoMessage.chat_type:type_name -> pb_enum.CHAT_TYPE
	7, // 3: pb_convo.ConvoMessage.msg_from:type_name -> pb_enum.MSG_FROM
	8, // 4: pb_convo.ConvoMessage.msg_type:type_name -> pb_enum.MSG_TYPE
	0, // 5: pb_convo.Convo.ConvoList:input_type -> pb_convo.ConvoListReq
	2, // 6: pb_convo.Convo.ConvoChatSeqList:input_type -> pb_convo.ConvoChatSeqListReq
	1, // 7: pb_convo.Convo.ConvoList:output_type -> pb_convo.ConvoListResp
	3, // 8: pb_convo.Convo.ConvoChatSeqList:output_type -> pb_convo.ConvoChatSeqListResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pb_convo_convo_proto_init() }
func file_pb_convo_convo_proto_init() {
	if File_pb_convo_convo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_convo_convo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvoListReq); i {
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
		file_pb_convo_convo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvoListResp); i {
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
		file_pb_convo_convo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvoChatSeqListReq); i {
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
		file_pb_convo_convo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvoChatSeqListResp); i {
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
		file_pb_convo_convo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvoChatSeq); i {
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
		file_pb_convo_convo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvoMessage); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_convo_convo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_convo_convo_proto_goTypes,
		DependencyIndexes: file_pb_convo_convo_proto_depIdxs,
		MessageInfos:      file_pb_convo_convo_proto_msgTypes,
	}.Build()
	File_pb_convo_convo_proto = out.File
	file_pb_convo_convo_proto_rawDesc = nil
	file_pb_convo_convo_proto_goTypes = nil
	file_pb_convo_convo_proto_depIdxs = nil
}
