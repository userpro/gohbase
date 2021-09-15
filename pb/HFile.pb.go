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
// source: HFile.proto

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

type CompactionEventTracker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompactedStoreFile [][]byte `protobuf:"bytes,1,rep,name=compacted_store_file,json=compactedStoreFile" json:"compacted_store_file,omitempty"`
}

func (x *CompactionEventTracker) Reset() {
	*x = CompactionEventTracker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HFile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionEventTracker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionEventTracker) ProtoMessage() {}

func (x *CompactionEventTracker) ProtoReflect() protoreflect.Message {
	mi := &file_HFile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionEventTracker.ProtoReflect.Descriptor instead.
func (*CompactionEventTracker) Descriptor() ([]byte, []int) {
	return file_HFile_proto_rawDescGZIP(), []int{0}
}

func (x *CompactionEventTracker) GetCompactedStoreFile() [][]byte {
	if x != nil {
		return x.CompactedStoreFile
	}
	return nil
}

// Map of name/values
type FileInfoProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapEntry []*BytesBytesPair `protobuf:"bytes,1,rep,name=map_entry,json=mapEntry" json:"map_entry,omitempty"`
}

func (x *FileInfoProto) Reset() {
	*x = FileInfoProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HFile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoProto) ProtoMessage() {}

func (x *FileInfoProto) ProtoReflect() protoreflect.Message {
	mi := &file_HFile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoProto.ProtoReflect.Descriptor instead.
func (*FileInfoProto) Descriptor() ([]byte, []int) {
	return file_HFile_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfoProto) GetMapEntry() []*BytesBytesPair {
	if x != nil {
		return x.MapEntry
	}
	return nil
}

// HFile file trailer
type FileTrailerProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileInfoOffset            *uint64 `protobuf:"varint,1,opt,name=file_info_offset,json=fileInfoOffset" json:"file_info_offset,omitempty"`
	LoadOnOpenDataOffset      *uint64 `protobuf:"varint,2,opt,name=load_on_open_data_offset,json=loadOnOpenDataOffset" json:"load_on_open_data_offset,omitempty"`
	UncompressedDataIndexSize *uint64 `protobuf:"varint,3,opt,name=uncompressed_data_index_size,json=uncompressedDataIndexSize" json:"uncompressed_data_index_size,omitempty"`
	TotalUncompressedBytes    *uint64 `protobuf:"varint,4,opt,name=total_uncompressed_bytes,json=totalUncompressedBytes" json:"total_uncompressed_bytes,omitempty"`
	DataIndexCount            *uint32 `protobuf:"varint,5,opt,name=data_index_count,json=dataIndexCount" json:"data_index_count,omitempty"`
	MetaIndexCount            *uint32 `protobuf:"varint,6,opt,name=meta_index_count,json=metaIndexCount" json:"meta_index_count,omitempty"`
	EntryCount                *uint64 `protobuf:"varint,7,opt,name=entry_count,json=entryCount" json:"entry_count,omitempty"`
	NumDataIndexLevels        *uint32 `protobuf:"varint,8,opt,name=num_data_index_levels,json=numDataIndexLevels" json:"num_data_index_levels,omitempty"`
	FirstDataBlockOffset      *uint64 `protobuf:"varint,9,opt,name=first_data_block_offset,json=firstDataBlockOffset" json:"first_data_block_offset,omitempty"`
	LastDataBlockOffset       *uint64 `protobuf:"varint,10,opt,name=last_data_block_offset,json=lastDataBlockOffset" json:"last_data_block_offset,omitempty"`
	ComparatorClassName       *string `protobuf:"bytes,11,opt,name=comparator_class_name,json=comparatorClassName" json:"comparator_class_name,omitempty"`
	CompressionCodec          *uint32 `protobuf:"varint,12,opt,name=compression_codec,json=compressionCodec" json:"compression_codec,omitempty"`
	EncryptionKey             []byte  `protobuf:"bytes,13,opt,name=encryption_key,json=encryptionKey" json:"encryption_key,omitempty"`
}

func (x *FileTrailerProto) Reset() {
	*x = FileTrailerProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HFile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileTrailerProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTrailerProto) ProtoMessage() {}

func (x *FileTrailerProto) ProtoReflect() protoreflect.Message {
	mi := &file_HFile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTrailerProto.ProtoReflect.Descriptor instead.
func (*FileTrailerProto) Descriptor() ([]byte, []int) {
	return file_HFile_proto_rawDescGZIP(), []int{2}
}

func (x *FileTrailerProto) GetFileInfoOffset() uint64 {
	if x != nil && x.FileInfoOffset != nil {
		return *x.FileInfoOffset
	}
	return 0
}

func (x *FileTrailerProto) GetLoadOnOpenDataOffset() uint64 {
	if x != nil && x.LoadOnOpenDataOffset != nil {
		return *x.LoadOnOpenDataOffset
	}
	return 0
}

func (x *FileTrailerProto) GetUncompressedDataIndexSize() uint64 {
	if x != nil && x.UncompressedDataIndexSize != nil {
		return *x.UncompressedDataIndexSize
	}
	return 0
}

func (x *FileTrailerProto) GetTotalUncompressedBytes() uint64 {
	if x != nil && x.TotalUncompressedBytes != nil {
		return *x.TotalUncompressedBytes
	}
	return 0
}

func (x *FileTrailerProto) GetDataIndexCount() uint32 {
	if x != nil && x.DataIndexCount != nil {
		return *x.DataIndexCount
	}
	return 0
}

func (x *FileTrailerProto) GetMetaIndexCount() uint32 {
	if x != nil && x.MetaIndexCount != nil {
		return *x.MetaIndexCount
	}
	return 0
}

func (x *FileTrailerProto) GetEntryCount() uint64 {
	if x != nil && x.EntryCount != nil {
		return *x.EntryCount
	}
	return 0
}

func (x *FileTrailerProto) GetNumDataIndexLevels() uint32 {
	if x != nil && x.NumDataIndexLevels != nil {
		return *x.NumDataIndexLevels
	}
	return 0
}

func (x *FileTrailerProto) GetFirstDataBlockOffset() uint64 {
	if x != nil && x.FirstDataBlockOffset != nil {
		return *x.FirstDataBlockOffset
	}
	return 0
}

func (x *FileTrailerProto) GetLastDataBlockOffset() uint64 {
	if x != nil && x.LastDataBlockOffset != nil {
		return *x.LastDataBlockOffset
	}
	return 0
}

func (x *FileTrailerProto) GetComparatorClassName() string {
	if x != nil && x.ComparatorClassName != nil {
		return *x.ComparatorClassName
	}
	return ""
}

func (x *FileTrailerProto) GetCompressionCodec() uint32 {
	if x != nil && x.CompressionCodec != nil {
		return *x.CompressionCodec
	}
	return 0
}

func (x *FileTrailerProto) GetEncryptionKey() []byte {
	if x != nil {
		return x.EncryptionKey
	}
	return nil
}

var File_HFile_proto protoreflect.FileDescriptor

var file_HFile_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x48, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x0b, 0x48, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x0d, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x09, 0x6d,
	0x61, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x8b, 0x05, 0x0a,
	0x10, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x18, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x6c,
	0x6f, 0x61, 0x64, 0x4f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x1c, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x19, 0x75, 0x6e, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x6e,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x6e, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x66, 0x69, 0x72, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x33, 0x0a,
	0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x42, 0x48, 0x0a, 0x31, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x0b, 0x48, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x88, 0x01,
	0x01, 0xa0, 0x01, 0x01,
}

var (
	file_HFile_proto_rawDescOnce sync.Once
	file_HFile_proto_rawDescData = file_HFile_proto_rawDesc
)

func file_HFile_proto_rawDescGZIP() []byte {
	file_HFile_proto_rawDescOnce.Do(func() {
		file_HFile_proto_rawDescData = protoimpl.X.CompressGZIP(file_HFile_proto_rawDescData)
	})
	return file_HFile_proto_rawDescData
}

var file_HFile_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_HFile_proto_goTypes = []interface{}{
	(*CompactionEventTracker)(nil), // 0: pb.CompactionEventTracker
	(*FileInfoProto)(nil),          // 1: pb.FileInfoProto
	(*FileTrailerProto)(nil),       // 2: pb.FileTrailerProto
	(*BytesBytesPair)(nil),         // 3: pb.BytesBytesPair
}
var file_HFile_proto_depIdxs = []int32{
	3, // 0: pb.FileInfoProto.map_entry:type_name -> pb.BytesBytesPair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HFile_proto_init() }
func file_HFile_proto_init() {
	if File_HFile_proto != nil {
		return
	}
	file_HBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HFile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactionEventTracker); i {
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
		file_HFile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfoProto); i {
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
		file_HFile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileTrailerProto); i {
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
			RawDescriptor: file_HFile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HFile_proto_goTypes,
		DependencyIndexes: file_HFile_proto_depIdxs,
		MessageInfos:      file_HFile_proto_msgTypes,
	}.Build()
	File_HFile_proto = out.File
	file_HFile_proto_rawDesc = nil
	file_HFile_proto_goTypes = nil
	file_HFile_proto_depIdxs = nil
}
