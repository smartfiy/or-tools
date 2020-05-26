package constraintsolver

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
// https://developers.google.com/optimization/routing/vrptw
////////////////////////////////////////////////////////////////////////////////

var (
	timeMatrix = [][]int64{
		{0, 6, 9, 8, 7, 3, 6, 2, 3, 2, 6, 6, 4, 4, 5, 9, 7},
		{6, 0, 8, 3, 2, 6, 8, 4, 8, 8, 13, 7, 5, 8, 12, 10, 14},
		{9, 8, 0, 11, 10, 6, 3, 9, 5, 8, 4, 15, 14, 13, 9, 18, 9},
		{8, 3, 11, 0, 1, 7, 10, 6, 10, 10, 14, 6, 7, 9, 14, 6, 16},
		{7, 2, 10, 1, 0, 6, 9, 4, 8, 9, 13, 4, 6, 8, 12, 8, 14},
		{3, 6, 6, 7, 6, 0, 2, 3, 2, 2, 7, 9, 7, 7, 6, 12, 8},
		{6, 8, 3, 10, 9, 2, 0, 6, 2, 5, 4, 12, 10, 10, 6, 15, 5},
		{2, 4, 9, 6, 4, 3, 6, 0, 4, 4, 8, 5, 4, 3, 7, 8, 10},
		{3, 8, 5, 10, 8, 2, 2, 4, 0, 3, 4, 9, 8, 7, 3, 13, 6},
		{2, 8, 8, 10, 9, 2, 5, 4, 3, 0, 4, 6, 5, 4, 3, 9, 5},
		{6, 13, 4, 14, 13, 7, 4, 8, 4, 4, 0, 10, 9, 8, 4, 13, 4},
		{6, 7, 15, 6, 4, 9, 12, 5, 9, 6, 10, 0, 1, 3, 7, 3, 10},
		{4, 5, 14, 7, 6, 7, 10, 4, 8, 5, 9, 1, 0, 2, 6, 4, 8},
		{4, 8, 13, 9, 8, 7, 10, 3, 7, 4, 8, 3, 2, 0, 4, 5, 6},
		{5, 12, 9, 14, 12, 6, 6, 7, 3, 3, 4, 7, 6, 4, 0, 9, 2},
		{9, 10, 18, 6, 8, 12, 15, 8, 13, 9, 13, 3, 4, 5, 9, 0, 9},
		{7, 14, 9, 16, 14, 8, 5, 10, 6, 5, 4, 10, 8, 6, 2, 9, 0},
	}

	timeWindows = [][]int64{
		{0, 5},   // depot
		{7, 12},  // 1
		{10, 15}, // 2
		{16, 18}, // 3
		{10, 13}, // 4
		{0, 5},   // 5
		{5, 10},  // 6
		{0, 4},   // 7
		{5, 10},  // 8
		{0, 3},   // 9
		{10, 16}, // 10
		{10, 15}, // 11
		{0, 5},   // 12
		{5, 10},  // 13
		{7, 8},   // 14
		{10, 15}, // 15
		{11, 15}, // 16
	}
)

type DataModelVRPTW struct {
	timeMatrix  [][]int64
	timeWindows [][]int64
	numVehicles int
	depot       int
}

////////////////////////////////////////////////////////////////////////////////

func TestConstraintSolver_VRPTW(t *testing.T) {
	// Instantiate the data problem.
	data := DataModelVRPTW{
		timeMatrix:  timeMatrix,
		timeWindows: timeWindows,
		numVehicles: 4,
		depot:       0,
	}

	// Create Routing Index Manager
	starts := []int{0, 0, 0, 0}
	ends := []int{0, 0, 0, 0}
	manager := NewRoutingIndexManager(len(data.timeMatrix), data.numVehicles,
		starts, ends)

	// Create Routing Model.
	routing := NewRoutingModel(manager, DefaultRoutingModelParameters())

	// Create and register a transit callback.
	f := func(fromIndex, toIndex int64) int64 {
		// Convert from routing variable Index to time matrix NodeIndex.
		fromNode := manager.IndexToNode(fromIndex)
		toNode := manager.IndexToNode(toIndex)
		return data.timeMatrix[fromNode][toNode]
	}
	w := NewGoRoutingTransitCallback2Wrapper(f)
	defer w.Delete()
	transitCallbackIndex := routing.RegisterTransitCallback(w.Wrap())

	// Define cost of each arc.
	routing.SetArcCostEvaluatorOfAllVehicles(transitCallbackIndex)

	// Add Time constraint.
	routing.AddDimension(transitCallbackIndex, // transit callback index
		int64(30), // allow waiting time
		int64(30), // maximum time per vehicle
		false,     // Don't force start cumul to zero
		"Time")
	timeDimension := routing.GetDimensionOrDie("Time")

	// Add time window constraints for each location except depot.
	for i := 1; i < len(data.timeWindows); i++ {
		index := manager.NodeToIndex(i)
		timeDimension.CumulVar(index).SetRange(data.timeWindows[i][0],
			data.timeWindows[i][1])
	}
	// Add time window constraints for each vehicle start node.
	for i := 0; i < data.numVehicles; i++ {
		index := routing.Start(i)
		timeDimension.CumulVar(index).SetRange(data.timeWindows[0][0],
			data.timeWindows[0][1])
	}

	// Instantiate route start and end times to produce feasible times.
	for i := 0; i < data.numVehicles; i++ {
		routing.AddVariableMinimizedByFinalizer(
			timeDimension.CumulVar(routing.Start(i)))
		routing.AddVariableMinimizedByFinalizer(
			timeDimension.CumulVar(routing.End(i)))
	}

	// Setting first solution heuristic.
	searchParameters := DefaultRoutingSearchParameters()
	searchParameters.FirstSolutionStrategy = FirstSolutionStrategy_PATH_CHEAPEST_ARC

	// Solve the problem.
	solution := routing.SolveWithParameters(searchParameters)

	// Print solution on console.
	printSolutionVRPTW(t, data, manager, routing, solution)
}

////////////////////////////////////////////////////////////////////////////////

// Print the solution.
// data: Data of the problem.
// manager: Index manager used.
// routing: Routing solver used.
// solution: Solution found by the solver.
func printSolutionVRPTW(t *testing.T, data DataModelVRPTW, manager RoutingIndexManager,
	routing RoutingModel, solution Assignment) {
	timeDimension := routing.GetDimensionOrDie("Time")
	var totalTime int64
	for vehicleId := 0; vehicleId < data.numVehicles; vehicleId++ {
		index := routing.Start(vehicleId)
		t.Logf("Route for vehicle %v:", vehicleId)
		var route string
		for !routing.IsEnd(index) {
			timeVar := timeDimension.CumulVar(index)
			route += fmt.Sprintf("%v Time(%v, %v) -> ", manager.IndexToNode(index), solution.Min(timeVar), solution.Max(timeVar))
			index = solution.Value(routing.NextVar(index))
		}
		timeVar := timeDimension.CumulVar(index)
		t.Logf("%v%v Time(%v, %v)", route, manager.IndexToNode(index), solution.Min(timeVar), solution.Max(timeVar))
		t.Logf("Time of the route: %vmin", solution.Min(timeVar))
		t.Log()
		totalTime += solution.Min(timeVar)
	}
	t.Logf("Total time of all routes: %vmin", totalTime)
	t.Log()
	t.Log("Advanced usage:")
	t.Logf("Problem solved in %vms", routing.Solver().WallTime())
}
