// Copyright 2010-2018 Google LLC
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

// Typemaps for Routing Index Types. This does not define any type wrappings,
// because these index types are are never exposed to the target language.
// Instead, indices are manipulated as native target language types (e.g. python
// int).
// This file is to be %included when wrapped objects need to use these typemaps.

%include "ortools/base/base.i"
%import "ortools/util/go/vector.i"
%import "ortools/util/go/function.i"

%{
#include "ortools/constraint_solver/routing_types.h"
%}

%module(directors="1") operations_research;

// This macro defines typemaps for IndexT, std::vector<IndexT> and
// std::vector<std::vector<IndexT>>.
%define DEFINE_INDEX_TYPE(ns, IndexT)

// Convert IndexT to (32-bit signed) integers.
%typemap(gotype) ns IndexT "int"
%typemap(goin) ns IndexT {
  $result = $1
}
%typemap(imtype) ns IndexT "int"
%typemap(in) ns IndexT {
  $1 = ns IndexT($input);
}
%typemap(out) ns IndexT {
  $result = $1.value();
}
%typemap(goout) ns IndexT {
  return int($1);
}

// Convert std::vector<ns IndexT> to/from int slices.
VECTOR_AS_GO_SLICE_NAMESPACE(ns, IndexT, int)
// TODO: 2d slice conversion?
// Convert std::vector<std::vector<IndexT>> to/from two-dimensional int slices.
// VECTOR_AS_GO_SLICE_NAMESPACE(ns, std::vector<IndexT>, int)

%enddef  // DEFINE_INDEX_TYPE

// This macro applies all typemaps for a given index type to a typedef.
// Normally we should not need that as SWIG is supposed to automatically apply
// all typemaps to typedef definitions (http://www.swig.org/Doc2.0/SWIGDocumentation.html#Typemaps_typedef_reductions),
// but this is not actually the case.
%define DEFINE_INDEX_TYPE_TYPEDEF(IndexT, NewIndexT)
%apply IndexT { NewIndexT };
%apply std::vector<IndexT> { std::vector<NewIndexT> };
%apply std::vector<IndexT>* { std::vector<NewIndexT>* };
%apply const std::vector<IndexT>& { std::vector<NewIndexT>& };
%apply const std::vector<std::vector<IndexT> >& { const std::vector<std::vector<NewIndexT> >& };
%enddef  // DEFINE_INDEX_TYPE_TYPEDEF

DEFINE_INDEX_TYPE(operations_research::, RoutingNodeIndex);
DEFINE_INDEX_TYPE(operations_research::, RoutingCostClassIndex);
DEFINE_INDEX_TYPE(operations_research::, RoutingDimensionIndex);
DEFINE_INDEX_TYPE(operations_research::, RoutingDisjunctionIndex);
DEFINE_INDEX_TYPE(operations_research::, RoutingVehicleClassIndex);

STD_FUNCTION_AS_GO(RoutingTransitCallback1, int64_t, int64_t);
STD_FUNCTION_AS_GO(RoutingTransitCallback2, int64_t, int64_t, int64_t);
