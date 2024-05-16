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

%include "ortools/base/base.i"

%{
#include <vector>
#include "ortools/base/types.h"
%}

%insert(go_header) %{
type swig_goslice struct { arr uintptr; n int; c int }
%}

%define _VECTOR_AS_GO_SLICE(ns, name, goname, gonameim, ref, deref)
%{
std::vector< ns name ref > name##SliceToVector(_goslice_ slice) {
    std::vector< ns name ref > v;
    for (int i = 0; i < slice.len; i++) {
        ns name ref a = (( ns name * ref )slice.array)[i];
        v.push_back(a);
    }
    return v;
}

_goslice_ vectorTo##name##Slice(const std::vector< ns name ref >& arr) {
    _goslice_ slice;
    size_t count = arr.size();
    ns name * ref go_arr = ( ns name * ref )malloc(sizeof( ns name ref) * count);
    slice.array = go_arr;
    slice.len = slice.cap = count;
    
    for (int i = 0; i < count; i++) {
        go_arr[i] = arr[i];
    }
    
    return slice;
}
%}

%insert(go_header) %{
func swigCopy##name##SliceIn(s []goname) []gonameim {
    newSlice := make([] gonameim, len(s))
    for i := range newSlice {
        newSlice[i] = gonameim(s[i])
    }
    return newSlice
}

func swigCopy##name##SliceOut(s *[]gonameim) []goname {
    newSlice := make([] goname, len(*s))
    for i := range newSlice {
        newSlice[i] = goname((*s)[i])
    }
    p := *(*swig_goslice)(unsafe.Pointer(s))
    Swig_free(p.arr)
    return newSlice
}
%}

%typemap(gotype) std::vector< ns name ref > "[]goname"
#if "gonameim" != "goname"
%typemap(imtype) std::vector< ns name ref > "[]gonameim"
#endif

#if "gonameim" != "goname"
%typemap(goin) std::vector< ns name ref > %{
    $result = swigCopy##name##SliceIn($input)
%}
#endif

%typemap(in) std::vector< ns name ref > %{
    $1 = name##SliceToVector($input);
%}

%typemap(out) std::vector< ns name ref > %{
    $result = vectorTo##name##Slice($1);
%}

#if "gonameim" != "goname"
%typemap(goout) std::vector< ns name ref > %{
    $result = swigCopy##name##SliceOut(&$1)
%}
#endif

%typemap(gotype) const std::vector< ns name ref >& "[]goname"
#if "gonameim" != "goname"
%typemap(imtype) const std::vector< ns name ref >& "[]gonameim"
#endif

#if "gonameim" != "goname"
%typemap(goin) const std::vector< ns name ref > & %{
    $result = swigCopy##name##SliceIn($input)
%}
#endif

%typemap(in) const std::vector< ns name ref > & %{
    $*1_ltype $1_arr;
    $1_arr = name##SliceToVector($input);
    $1 = &$1_arr;
%}

%typemap(out) const std::vector< ns name ref > & %{
    $result = vectorTo##name##Slice(*$1);
%}

#if "gonameim" != "goname"
%typemap(goout) const std::vector< ns name ref > & %{
    $result = swigCopy##name##SliceOut(&$1)
%}
#endif

%insert(go_header) %{
type name##SliceWithPointer struct {
    slice []gonameim
    ptr uintptr
}
%}

%enddef

#define nothing
#define VECTOR_AS_GO_SLICE(name, goname, gonameim) _VECTOR_AS_GO_SLICE(nothing, name, goname, gonameim, nothing, *)
#define VECTOR_AS_GO_SLICE_NAMESPACE(ns, name, goname, gonameim) _VECTOR_AS_GO_SLICE(ns, name, goname, gonameim, nothing, *)

VECTOR_AS_GO_SLICE(int, int, C.int)
VECTOR_AS_GO_SLICE(int64_t, int64, int64)
