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

#include "ortools/algorithms/set_cover_invariant.h"

#include <algorithm>
#include <limits>
#include <vector>

#include "absl/log/check.h"
#include "absl/types/span.h"
#include "ortools/algorithms/set_cover_model.h"
#include "ortools/base/logging.h"

namespace operations_research {
// Note: in many of the member functions, variables have "crypterse" names
// to avoid confusing them with member data. For example mrgnl_impcts is used
// to avoid confusion with marginal_impacts_.
void SetCoverInvariant::Initialize() {
  DCHECK(model_->ComputeFeasibility());
  model_->CreateSparseRowView();

  const SubsetIndex num_subsets(model_->num_subsets());
  is_selected_.assign(num_subsets, false);
  is_removable_.assign(num_subsets, false);
  marginal_impacts_.assign(num_subsets, ElementIndex(0));

  const SparseColumnView& columns = model_->columns();
  for (SubsetIndex subset(0); subset < num_subsets; ++subset) {
    marginal_impacts_[subset] = columns[subset].size().value();
  }
  const ElementIndex num_elements(model_->num_elements());
  coverage_.assign(num_elements, SubsetIndex(0));
  cost_ = 0.0;
  num_elements_covered_ = ElementIndex(0);
}

bool SetCoverInvariant::CheckConsistency() const {
  CHECK(CheckCoverageAndMarginalImpacts(is_selected_));
  CHECK(CheckIsRemovable());
  return true;
}

void SetCoverInvariant::LoadSolution(const SubsetBoolVector& c) {
  is_selected_ = c;
  MakeDataConsistent();
}

bool SetCoverInvariant::CheckSolution() const {
  bool is_ok = true;

  const ElementToSubsetVector cvrg = ComputeCoverage(is_selected_);
  const ElementIndex num_elements(model_->num_elements());
  for (ElementIndex element(0); element < num_elements; ++element) {
    if (cvrg[element] == 0) {
      LOG(ERROR) << "Recomputed coverage_ for element " << element << " = 0";
      is_ok = false;
    }
  }

  const Cost recomputed_cost = ComputeCost(is_selected_);
  if (cost_ != recomputed_cost) {
    LOG(ERROR) << "Cost = " << cost_
               << ", while recomputed cost_ = " << recomputed_cost;
    is_ok = false;
  }
  return is_ok;
}

bool SetCoverInvariant::CheckCoverageAgainstSolution(
    const SubsetBoolVector& choices) const {
  const SubsetIndex num_subsets(model_->num_subsets());
  DCHECK_EQ(num_subsets, choices.size());
  const ElementToSubsetVector cvrg = ComputeCoverage(choices);
  bool is_ok = true;
  const ElementIndex num_elements(model_->num_elements());
  for (ElementIndex element(0); element < num_elements; ++element) {
    if (coverage_[element] != cvrg[element]) {
      LOG(ERROR) << "Recomputed coverage_ for element " << element << " = "
                 << cvrg[element]
                 << ", while updated coverage_ = " << coverage_[element];
      is_ok = false;
    }
  }
  return is_ok;
}

bool SetCoverInvariant::CheckCoverageAndMarginalImpacts(
    const SubsetBoolVector& choices) const {
  const SubsetIndex num_subsets(model_->num_subsets());
  DCHECK_EQ(num_subsets, choices.size());
  const ElementToSubsetVector cvrg = ComputeCoverage(choices);
  bool is_ok = CheckCoverageAgainstSolution(choices);
  const SubsetToElementVector mrgnl_impcts = ComputeMarginalImpacts(cvrg);
  for (SubsetIndex subset(0); subset < num_subsets; ++subset) {
    if (marginal_impacts_[subset] != mrgnl_impcts[subset]) {
      LOG(ERROR) << "Recomputed marginal impact for subset " << subset << " = "
                 << mrgnl_impcts[subset] << ", while updated marginal impact = "
                 << marginal_impacts_[subset];
      is_ok = false;
    }
  }
  return is_ok;
}

// Used only once, for testing. TODO(user): Merge with
// CheckCoverageAndMarginalImpacts.
SubsetToElementVector SetCoverInvariant::ComputeMarginalImpacts(
    const ElementToSubsetVector& cvrg) const {
  const ElementIndex num_elements(model_->num_elements());
  DCHECK_EQ(num_elements, cvrg.size());
  const SparseColumnView& columns = model_->columns();
  const SubsetIndex num_subsets(model_->num_subsets());
  SubsetToElementVector mrgnl_impcts(num_subsets, ElementIndex(0));
  for (SubsetIndex subset(0); subset < num_subsets; ++subset) {
    for (ElementIndex element : columns[subset]) {
      if (cvrg[element] == 0) {
        ++mrgnl_impcts[subset];
      }
    }
    DCHECK_LE(mrgnl_impcts[subset], columns[subset].size().value());
    DCHECK_GE(mrgnl_impcts[subset], 0);
  }
  return mrgnl_impcts;
}

Cost SetCoverInvariant::ComputeCost(const SubsetBoolVector& c) const {
  DCHECK_EQ(c.size(), model_->num_subsets());
  Cost recomputed_cost = 0;
  const SubsetCostVector& subset_costs = model_->subset_costs();
  for (SubsetIndex subset(0); bool b : c) {
    if (b) {
      recomputed_cost += subset_costs[subset];
    }
    ++subset;
  }
  return recomputed_cost;
}

ElementIndex SetCoverInvariant::ComputeNumElementsCovered(
    const ElementToSubsetVector& cvrg) const {
  // Use "crypterse" naming to avoid confusing with num_elements_.
  int num_elmnts_cvrd = 0;
  for (ElementIndex element(0); element < model_->num_elements(); ++element) {
    if (cvrg[element] >= 1) {
      ++num_elmnts_cvrd;
    }
  }
  return ElementIndex(num_elmnts_cvrd);
}

ElementToSubsetVector SetCoverInvariant::ComputeCoverage(
    const SubsetBoolVector& choices) const {
  const ElementIndex num_elements(model_->num_elements());
  const SparseRowView& rows = model_->rows();
  // Use "crypterse" naming to avoid confusing with coverage_.
  ElementToSubsetVector cvrg(num_elements, SubsetIndex(0));
  for (ElementIndex element(0); element < num_elements; ++element) {
    for (SubsetIndex subset : rows[element]) {
      if (choices[subset]) {
        ++cvrg[element];
      }
    }
    DCHECK_LE(cvrg[element], rows[element].size().value());
    DCHECK_GE(cvrg[element], 0);
  }
  return cvrg;
}

bool SetCoverInvariant::CheckSingleSubsetCoverage(SubsetIndex subset) const {
  ElementToSubsetVector cvrg = ComputeSingleSubsetCoverage(subset);
  const SparseColumnView& columns = model_->columns();
  for (const ElementIndex element : columns[subset]) {
    DCHECK_EQ(coverage_[element], cvrg[element]) << " Element = " << element;
  }
  return true;
}

// Used only once, for testing. TODO(user): Merge with
// CheckSingleSubsetCoverage.
ElementToSubsetVector SetCoverInvariant::ComputeSingleSubsetCoverage(
    SubsetIndex subset) const {
  const SparseColumnView& columns = model_->columns();
  const ElementIndex num_elements(model_->num_elements());
  // Use "crypterse" naming to avoid confusing with cvrg.
  ElementToSubsetVector cvrg(num_elements, SubsetIndex(0));
  const SparseRowView& rows = model_->rows();
  for (const ElementIndex element : columns[subset]) {
    for (SubsetIndex subset : rows[element]) {
      if (is_selected_[subset]) {
        ++cvrg[element];
      }
    }
    DCHECK_LE(cvrg[element], rows[element].size().value());
    DCHECK_GE(cvrg[element], 0);
  }
  return cvrg;
}

std::vector<SubsetIndex> SetCoverInvariant::Toggle(SubsetIndex subset,
                                                   bool value) {
  // Note: "if p then q" is also "not(p) or q", or p <= q (p LE q).
  // If selected, then is_removable, to make sure we still have a solution.
  DCHECK(is_selected_[subset] <= is_removable_[subset]);
  // If value, then marginal_impact > 0, to not increase the cost.
  DCHECK((value <= (marginal_impacts_[subset] > 0)));
  return UnsafeToggle(subset, value);
}

std::vector<SubsetIndex> SetCoverInvariant::UnsafeToggle(SubsetIndex subset,
                                                         bool value) {
  // We allow to deselect a non-removable subset, but at least the
  // Toggle should be a real change.
  DCHECK_NE(is_selected_[subset], value);
  // If selected, then marginal_impact == 0.
  DCHECK(is_selected_[subset] <= (marginal_impacts_[subset] == 0));
  DVLOG(1) << (value ? "S" : "Des") << "electing subset " << subset;
  const SubsetCostVector& subset_costs = model_->subset_costs();
  cost_ += value ? subset_costs[subset] : -subset_costs[subset];
  is_selected_[subset] = value;
  UpdateCoverage(subset, value);
  const std::vector<SubsetIndex> impacted_subsets =
      ComputeImpactedSubsets(subset);
  UpdateIsRemovable(impacted_subsets);
  UpdateMarginalImpacts(impacted_subsets);
  DCHECK((is_selected_[subset] <= (marginal_impacts_[subset] == 0)));
  return impacted_subsets;
}

void SetCoverInvariant::UpdateCoverage(SubsetIndex subset, bool value) {
  const SparseColumnView& columns = model_->columns();
  const SparseRowView& rows = model_->rows();
  const int delta = value ? 1 : -1;
  for (const ElementIndex element : columns[subset]) {
    DVLOG(2) << "Coverage of element " << element << " changed from "
             << coverage_[element] << " to " << coverage_[element] + delta;
    coverage_[element] += delta;
    DCHECK_GE(coverage_[element], 0);
    DCHECK_LE(coverage_[element], rows[element].size().value());
    if (coverage_[element] == 1) {
      ++num_elements_covered_;
    } else if (coverage_[element] == 0) {
      --num_elements_covered_;
    }
  }
  DCHECK(CheckSingleSubsetCoverage(subset));
}

// Compute the impact of the change in the assignment for each subset
// containing element. Be careful to add the elements only once.
std::vector<SubsetIndex> SetCoverInvariant::ComputeImpactedSubsets(
    SubsetIndex subset) const {
  const SparseColumnView& columns = model_->columns();
  const SparseRowView& rows = model_->rows();
  SubsetBoolVector subset_seen(columns.size(), false);
  std::vector<SubsetIndex> impacted_subsets;
  impacted_subsets.reserve(columns.size().value());
  for (const ElementIndex element : columns[subset]) {
    for (const SubsetIndex subset : rows[element]) {
      if (!subset_seen[subset]) {
        subset_seen[subset] = true;
        impacted_subsets.push_back(subset);
      }
    }
  }
  DCHECK_LE(impacted_subsets.size(), model_->num_subsets());
  // Testing has shown there is no gain in sorting impacted_subsets.
  return impacted_subsets;
}

bool SetCoverInvariant::ComputeIsRemovable(SubsetIndex subset) const {
  DCHECK(CheckSingleSubsetCoverage(subset));
  const SparseColumnView& columns = model_->columns();
  for (const ElementIndex element : columns[subset]) {
    if (coverage_[element] <= 1) {
      return false;
    }
  }
  return true;
}

void SetCoverInvariant::UpdateIsRemovable(
    absl::Span<const SubsetIndex> impacted_subsets) {
  for (const SubsetIndex subset : impacted_subsets) {
    is_removable_[subset] = ComputeIsRemovable(subset);
  }
}

SubsetBoolVector SetCoverInvariant::ComputeIsRemovable(
    const ElementToSubsetVector& cvrg) const {
  DCHECK(CheckCoverageAgainstSolution(is_selected_));
  const SubsetIndex num_subsets(model_->num_subsets());
  SubsetBoolVector is_rmvble(num_subsets, true);
  const SparseRowView& rows = model_->rows();
  for (ElementIndex element(0); element < rows.size(); ++element) {
    if (cvrg[element] <= 1) {
      for (const SubsetIndex subset : rows[element]) {
        is_rmvble[subset] = false;
      }
    }
  }
  for (SubsetIndex subset(0); subset < num_subsets; ++subset) {
    DCHECK_EQ(is_rmvble[subset], ComputeIsRemovable(subset));
  }
  return is_rmvble;
}

bool SetCoverInvariant::CheckIsRemovable() const {
  const SubsetBoolVector is_rmvble = ComputeIsRemovable(coverage_);
  const SubsetIndex num_subsets(model_->num_subsets());
  for (SubsetIndex subset(0); subset < num_subsets; ++subset) {
    DCHECK_EQ(is_rmvble[subset], ComputeIsRemovable(subset));
  }
  return true;
}

void SetCoverInvariant::UpdateMarginalImpacts(
    absl::Span<const SubsetIndex> impacted_subsets) {
  const SparseColumnView& columns = model_->columns();
  for (const SubsetIndex subset : impacted_subsets) {
    ElementIndex impact(0);
    for (const ElementIndex element : columns[subset]) {
      if (coverage_[element] == 0) {
        ++impact;
      }
    }
    DVLOG(2) << "Changing impact of subset " << subset << " from "
             << marginal_impacts_[subset] << " to " << impact;
    marginal_impacts_[subset] = impact;
    DCHECK_LE(marginal_impacts_[subset], columns[subset].size().value());
    DCHECK_GE(marginal_impacts_[subset], 0);
  }
  DCHECK(CheckCoverageAndMarginalImpacts(is_selected_));
}

std::vector<SubsetIndex> SetCoverInvariant::ComputeSettableSubsets() const {
  SubsetBoolVector subset_seen(model_->num_subsets(), false);
  std::vector<SubsetIndex> focus;
  focus.reserve(model_->num_subsets().value());
  const SparseRowView& rows = model_->rows();
  for (ElementIndex element(0); element < rows.size(); ++element) {
    if (coverage_[element] >= 1) continue;
    for (const SubsetIndex subset : rows[element]) {
      if (!is_selected_[subset]) continue;
      if (subset_seen[subset]) continue;
      subset_seen[subset] = true;
      focus.push_back(subset);
    }
  }
  DCHECK_LE(focus.size(), model_->num_subsets());
  // Testing has shown there is no gain in sorting focus.
  return focus;
}

std::vector<SubsetIndex> SetCoverInvariant::ComputeResettableSubsets() const {
  SubsetBoolVector subset_seen(model_->num_subsets(), false);
  std::vector<SubsetIndex> focus;
  focus.reserve(model_->num_subsets().value());
  const SparseRowView& rows = model_->rows();
  for (ElementIndex element(0); element < rows.size(); ++element) {
    if (coverage_[element] < 1) continue;
    for (const SubsetIndex subset : rows[element]) {
      if (!is_selected_[subset]) continue;
      if (subset_seen[subset]) continue;
      subset_seen[subset] = true;
      focus.push_back(subset);
    }
  }
  DCHECK_LE(focus.size(), model_->num_subsets());
  // Testing has shown there is no gain in sorting focus.
  return focus;
}

SetCoverSolutionResponse SetCoverInvariant::ExportSolutionAsProto() const {
  SetCoverSolutionResponse message;
  message.set_num_subsets(is_selected_.size().value());
  Cost lower_bound = std::numeric_limits<Cost>::max();
  for (SubsetIndex subset(0); subset < model_->num_subsets(); ++subset) {
    if (is_selected_[subset]) {
      message.add_subset(subset.value());
    }
    lower_bound = std::min(model_->subset_costs()[subset], lower_bound);
  }
  message.set_cost(cost_);
  message.set_cost_lower_bound(lower_bound);
  return message;
}

void SetCoverInvariant::ImportSolutionFromProto(
    const SetCoverSolutionResponse& message) {
  is_selected_.resize(SubsetIndex(message.num_subsets()), false);
  for (auto s : message.subset()) {
    is_selected_[SubsetIndex(s)] = true;
  }
  MakeDataConsistent();
  Cost cost = message.cost();
  CHECK_EQ(cost, cost_);
}

}  // namespace operations_research
