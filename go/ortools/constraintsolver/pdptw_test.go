package constraintsolver

import (
	"fmt"
	"testing"
)

type DataModelPDPTW struct {
	timeMatrix       [][]int64
	timeWindows      [][]int64
	pickupDeliveries [][]int
	numVehicles      int
	depot            int
}

////////////////////////////////////////////////////////////////////////////////

func TestConstraintSolver_PDPTW(t *testing.T) {

	// Solvable pickup deliveries
	solvablePickupDeliveries := [][]int{
		{6, 2},
		{9, 14},
		{5, 10},
	}

	// Instantiate the data problem.
	data := DataModelPDPTW{
		timeMatrix:       timeMatrix,
		timeWindows:      timeWindows,
		pickupDeliveries: solvablePickupDeliveries,
		numVehicles:      4,
		depot:            0,
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
		int64(30),  // allow waiting time
		int64(360), // maximum time per vehicle
		false,      // Don't force start cumul to zero
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

	// Define Transportation Requests.
	solver := routing.Solver()
	for _, request := range data.pickupDeliveries {
		pickupIndex := manager.NodeToIndex(request[0])
		deliveryIndex := manager.NodeToIndex(request[1])
		routing.AddPickupAndDelivery(pickupIndex, deliveryIndex)
		solver.AddConstraint(
			solver.MakeEquality(
				SwigcptrIntExpr(routing.VehicleVar(pickupIndex).Swigcptr()),
				SwigcptrIntExpr(routing.VehicleVar(deliveryIndex).Swigcptr())))
		solver.AddConstraint(
			solver.MakeLessOrEqual(
				SwigcptrIntExpr(timeDimension.CumulVar(pickupIndex).Swigcptr()),
				SwigcptrIntExpr(timeDimension.CumulVar(deliveryIndex).Swigcptr())))
	}

	// Setting first solution heuristic.
	searchParameters := DefaultRoutingSearchParameters()
	// searchParameters.TimeLimit = &duration.Duration{Seconds: 15}
	// searchParameters.SolutionLimit = 10
	searchParameters.FirstSolutionStrategy = FirstSolutionStrategy_PATH_CHEAPEST_ARC

	// Solve the problem.
	solution := routing.SolveWithParameters(searchParameters)

	// Print solution on console.
	printSolutionPDPTW(t, data, manager, routing, solution)
}

func TestConstraintSolver_PDPTW_SingleVehicle(t *testing.T) {

	// 0. Route End

	// 1. Vehicle Loc
	// 1,1 : current vehicle loc - 0 mins
	// 1,2 : do pkg A - 7 mins
	// 1,3 : pu pkg B - 4 mins
	// 1,4 : do pkg B - 2 mins

	// 2. Drop-off A (due in 17 mins)
	// 2,1 : current vehicle loc - 7 mins
	// 2,2 : do pkg A - 0 mins
	// 2,3 : pu pkg B - 5 mins
	// 2,4 : do pkg B - 9 mins

	// 3. Pick-up B (due in 5 mins)
	// 3,1 : current vehicle loc - 4 mins
	// 3,2 : do pkg A - 5 mins
	// 3,3 : pu pkg B - 0 mins
	// 3,4 : do pkg B - 3 mins

	// 4. Drop-off B (due in 10 mins)
	// 4,1 : current vehicle loc - 2 mins
	// 4,2 : do pkg A - 9 mins
	// 4,3 : pu pkg B - 3 mins
	// 4,4 : do pkg B - 0 mins

	svTimeMatrix := [][]int64{
		{0, 0, 0, 0, 0},
		{0, 0, 7, 4, 2},
		{0, 7, 0, 5, 9},
		{0, 4, 5, 0, 3},
		{0, 2, 9, 3, 0},
	}

	svTimeWindows := [][]int64{
		{0, 0},
		{0, 0},
		{1, 17},
		{1, 5},
		{1, 10},
	}

	svPickupDeliveries := [][]int{
		{3, 4},
	}

	// Instantiate the data problem.
	data := DataModelPDPTW{
		timeMatrix:       svTimeMatrix,
		timeWindows:      svTimeWindows,
		pickupDeliveries: svPickupDeliveries,
		numVehicles:      1,
		depot:            1,
	}

	// Create Routing Index Manager
	starts := []int{1}
	ends := []int{0}
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
		true,      // Force start cumul to zero
		"Time")
	timeDimension := routing.GetDimensionOrDie("Time")

	// Add time window constraints for each location except depot and cur
	// vehicle location.
	for i := 2; i < len(data.timeWindows); i++ {
		index := manager.NodeToIndex(i)
		timeDimension.CumulVar(index).SetRange(data.timeWindows[i][0],
			data.timeWindows[i][1])
		// timeDimension.CumulVar(index).SetMax(data.timeWindows[i][1])
	}
	// Add time window constraints for each vehicle start node.
	// for i := 0; i < data.numVehicles; i++ {
	// 	index := routing.Start(i)
	// 	timeDimension.CumulVar(index).SetRange(data.timeWindows[0][0],
	// 		data.timeWindows[0][1])
	// }

	// Instantiate route start and end times to produce feasible times.
	for i := 0; i < data.numVehicles; i++ {
		routing.AddVariableMinimizedByFinalizer(
			timeDimension.CumulVar(routing.Start(i)))
		routing.AddVariableMinimizedByFinalizer(
			timeDimension.CumulVar(routing.End(i)))
	}

	// Define Transportation Requests.
	solver := routing.Solver()
	for _, request := range data.pickupDeliveries {
		pickupIndex := manager.NodeToIndex(request[0])
		deliveryIndex := manager.NodeToIndex(request[1])
		routing.AddPickupAndDelivery(pickupIndex, deliveryIndex)
		solver.AddConstraint(
			solver.MakeEquality(
				SwigcptrIntExpr(routing.VehicleVar(pickupIndex).Swigcptr()),
				SwigcptrIntExpr(routing.VehicleVar(deliveryIndex).Swigcptr())))
		solver.AddConstraint(
			solver.MakeLessOrEqual(
				SwigcptrIntExpr(timeDimension.CumulVar(pickupIndex).Swigcptr()),
				SwigcptrIntExpr(timeDimension.CumulVar(deliveryIndex).Swigcptr())))
	}

	// Setting first solution heuristic.
	searchParameters := DefaultRoutingSearchParameters()
	searchParameters.FirstSolutionStrategy = FirstSolutionStrategy_PARALLEL_CHEAPEST_INSERTION

	// Solve the problem.
	solution := routing.SolveWithParameters(searchParameters)

	// Print solution on console.
	printSolutionPDPTW(t, data, manager, routing, solution)
}

////////////////////////////////////////////////////////////////////////////////

// Print the solution.
// data: Data of the problem.
// manager: Index manager used.
// routing: Routing solver used.
// solution: Solution found by the solver.
func printSolutionPDPTW(t *testing.T, data DataModelPDPTW, manager RoutingIndexManager,
	routing RoutingModel, solution Assignment) {

	t.Logf("Solver status: %v", routing.GetStatus())

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
