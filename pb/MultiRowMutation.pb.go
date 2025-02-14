//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: MultiRowMutation.proto

package pb

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

type MultiRowMutationProcessorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MultiRowMutationProcessorRequest) Reset() {
	*x = MultiRowMutationProcessorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultiRowMutation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiRowMutationProcessorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiRowMutationProcessorRequest) ProtoMessage() {}

func (x *MultiRowMutationProcessorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MultiRowMutation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiRowMutationProcessorRequest.ProtoReflect.Descriptor instead.
func (*MultiRowMutationProcessorRequest) Descriptor() ([]byte, []int) {
	return file_MultiRowMutation_proto_rawDescGZIP(), []int{0}
}

type MultiRowMutationProcessorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MultiRowMutationProcessorResponse) Reset() {
	*x = MultiRowMutationProcessorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultiRowMutation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiRowMutationProcessorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiRowMutationProcessorResponse) ProtoMessage() {}

func (x *MultiRowMutationProcessorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MultiRowMutation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiRowMutationProcessorResponse.ProtoReflect.Descriptor instead.
func (*MultiRowMutationProcessorResponse) Descriptor() ([]byte, []int) {
	return file_MultiRowMutation_proto_rawDescGZIP(), []int{1}
}

type MutateRowsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MutationRequest []*MutationProto `protobuf:"bytes,1,rep,name=mutation_request,json=mutationRequest" json:"mutation_request,omitempty"`
	NonceGroup      *uint64          `protobuf:"varint,2,opt,name=nonce_group,json=nonceGroup" json:"nonce_group,omitempty"`
	Nonce           *uint64          `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	Region          *RegionSpecifier `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
	Condition       []*Condition     `protobuf:"bytes,5,rep,name=condition" json:"condition,omitempty"`
}

func (x *MutateRowsRequest) Reset() {
	*x = MutateRowsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultiRowMutation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateRowsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateRowsRequest) ProtoMessage() {}

func (x *MutateRowsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MultiRowMutation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateRowsRequest.ProtoReflect.Descriptor instead.
func (*MutateRowsRequest) Descriptor() ([]byte, []int) {
	return file_MultiRowMutation_proto_rawDescGZIP(), []int{2}
}

func (x *MutateRowsRequest) GetMutationRequest() []*MutationProto {
	if x != nil {
		return x.MutationRequest
	}
	return nil
}

func (x *MutateRowsRequest) GetNonceGroup() uint64 {
	if x != nil && x.NonceGroup != nil {
		return *x.NonceGroup
	}
	return 0
}

func (x *MutateRowsRequest) GetNonce() uint64 {
	if x != nil && x.Nonce != nil {
		return *x.Nonce
	}
	return 0
}

func (x *MutateRowsRequest) GetRegion() *RegionSpecifier {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *MutateRowsRequest) GetCondition() []*Condition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type MutateRowsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processed *bool `protobuf:"varint,1,opt,name=processed" json:"processed,omitempty"`
}

func (x *MutateRowsResponse) Reset() {
	*x = MutateRowsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultiRowMutation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateRowsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateRowsResponse) ProtoMessage() {}

func (x *MutateRowsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MultiRowMutation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateRowsResponse.ProtoReflect.Descriptor instead.
func (*MutateRowsResponse) Descriptor() ([]byte, []int) {
	return file_MultiRowMutation_proto_rawDescGZIP(), []int{3}
}

func (x *MutateRowsResponse) GetProcessed() bool {
	if x != nil && x.Processed != nil {
		return *x.Processed
	}
	return false
}

var File_MultiRowMutation_proto protoreflect.FileDescriptor

var file_MultiRowMutation_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x6f, 0x77, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x48, 0x42, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x20, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x52, 0x6f, 0x77, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x21, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x52, 0x6f, 0x77, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xe2, 0x01, 0x0a, 0x11, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x10, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x0f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x12, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x32, 0x56, 0x0a, 0x17, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x52, 0x6f, 0x77, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x77, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x53, 0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61,
	0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x16, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x6f, 0x77,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01,
	0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_MultiRowMutation_proto_rawDescOnce sync.Once
	file_MultiRowMutation_proto_rawDescData = file_MultiRowMutation_proto_rawDesc
)

func file_MultiRowMutation_proto_rawDescGZIP() []byte {
	file_MultiRowMutation_proto_rawDescOnce.Do(func() {
		file_MultiRowMutation_proto_rawDescData = protoimpl.X.CompressGZIP(file_MultiRowMutation_proto_rawDescData)
	})
	return file_MultiRowMutation_proto_rawDescData
}

var file_MultiRowMutation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_MultiRowMutation_proto_goTypes = []interface{}{
	(*MultiRowMutationProcessorRequest)(nil),  // 0: pb.MultiRowMutationProcessorRequest
	(*MultiRowMutationProcessorResponse)(nil), // 1: pb.MultiRowMutationProcessorResponse
	(*MutateRowsRequest)(nil),                 // 2: pb.MutateRowsRequest
	(*MutateRowsResponse)(nil),                // 3: pb.MutateRowsResponse
	(*MutationProto)(nil),                     // 4: pb.MutationProto
	(*RegionSpecifier)(nil),                   // 5: pb.RegionSpecifier
	(*Condition)(nil),                         // 6: pb.Condition
}
var file_MultiRowMutation_proto_depIdxs = []int32{
	4, // 0: pb.MutateRowsRequest.mutation_request:type_name -> pb.MutationProto
	5, // 1: pb.MutateRowsRequest.region:type_name -> pb.RegionSpecifier
	6, // 2: pb.MutateRowsRequest.condition:type_name -> pb.Condition
	2, // 3: pb.MultiRowMutationService.MutateRows:input_type -> pb.MutateRowsRequest
	3, // 4: pb.MultiRowMutationService.MutateRows:output_type -> pb.MutateRowsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_MultiRowMutation_proto_init() }
func file_MultiRowMutation_proto_init() {
	if File_MultiRowMutation_proto != nil {
		return
	}
	file_Client_proto_init()
	file_HBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MultiRowMutation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiRowMutationProcessorRequest); i {
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
		file_MultiRowMutation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiRowMutationProcessorResponse); i {
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
		file_MultiRowMutation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateRowsRequest); i {
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
		file_MultiRowMutation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateRowsResponse); i {
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
			RawDescriptor: file_MultiRowMutation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MultiRowMutation_proto_goTypes,
		DependencyIndexes: file_MultiRowMutation_proto_depIdxs,
		MessageInfos:      file_MultiRowMutation_proto_msgTypes,
	}.Build()
	File_MultiRowMutation_proto = out.File
	file_MultiRowMutation_proto_rawDesc = nil
	file_MultiRowMutation_proto_goTypes = nil
	file_MultiRowMutation_proto_depIdxs = nil
}
