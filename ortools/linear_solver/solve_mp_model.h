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

// Functions for solving optimization models defined by MPModelRequest.
//
// See linear_solver.proto for further documentation.

#ifndef OR_TOOLS_LINEAR_SOLVER_SOLVE_MP_MODEL_H_
#define OR_TOOLS_LINEAR_SOLVER_SOLVE_MP_MODEL_H_

#include <atomic>
#include <string>

#include "ortools/linear_solver/linear_solver.pb.h"

namespace operations_research {

/**
 * Solves the model encoded by a MPModelRequest protocol buffer and returns the
 * solution encoded as a MPSolutionResponse. The solve is stopped prematurely
 * if interrupt is non-null at set to true during (or before) solving.
 * Interruption is only supported if SolverTypeSupportsInterruption() returns
 * true for the requested solver. Passing a non-null interruption with any
 * other solver type immediately returns an MPSOLVER_INCOMPATIBLE_OPTIONS
 * error.
 */
MPSolutionResponse SolveMPModel(
    const MPModelRequest& model_request,
    // `interrupt` is non-const because the internal
    // solver may set it to true itself, in some cases.
    std::atomic<bool>* interrupt = nullptr);

bool SolverTypeSupportsInterruption(MPModelRequest::SolverType solver);

// Gives some brief (a few lines, at most) human-readable information about
// the given request, suitable for debug logging.
std::string MPModelRequestLoggingInfo(const MPModelRequest& request);

}  // namespace operations_research

#endif  // OR_TOOLS_LINEAR_SOLVER_SOLVE_MP_MODEL_H_
