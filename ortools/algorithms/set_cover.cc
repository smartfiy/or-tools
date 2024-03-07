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

#include "ortools/algorithms/set_cover.h"

#include <algorithm>
#include <cstddef>
#include <iterator>
#include <limits>
#include <numeric>
#include <vector>

#include "absl/container/flat_hash_set.h"
#include "absl/log/check.h"
#include "absl/random/random.h"
#include "absl/types/span.h"
#include "ortools/algorithms/set_cover_invariant.h"
#include "ortools/algorithms/set_cover_model.h"
#include "ortools/algorithms/set_cover_utils.h"
#include "ortools/base/logging.h"

namespace operations_research {

constexpr SubsetIndex kNotFound(-1);

// TrivialSolutionGenerator.

bool TrivialSolutionGenerator::NextSolution() {
  return NextSolution(inv_->model()->all_subsets());
}

bool TrivialSolutionGenerator::NextSolution(
    absl::Span<const SubsetIndex> focus) {
  const SubsetIndex num_subsets(inv_->model()->num_subsets());
  SubsetBoolVector choices(num_subsets, false);
  for (const SubsetIndex subset : focus) {
    choices[subset] = true;
  }
  inv_->LoadSolution(choices);
  return true;
}

// RandomSolutionGenerator.

bool RandomSolutionGenerator::NextSolution() {
  return NextSolution(inv_->model()->all_subsets());
}

bool RandomSolutionGenerator::NextSolution(
    const std::vector<SubsetIndex>& focus) {
  std::vector<SubsetIndex> shuffled = focus;
  std::shuffle(shuffled.begin(), shuffled.end(), absl::BitGen());
  for (const SubsetIndex subset : shuffled) {
    if (inv_->is_selected()[subset]) continue;
    if (inv_->marginal_impacts()[subset] != 0) {
      inv_->Toggle(subset, true);
    }
  }
  DCHECK(inv_->CheckConsistency());
  return true;
}

// GreedySolutionGenerator.

void GreedySolutionGenerator::UpdatePriorities(
    const std::vector<SubsetIndex>& impacted_subsets,
    const SubsetCostVector& costs) {
  for (const SubsetIndex subset : impacted_subsets) {
    const ElementIndex marginal_impact(inv_->marginal_impacts()[subset]);
    if (marginal_impact != 0) {
      const Cost marginal_cost_increase =
          costs[subset] / marginal_impact.value();
      pq_.ChangePriority(subset, -marginal_cost_increase);
    } else {
      pq_.Remove(subset);
    }
  }
}

bool GreedySolutionGenerator::NextSolution() {
  return NextSolution(inv_->model()->all_subsets(),
                      inv_->model()->subset_costs());
}

bool GreedySolutionGenerator::NextSolution(
    const std::vector<SubsetIndex>& focus) {
  return NextSolution(focus, inv_->model()->subset_costs());
}

bool GreedySolutionGenerator::NextSolution(
    const std::vector<SubsetIndex>& focus, const SubsetCostVector& costs) {
  inv_->MakeDataConsistent();

  // The priority is the minimum marginal cost increase. Since the
  // priority queue returns the smallest value, we use the opposite.
  // TODO(user): build in O(N) instead of O(n lg(N)).
  for (const SubsetIndex subset : focus) {
    if (!inv_->is_selected()[subset] && inv_->marginal_impacts()[subset] != 0) {
      const Cost marginal_cost_increase =
          costs[subset] / inv_->marginal_impacts()[subset].value();
      pq_.Add(subset, -marginal_cost_increase);
    }
  }
  const ElementIndex num_elements(inv_->model()->num_elements());
  ElementIndex num_elements_covered(inv_->num_elements_covered());
  while (num_elements_covered < num_elements && !pq_.IsEmpty()) {
    const SubsetIndex best_subset = pq_.TopSubset();
    DVLOG(1) << "Best subset: " << best_subset.value()
             << " Priority = " << pq_.Priority(best_subset)
             << " queue size = " << pq_.Size();
    const std::vector<SubsetIndex> impacted_subsets =
        inv_->Toggle(best_subset, true);
    UpdatePriorities(impacted_subsets, costs);
    num_elements_covered = inv_->num_elements_covered();
    DVLOG(1) << "Cost = " << inv_->cost() << " num_uncovered_elements = "
             << num_elements - num_elements_covered;
  }
  DCHECK(pq_.IsEmpty());
  DCHECK(inv_->CheckConsistency());
  DCHECK(inv_->CheckSolution());
  return true;
}

// SteepestSearch.

void SteepestSearch::UpdatePriorities(
    absl::Span<const SubsetIndex> impacted_subsets) {
  // Update priority queue. Since best_subset is in impacted_subsets, it will
  // be removed.
  for (const SubsetIndex subset : impacted_subsets) {
    pq_.Remove(subset);
  }
}

bool SteepestSearch::NextSolution(int num_iterations) {
  return NextSolution(inv_->model()->all_subsets(),
                      inv_->model()->subset_costs(), num_iterations);
}

bool SteepestSearch::NextSolution(absl::Span<const SubsetIndex> focus,
                                  int num_iterations) {
  return NextSolution(focus, inv_->model()->subset_costs(), num_iterations);
}

bool SteepestSearch::NextSolution(absl::Span<const SubsetIndex> focus,
                                  const SubsetCostVector& costs,
                                  int num_iterations) {
  // Return false if inv_ contains no solution.
  if (!inv_->CheckSolution()) return false;
  // Create priority queue with cost of using a subset, by decreasing order.
  // Do it only for removable subsets.
  for (const SubsetIndex subset : focus) {
    // The priority is the gain from removing the subset from the solution.
    if (inv_->is_selected()[subset] && inv_->is_removable()[subset]) {
      pq_.Add(subset, costs[subset]);
    }
  }
  for (int iteration = 0; iteration < num_iterations && !pq_.IsEmpty();
       ++iteration) {
    const SubsetIndex best_subset = pq_.TopSubset();
    DCHECK_GT(costs[best_subset], 0.0);
    DCHECK(inv_->is_removable()[best_subset]);
    DCHECK(inv_->is_selected()[best_subset]);
    const std::vector<SubsetIndex> impacted_subsets =
        inv_->Toggle(best_subset, false);
    UpdatePriorities(impacted_subsets);
    DVLOG(1) << "Cost = " << inv_->cost();
  }
  DCHECK(inv_->CheckConsistency());
  DCHECK(inv_->CheckSolution());
  return true;
}

// Guided Tabu Search

void GuidedTabuSearch::Initialize() {
  const SparseColumnView& columns = inv_->model()->columns();
  const SubsetCostVector& subset_costs = inv_->model()->subset_costs();
  times_penalized_.AssignToZero(columns.size());
  augmented_costs_ = subset_costs;
  utilities_ = subset_costs;
}

namespace {
bool FlipCoin() {
  // TODO(user): use STL for repeatable testing.
  return absl::Bernoulli(absl::BitGen(), 0.5);
}
}  // namespace

void GuidedTabuSearch::UpdatePenalties(absl::Span<const SubsetIndex> focus) {
  const SubsetCostVector& subset_costs = inv_->model()->subset_costs();
  Cost max_utility = -1.0;
  for (const SubsetIndex subset : focus) {
    if (inv_->is_selected()[subset]) {
      max_utility = std::max(max_utility, utilities_[subset]);
    }
  }
  const double epsilon_utility = epsilon_ * max_utility;
  for (const SubsetIndex subset : focus) {
    if (inv_->is_selected()[subset]) {
      const double utility = utilities_[subset];
      if ((max_utility - utility <= epsilon_utility) && FlipCoin()) {
        ++times_penalized_[subset];
        const int times_penalized = times_penalized_[subset];
        const Cost cost =
            subset_costs[subset];  // / columns[subset].size().value();
        utilities_[subset] = cost / (1 + times_penalized);
        augmented_costs_[subset] =
            cost * (1 + penalty_factor_ * times_penalized);
      }
    }
  }
}

bool GuidedTabuSearch::NextSolution(int num_iterations) {
  return NextSolution(inv_->model()->all_subsets(), num_iterations);
}

bool GuidedTabuSearch::NextSolution(const std::vector<SubsetIndex>& focus,
                                    int num_iterations) {
  const SubsetCostVector& subset_costs = inv_->model()->subset_costs();
  constexpr Cost kMaxPossibleCost = std::numeric_limits<Cost>::max();
  Cost best_cost = inv_->cost();
  SubsetBoolVector best_choices = inv_->is_selected();
  Cost augmented_cost =
      std::accumulate(augmented_costs_.begin(), augmented_costs_.end(), 0.0);
  for (int iteration = 0; iteration < num_iterations; ++iteration) {
    Cost best_delta = kMaxPossibleCost;
    SubsetIndex best_subset = kNotFound;
    for (const SubsetIndex subset : focus) {
      const Cost delta = augmented_costs_[subset];
      DVLOG(1) << "Subset, " << subset.value() << ", at ,"
               << inv_->is_selected()[subset] << ", is removable =, "
               << inv_->is_removable()[subset] << ", delta =, " << delta
               << ", best_delta =, " << best_delta;
      if (inv_->is_selected()[subset]) {
        // Try to remove subset from solution, if the gain from removing is
        // worth it:
        if (-delta < best_delta &&
            // and it can be removed, and
            inv_->is_removable()[subset] &&
            // it is not Tabu OR decreases the actual cost (aspiration):
            (!tabu_list_.Contains(subset) ||
             inv_->cost() - subset_costs[subset] < best_cost)) {
          best_delta = -delta;
          best_subset = subset;
        }
      } else {
        // Try to use subset in solution, if its penalized delta is good.
        if (delta < best_delta) {
          // The limit kMaxPossibleCost is ill-defined,
          // there is always a best_subset. Is it intended?
          if (!tabu_list_.Contains(subset)) {
            best_delta = delta;
            best_subset = subset;
          }
        }
      }
    }
    if (best_subset == kNotFound) {  // Local minimum reached.
      inv_->LoadSolution(best_choices);
      return true;
    }
    DVLOG(1) << "Best subset, " << best_subset.value() << ", at ,"
             << inv_->is_selected()[best_subset] << ", is removable = ,"
             << inv_->is_removable()[best_subset] << ", best_delta = ,"
             << best_delta;

    UpdatePenalties(focus);
    tabu_list_.Add(best_subset);
    const std::vector<SubsetIndex> impacted_subsets =
        inv_->UnsafeToggle(best_subset, !inv_->is_selected()[best_subset]);
    // TODO(user): make the cost computation incremental.
    augmented_cost =
        std::accumulate(augmented_costs_.begin(), augmented_costs_.end(), 0.0);

    DVLOG(1) << "Iteration, " << iteration << ", current cost = ,"
             << inv_->cost() << ", best cost = ," << best_cost
             << ", penalized cost = ," << augmented_cost;
    if (inv_->cost() < best_cost) {
      LOG(INFO) << "Updated best cost, " << "Iteration, " << iteration
                << ", current cost = ," << inv_->cost() << ", best cost = ,"
                << best_cost << ", penalized cost = ," << augmented_cost;
      best_cost = inv_->cost();
      best_choices = inv_->is_selected();
    }
  }
  inv_->LoadSolution(best_choices);
  DCHECK(inv_->CheckConsistency());
  DCHECK(inv_->CheckSolution());
  return true;
}

namespace {
void SampleSubsets(std::vector<SubsetIndex>* list, std::size_t num_subsets) {
  num_subsets = std::min(num_subsets, list->size());
  CHECK_GE(num_subsets, 0);
  std::shuffle(list->begin(), list->end(), absl::BitGen());
  list->resize(num_subsets);
}
}  // namespace

std::vector<SubsetIndex> ClearRandomSubsets(std::size_t num_subsets,
                                            SetCoverInvariant* inv) {
  return ClearRandomSubsets(inv->model()->all_subsets(), num_subsets, inv);
}

std::vector<SubsetIndex> ClearRandomSubsets(absl::Span<const SubsetIndex> focus,
                                            std::size_t num_subsets,
                                            SetCoverInvariant* inv) {
  num_subsets = std::min(num_subsets, focus.size());
  CHECK_GE(num_subsets, 0);
  std::vector<SubsetIndex> chosen_indices;
  for (const SubsetIndex subset : focus) {
    if (inv->is_selected()[subset]) {
      chosen_indices.push_back(subset);
    }
  }
  SampleSubsets(&chosen_indices, num_subsets);
  for (const SubsetIndex subset : chosen_indices) {
    // Use UnsafeToggle because we allow non-solutions.
    inv->UnsafeToggle(subset, false);
  }
  return chosen_indices;
}

std::vector<SubsetIndex> ClearMostCoveredElements(std::size_t num_subsets,
                                                  SetCoverInvariant* inv) {
  return ClearMostCoveredElements(inv->model()->all_subsets(), num_subsets,
                                  inv);
}

std::vector<SubsetIndex> ClearMostCoveredElements(
    absl::Span<const SubsetIndex> focus, std::size_t num_subsets,
    SetCoverInvariant* inv) {
  // This is the vector we will return.
  std::vector<SubsetIndex> chosen_indices;

  const ElementToSubsetVector& coverage = inv->coverage();

  // Compute a permutation of the element indices by decreasing order of
  // coverage by element.
  std::vector<ElementIndex> permutation(coverage.size().value());
  std::iota(permutation.begin(), permutation.end(), 0);
  std::sort(permutation.begin(), permutation.end(),
            [&coverage](ElementIndex i, ElementIndex j) {
              return coverage[i] > coverage[j];
            });

  // Now, for the elements that are over-covered (coverage > 1), collect the
  // sets that are used.
  absl::flat_hash_set<SubsetIndex> used_subsets_collection;
  for (ElementIndex element : permutation) {
    if (coverage[element] <= 1) break;
    for (SubsetIndex subset : inv->model()->rows()[element]) {
      if (inv->is_selected()[subset]) {
        used_subsets_collection.insert(subset);
      }
    }
  }

  // Now the impacted subset is a vector representation of the flat_hash_set
  // collection.
  std::vector<SubsetIndex> impacted_subsets(used_subsets_collection.begin(),
                                            used_subsets_collection.end());
  // Sort the impacted subsets to be able to intersect the vector later.
  std::sort(impacted_subsets.begin(), impacted_subsets.end());

  // chosen_indices = focus ⋂ impacted_subsets
  std::set_intersection(focus.begin(), focus.end(), impacted_subsets.begin(),
                        impacted_subsets.end(),
                        std::back_inserter(chosen_indices));

  std::shuffle(chosen_indices.begin(), chosen_indices.end(), absl::BitGen());
  chosen_indices.resize(std::min(chosen_indices.size(), num_subsets));

  // Sort before traversing indices (and memory) in order.
  std::sort(chosen_indices.begin(), chosen_indices.end());
  for (const SubsetIndex subset : chosen_indices) {
    // Use UnsafeToggle because we allow non-solutions.
    inv->UnsafeToggle(subset, false);
  }
  return chosen_indices;
}

}  // namespace operations_research
