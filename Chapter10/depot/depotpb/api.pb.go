// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: depotpb/api.proto

package depotpb

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

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StoreId   string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Quantity  int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ShoppingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string           `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Stops         map[string]*Stop `protobuf:"bytes,3,rep,name=stops,proto3" json:"stops,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AssignedBotId string           `protobuf:"bytes,4,opt,name=assigned_bot_id,json=assignedBotId,proto3" json:"assigned_bot_id,omitempty"`
	Status        string           `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ShoppingList) Reset() {
	*x = ShoppingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShoppingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingList) ProtoMessage() {}

func (x *ShoppingList) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingList.ProtoReflect.Descriptor instead.
func (*ShoppingList) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{1}
}

func (x *ShoppingList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShoppingList) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ShoppingList) GetStops() map[string]*Stop {
	if x != nil {
		return x.Stops
	}
	return nil
}

func (x *ShoppingList) GetAssignedBotId() string {
	if x != nil {
		return x.AssignedBotId
	}
	return ""
}

func (x *ShoppingList) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreName     string           `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	StoreLocation string           `protobuf:"bytes,2,opt,name=store_location,json=storeLocation,proto3" json:"store_location,omitempty"`
	Items         map[string]*Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{2}
}

func (x *Stop) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Stop) GetStoreLocation() string {
	if x != nil {
		return x.StoreLocation
	}
	return ""
}

func (x *Stop) GetItems() map[string]*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateShoppingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string       `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Items   []*OrderItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateShoppingListRequest) Reset() {
	*x = CreateShoppingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShoppingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShoppingListRequest) ProtoMessage() {}

func (x *CreateShoppingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShoppingListRequest.ProtoReflect.Descriptor instead.
func (*CreateShoppingListRequest) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateShoppingListRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateShoppingListRequest) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateShoppingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateShoppingListResponse) Reset() {
	*x = CreateShoppingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShoppingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShoppingListResponse) ProtoMessage() {}

func (x *CreateShoppingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShoppingListResponse.ProtoReflect.Descriptor instead.
func (*CreateShoppingListResponse) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreateShoppingListResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelShoppingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelShoppingListRequest) Reset() {
	*x = CancelShoppingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelShoppingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShoppingListRequest) ProtoMessage() {}

func (x *CancelShoppingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShoppingListRequest.ProtoReflect.Descriptor instead.
func (*CancelShoppingListRequest) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{6}
}

func (x *CancelShoppingListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelShoppingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelShoppingListResponse) Reset() {
	*x = CancelShoppingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelShoppingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShoppingListResponse) ProtoMessage() {}

func (x *CancelShoppingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShoppingListResponse.ProtoReflect.Descriptor instead.
func (*CancelShoppingListResponse) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{7}
}

type AssignShoppingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BotId string `protobuf:"bytes,2,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
}

func (x *AssignShoppingListRequest) Reset() {
	*x = AssignShoppingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignShoppingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignShoppingListRequest) ProtoMessage() {}

func (x *AssignShoppingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignShoppingListRequest.ProtoReflect.Descriptor instead.
func (*AssignShoppingListRequest) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{8}
}

func (x *AssignShoppingListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssignShoppingListRequest) GetBotId() string {
	if x != nil {
		return x.BotId
	}
	return ""
}

type AssignShoppingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssignShoppingListResponse) Reset() {
	*x = AssignShoppingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignShoppingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignShoppingListResponse) ProtoMessage() {}

func (x *AssignShoppingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignShoppingListResponse.ProtoReflect.Descriptor instead.
func (*AssignShoppingListResponse) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{9}
}

type CompleteShoppingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CompleteShoppingListRequest) Reset() {
	*x = CompleteShoppingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteShoppingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteShoppingListRequest) ProtoMessage() {}

func (x *CompleteShoppingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteShoppingListRequest.ProtoReflect.Descriptor instead.
func (*CompleteShoppingListRequest) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{10}
}

func (x *CompleteShoppingListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CompleteShoppingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteShoppingListResponse) Reset() {
	*x = CompleteShoppingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depotpb_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteShoppingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteShoppingListResponse) ProtoMessage() {}

func (x *CompleteShoppingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_depotpb_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteShoppingListResponse.ProtoReflect.Descriptor instead.
func (*CompleteShoppingListResponse) Descriptor() ([]byte, []int) {
	return file_depotpb_api_proto_rawDescGZIP(), []int{11}
}

var File_depotpb_api_proto protoreflect.FileDescriptor

var file_depotpb_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x22, 0x61, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0xfa, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x74,
	0x6f, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x47, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc5, 0x01, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x47, 0x0a, 0x0a, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x36, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x60, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2c,
	0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x19,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x19, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x1b, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x98, 0x03, 0x0a, 0x0c, 0x44, 0x65, 0x70,
	0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x22, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x22, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x22, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x14,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x78, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74,
	0x70, 0x62, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23,
	0x65, 0x64, 0x61, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x65,
	0x70, 0x6f, 0x74, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x6f,
	0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x44, 0x65, 0x70, 0x6f,
	0x74, 0x70, 0x62, 0xca, 0x02, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0xe2, 0x02, 0x13,
	0x44, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_depotpb_api_proto_rawDescOnce sync.Once
	file_depotpb_api_proto_rawDescData = file_depotpb_api_proto_rawDesc
)

func file_depotpb_api_proto_rawDescGZIP() []byte {
	file_depotpb_api_proto_rawDescOnce.Do(func() {
		file_depotpb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_depotpb_api_proto_rawDescData)
	})
	return file_depotpb_api_proto_rawDescData
}

var file_depotpb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_depotpb_api_proto_goTypes = []interface{}{
	(*OrderItem)(nil),                    // 0: depotpb.OrderItem
	(*ShoppingList)(nil),                 // 1: depotpb.ShoppingList
	(*Stop)(nil),                         // 2: depotpb.Stop
	(*Item)(nil),                         // 3: depotpb.Item
	(*CreateShoppingListRequest)(nil),    // 4: depotpb.CreateShoppingListRequest
	(*CreateShoppingListResponse)(nil),   // 5: depotpb.CreateShoppingListResponse
	(*CancelShoppingListRequest)(nil),    // 6: depotpb.CancelShoppingListRequest
	(*CancelShoppingListResponse)(nil),   // 7: depotpb.CancelShoppingListResponse
	(*AssignShoppingListRequest)(nil),    // 8: depotpb.AssignShoppingListRequest
	(*AssignShoppingListResponse)(nil),   // 9: depotpb.AssignShoppingListResponse
	(*CompleteShoppingListRequest)(nil),  // 10: depotpb.CompleteShoppingListRequest
	(*CompleteShoppingListResponse)(nil), // 11: depotpb.CompleteShoppingListResponse
	nil,                                  // 12: depotpb.ShoppingList.StopsEntry
	nil,                                  // 13: depotpb.Stop.ItemsEntry
}
var file_depotpb_api_proto_depIdxs = []int32{
	12, // 0: depotpb.ShoppingList.stops:type_name -> depotpb.ShoppingList.StopsEntry
	13, // 1: depotpb.Stop.items:type_name -> depotpb.Stop.ItemsEntry
	0,  // 2: depotpb.CreateShoppingListRequest.items:type_name -> depotpb.OrderItem
	2,  // 3: depotpb.ShoppingList.StopsEntry.value:type_name -> depotpb.Stop
	3,  // 4: depotpb.Stop.ItemsEntry.value:type_name -> depotpb.Item
	4,  // 5: depotpb.DepotService.CreateShoppingList:input_type -> depotpb.CreateShoppingListRequest
	6,  // 6: depotpb.DepotService.CancelShoppingList:input_type -> depotpb.CancelShoppingListRequest
	8,  // 7: depotpb.DepotService.AssignShoppingList:input_type -> depotpb.AssignShoppingListRequest
	10, // 8: depotpb.DepotService.CompleteShoppingList:input_type -> depotpb.CompleteShoppingListRequest
	5,  // 9: depotpb.DepotService.CreateShoppingList:output_type -> depotpb.CreateShoppingListResponse
	7,  // 10: depotpb.DepotService.CancelShoppingList:output_type -> depotpb.CancelShoppingListResponse
	9,  // 11: depotpb.DepotService.AssignShoppingList:output_type -> depotpb.AssignShoppingListResponse
	11, // 12: depotpb.DepotService.CompleteShoppingList:output_type -> depotpb.CompleteShoppingListResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_depotpb_api_proto_init() }
func file_depotpb_api_proto_init() {
	if File_depotpb_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_depotpb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_depotpb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShoppingList); i {
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
		file_depotpb_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop); i {
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
		file_depotpb_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_depotpb_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShoppingListRequest); i {
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
		file_depotpb_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShoppingListResponse); i {
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
		file_depotpb_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelShoppingListRequest); i {
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
		file_depotpb_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelShoppingListResponse); i {
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
		file_depotpb_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignShoppingListRequest); i {
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
		file_depotpb_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignShoppingListResponse); i {
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
		file_depotpb_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteShoppingListRequest); i {
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
		file_depotpb_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteShoppingListResponse); i {
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
			RawDescriptor: file_depotpb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_depotpb_api_proto_goTypes,
		DependencyIndexes: file_depotpb_api_proto_depIdxs,
		MessageInfos:      file_depotpb_api_proto_msgTypes,
	}.Build()
	File_depotpb_api_proto = out.File
	file_depotpb_api_proto_rawDesc = nil
	file_depotpb_api_proto_goTypes = nil
	file_depotpb_api_proto_depIdxs = nil
}
