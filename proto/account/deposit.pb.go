// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: deposit.proto

package account

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

type DepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deposit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{0}
}

func (x *DepositRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DepositResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DepositResponse) Reset() {
	*x = DepositResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deposit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositResponse) ProtoMessage() {}

func (x *DepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositResponse.ProtoReflect.Descriptor instead.
func (*DepositResponse) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{1}
}

func (x *DepositResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type GetDepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDepositRequest) Reset() {
	*x = GetDepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deposit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepositRequest) ProtoMessage() {}

func (x *GetDepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepositRequest.ProtoReflect.Descriptor instead.
func (*GetDepositRequest) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{2}
}

type GetDepositResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalDeposit float32 `protobuf:"fixed32,1,opt,name=totalDeposit,proto3" json:"totalDeposit,omitempty"`
}

func (x *GetDepositResponse) Reset() {
	*x = GetDepositResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deposit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepositResponse) ProtoMessage() {}

func (x *GetDepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepositResponse.ProtoReflect.Descriptor instead.
func (*GetDepositResponse) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{3}
}

func (x *GetDepositResponse) GetTotalDeposit() float32 {
	if x != nil {
		return x.TotalDeposit
	}
	return 0
}

var File_deposit_proto protoreflect.FileDescriptor

var file_deposit_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x32, 0x99, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deposit_proto_rawDescOnce sync.Once
	file_deposit_proto_rawDescData = file_deposit_proto_rawDesc
)

func file_deposit_proto_rawDescGZIP() []byte {
	file_deposit_proto_rawDescOnce.Do(func() {
		file_deposit_proto_rawDescData = protoimpl.X.CompressGZIP(file_deposit_proto_rawDescData)
	})
	return file_deposit_proto_rawDescData
}

var file_deposit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_deposit_proto_goTypes = []interface{}{
	(*DepositRequest)(nil),     // 0: account.DepositRequest
	(*DepositResponse)(nil),    // 1: account.DepositResponse
	(*GetDepositRequest)(nil),  // 2: account.GetDepositRequest
	(*GetDepositResponse)(nil), // 3: account.GetDepositResponse
}
var file_deposit_proto_depIdxs = []int32{
	0, // 0: account.DepositService.Deposit:input_type -> account.DepositRequest
	2, // 1: account.DepositService.GetDeposit:input_type -> account.GetDepositRequest
	1, // 2: account.DepositService.Deposit:output_type -> account.DepositResponse
	3, // 3: account.DepositService.GetDeposit:output_type -> account.GetDepositResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deposit_proto_init() }
func file_deposit_proto_init() {
	if File_deposit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deposit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest); i {
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
		file_deposit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositResponse); i {
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
		file_deposit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDepositRequest); i {
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
		file_deposit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDepositResponse); i {
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
			RawDescriptor: file_deposit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deposit_proto_goTypes,
		DependencyIndexes: file_deposit_proto_depIdxs,
		MessageInfos:      file_deposit_proto_msgTypes,
	}.Build()
	File_deposit_proto = out.File
	file_deposit_proto_rawDesc = nil
	file_deposit_proto_goTypes = nil
	file_deposit_proto_depIdxs = nil
}