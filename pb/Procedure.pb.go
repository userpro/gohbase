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
// source: Procedure.proto

package pb

import (
	any "github.com/golang/protobuf/ptypes/any"
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

type ProcedureState int32

const (
	ProcedureState_INITIALIZING    ProcedureState = 1 // Procedure in construction, not yet added to the executor
	ProcedureState_RUNNABLE        ProcedureState = 2 // Procedure added to the executor, and ready to be executed
	ProcedureState_WAITING         ProcedureState = 3 // The procedure is waiting on children to be completed
	ProcedureState_WAITING_TIMEOUT ProcedureState = 4 // The procedure is waiting a timout or an external event
	ProcedureState_ROLLEDBACK      ProcedureState = 5 // The procedure failed and was rolledback
	ProcedureState_SUCCESS         ProcedureState = 6 // The procedure execution is completed successfully.
	ProcedureState_FAILED          ProcedureState = 7 // The procedure execution is failed, may need to rollback
)

// Enum value maps for ProcedureState.
var (
	ProcedureState_name = map[int32]string{
		1: "INITIALIZING",
		2: "RUNNABLE",
		3: "WAITING",
		4: "WAITING_TIMEOUT",
		5: "ROLLEDBACK",
		6: "SUCCESS",
		7: "FAILED",
	}
	ProcedureState_value = map[string]int32{
		"INITIALIZING":    1,
		"RUNNABLE":        2,
		"WAITING":         3,
		"WAITING_TIMEOUT": 4,
		"ROLLEDBACK":      5,
		"SUCCESS":         6,
		"FAILED":          7,
	}
)

func (x ProcedureState) Enum() *ProcedureState {
	p := new(ProcedureState)
	*p = x
	return p
}

func (x ProcedureState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcedureState) Descriptor() protoreflect.EnumDescriptor {
	return file_Procedure_proto_enumTypes[0].Descriptor()
}

func (ProcedureState) Type() protoreflect.EnumType {
	return &file_Procedure_proto_enumTypes[0]
}

func (x ProcedureState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProcedureState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProcedureState(num)
	return nil
}

// Deprecated: Use ProcedureState.Descriptor instead.
func (ProcedureState) EnumDescriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{0}
}

type ProcedureWALEntry_Type int32

const (
	ProcedureWALEntry_PROCEDURE_WAL_EOF     ProcedureWALEntry_Type = 1
	ProcedureWALEntry_PROCEDURE_WAL_INIT    ProcedureWALEntry_Type = 2
	ProcedureWALEntry_PROCEDURE_WAL_INSERT  ProcedureWALEntry_Type = 3
	ProcedureWALEntry_PROCEDURE_WAL_UPDATE  ProcedureWALEntry_Type = 4
	ProcedureWALEntry_PROCEDURE_WAL_DELETE  ProcedureWALEntry_Type = 5
	ProcedureWALEntry_PROCEDURE_WAL_COMPACT ProcedureWALEntry_Type = 6
)

// Enum value maps for ProcedureWALEntry_Type.
var (
	ProcedureWALEntry_Type_name = map[int32]string{
		1: "PROCEDURE_WAL_EOF",
		2: "PROCEDURE_WAL_INIT",
		3: "PROCEDURE_WAL_INSERT",
		4: "PROCEDURE_WAL_UPDATE",
		5: "PROCEDURE_WAL_DELETE",
		6: "PROCEDURE_WAL_COMPACT",
	}
	ProcedureWALEntry_Type_value = map[string]int32{
		"PROCEDURE_WAL_EOF":     1,
		"PROCEDURE_WAL_INIT":    2,
		"PROCEDURE_WAL_INSERT":  3,
		"PROCEDURE_WAL_UPDATE":  4,
		"PROCEDURE_WAL_DELETE":  5,
		"PROCEDURE_WAL_COMPACT": 6,
	}
)

func (x ProcedureWALEntry_Type) Enum() *ProcedureWALEntry_Type {
	p := new(ProcedureWALEntry_Type)
	*p = x
	return p
}

func (x ProcedureWALEntry_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcedureWALEntry_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_Procedure_proto_enumTypes[1].Descriptor()
}

func (ProcedureWALEntry_Type) Type() protoreflect.EnumType {
	return &file_Procedure_proto_enumTypes[1]
}

func (x ProcedureWALEntry_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProcedureWALEntry_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProcedureWALEntry_Type(num)
	return nil
}

// Deprecated: Use ProcedureWALEntry_Type.Descriptor instead.
func (ProcedureWALEntry_Type) EnumDescriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{6, 0}
}

//*
// Procedure metadata, serialized by the ProcedureStore to be able to recover the old state.
type Procedure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal "static" state
	ClassName     *string `protobuf:"bytes,1,req,name=class_name,json=className" json:"class_name,omitempty"` // full classname to be able to instantiate the procedure
	ParentId      *uint64 `protobuf:"varint,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`   // parent if not a root-procedure otherwise not set
	ProcId        *uint64 `protobuf:"varint,3,req,name=proc_id,json=procId" json:"proc_id,omitempty"`
	SubmittedTime *uint64 `protobuf:"varint,4,req,name=submitted_time,json=submittedTime" json:"submitted_time,omitempty"`
	Owner         *string `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
	// internal "runtime" state
	State      *ProcedureState `protobuf:"varint,6,req,name=state,enum=pb.ProcedureState" json:"state,omitempty"`
	StackId    []uint32        `protobuf:"varint,7,rep,name=stack_id,json=stackId" json:"stack_id,omitempty"` // stack indices in case the procedure was running
	LastUpdate *uint64         `protobuf:"varint,8,req,name=last_update,json=lastUpdate" json:"last_update,omitempty"`
	Timeout    *uint32         `protobuf:"varint,9,opt,name=timeout" json:"timeout,omitempty"`
	// user state/results
	Exception    *ForeignExceptionMessage `protobuf:"bytes,10,opt,name=exception" json:"exception,omitempty"`
	Result       []byte                   `protobuf:"bytes,11,opt,name=result" json:"result,omitempty"`                                 // opaque (user) result structure
	StateData    []byte                   `protobuf:"bytes,12,opt,name=state_data,json=stateData" json:"state_data,omitempty"`          // opaque (user) procedure internal-state - OBSOLATE
	StateMessage []*any.Any               `protobuf:"bytes,15,rep,name=state_message,json=stateMessage" json:"state_message,omitempty"` // opaque (user) procedure internal-state
	// Nonce to prevent same procedure submit by multiple times
	NonceGroup *uint64 `protobuf:"varint,13,opt,name=nonce_group,json=nonceGroup,def=0" json:"nonce_group,omitempty"`
	Nonce      *uint64 `protobuf:"varint,14,opt,name=nonce,def=0" json:"nonce,omitempty"`
	// whether the procedure has held the lock
	Locked *bool `protobuf:"varint,16,opt,name=locked,def=0" json:"locked,omitempty"`
	// whether the procedure need to be bypassed
	Bypass *bool `protobuf:"varint,17,opt,name=bypass,def=0" json:"bypass,omitempty"`
}

// Default values for Procedure fields.
const (
	Default_Procedure_NonceGroup = uint64(0)
	Default_Procedure_Nonce      = uint64(0)
	Default_Procedure_Locked     = bool(false)
	Default_Procedure_Bypass     = bool(false)
)

func (x *Procedure) Reset() {
	*x = Procedure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Procedure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Procedure) ProtoMessage() {}

func (x *Procedure) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Procedure.ProtoReflect.Descriptor instead.
func (*Procedure) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{0}
}

func (x *Procedure) GetClassName() string {
	if x != nil && x.ClassName != nil {
		return *x.ClassName
	}
	return ""
}

func (x *Procedure) GetParentId() uint64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *Procedure) GetProcId() uint64 {
	if x != nil && x.ProcId != nil {
		return *x.ProcId
	}
	return 0
}

func (x *Procedure) GetSubmittedTime() uint64 {
	if x != nil && x.SubmittedTime != nil {
		return *x.SubmittedTime
	}
	return 0
}

func (x *Procedure) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

func (x *Procedure) GetState() ProcedureState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ProcedureState_INITIALIZING
}

func (x *Procedure) GetStackId() []uint32 {
	if x != nil {
		return x.StackId
	}
	return nil
}

func (x *Procedure) GetLastUpdate() uint64 {
	if x != nil && x.LastUpdate != nil {
		return *x.LastUpdate
	}
	return 0
}

func (x *Procedure) GetTimeout() uint32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *Procedure) GetException() *ForeignExceptionMessage {
	if x != nil {
		return x.Exception
	}
	return nil
}

func (x *Procedure) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Procedure) GetStateData() []byte {
	if x != nil {
		return x.StateData
	}
	return nil
}

func (x *Procedure) GetStateMessage() []*any.Any {
	if x != nil {
		return x.StateMessage
	}
	return nil
}

func (x *Procedure) GetNonceGroup() uint64 {
	if x != nil && x.NonceGroup != nil {
		return *x.NonceGroup
	}
	return Default_Procedure_NonceGroup
}

func (x *Procedure) GetNonce() uint64 {
	if x != nil && x.Nonce != nil {
		return *x.Nonce
	}
	return Default_Procedure_Nonce
}

func (x *Procedure) GetLocked() bool {
	if x != nil && x.Locked != nil {
		return *x.Locked
	}
	return Default_Procedure_Locked
}

func (x *Procedure) GetBypass() bool {
	if x != nil && x.Bypass != nil {
		return *x.Bypass
	}
	return Default_Procedure_Bypass
}

//*
// SequentialProcedure data
type SequentialProcedureData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Executed *bool `protobuf:"varint,1,req,name=executed" json:"executed,omitempty"`
}

func (x *SequentialProcedureData) Reset() {
	*x = SequentialProcedureData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SequentialProcedureData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequentialProcedureData) ProtoMessage() {}

func (x *SequentialProcedureData) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequentialProcedureData.ProtoReflect.Descriptor instead.
func (*SequentialProcedureData) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{1}
}

func (x *SequentialProcedureData) GetExecuted() bool {
	if x != nil && x.Executed != nil {
		return *x.Executed
	}
	return false
}

//*
// StateMachineProcedure data
type StateMachineProcedureData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State []uint32 `protobuf:"varint,1,rep,name=state" json:"state,omitempty"`
}

func (x *StateMachineProcedureData) Reset() {
	*x = StateMachineProcedureData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateMachineProcedureData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateMachineProcedureData) ProtoMessage() {}

func (x *StateMachineProcedureData) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateMachineProcedureData.ProtoReflect.Descriptor instead.
func (*StateMachineProcedureData) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{2}
}

func (x *StateMachineProcedureData) GetState() []uint32 {
	if x != nil {
		return x.State
	}
	return nil
}

//*
// Procedure WAL header
type ProcedureWALHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Type      *uint32 `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	LogId     *uint64 `protobuf:"varint,3,req,name=log_id,json=logId" json:"log_id,omitempty"`
	MinProcId *uint64 `protobuf:"varint,4,req,name=min_proc_id,json=minProcId" json:"min_proc_id,omitempty"`
}

func (x *ProcedureWALHeader) Reset() {
	*x = ProcedureWALHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureWALHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureWALHeader) ProtoMessage() {}

func (x *ProcedureWALHeader) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureWALHeader.ProtoReflect.Descriptor instead.
func (*ProcedureWALHeader) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{3}
}

func (x *ProcedureWALHeader) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *ProcedureWALHeader) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *ProcedureWALHeader) GetLogId() uint64 {
	if x != nil && x.LogId != nil {
		return *x.LogId
	}
	return 0
}

func (x *ProcedureWALHeader) GetMinProcId() uint64 {
	if x != nil && x.MinProcId != nil {
		return *x.MinProcId
	}
	return 0
}

//*
// Procedure WAL trailer
type ProcedureWALTrailer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	TrackerPos *uint64 `protobuf:"varint,2,req,name=tracker_pos,json=trackerPos" json:"tracker_pos,omitempty"`
}

func (x *ProcedureWALTrailer) Reset() {
	*x = ProcedureWALTrailer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureWALTrailer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureWALTrailer) ProtoMessage() {}

func (x *ProcedureWALTrailer) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureWALTrailer.ProtoReflect.Descriptor instead.
func (*ProcedureWALTrailer) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{4}
}

func (x *ProcedureWALTrailer) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *ProcedureWALTrailer) GetTrackerPos() uint64 {
	if x != nil && x.TrackerPos != nil {
		return *x.TrackerPos
	}
	return 0
}

type ProcedureStoreTracker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node []*ProcedureStoreTracker_TrackerNode `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
}

func (x *ProcedureStoreTracker) Reset() {
	*x = ProcedureStoreTracker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureStoreTracker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureStoreTracker) ProtoMessage() {}

func (x *ProcedureStoreTracker) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureStoreTracker.ProtoReflect.Descriptor instead.
func (*ProcedureStoreTracker) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{5}
}

func (x *ProcedureStoreTracker) GetNode() []*ProcedureStoreTracker_TrackerNode {
	if x != nil {
		return x.Node
	}
	return nil
}

type ProcedureWALEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      *ProcedureWALEntry_Type `protobuf:"varint,1,req,name=type,enum=pb.ProcedureWALEntry_Type" json:"type,omitempty"`
	Procedure []*Procedure            `protobuf:"bytes,2,rep,name=procedure" json:"procedure,omitempty"`
	ProcId    *uint64                 `protobuf:"varint,3,opt,name=proc_id,json=procId" json:"proc_id,omitempty"`
	ChildId   []uint64                `protobuf:"varint,4,rep,name=child_id,json=childId" json:"child_id,omitempty"`
}

func (x *ProcedureWALEntry) Reset() {
	*x = ProcedureWALEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureWALEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureWALEntry) ProtoMessage() {}

func (x *ProcedureWALEntry) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureWALEntry.ProtoReflect.Descriptor instead.
func (*ProcedureWALEntry) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{6}
}

func (x *ProcedureWALEntry) GetType() ProcedureWALEntry_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ProcedureWALEntry_PROCEDURE_WAL_EOF
}

func (x *ProcedureWALEntry) GetProcedure() []*Procedure {
	if x != nil {
		return x.Procedure
	}
	return nil
}

func (x *ProcedureWALEntry) GetProcId() uint64 {
	if x != nil && x.ProcId != nil {
		return *x.ProcId
	}
	return 0
}

func (x *ProcedureWALEntry) GetChildId() []uint64 {
	if x != nil {
		return x.ChildId
	}
	return nil
}

type ProcedureStoreTracker_TrackerNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartId *uint64  `protobuf:"varint,1,req,name=start_id,json=startId" json:"start_id,omitempty"`
	Updated []uint64 `protobuf:"varint,2,rep,name=updated" json:"updated,omitempty"`
	Deleted []uint64 `protobuf:"varint,3,rep,name=deleted" json:"deleted,omitempty"`
}

func (x *ProcedureStoreTracker_TrackerNode) Reset() {
	*x = ProcedureStoreTracker_TrackerNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Procedure_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureStoreTracker_TrackerNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureStoreTracker_TrackerNode) ProtoMessage() {}

func (x *ProcedureStoreTracker_TrackerNode) ProtoReflect() protoreflect.Message {
	mi := &file_Procedure_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureStoreTracker_TrackerNode.ProtoReflect.Descriptor instead.
func (*ProcedureStoreTracker_TrackerNode) Descriptor() ([]byte, []int) {
	return file_Procedure_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ProcedureStoreTracker_TrackerNode) GetStartId() uint64 {
	if x != nil && x.StartId != nil {
		return *x.StartId
	}
	return 0
}

func (x *ProcedureStoreTracker_TrackerNode) GetUpdated() []uint64 {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *ProcedureStoreTracker_TrackerNode) GetDeleted() []uint64 {
	if x != nil {
		return x.Deleted
	}
	return nil
}

var File_Procedure_proto protoreflect.FileDescriptor

var file_Procedure_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x72, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x04,
	0x52, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x0b, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x0a, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x3a,
	0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x06, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x22, 0x35, 0x0a,
	0x17, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x64, 0x22, 0x31, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x79, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x57, 0x41, 0x4c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6c,
	0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x49, 0x64, 0x22, 0x50, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x57,
	0x41, 0x4c, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75,
	0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x39,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x5c, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xc5, 0x02, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x57, 0x41, 0x4c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x57, 0x41, 0x4c, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x72, 0x6f,
	0x63, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x9e,
	0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x43, 0x45,
	0x44, 0x55, 0x52, 0x45, 0x5f, 0x57, 0x41, 0x4c, 0x5f, 0x45, 0x4f, 0x46, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x44, 0x55, 0x52, 0x45, 0x5f, 0x57, 0x41, 0x4c, 0x5f,
	0x49, 0x4e, 0x49, 0x54, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x44,
	0x55, 0x52, 0x45, 0x5f, 0x57, 0x41, 0x4c, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x44, 0x55, 0x52, 0x45, 0x5f, 0x57, 0x41,
	0x4c, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x44, 0x55, 0x52, 0x45, 0x5f, 0x57, 0x41, 0x4c, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x44, 0x55, 0x52,
	0x45, 0x5f, 0x57, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x10, 0x06, 0x2a,
	0x7b, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x55, 0x4e, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x13,
	0x0a, 0x0f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x42, 0x41, 0x43,
	0x4b, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x06,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x42, 0x4c, 0x0a, 0x31,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x48, 0x01, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_Procedure_proto_rawDescOnce sync.Once
	file_Procedure_proto_rawDescData = file_Procedure_proto_rawDesc
)

func file_Procedure_proto_rawDescGZIP() []byte {
	file_Procedure_proto_rawDescOnce.Do(func() {
		file_Procedure_proto_rawDescData = protoimpl.X.CompressGZIP(file_Procedure_proto_rawDescData)
	})
	return file_Procedure_proto_rawDescData
}

var file_Procedure_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Procedure_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_Procedure_proto_goTypes = []interface{}{
	(ProcedureState)(0),                       // 0: pb.ProcedureState
	(ProcedureWALEntry_Type)(0),               // 1: pb.ProcedureWALEntry.Type
	(*Procedure)(nil),                         // 2: pb.Procedure
	(*SequentialProcedureData)(nil),           // 3: pb.SequentialProcedureData
	(*StateMachineProcedureData)(nil),         // 4: pb.StateMachineProcedureData
	(*ProcedureWALHeader)(nil),                // 5: pb.ProcedureWALHeader
	(*ProcedureWALTrailer)(nil),               // 6: pb.ProcedureWALTrailer
	(*ProcedureStoreTracker)(nil),             // 7: pb.ProcedureStoreTracker
	(*ProcedureWALEntry)(nil),                 // 8: pb.ProcedureWALEntry
	(*ProcedureStoreTracker_TrackerNode)(nil), // 9: pb.ProcedureStoreTracker.TrackerNode
	(*ForeignExceptionMessage)(nil),           // 10: pb.ForeignExceptionMessage
	(*any.Any)(nil),                           // 11: google.protobuf.Any
}
var file_Procedure_proto_depIdxs = []int32{
	0,  // 0: pb.Procedure.state:type_name -> pb.ProcedureState
	10, // 1: pb.Procedure.exception:type_name -> pb.ForeignExceptionMessage
	11, // 2: pb.Procedure.state_message:type_name -> google.protobuf.Any
	9,  // 3: pb.ProcedureStoreTracker.node:type_name -> pb.ProcedureStoreTracker.TrackerNode
	1,  // 4: pb.ProcedureWALEntry.type:type_name -> pb.ProcedureWALEntry.Type
	2,  // 5: pb.ProcedureWALEntry.procedure:type_name -> pb.Procedure
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_Procedure_proto_init() }
func file_Procedure_proto_init() {
	if File_Procedure_proto != nil {
		return
	}
	file_ErrorHandling_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Procedure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Procedure); i {
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
		file_Procedure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SequentialProcedureData); i {
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
		file_Procedure_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateMachineProcedureData); i {
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
		file_Procedure_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureWALHeader); i {
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
		file_Procedure_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureWALTrailer); i {
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
		file_Procedure_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureStoreTracker); i {
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
		file_Procedure_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureWALEntry); i {
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
		file_Procedure_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureStoreTracker_TrackerNode); i {
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
			RawDescriptor: file_Procedure_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Procedure_proto_goTypes,
		DependencyIndexes: file_Procedure_proto_depIdxs,
		EnumInfos:         file_Procedure_proto_enumTypes,
		MessageInfos:      file_Procedure_proto_msgTypes,
	}.Build()
	File_Procedure_proto = out.File
	file_Procedure_proto_rawDesc = nil
	file_Procedure_proto_goTypes = nil
	file_Procedure_proto_depIdxs = nil
}
