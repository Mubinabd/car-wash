// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/notification.proto

package genproto

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

type AddNotificationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	BookingId string `protobuf:"bytes,3,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *AddNotificationReq) Reset() {
	*x = AddNotificationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNotificationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNotificationReq) ProtoMessage() {}

func (x *AddNotificationReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNotificationReq.ProtoReflect.Descriptor instead.
func (*AddNotificationReq) Descriptor() ([]byte, []int) {
	return file_protos_notification_proto_rawDescGZIP(), []int{0}
}

func (x *AddNotificationReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddNotificationReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddNotificationReq) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	IsRead    bool   `protobuf:"varint,5,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	BookingId string `protobuf:"bytes,6,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_protos_notification_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Notification) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *Notification) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

type GetNotificationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetNotificationsReq) Reset() {
	*x = GetNotificationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsReq) ProtoMessage() {}

func (x *GetNotificationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsReq.ProtoReflect.Descriptor instead.
func (*GetNotificationsReq) Descriptor() ([]byte, []int) {
	return file_protos_notification_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotificationsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type MarkNotificationAsReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MarkNotificationAsReadReq) Reset() {
	*x = MarkNotificationAsReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkNotificationAsReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkNotificationAsReadReq) ProtoMessage() {}

func (x *MarkNotificationAsReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkNotificationAsReadReq.ProtoReflect.Descriptor instead.
func (*MarkNotificationAsReadReq) Descriptor() ([]byte, []int) {
	return file_protos_notification_proto_rawDescGZIP(), []int{3}
}

func (x *MarkNotificationAsReadReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MarkNotificationAsReadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MarkNotificationAsReadResp) Reset() {
	*x = MarkNotificationAsReadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkNotificationAsReadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkNotificationAsReadResp) ProtoMessage() {}

func (x *MarkNotificationAsReadResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkNotificationAsReadResp.ProtoReflect.Descriptor instead.
func (*MarkNotificationAsReadResp) Descriptor() ([]byte, []int) {
	return file_protos_notification_proto_rawDescGZIP(), []int{4}
}

func (x *MarkNotificationAsReadResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MarkNotificationAsReadResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_notification_proto protoreflect.FileDescriptor

var file_protos_notification_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x72,
	0x5f, 0x77, 0x61, 0x73, 0x68, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a,
	0x19, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x1a, 0x4d, 0x61,
	0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x87, 0x02, 0x0a,
	0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61,
	0x73, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x72,
	0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x77, 0x61, 0x73, 0x68, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x63, 0x0a, 0x16, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x23, 0x2e, 0x63, 0x61,
	0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x24, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_notification_proto_rawDescOnce sync.Once
	file_protos_notification_proto_rawDescData = file_protos_notification_proto_rawDesc
)

func file_protos_notification_proto_rawDescGZIP() []byte {
	file_protos_notification_proto_rawDescOnce.Do(func() {
		file_protos_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_notification_proto_rawDescData)
	})
	return file_protos_notification_proto_rawDescData
}

var file_protos_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_notification_proto_goTypes = []any{
	(*AddNotificationReq)(nil),         // 0: car_wash.AddNotificationReq
	(*Notification)(nil),               // 1: car_wash.Notification
	(*GetNotificationsReq)(nil),        // 2: car_wash.GetNotificationsReq
	(*MarkNotificationAsReadReq)(nil),  // 3: car_wash.MarkNotificationAsReadReq
	(*MarkNotificationAsReadResp)(nil), // 4: car_wash.MarkNotificationAsReadResp
	(*Empty)(nil),                      // 5: car_wash.Empty
}
var file_protos_notification_proto_depIdxs = []int32{
	0, // 0: car_wash.NotificationService.AddNotification:input_type -> car_wash.AddNotificationReq
	2, // 1: car_wash.NotificationService.GetNotifications:input_type -> car_wash.GetNotificationsReq
	3, // 2: car_wash.NotificationService.MarkNotificationAsRead:input_type -> car_wash.MarkNotificationAsReadReq
	5, // 3: car_wash.NotificationService.AddNotification:output_type -> car_wash.Empty
	1, // 4: car_wash.NotificationService.GetNotifications:output_type -> car_wash.Notification
	4, // 5: car_wash.NotificationService.MarkNotificationAsRead:output_type -> car_wash.MarkNotificationAsReadResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_notification_proto_init() }
func file_protos_notification_proto_init() {
	if File_protos_notification_proto != nil {
		return
	}
	file_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_notification_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddNotificationReq); i {
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
		file_protos_notification_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Notification); i {
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
		file_protos_notification_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetNotificationsReq); i {
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
		file_protos_notification_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MarkNotificationAsReadReq); i {
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
		file_protos_notification_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MarkNotificationAsReadResp); i {
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
			RawDescriptor: file_protos_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_notification_proto_goTypes,
		DependencyIndexes: file_protos_notification_proto_depIdxs,
		MessageInfos:      file_protos_notification_proto_msgTypes,
	}.Build()
	File_protos_notification_proto = out.File
	file_protos_notification_proto_rawDesc = nil
	file_protos_notification_proto_goTypes = nil
	file_protos_notification_proto_depIdxs = nil
}