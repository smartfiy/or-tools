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

// TODO(user): Refactor this file to adhere to the SWIG style guide.

%include "ortools/util/go/cgo.i"

%include "exception.i"
%include "std_vector.i"
%include "std_common.i"
%include "std_string.i"

%include "ortools/base/base.i"
%include "ortools/util/go/vector.i"
%include "ortools/util/go/proto.i"

// We need to forward-declare the proto here, so that PROTO_INPUT involving it
// works correctly. The order matters very much: this declaration needs to be
// before the %{ #include ".../constraint_solver.h" %}.
namespace operations_research {
class ConstraintSolverParameters;
class RegularLimitParameters;
}  // namespace operations_research

%module(directors="1") operations_research;
#pragma SWIG nowarn=473

%{
#include <setjmp.h>

#include <string>
#include <vector>
#include <functional>

#include "ortools/base/integral_types.h"
#include "ortools/constraint_solver/constraint_solver.h"
#include "ortools/constraint_solver/constraint_solveri.h"
#include "ortools/constraint_solver/search_limit.pb.h"
#include "ortools/constraint_solver/solver_parameters.pb.h"

namespace operations_research {
class LocalSearchPhaseParameters {
 public:
  LocalSearchPhaseParameters() {}
  ~LocalSearchPhaseParameters() {}
};
}  // namespace operations_research

struct FailureProtect {
  jmp_buf exception_buffer;
  void JumpBack() {
    longjmp(exception_buffer, 1);
  }
};
%}

// ############ BEGIN DUPLICATED CODE BLOCK ############
// IMPORTANT: keep this code block in sync with the .i
// files in ../python and ../csharp.

// Protect from failure.
%define PROTECT_FROM_FAILURE(Method, GetSolver)
%exception Method {
  operations_research::Solver* const solver = GetSolver;
  FailureProtect protect;
  solver->set_fail_intercept([&protect]() { protect.JumpBack(); });
  if (setjmp(protect.exception_buffer) == 0) {
    $action
    solver->clear_fail_intercept();
  } else {
    solver->clear_fail_intercept();
    _swig_gopanic("CP Solver fail");
  }
}
%enddef  // PROTECT_FROM_FAILURE

namespace operations_research {
PROTECT_FROM_FAILURE(IntExpr::SetValue(int64 v), arg1->solver());
PROTECT_FROM_FAILURE(IntExpr::SetMin(int64 v), arg1->solver());
PROTECT_FROM_FAILURE(IntExpr::SetMax(int64 v), arg1->solver());
PROTECT_FROM_FAILURE(IntExpr::SetRange(int64 l, int64 u), arg1->solver());
PROTECT_FROM_FAILURE(IntVar::RemoveValue(int64 v), arg1->solver());
PROTECT_FROM_FAILURE(IntVar::RemoveValues(const std::vector<int64>& values),
                     arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetStartMin(int64 m), arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetStartMax(int64 m), arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetStartRange(int64 mi, int64 ma),
                     arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetDurationMin(int64 m), arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetDurationMax(int64 m), arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetDurationRange(int64 mi, int64 ma),
                     arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetEndMin(int64 m), arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetEndMax(int64 m), arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetEndRange(int64 mi, int64 ma),
                     arg1->solver());
PROTECT_FROM_FAILURE(IntervalVar::SetPerformed(bool val), arg1->solver());
PROTECT_FROM_FAILURE(Solver::AddConstraint(Constraint* const c), arg1);
PROTECT_FROM_FAILURE(Solver::Fail(), arg1);
#undef PROTECT_FROM_FAILURE
}  // namespace operations_research

// ############ END DUPLICATED CODE BLOCK ############

%apply int64 * INOUT { int64 *const marker };
%apply int64 * OUTPUT { int64 *l, int64 *u, int64 *value };

// Since knapsack_solver.i and constraint_solver.i both need to
// instantiate the vector template, but their go_wrap.cc
// files end up being compiled into the same .dll, we must name the
// vector template differently.

// TupleSet depends on the previous typemaps
%include "ortools/util/go/tuple_set.i"

// Renaming
namespace operations_research {

// Decision
%feature("director") Decision;
%unignore Decision;
// Methods:
%rename (ApplyWrapper) Decision::Apply;
%rename (RefuteWrapper) Decision::Refute;

// DecisionBuilder
%feature("director") DecisionBuilder;
%unignore DecisionBuilder;
// Methods:
%rename (NextWrapper) DecisionBuilder::Next;

// SymmetryBreaker
// %feature("director") SymmetryBreaker;
// %unignore SymmetryBreaker;

// UnsortedNullableRevBitset
// TODO(corentinl) To removed from constraint_solveri.h (only use by table.cc)
%ignore UnsortedNullableRevBitset;

// Assignment
%unignore Assignment;
// Ignored:
%ignore Assignment::Load;
%ignore Assignment::Save;

// template AssignmentContainer<>
// Ignored:
%ignore AssignmentContainer::MutableElement;
%ignore AssignmentContainer::MutableElementOrNull;
%ignore AssignmentContainer::ElementPtrOrNull;
%ignore AssignmentContainer::elements;

// AssignmentElement
%unignore AssignmentElement;
// Methods:
%unignore AssignmentElement::Activate;
%unignore AssignmentElement::Deactivate;
%unignore AssignmentElement::Activated;

// IntVarElement
%unignore IntVarElement;
// Ignored:
%ignore IntVarElement::LoadFromProto;
%ignore IntVarElement::WriteToProto;

// IntervalVarElement
%unignore IntervalVarElement;
// Ignored:
%ignore IntervalVarElement::LoadFromProto;
%ignore IntervalVarElement::WriteToProto;

// SequenceVarElement
%unignore SequenceVarElement;
// Ignored:
%ignore SequenceVarElement::LoadFromProto;
%ignore SequenceVarElement::WriteToProto;

// SolutionCollector
%feature("director") SolutionCollector;
%unignore SolutionCollector;

// Solver
%unignore Solver;
// Ignored:
%ignore Solver::SearchLogParameters;
%ignore Solver::ActiveSearch;
%ignore Solver::SetSearchContext;
%ignore Solver::SearchContext;
%ignore Solver::MakeSearchLog(SearchLogParameters parameters);
%ignore Solver::MakeIntVarArray;
%ignore Solver::MakeBoolVarArray;
%ignore Solver::MakeFixedDurationIntervalVarArray;
%ignore Solver::SetBranchSelector;
%ignore Solver::MakeApplyBranchSelector;
%ignore Solver::MakeAtMost;
%ignore Solver::Now;
%ignore Solver::demon_profiler;
%ignore Solver::set_fail_intercept;
%ignore Solver::tmp_vector_;
// Methods:
// %rename (Add) Solver::AddConstraint;

// IntExpr
%unignore IntExpr;

// IntVar
%unignore IntVar;
// Ignored:
%ignore IntVar::MakeDomainIterator;
%ignore IntVar::MakeHoleIterator;
// Methods:
%extend IntVar {
  IntVarIterator* GetDomain() {
    return $self->MakeDomainIterator(false);
  }
  IntVarIterator* GetHoles() {
    return $self->MakeHoleIterator(false);
  }
}

// IntervalVar
%unignore IntervalVar;
// Extend IntervalVar with an intuitive API to create precedence constraints.
%extend IntervalVar {
  Constraint* EndsAfterEnd(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::ENDS_AFTER_END, other);
  }
  Constraint* EndsAfterStart(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::ENDS_AFTER_START, other);
  }
  Constraint* EndsAtEnd(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::ENDS_AT_END, other);
  }
  Constraint* EndsAtStart(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::ENDS_AT_START, other);
  }
  Constraint* StartsAfterEnd(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::STARTS_AFTER_END, other);
  }
  Constraint* StartsAfterStart(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::STARTS_AFTER_START, other);
  }
  Constraint* StartsAtEnd(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::STARTS_AT_END, other);
  }
  Constraint* StartsAtStart(IntervalVar* other) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::STARTS_AT_START, other);
  }
  Constraint* EndsAfterEndWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::ENDS_AFTER_END, other, delay);
  }
  Constraint* EndsAfterStartWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::ENDS_AFTER_START, other, delay);
  }
  Constraint* EndsAtEndWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::ENDS_AT_END, other, delay);
  }
  Constraint* EndsAtStartWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::ENDS_AT_START, other, delay);
  }
  Constraint* StartsAfterEndWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::STARTS_AFTER_END, other, delay);
  }
  Constraint* StartsAfterStartWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::STARTS_AFTER_START, other, delay);
  }
  Constraint* StartsAtEndWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::STARTS_AT_END, other, delay);
  }
  Constraint* StartsAtStartWithDelay(IntervalVar* other, int64 delay) {
    return $self->solver()->MakeIntervalVarRelationWithDelay($self, operations_research::Solver::STARTS_AT_START, other, delay);
  }
  Constraint* EndsAfter(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::ENDS_AFTER, date);
  }
  Constraint* EndsAt(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::ENDS_AT, date);
  }
  Constraint* EndsBefore(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::ENDS_BEFORE, date);
  }
  Constraint* StartsAfter(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::STARTS_AFTER, date);
  }
  Constraint* StartsAt(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::STARTS_AT, date);
  }
  Constraint* StartsBefore(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::STARTS_BEFORE, date);
  }
  Constraint* CrossesDate(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::CROSS_DATE, date);
  }
  Constraint* AvoidsDate(int64 date) {
    return $self->solver()->MakeIntervalVarRelation($self, operations_research::Solver::AVOID_DATE, date);
  }
  IntervalVar* RelaxedMax() {
    return $self->solver()->MakeIntervalRelaxedMax($self);
  }
  IntervalVar* RelaxedMin() {
    return $self->solver()->MakeIntervalRelaxedMin($self);
  }
}

// OptimizeVar
%feature("director") OptimizeVar;
%unignore OptimizeVar;
// Methods:
%unignore OptimizeVar::ApplyBound;
%unignore OptimizeVar::Print;
%unignore OptimizeVar::Var;

// SequenceVar
%unignore SequenceVar;
// Ignored:
%ignore SequenceVar::ComputePossibleFirstsAndLasts;
%ignore SequenceVar::FillSequence;

// Constraint
%feature("director") Constraint;
%unignore Constraint;
// Ignored:
%ignore Constraint::PostAndPropagate;
// Methods:
%rename (InitialPropagateWrapper) Constraint::InitialPropagate;
%feature ("nodirector") Constraint::Accept;
%feature ("nodirector") Constraint::Var;
%feature ("nodirector") Constraint::IsCastConstraint;

// DisjunctiveConstraint
%unignore DisjunctiveConstraint;
// Methods:
%rename (SequenceVar) DisjunctiveConstraint::MakeSequenceVar;

// Pack
%unignore Pack;

// PropagationBaseObject
%unignore PropagationBaseObject;
// Ignored:
%ignore PropagationBaseObject::ExecuteAll;
%ignore PropagationBaseObject::EnqueueAll;
%ignore PropagationBaseObject::set_action_on_fail;

// SearchMonitor
%feature("director") SearchMonitor;
%unignore SearchMonitor;

// SearchLimit
%feature("director") SearchLimit;
%unignore SearchLimit;
// Methods:
%rename (IsCrossed) SearchLimit::crossed;

// RegularLimit
%feature("director") RegularLimit;
%unignore RegularLimit;
%ignore RegularLimit::duration_limit;
%ignore RegularLimit::AbsoluteSolverDeadline;

// Searchlog
%unignore SearchLog;
// Ignored:
// No custom wrapping for this method, we simply ignore it.
%ignore SearchLog::SearchLog(
    Solver* const s, OptimizeVar* const obj, IntVar* const var,
    double scaling_factor, double offset,
    std::function<std::string()> display_callback, int period);
// Methods:
%unignore SearchLog::Maintain;
%unignore SearchLog::OutputDecision;

// IntVarLocalSearchHandler
%ignore IntVarLocalSearchHandler;

// SequenceVarLocalSearchHandler
%ignore SequenceVarLocalSearchHandler;

// LocalSearchOperator
%feature("director") LocalSearchOperator;
%unignore LocalSearchOperator;
// Methods:
%unignore LocalSearchOperator::MakeNextNeighbor;
%unignore LocalSearchOperator::Reset;
%unignore LocalSearchOperator::Start;

// VarLocalSearchOperator<>
%unignore VarLocalSearchOperator;
// Ignored:
%ignore VarLocalSearchOperator::Start;
%ignore VarLocalSearchOperator::ApplyChanges;
%ignore VarLocalSearchOperator::RevertChanges;
%ignore VarLocalSearchOperator::SkipUnchanged;
// Methods:
%unignore VarLocalSearchOperator::Size;
%unignore VarLocalSearchOperator::Value;
%unignore VarLocalSearchOperator::IsIncremental;
%unignore VarLocalSearchOperator::OnStart;
%unignore VarLocalSearchOperator::OldValue;
%unignore VarLocalSearchOperator::SetValue;
%unignore VarLocalSearchOperator::Var;
%unignore VarLocalSearchOperator::Activated;
%unignore VarLocalSearchOperator::Activate;
%unignore VarLocalSearchOperator::Deactivate;
%unignore VarLocalSearchOperator::AddVars;

// IntVarLocalSearchOperator
%feature("director") IntVarLocalSearchOperator;
%unignore IntVarLocalSearchOperator;
// Ignored:
%ignore IntVarLocalSearchOperator::MakeNextNeighbor;
// Methods:
%unignore IntVarLocalSearchOperator::Size;
%unignore IntVarLocalSearchOperator::MakeOneNeighbor;
%unignore IntVarLocalSearchOperator::Value;
%unignore IntVarLocalSearchOperator::IsIncremental;
%unignore IntVarLocalSearchOperator::OnStart;
%unignore IntVarLocalSearchOperator::OldValue;
%unignore IntVarLocalSearchOperator::SetValue;
%unignore IntVarLocalSearchOperator::Var;
%unignore IntVarLocalSearchOperator::Activated;
%unignore IntVarLocalSearchOperator::Activate;
%unignore IntVarLocalSearchOperator::Deactivate;
%unignore IntVarLocalSearchOperator::AddVars;

// BaseLns
%feature("director") BaseLns;
%unignore BaseLns;
// Methods:
%unignore BaseLns::InitFragments;
%unignore BaseLns::NextFragment;
%feature ("nodirector") BaseLns::OnStart;
%feature ("nodirector") BaseLns::SkipUnchanged;
%feature ("nodirector") BaseLns::MakeOneNeighbor;
%unignore BaseLns::IsIncremental;
%unignore BaseLns::AppendToFragment;
%unignore BaseLns::FragmentSize;

// ChangeValue
%feature("director") ChangeValue;
%unignore ChangeValue;
// Methods:
%unignore ChangeValue::ModifyValue;

// SequenceVarLocalSearchOperator
%feature("director") SequenceVarLocalSearchOperator;
%unignore SequenceVarLocalSearchOperator;
// Ignored:
%ignore SequenceVarLocalSearchOperator::SetBackwardSequence;
%ignore SequenceVarLocalSearchOperator::SetForwardSequence;
// Methods:
%unignore SequenceVarLocalSearchOperator::OldSequence;
%unignore SequenceVarLocalSearchOperator::Sequence;
%unignore SequenceVarLocalSearchOperator::Start;

// PathOperator
%feature("director") PathOperator;
%unignore PathOperator;
// Ignored:
%ignore PathOperator::PathOperator;
%ignore PathOperator::Next;
%ignore PathOperator::Path;
%ignore PathOperator::SkipUnchanged;
%ignore PathOperator::number_of_nexts;
// Methods:
%unignore PathOperator::MakeNeighbor;

// LocalSearchFilter
%feature("director") LocalSearchFilter;
%unignore LocalSearchFilter;
// Methods:
%unignore LocalSearchFilter::Accept;
%unignore LocalSearchFilter::Synchronize;
%unignore LocalSearchFilter::IsIncremental;

// IntVarLocalSearchFilter
%feature("director") IntVarLocalSearchFilter;
%unignore IntVarLocalSearchFilter;
// Ignored:
%ignore IntVarLocalSearchFilter::FindIndex;
%ignore IntVarLocalSearchFilter::IntVarLocalSearchFilter(
    const std::vector<IntVar*>& vars,
    Solver::ObjectiveWatcher objective_callback);
%ignore IntVarLocalSearchFilter::IsVarSynced;
// Methods:
%feature("nodirector") IntVarLocalSearchFilter::Synchronize;  // Inherited.
%unignore IntVarLocalSearchFilter::AddVars;  // Inherited.
%unignore IntVarLocalSearchFilter::IsIncremental;
%unignore IntVarLocalSearchFilter::OnSynchronize;
%unignore IntVarLocalSearchFilter::Size;
%unignore IntVarLocalSearchFilter::Start;
%unignore IntVarLocalSearchFilter::Value;
%unignore IntVarLocalSearchFilter::Var;  // Inherited.
// Extend IntVarLocalSearchFilter with an intuitive API.
%extend IntVarLocalSearchFilter {
  int Index(IntVar* const var) {
    int64 index = -1;
    $self->FindIndex(var, &index);
    return index;
  }
}

// Demon
%feature("director") Demon;
%unignore Demon;
// Methods:
%feature("nodirector") Demon::inhibit;
%feature("nodirector") Demon::desinhibit;
%rename (RunWrapper) Demon::Run;
%rename (Inhibit) Demon::inhibit;
%rename (Desinhibit) Demon::desinhibit;

class LocalSearchPhaseParameters {
 public:
  LocalSearchPhaseParameters();
  ~LocalSearchPhaseParameters();
};

}  // namespace operations_research

%define CONVERT_VECTOR(CTYPE, TYPE)
SWIG_STD_VECTOR_ENHANCED(CTYPE*);
%template(TYPE ## Vector) std::vector<CTYPE*>;
%enddef  // CONVERT_VECTOR

CONVERT_VECTOR(operations_research::IntVar, IntVar)
CONVERT_VECTOR(operations_research::SearchMonitor, SearchMonitor)
CONVERT_VECTOR(operations_research::DecisionBuilder, DecisionBuilder)
CONVERT_VECTOR(operations_research::IntervalVar, IntervalVar)
CONVERT_VECTOR(operations_research::SequenceVar, SequenceVar)
CONVERT_VECTOR(operations_research::LocalSearchOperator, LocalSearchOperator)
CONVERT_VECTOR(operations_research::LocalSearchFilter, LocalSearchFilter)
CONVERT_VECTOR(operations_research::SymmetryBreaker, SymmetryBreaker)

#undef CONVERT_VECTOR

// Generic rename rule
%rename("%(camelcase)s", %$isfunction) "";
%rename (ToString) *::DebugString;
%rename (solver) *::solver;

// Protobuf support
PROTO_INPUT(operations_research::ConstraintSolverParameters,
            ConstraintSolverParameters,
            parameters)
PROTO2_RETURN(operations_research::ConstraintSolverParameters,
              ConstraintSolverParameters)

PROTO_INPUT(operations_research::RegularLimitParameters,
            RegularLimitParameters,
            proto)
PROTO2_RETURN(operations_research::RegularLimitParameters,
              RegularLimitParameters)

PROTO_INPUT(operations_research::CpModel,
            CpModel,
            proto)
PROTO2_RETURN(operations_research::CpModel,
              CpModel)

namespace operations_research {
// Globals
// Ignored:
%ignore FillValues;
}  // namespace operations_research

// Wrap cp includes
// TODO(user): Replace with %ignoreall/%unignoreall
//swiglint: disable include-h-allglobals
%include "ortools/constraint_solver/constraint_solver.h"
%include "ortools/constraint_solver/constraint_solveri.h"

namespace operations_research {
%template(RevInteger) Rev<int64>;
%template(RevBool) Rev<bool>;
typedef Assignment::AssignmentContainer AssignmentContainer;
%template(AssignmentIntContainer) AssignmentContainer<IntVar, IntVarElement>;
%template(AssignmentIntervalContainer) AssignmentContainer<IntervalVar, IntervalVarElement>;
%template(AssignmentSequenceContainer) AssignmentContainer<SequenceVar, SequenceVarElement>;
}
