// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: grpc/proto.proto

package proto

import (
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

type AckStatus int32

const (
	AckStatus_FAIL      AckStatus = 0
	AckStatus_SUCCESS   AckStatus = 1
	AckStatus_EXCEPTION AckStatus = 2
)

// Enum value maps for AckStatus.
var (
	AckStatus_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
		2: "EXCEPTION",
	}
	AckStatus_value = map[string]int32{
		"FAIL":      0,
		"SUCCESS":   1,
		"EXCEPTION": 2,
	}
)

func (x AckStatus) Enum() *AckStatus {
	p := new(AckStatus)
	*p = x
	return p
}

func (x AckStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AckStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_proto_proto_enumTypes[0].Descriptor()
}

func (AckStatus) Type() protoreflect.EnumType {
	return &file_grpc_proto_proto_enumTypes[0]
}

func (x AckStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AckStatus.Descriptor instead.
func (AckStatus) EnumDescriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

type TimesBiddedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TimesBiddedRequest) Reset() {
	*x = TimesBiddedRequest{}
	mi := &file_grpc_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimesBiddedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimesBiddedRequest) ProtoMessage() {}

func (x *TimesBiddedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimesBiddedRequest.ProtoReflect.Descriptor instead.
func (*TimesBiddedRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

type TimesBiddedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimesBidded int32 `protobuf:"varint,1,opt,name=timesBidded,proto3" json:"timesBidded,omitempty"`
}

func (x *TimesBiddedResponse) Reset() {
	*x = TimesBiddedResponse{}
	mi := &file_grpc_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimesBiddedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimesBiddedResponse) ProtoMessage() {}

func (x *TimesBiddedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimesBiddedResponse.ProtoReflect.Descriptor instead.
func (*TimesBiddedResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{1}
}

func (x *TimesBiddedResponse) GetTimesBidded() int32 {
	if x != nil {
		return x.TimesBidded
	}
	return 0
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID string `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	mi := &file_grpc_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{2}
}

func (x *BidRequest) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack     AckStatus `protobuf:"varint,1,opt,name=ack,proto3,enum=AckStatus" json:"ack,omitempty"`
	Comment string    `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	mi := &file_grpc_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{3}
}

func (x *BidResponse) GetAck() AckStatus {
	if x != nil {
		return x.Ack
	}
	return AckStatus_FAIL
}

func (x *BidResponse) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	mi := &file_grpc_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{4}
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner string `protobuf:"bytes,1,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	mi := &file_grpc_proto_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{5}
}

func (x *ResultResponse) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

type HighestBidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HighestBidRequest) Reset() {
	*x = HighestBidRequest{}
	mi := &file_grpc_proto_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HighestBidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HighestBidRequest) ProtoMessage() {}

func (x *HighestBidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HighestBidRequest.ProtoReflect.Descriptor instead.
func (*HighestBidRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{6}
}

type HighestBidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighesBid int32 `protobuf:"varint,1,opt,name=HighesBid,proto3" json:"HighesBid,omitempty"`
}

func (x *HighestBidResponse) Reset() {
	*x = HighestBidResponse{}
	mi := &file_grpc_proto_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HighestBidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HighestBidResponse) ProtoMessage() {}

func (x *HighestBidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HighestBidResponse.ProtoReflect.Descriptor instead.
func (*HighestBidResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{7}
}

func (x *HighestBidResponse) GetHighesBid() int32 {
	if x != nil {
		return x.HighesBid
	}
	return 0
}

var File_grpc_proto_proto protoreflect.FileDescriptor

var file_grpc_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x69, 0x64, 0x64, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x13, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x42, 0x69, 0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x69, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x69, 0x64, 0x64, 0x65,
	0x64, 0x22, 0x3c, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x45, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x41, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x22, 0x13, 0x0a, 0x11, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x12, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x42, 0x69, 0x64, 0x2a, 0x31, 0x0a, 0x09, 0x41, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0xdc, 0x01,
	0x0a, 0x0e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x12, 0x0b, 0x2e, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x42, 0x69, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x12, 0x2e, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x48, 0x69, 0x67,
	0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x69, 0x64, 0x64, 0x65,
	0x64, 0x12, 0x13, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x69, 0x64, 0x64, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x69,
	0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12,
	0x68, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x35, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_proto_rawDescOnce sync.Once
	file_grpc_proto_proto_rawDescData = file_grpc_proto_proto_rawDesc
)

func file_grpc_proto_proto_rawDescGZIP() []byte {
	file_grpc_proto_proto_rawDescOnce.Do(func() {
		file_grpc_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_proto_rawDescData)
	})
	return file_grpc_proto_proto_rawDescData
}

var file_grpc_proto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_proto_proto_goTypes = []any{
	(AckStatus)(0),              // 0: AckStatus
	(*TimesBiddedRequest)(nil),  // 1: TimesBiddedRequest
	(*TimesBiddedResponse)(nil), // 2: TimesBiddedResponse
	(*BidRequest)(nil),          // 3: BidRequest
	(*BidResponse)(nil),         // 4: BidResponse
	(*ResultRequest)(nil),       // 5: ResultRequest
	(*ResultResponse)(nil),      // 6: ResultResponse
	(*HighestBidRequest)(nil),   // 7: HighestBidRequest
	(*HighestBidResponse)(nil),  // 8: HighestBidResponse
}
var file_grpc_proto_proto_depIdxs = []int32{
	0, // 0: BidResponse.ack:type_name -> AckStatus
	3, // 1: BiddingService.PlaceBid:input_type -> BidRequest
	5, // 2: BiddingService.GetResult:input_type -> ResultRequest
	7, // 3: BiddingService.GetHighestBid:input_type -> HighestBidRequest
	1, // 4: BiddingService.GetTimesBidded:input_type -> TimesBiddedRequest
	4, // 5: BiddingService.PlaceBid:output_type -> BidResponse
	6, // 6: BiddingService.GetResult:output_type -> ResultResponse
	8, // 7: BiddingService.GetHighestBid:output_type -> HighestBidResponse
	2, // 8: BiddingService.GetTimesBidded:output_type -> TimesBiddedResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_proto_init() }
func file_grpc_proto_proto_init() {
	if File_grpc_proto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_proto_goTypes,
		DependencyIndexes: file_grpc_proto_proto_depIdxs,
		EnumInfos:         file_grpc_proto_proto_enumTypes,
		MessageInfos:      file_grpc_proto_proto_msgTypes,
	}.Build()
	File_grpc_proto_proto = out.File
	file_grpc_proto_proto_rawDesc = nil
	file_grpc_proto_proto_goTypes = nil
	file_grpc_proto_proto_depIdxs = nil
}
