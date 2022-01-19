// Copyright 2010-2021 Google LLC
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

// This file contains protocol buffers for all parameters of the CP solver.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: ortools/constraint_solver/solver_parameters.proto

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

//
// Internal parameters of the solver.
//
type ConstraintSolverParameters_TrailCompression int32

const (
	ConstraintSolverParameters_NO_COMPRESSION     ConstraintSolverParameters_TrailCompression = 0
	ConstraintSolverParameters_COMPRESS_WITH_ZLIB ConstraintSolverParameters_TrailCompression = 1
)

// Enum value maps for ConstraintSolverParameters_TrailCompression.
var (
	ConstraintSolverParameters_TrailCompression_name = map[int32]string{
		0: "NO_COMPRESSION",
		1: "COMPRESS_WITH_ZLIB",
	}
	ConstraintSolverParameters_TrailCompression_value = map[string]int32{
		"NO_COMPRESSION":     0,
		"COMPRESS_WITH_ZLIB": 1,
	}
)

func (x ConstraintSolverParameters_TrailCompression) Enum() *ConstraintSolverParameters_TrailCompression {
	p := new(ConstraintSolverParameters_TrailCompression)
	*p = x
	return p
}

func (x ConstraintSolverParameters_TrailCompression) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConstraintSolverParameters_TrailCompression) Descriptor() protoreflect.EnumDescriptor {
	return file_ortools_constraint_solver_solver_parameters_proto_enumTypes[0].Descriptor()
}

func (ConstraintSolverParameters_TrailCompression) Type() protoreflect.EnumType {
	return &file_ortools_constraint_solver_solver_parameters_proto_enumTypes[0]
}

func (x ConstraintSolverParameters_TrailCompression) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConstraintSolverParameters_TrailCompression.Descriptor instead.
func (ConstraintSolverParameters_TrailCompression) EnumDescriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_solver_parameters_proto_rawDescGZIP(), []int{0, 0}
}

// Solver parameters.
type ConstraintSolverParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This parameter indicates if the solver should compress the trail
	// during the search. No compression means that the solver will be faster,
	// but will use more memory.
	CompressTrail ConstraintSolverParameters_TrailCompression `protobuf:"varint,1,opt,name=compress_trail,json=compressTrail,proto3,enum=operations_research.ConstraintSolverParameters_TrailCompression" json:"compress_trail,omitempty"`
	// This parameter indicates the default size of a block of the trail.
	// Compression applies at the block level.
	TrailBlockSize int32 `protobuf:"varint,2,opt,name=trail_block_size,json=trailBlockSize,proto3" json:"trail_block_size,omitempty"`
	// When a sum/min/max operation is applied on a large array, this
	// array is recursively split into blocks of size 'array_split_size'.
	ArraySplitSize int32 `protobuf:"varint,3,opt,name=array_split_size,json=arraySplitSize,proto3" json:"array_split_size,omitempty"`
	// This parameters indicates if the solver should store the names of
	// the objets it manages.
	StoreNames bool `protobuf:"varint,4,opt,name=store_names,json=storeNames,proto3" json:"store_names,omitempty"`
	// Create names for cast variables.
	NameCastVariables bool `protobuf:"varint,5,opt,name=name_cast_variables,json=nameCastVariables,proto3" json:"name_cast_variables,omitempty"`
	// Should anonymous variables be given a name.
	NameAllVariables bool `protobuf:"varint,6,opt,name=name_all_variables,json=nameAllVariables,proto3" json:"name_all_variables,omitempty"`
	// Activate propagation profiling.
	ProfilePropagation bool `protobuf:"varint,7,opt,name=profile_propagation,json=profilePropagation,proto3" json:"profile_propagation,omitempty"`
	// Export propagation profiling data to file.
	ProfileFile string `protobuf:"bytes,8,opt,name=profile_file,json=profileFile,proto3" json:"profile_file,omitempty"`
	// Activate local search profiling.
	ProfileLocalSearch bool `protobuf:"varint,16,opt,name=profile_local_search,json=profileLocalSearch,proto3" json:"profile_local_search,omitempty"`
	// Print local search profiling data after solving.
	PrintLocalSearchProfile bool `protobuf:"varint,17,opt,name=print_local_search_profile,json=printLocalSearchProfile,proto3" json:"print_local_search_profile,omitempty"`
	// Activate propagate tracing.
	TracePropagation bool `protobuf:"varint,9,opt,name=trace_propagation,json=tracePropagation,proto3" json:"trace_propagation,omitempty"`
	// Trace search.
	TraceSearch bool `protobuf:"varint,10,opt,name=trace_search,json=traceSearch,proto3" json:"trace_search,omitempty"`
	// Print the model before solving.
	PrintModel bool `protobuf:"varint,11,opt,name=print_model,json=printModel,proto3" json:"print_model,omitempty"`
	// Print model statistics before solving.
	PrintModelStats bool `protobuf:"varint,12,opt,name=print_model_stats,json=printModelStats,proto3" json:"print_model_stats,omitempty"`
	// Print added constraints.
	PrintAddedConstraints bool `protobuf:"varint,13,opt,name=print_added_constraints,json=printAddedConstraints,proto3" json:"print_added_constraints,omitempty"`
	DisableSolve          bool `protobuf:"varint,15,opt,name=disable_solve,json=disableSolve,proto3" json:"disable_solve,omitempty"`
	//
	// Control the implementation of the table constraint.
	//
	UseSmallTable bool `protobuf:"varint,101,opt,name=use_small_table,json=useSmallTable,proto3" json:"use_small_table,omitempty"`
	//
	// Control the propagation of the cumulative constraint.
	//
	UseCumulativeEdgeFinder    bool  `protobuf:"varint,105,opt,name=use_cumulative_edge_finder,json=useCumulativeEdgeFinder,proto3" json:"use_cumulative_edge_finder,omitempty"`
	UseCumulativeTimeTable     bool  `protobuf:"varint,106,opt,name=use_cumulative_time_table,json=useCumulativeTimeTable,proto3" json:"use_cumulative_time_table,omitempty"`
	UseCumulativeTimeTableSync bool  `protobuf:"varint,112,opt,name=use_cumulative_time_table_sync,json=useCumulativeTimeTableSync,proto3" json:"use_cumulative_time_table_sync,omitempty"`
	UseSequenceHighDemandTasks bool  `protobuf:"varint,107,opt,name=use_sequence_high_demand_tasks,json=useSequenceHighDemandTasks,proto3" json:"use_sequence_high_demand_tasks,omitempty"`
	UseAllPossibleDisjunctions bool  `protobuf:"varint,108,opt,name=use_all_possible_disjunctions,json=useAllPossibleDisjunctions,proto3" json:"use_all_possible_disjunctions,omitempty"`
	MaxEdgeFinderSize          int32 `protobuf:"varint,109,opt,name=max_edge_finder_size,json=maxEdgeFinderSize,proto3" json:"max_edge_finder_size,omitempty"`
	//
	// Control the propagation of the diffn constraint.
	//
	DiffnUseCumulative bool `protobuf:"varint,110,opt,name=diffn_use_cumulative,json=diffnUseCumulative,proto3" json:"diffn_use_cumulative,omitempty"`
	//
	// Control the implementation of the element constraint.
	//
	UseElementRmq bool `protobuf:"varint,111,opt,name=use_element_rmq,json=useElementRmq,proto3" json:"use_element_rmq,omitempty"`
	//
	// Skip locally optimal pairs of paths in PathOperators. Setting this
	// parameter to true might skip valid neighbors if there are constraints
	// linking paths together (such as precedences). In any other case this
	// should only speed up the search without omitting any neighbors.
	//
	SkipLocallyOptimalPaths bool `protobuf:"varint,113,opt,name=skip_locally_optimal_paths,json=skipLocallyOptimalPaths,proto3" json:"skip_locally_optimal_paths,omitempty"`
	//
	// Control the behavior of local search.
	//
	CheckSolutionPeriod int32 `protobuf:"varint,114,opt,name=check_solution_period,json=checkSolutionPeriod,proto3" json:"check_solution_period,omitempty"`
}

func (x *ConstraintSolverParameters) Reset() {
	*x = ConstraintSolverParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ortools_constraint_solver_solver_parameters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstraintSolverParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstraintSolverParameters) ProtoMessage() {}

func (x *ConstraintSolverParameters) ProtoReflect() protoreflect.Message {
	mi := &file_ortools_constraint_solver_solver_parameters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstraintSolverParameters.ProtoReflect.Descriptor instead.
func (*ConstraintSolverParameters) Descriptor() ([]byte, []int) {
	return file_ortools_constraint_solver_solver_parameters_proto_rawDescGZIP(), []int{0}
}

func (x *ConstraintSolverParameters) GetCompressTrail() ConstraintSolverParameters_TrailCompression {
	if x != nil {
		return x.CompressTrail
	}
	return ConstraintSolverParameters_NO_COMPRESSION
}

func (x *ConstraintSolverParameters) GetTrailBlockSize() int32 {
	if x != nil {
		return x.TrailBlockSize
	}
	return 0
}

func (x *ConstraintSolverParameters) GetArraySplitSize() int32 {
	if x != nil {
		return x.ArraySplitSize
	}
	return 0
}

func (x *ConstraintSolverParameters) GetStoreNames() bool {
	if x != nil {
		return x.StoreNames
	}
	return false
}

func (x *ConstraintSolverParameters) GetNameCastVariables() bool {
	if x != nil {
		return x.NameCastVariables
	}
	return false
}

func (x *ConstraintSolverParameters) GetNameAllVariables() bool {
	if x != nil {
		return x.NameAllVariables
	}
	return false
}

func (x *ConstraintSolverParameters) GetProfilePropagation() bool {
	if x != nil {
		return x.ProfilePropagation
	}
	return false
}

func (x *ConstraintSolverParameters) GetProfileFile() string {
	if x != nil {
		return x.ProfileFile
	}
	return ""
}

func (x *ConstraintSolverParameters) GetProfileLocalSearch() bool {
	if x != nil {
		return x.ProfileLocalSearch
	}
	return false
}

func (x *ConstraintSolverParameters) GetPrintLocalSearchProfile() bool {
	if x != nil {
		return x.PrintLocalSearchProfile
	}
	return false
}

func (x *ConstraintSolverParameters) GetTracePropagation() bool {
	if x != nil {
		return x.TracePropagation
	}
	return false
}

func (x *ConstraintSolverParameters) GetTraceSearch() bool {
	if x != nil {
		return x.TraceSearch
	}
	return false
}

func (x *ConstraintSolverParameters) GetPrintModel() bool {
	if x != nil {
		return x.PrintModel
	}
	return false
}

func (x *ConstraintSolverParameters) GetPrintModelStats() bool {
	if x != nil {
		return x.PrintModelStats
	}
	return false
}

func (x *ConstraintSolverParameters) GetPrintAddedConstraints() bool {
	if x != nil {
		return x.PrintAddedConstraints
	}
	return false
}

func (x *ConstraintSolverParameters) GetDisableSolve() bool {
	if x != nil {
		return x.DisableSolve
	}
	return false
}

func (x *ConstraintSolverParameters) GetUseSmallTable() bool {
	if x != nil {
		return x.UseSmallTable
	}
	return false
}

func (x *ConstraintSolverParameters) GetUseCumulativeEdgeFinder() bool {
	if x != nil {
		return x.UseCumulativeEdgeFinder
	}
	return false
}

func (x *ConstraintSolverParameters) GetUseCumulativeTimeTable() bool {
	if x != nil {
		return x.UseCumulativeTimeTable
	}
	return false
}

func (x *ConstraintSolverParameters) GetUseCumulativeTimeTableSync() bool {
	if x != nil {
		return x.UseCumulativeTimeTableSync
	}
	return false
}

func (x *ConstraintSolverParameters) GetUseSequenceHighDemandTasks() bool {
	if x != nil {
		return x.UseSequenceHighDemandTasks
	}
	return false
}

func (x *ConstraintSolverParameters) GetUseAllPossibleDisjunctions() bool {
	if x != nil {
		return x.UseAllPossibleDisjunctions
	}
	return false
}

func (x *ConstraintSolverParameters) GetMaxEdgeFinderSize() int32 {
	if x != nil {
		return x.MaxEdgeFinderSize
	}
	return 0
}

func (x *ConstraintSolverParameters) GetDiffnUseCumulative() bool {
	if x != nil {
		return x.DiffnUseCumulative
	}
	return false
}

func (x *ConstraintSolverParameters) GetUseElementRmq() bool {
	if x != nil {
		return x.UseElementRmq
	}
	return false
}

func (x *ConstraintSolverParameters) GetSkipLocallyOptimalPaths() bool {
	if x != nil {
		return x.SkipLocallyOptimalPaths
	}
	return false
}

func (x *ConstraintSolverParameters) GetCheckSolutionPeriod() int32 {
	if x != nil {
		return x.CheckSolutionPeriod
	}
	return 0
}

var File_ortools_constraint_solver_solver_parameters_proto protoreflect.FileDescriptor

var file_ortools_constraint_solver_solver_parameters_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6f, 0x72, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0xd4, 0x0b, 0x0a, 0x1a, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x67, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x40, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x53, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6c,
	0x12, 0x28, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x69,
	0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x72, 0x72, 0x61, 0x79, 0x53, 0x70, 0x6c, 0x69, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x61,
	0x73, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x73, 0x74, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x6c, 0x6c, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x3b, 0x0a, 0x1a, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x10, 0x74, 0x72, 0x61, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65,
	0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x53, 0x6d,
	0x61, 0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x75, 0x73, 0x65, 0x5f,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x64, 0x67, 0x65, 0x5f,
	0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x69, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x75, 0x73,
	0x65, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x45, 0x64, 0x67, 0x65, 0x46,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x73, 0x65, 0x43, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x42, 0x0a, 0x1e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x18, 0x70, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x75, 0x73, 0x65, 0x43, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x42, 0x0a, 0x1e, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x75, 0x73,
	0x65, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x69, 0x67, 0x68, 0x44, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x75, 0x73, 0x65, 0x5f,
	0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x73,
	0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1a, 0x75, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x44,
	0x69, 0x73, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x6d,
	0x61, 0x78, 0x5f, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x45, 0x64,
	0x67, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x30, 0x0a, 0x14,
	0x64, 0x69, 0x66, 0x66, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x69, 0x66, 0x66,
	0x6e, 0x55, 0x73, 0x65, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6d,
	0x71, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x6d, 0x71, 0x12, 0x3b, 0x0a, 0x1a, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x18, 0x71, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x73, 0x6b, 0x69, 0x70,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x72, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x13, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x3e, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x69, 0x6c,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x4e,
	0x4f, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x5a, 0x4c, 0x49, 0x42, 0x10, 0x01, 0x4a, 0x04, 0x08, 0x64, 0x10, 0x65, 0x4a, 0x04, 0x08,
	0x66, 0x10, 0x67, 0x4a, 0x04, 0x08, 0x67, 0x10, 0x68, 0x4a, 0x04, 0x08, 0x68, 0x10, 0x69, 0x42,
	0x9b, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6f,
	0x72, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x73, 0x70, 0x61, 0x63, 0x65, 0x74, 0x65,
	0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x2d, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x6f, 0x72, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x6f, 0x72, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0xaa, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ortools_constraint_solver_solver_parameters_proto_rawDescOnce sync.Once
	file_ortools_constraint_solver_solver_parameters_proto_rawDescData = file_ortools_constraint_solver_solver_parameters_proto_rawDesc
)

func file_ortools_constraint_solver_solver_parameters_proto_rawDescGZIP() []byte {
	file_ortools_constraint_solver_solver_parameters_proto_rawDescOnce.Do(func() {
		file_ortools_constraint_solver_solver_parameters_proto_rawDescData = protoimpl.X.CompressGZIP(file_ortools_constraint_solver_solver_parameters_proto_rawDescData)
	})
	return file_ortools_constraint_solver_solver_parameters_proto_rawDescData
}

var file_ortools_constraint_solver_solver_parameters_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ortools_constraint_solver_solver_parameters_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ortools_constraint_solver_solver_parameters_proto_goTypes = []interface{}{
	(ConstraintSolverParameters_TrailCompression)(0), // 0: operations_research.ConstraintSolverParameters.TrailCompression
	(*ConstraintSolverParameters)(nil),               // 1: operations_research.ConstraintSolverParameters
}
var file_ortools_constraint_solver_solver_parameters_proto_depIdxs = []int32{
	0, // 0: operations_research.ConstraintSolverParameters.compress_trail:type_name -> operations_research.ConstraintSolverParameters.TrailCompression
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ortools_constraint_solver_solver_parameters_proto_init() }
func file_ortools_constraint_solver_solver_parameters_proto_init() {
	if File_ortools_constraint_solver_solver_parameters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ortools_constraint_solver_solver_parameters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstraintSolverParameters); i {
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
			RawDescriptor: file_ortools_constraint_solver_solver_parameters_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ortools_constraint_solver_solver_parameters_proto_goTypes,
		DependencyIndexes: file_ortools_constraint_solver_solver_parameters_proto_depIdxs,
		EnumInfos:         file_ortools_constraint_solver_solver_parameters_proto_enumTypes,
		MessageInfos:      file_ortools_constraint_solver_solver_parameters_proto_msgTypes,
	}.Build()
	File_ortools_constraint_solver_solver_parameters_proto = out.File
	file_ortools_constraint_solver_solver_parameters_proto_rawDesc = nil
	file_ortools_constraint_solver_solver_parameters_proto_goTypes = nil
	file_ortools_constraint_solver_solver_parameters_proto_depIdxs = nil
}
