// Copyright 2010-2024 Google LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Protocol buffer used to parametrize an iterated local search (ILS) approach.
// ILS is an iterative metaheuristic in which every iteration consists in
// performing a perturbation followed by an improvement step on a reference
// solution to generate a neighbor solution.
// The neighbor solution is accepted as the new reference solution according
// to an acceptance criterion.
// The best found solution is eventually returned.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: ortools/constraint_solver/routing_ils.proto

package constraintsolver

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

type RuinStrategy_Value int32

const (
	// Unspecified value.
	RuinStrategy_UNSET RuinStrategy_Value = 0
	// Removes a number of spatially close routes.
	RuinStrategy_SPATIALLY_CLOSE_ROUTES_REMOVAL RuinStrategy_Value = 1
)

// Enum value maps for RuinStrategy_Value.
var (
	RuinStrategy_Value_name = map[int32]string{
		0: "UNSET",
		1: "SPATIALLY_CLOSE_ROUTES_REMOVAL",
	}
	RuinStrategy_Value_value = map[string]int32{
		"UNSET":                          0,
		"SPATIALLY_CLOSE_ROUTES_REMOVAL": 1,
	}
)

func (x RuinStrategy_Value) Enum() *RuinStrategy_Value {
	p := new(RuinStrategy_Value)
	*p = x
	return p
}

func (x RuinStrategy_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RuinStrategy_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_ortools_constraint_solver_routing_ils_proto_enumTypes[0].Descriptor()
}

func (RuinStrategy_Value) Type() protoreflect.EnumType {
	return &file_ortools_constraint_solver_routing_ils_proto_enumTypes[0]
}

func (x RuinStrategy_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RuinStrategy_Value.Descriptor instead.
func (RuinStrategy_Value) EnumDescriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{0, 0}
}

type PerturbationStrategy_Value int32

const (
	// Unspecified value.
	PerturbationStrategy_UNSET PerturbationStrategy_Value = 0
	// Performs a perturbation in a ruin and recreate fashion.
	PerturbationStrategy_RUIN_AND_RECREATE PerturbationStrategy_Value = 1
)

// Enum value maps for PerturbationStrategy_Value.
var (
	PerturbationStrategy_Value_name = map[int32]string{
		0: "UNSET",
		1: "RUIN_AND_RECREATE",
	}
	PerturbationStrategy_Value_value = map[string]int32{
		"UNSET":             0,
		"RUIN_AND_RECREATE": 1,
	}
)

func (x PerturbationStrategy_Value) Enum() *PerturbationStrategy_Value {
	p := new(PerturbationStrategy_Value)
	*p = x
	return p
}

func (x PerturbationStrategy_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PerturbationStrategy_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_ortools_constraint_solver_routing_ils_proto_enumTypes[1].Descriptor()
}

func (PerturbationStrategy_Value) Type() protoreflect.EnumType {
	return &file_ortools_constraint_solver_routing_ils_proto_enumTypes[1]
}

func (x PerturbationStrategy_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PerturbationStrategy_Value.Descriptor instead.
func (PerturbationStrategy_Value) EnumDescriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{2, 0}
}

type AcceptanceStrategy_Value int32

const (
	// Unspecified value.
	AcceptanceStrategy_UNSET AcceptanceStrategy_Value = 0
	// Accept only solutions that are improving with respect to the reference
	// one.
	AcceptanceStrategy_GREEDY_DESCENT AcceptanceStrategy_Value = 1
)

// Enum value maps for AcceptanceStrategy_Value.
var (
	AcceptanceStrategy_Value_name = map[int32]string{
		0: "UNSET",
		1: "GREEDY_DESCENT",
	}
	AcceptanceStrategy_Value_value = map[string]int32{
		"UNSET":          0,
		"GREEDY_DESCENT": 1,
	}
)

func (x AcceptanceStrategy_Value) Enum() *AcceptanceStrategy_Value {
	p := new(AcceptanceStrategy_Value)
	*p = x
	return p
}

func (x AcceptanceStrategy_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcceptanceStrategy_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_ortools_constraint_solver_routing_ils_proto_enumTypes[2].Descriptor()
}

func (AcceptanceStrategy_Value) Type() protoreflect.EnumType {
	return &file_ortools_constraint_solver_routing_ils_proto_enumTypes[2]
}

func (x AcceptanceStrategy_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcceptanceStrategy_Value.Descriptor instead.
func (AcceptanceStrategy_Value) EnumDescriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{3, 0}
}

// Ruin strategies, used in perturbation based on ruin and recreate approaches.
type RuinStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RuinStrategy) Reset() {
	*x = RuinStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuinStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuinStrategy) ProtoMessage() {}

func (x *RuinStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuinStrategy.ProtoReflect.Descriptor instead.
func (*RuinStrategy) Descriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{0}
}

// Parameters to configure a perturbation based on a ruin and recreate approach.
type RuinRecreateParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Strategy defining how a reference solution is ruined.
	RuinStrategy RuinStrategy_Value `protobuf:"varint,1,opt,name=ruin_strategy,json=ruinStrategy,proto3,enum=operations_research.RuinStrategy_Value" json:"ruin_strategy,omitempty"`
	// Strategy defining how a reference solution is recreated.
	RecreateStrategy FirstSolutionStrategy_Value `protobuf:"varint,2,opt,name=recreate_strategy,json=recreateStrategy,proto3,enum=operations_research.FirstSolutionStrategy_Value" json:"recreate_strategy,omitempty"`
	// Number of routes removed during a ruin application defined on routes.
	NumRuinedRoutes uint32 `protobuf:"varint,3,opt,name=num_ruined_routes,json=numRuinedRoutes,proto3" json:"num_ruined_routes,omitempty"`
}

func (x *RuinRecreateParameters) Reset() {
	*x = RuinRecreateParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuinRecreateParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuinRecreateParameters) ProtoMessage() {}

func (x *RuinRecreateParameters) ProtoReflect() protoreflect.Message {
	mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuinRecreateParameters.ProtoReflect.Descriptor instead.
func (*RuinRecreateParameters) Descriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{1}
}

func (x *RuinRecreateParameters) GetRuinStrategy() RuinStrategy_Value {
	if x != nil {
		return x.RuinStrategy
	}
	return RuinStrategy_UNSET
}

func (x *RuinRecreateParameters) GetRecreateStrategy() FirstSolutionStrategy_Value {
	if x != nil {
		return x.RecreateStrategy
	}
	return FirstSolutionStrategy_UNSET
}

func (x *RuinRecreateParameters) GetNumRuinedRoutes() uint32 {
	if x != nil {
		return x.NumRuinedRoutes
	}
	return 0
}

// Defines how a reference solution is perturbed.
type PerturbationStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PerturbationStrategy) Reset() {
	*x = PerturbationStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerturbationStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerturbationStrategy) ProtoMessage() {}

func (x *PerturbationStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerturbationStrategy.ProtoReflect.Descriptor instead.
func (*PerturbationStrategy) Descriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{2}
}

// Determines when a neighbor solution, obtained by the application of a
// perturbation and improvement step to a reference solution, is used to
// replace the reference solution.
type AcceptanceStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcceptanceStrategy) Reset() {
	*x = AcceptanceStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptanceStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptanceStrategy) ProtoMessage() {}

func (x *AcceptanceStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptanceStrategy.ProtoReflect.Descriptor instead.
func (*AcceptanceStrategy) Descriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{3}
}

// Specifies the behavior of a search based on ILS.
type IteratedLocalSearchParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Determines how a reference solution S is perturbed to obtain a neighbor
	// solution S'.
	PerturbationStrategy PerturbationStrategy_Value `protobuf:"varint,1,opt,name=perturbation_strategy,json=perturbationStrategy,proto3,enum=operations_research.PerturbationStrategy_Value" json:"perturbation_strategy,omitempty"`
	// Parameters to customize a ruin and recreate perturbation.
	RuinRecreateParameters *RuinRecreateParameters `protobuf:"bytes,2,opt,name=ruin_recreate_parameters,json=ruinRecreateParameters,proto3" json:"ruin_recreate_parameters,omitempty"`
	// Determines whether solution S', obtained from the perturbation, should be
	// optimized with a local search application.
	ImprovePerturbedSolution bool `protobuf:"varint,3,opt,name=improve_perturbed_solution,json=improvePerturbedSolution,proto3" json:"improve_perturbed_solution,omitempty"`
	// Determines when the neighbor solution S', possibly improved if
	// `improve_perturbed_solution` is true, replaces the reference solution S.
	AcceptanceStrategy AcceptanceStrategy_Value `protobuf:"varint,4,opt,name=acceptance_strategy,json=acceptanceStrategy,proto3,enum=operations_research.AcceptanceStrategy_Value" json:"acceptance_strategy,omitempty"`
}

func (x *IteratedLocalSearchParameters) Reset() {
	*x = IteratedLocalSearchParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IteratedLocalSearchParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IteratedLocalSearchParameters) ProtoMessage() {}

func (x *IteratedLocalSearchParameters) ProtoReflect() protoreflect.Message {
	mi := &file_ortools_constraint_solver_routing_ils_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IteratedLocalSearchParameters.ProtoReflect.Descriptor instead.
func (*IteratedLocalSearchParameters) Descriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP(), []int{4}
}

func (x *IteratedLocalSearchParameters) GetPerturbationStrategy() PerturbationStrategy_Value {
	if x != nil {
		return x.PerturbationStrategy
	}
	return PerturbationStrategy_UNSET
}

func (x *IteratedLocalSearchParameters) GetRuinRecreateParameters() *RuinRecreateParameters {
	if x != nil {
		return x.RuinRecreateParameters
	}
	return nil
}

func (x *IteratedLocalSearchParameters) GetImprovePerturbedSolution() bool {
	if x != nil {
		return x.ImprovePerturbedSolution
	}
	return false
}

func (x *IteratedLocalSearchParameters) GetAcceptanceStrategy() AcceptanceStrategy_Value {
	if x != nil {
		return x.AcceptanceStrategy
	}
	return AcceptanceStrategy_UNSET
}

var File_ortools_constraint_solver_routing_ils_proto protoreflect.FileDescriptor

var file_ortools_constraint_solver_routing_ils_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6f, 0x72, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x1a, 0x2d, 0x6f, 0x72, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x52, 0x75, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x22, 0x36, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e,
	0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x50, 0x41, 0x54, 0x49, 0x41, 0x4c,
	0x4c, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x53, 0x5f,
	0x52, 0x45, 0x4d, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x01, 0x22, 0xf1, 0x01, 0x0a, 0x16, 0x52, 0x75,
	0x69, 0x6e, 0x52, 0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x72, 0x75, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x52, 0x75, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x72, 0x75, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x12, 0x5d, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x10, 0x72, 0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x75, 0x69, 0x6e, 0x65, 0x64, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75,
	0x6d, 0x52, 0x75, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x41, 0x0a,
	0x14, 0x50, 0x65, 0x72, 0x74, 0x75, 0x72, 0x62, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x29, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x55, 0x49,
	0x4e, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01,
	0x22, 0x3c, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x26, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x47, 0x52,
	0x45, 0x45, 0x44, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x22, 0x8a,
	0x03, 0x0a, 0x1d, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x64, 0x0a, 0x15, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x62, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x74, 0x75, 0x72, 0x62, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x14, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x62, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x65, 0x0a, 0x18, 0x72, 0x75, 0x69, 0x6e, 0x5f, 0x72,
	0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x52,
	0x75, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x16, 0x72, 0x75, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a,
	0x1a, 0x69, 0x6d, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x62,
	0x65, 0x64, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x18, 0x69, 0x6d, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x62, 0x65, 0x64, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x13, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x42, 0x46, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x6f,
	0x72, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ortools_constraint_solver_routing_ils_proto_rawDescOnce sync.Once
	file_ortools_constraint_solver_routing_ils_proto_rawDescData = file_ortools_constraint_solver_routing_ils_proto_rawDesc
)

func file_ortools_constraint_solver_routing_ils_proto_rawDescGZIP() []byte {
	file_ortools_constraint_solver_routing_ils_proto_rawDescOnce.Do(func() {
		file_ortools_constraint_solver_routing_ils_proto_rawDescData = protoimpl.X.CompressGZIP(file_ortools_constraint_solver_routing_ils_proto_rawDescData)
	})
	return file_ortools_constraint_solver_routing_ils_proto_rawDescData
}

var file_ortools_constraint_solver_routing_ils_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_ortools_constraint_solver_routing_ils_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ortools_constraint_solver_routing_ils_proto_goTypes = []interface{}{
	(RuinStrategy_Value)(0),               // 0: operations_research.RuinStrategy.Value
	(PerturbationStrategy_Value)(0),       // 1: operations_research.PerturbationStrategy.Value
	(AcceptanceStrategy_Value)(0),         // 2: operations_research.AcceptanceStrategy.Value
	(*RuinStrategy)(nil),                  // 3: operations_research.RuinStrategy
	(*RuinRecreateParameters)(nil),        // 4: operations_research.RuinRecreateParameters
	(*PerturbationStrategy)(nil),          // 5: operations_research.PerturbationStrategy
	(*AcceptanceStrategy)(nil),            // 6: operations_research.AcceptanceStrategy
	(*IteratedLocalSearchParameters)(nil), // 7: operations_research.IteratedLocalSearchParameters
	(FirstSolutionStrategy_Value)(0),      // 8: operations_research.FirstSolutionStrategy.Value
}
var file_ortools_constraint_solver_routing_ils_proto_depIdxs = []int32{
	0, // 0: operations_research.RuinRecreateParameters.ruin_strategy:type_name -> operations_research.RuinStrategy.Value
	8, // 1: operations_research.RuinRecreateParameters.recreate_strategy:type_name -> operations_research.FirstSolutionStrategy.Value
	1, // 2: operations_research.IteratedLocalSearchParameters.perturbation_strategy:type_name -> operations_research.PerturbationStrategy.Value
	4, // 3: operations_research.IteratedLocalSearchParameters.ruin_recreate_parameters:type_name -> operations_research.RuinRecreateParameters
	2, // 4: operations_research.IteratedLocalSearchParameters.acceptance_strategy:type_name -> operations_research.AcceptanceStrategy.Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ortools_constraint_solver_routing_ils_proto_init() }
func file_ortools_constraint_solver_routing_ils_proto_init() {
	if File_ortools_constraint_solver_routing_ils_proto != nil {
		return
	}
	file_ortools_constraint_solver_routing_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ortools_constraint_solver_routing_ils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuinStrategy); i {
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
		file_ortools_constraint_solver_routing_ils_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuinRecreateParameters); i {
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
		file_ortools_constraint_solver_routing_ils_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerturbationStrategy); i {
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
		file_ortools_constraint_solver_routing_ils_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptanceStrategy); i {
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
		file_ortools_constraint_solver_routing_ils_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IteratedLocalSearchParameters); i {
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
			RawDescriptor: file_ortools_constraint_solver_routing_ils_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ortools_constraint_solver_routing_ils_proto_goTypes,
		DependencyIndexes: file_ortools_constraint_solver_routing_ils_proto_depIdxs,
		EnumInfos:         file_ortools_constraint_solver_routing_ils_proto_enumTypes,
		MessageInfos:      file_ortools_constraint_solver_routing_ils_proto_msgTypes,
	}.Build()
	File_ortools_constraint_solver_routing_ils_proto = out.File
	file_ortools_constraint_solver_routing_ils_proto_rawDesc = nil
	file_ortools_constraint_solver_routing_ils_proto_goTypes = nil
	file_ortools_constraint_solver_routing_ils_proto_depIdxs = nil
}
