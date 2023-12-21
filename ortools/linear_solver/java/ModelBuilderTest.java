// Copyright 2010-2022 Google LLC
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

package com.google.ortools.modelbuilder;

import static com.google.common.truth.Truth.assertThat;

import com.google.ortools.Loader;
import java.time.Duration;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public final class ModelBuilderTest {
  @BeforeEach
  public void setUp() {
    Loader.loadNativeLibraries();
  }

  @Test
  public void testEnforcedLinearApi() {
    ModelBuilder model = new ModelBuilder();
    model.setName("minimal enforced linear test");
    double infinity = Double.POSITIVE_INFINITY;
    Variable x = model.newNumVar(0.0, infinity, "x");
    Variable y = model.newNumVar(0.0, infinity, "y");
    Variable z = model.newBoolVar("z");

    assertThat(model.numVariables()).isEqualTo(3);

    EnforcedLinearConstraint c0 = model.addEnforcedGreaterOrEqual(
        LinearExpr.newBuilder().add(x).addTerm(y, 2.0), 10.0, z, false);
    assertThat(c0.getLowerBound()).isEqualTo(10.0);
    assertThat(c0.getIndicatorVariable().getIndex()).isEqualTo(z.getIndex());
    assertThat(c0.getIndicatorValue()).isFalse();
  }

  @Test
  public void runMinimalLinearExample_ok() {
    final String name = "minimal_linear_example";
    ModelBuilder model = new ModelBuilder();
    model.setName(name);
    double infinity = Double.POSITIVE_INFINITY;
    Variable x1 = model.newNumVar(0.0, infinity, "x1");
    Variable x2 = model.newNumVar(0.0, infinity, "x2");
    Variable x3 = model.newNumVar(0.0, infinity, "x3");

    assertThat(model.numVariables()).isEqualTo(3);
    assertThat(x1.getIntegrality()).isFalse();
    assertThat(x1.getLowerBound()).isEqualTo(0.0);
    assertThat(x2.getUpperBound()).isEqualTo(infinity);
    x1.setLowerBound(1.0);
    assertThat(x1.getLowerBound()).isEqualTo(1.0);

    LinearConstraint c0 = model.addLessOrEqual(LinearExpr.sum(new Variable[] {x1, x2, x3}), 100.0);
    assertThat(c0.getUpperBound()).isEqualTo(100.0);
    LinearConstraint c1 =
        model
            .addLessOrEqual(
                LinearExpr.newBuilder().addTerm(x1, 10.0).addTerm(x2, 4.0).addTerm(x3, 5.0), 600.0)
            .withName("c1");
    assertThat(c1.getName()).isEqualTo("c1");
    LinearConstraint c2 = model.addLessOrEqual(
        LinearExpr.newBuilder().addTerm(x1, 2.0).addTerm(x2, 2.0).addTerm(x3, 6.0), 300.0);
    assertThat(c2.getUpperBound()).isEqualTo(300.0);

    model.maximize(
        LinearExpr.weightedSum(new Variable[] {x1, x2, x3}, new double[] {10.0, 6, 4.0}));
    assertThat(x3.getObjectiveCoefficient()).isEqualTo(4.0);
    assertThat(model.getObjectiveOffset()).isEqualTo(0.0);
    model.setObjectiveOffset(-5.5);
    assertThat(model.getObjectiveOffset()).isEqualTo(-5.5);

    ModelSolver solver = new ModelSolver("glop");
    assertThat(solver.solverIsSupported()).isTrue();
    solver.setTimeLimit(Duration.ofSeconds(1));
    assertThat(solver.solve(model)).isEqualTo(SolveStatus.OPTIMAL);

    assertThat(solver.getObjectiveValue())
        .isWithin(1e-5)
        .of(733.333333 + model.getObjectiveOffset());
    assertThat(solver.getValue(x1)).isWithin(1e-5).of(33.333333);
    assertThat(solver.getValue(x2)).isWithin(1e-5).of(66.6666673);
    assertThat(solver.getValue(x3)).isWithin(1e-5).of(0.0);

    double dualObjectiveValue = solver.getDualValue(c0) * c0.getUpperBound()
        + solver.getDualValue(c1) * c1.getUpperBound()
        + solver.getDualValue(c2) * c2.getUpperBound() + model.getObjectiveOffset();
    assertThat(solver.getObjectiveValue()).isWithin(1e-5).of(dualObjectiveValue);

    assertThat(solver.getReducedCost(x1)).isWithin(1e-5).of(0.0);
    assertThat(solver.getReducedCost(x2)).isWithin(1e-5).of(0.0);
    assertThat(solver.getReducedCost(x3))
        .isWithin(1e-5)
        .of(4.0 - 1.0 * solver.getDualValue(c0) - 5.0 * solver.getDualValue(c1));

    assertThat(solver.getActivity(c0)).isWithin(1e-5).of(100.0);
    assertThat(solver.getActivity(c1)).isWithin(1e-5).of(600.0);
    assertThat(solver.getActivity(c2)).isWithin(1e-5).of(200.0);

    assertThat(model.exportToLpString(false)).contains(name);
    assertThat(model.exportToMpsString(false)).contains(name);
  }

  @Test
  public void importFromMpsString() {
    ModelBuilder model = new ModelBuilder();
    String mpsData = "* Generated by MPModelProtoExporter\n"
        + "*   Name             : SupportedMaximizationProblem\n"
        + "*   Format           : Free\n"
        + "*   Constraints      : 0\n"
        + "*   Variables        : 1\n"
        + "*     Binary         : 0\n"
        + "*     Integer        : 0\n"
        + "*     Continuous     : 1\n"
        + "NAME          SupportedMaximizationProblem\n"
        + "OBJSENSE\n"
        + "  MAX\n"
        + "ROWS\n"
        + " N  COST\n"
        + "COLUMNS\n"
        + "    X_ONE   COST         1\n"
        + "BOUNDS\n"
        + " UP BOUND   X_ONE        4\n"
        + "ENDATA";
    assertThat(model.importFromMpsString(mpsData)).isTrue();
    assertThat(model.getName()).isEqualTo("SupportedMaximizationProblem");
  }

  @Test
  public void importFromLpString() {
    ModelBuilder model = new ModelBuilder();
    String lpData = "min: x + y;\n"
        + "bin: b1, b2, b3;\n"
        + "1 <= x <= 42;\n"
        + "constraint_num1: 5 b1 + 3b2 + x <= 7;\n"
        + "4 y + b2 - 3 b3 <= 2;\n"
        + "constraint_num2: -4 b1 + b2 - 3 z <= -2;\n";
    assertThat(model.importFromLpString(lpData)).isTrue();
    assertThat(model.numVariables()).isEqualTo(6);
    assertThat(model.numConstraints()).isEqualTo(3);
    assertThat(model.varFromIndex(0).getLowerBound()).isEqualTo(1.0);
    assertThat(model.varFromIndex(0).getUpperBound()).isEqualTo(42.0);
    assertThat(model.varFromIndex(0).getName()).isEqualTo("x");
  }
}
