package constraintsolver

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
// https://developers.google.com/optimization/routing/pickup_delivery
////////////////////////////////////////////////////////////////////////////////

var (
	distanceMatrix = [][]int64{
		{0, 548, 776, 696, 582, 274, 502, 194, 308, 194, 536, 502, 388, 354, 468, 776, 662},
		{548, 0, 684, 308, 194, 502, 730, 354, 696, 742, 1084, 594, 480, 674, 1016, 868, 1210},
		{776, 684, 0, 992, 878, 502, 274, 810, 468, 742, 400, 1278, 1164, 1130, 788, 1552, 754},
		{696, 308, 992, 0, 114, 650, 878, 502, 844, 890, 1232, 514, 628, 822, 1164, 560, 1358},
		{582, 194, 878, 114, 0, 536, 764, 388, 730, 776, 1118, 400, 514, 708, 1050, 674, 1244},
		{274, 502, 502, 650, 536, 0, 228, 308, 194, 240, 582, 776, 662, 628, 514, 1050, 708},
		{502, 730, 274, 878, 764, 228, 0, 536, 194, 468, 354, 1004, 890, 856, 514, 1278, 480},
		{194, 354, 810, 502, 388, 308, 536, 0, 342, 388, 730, 468, 354, 320, 662, 742, 856},
		{308, 696, 468, 844, 730, 194, 194, 342, 0, 274, 388, 810, 696, 662, 320, 1084, 514},
		{194, 742, 742, 890, 776, 240, 468, 388, 274, 0, 342, 536, 422, 388, 274, 810, 468},
		{536, 1084, 400, 1232, 1118, 582, 354, 730, 388, 342, 0, 878, 764, 730, 388, 1152, 354},
		{502, 594, 1278, 514, 400, 776, 1004, 468, 810, 536, 878, 0, 114, 308, 650, 274, 844},
		{388, 480, 1164, 628, 514, 662, 890, 354, 696, 422, 764, 114, 0, 194, 536, 388, 730},
		{354, 674, 1130, 822, 708, 628, 856, 320, 662, 388, 730, 308, 194, 0, 342, 422, 536},
		{468, 1016, 788, 1164, 1050, 514, 514, 662, 320, 274, 388, 650, 536, 342, 0, 764, 194},
		{776, 868, 1552, 560, 674, 1050, 1278, 742, 1084, 810, 1152, 274, 388, 422, 764, 0, 798},
		{662, 1210, 754, 1358, 1244, 708, 480, 856, 514, 468, 354, 844, 730, 536, 194, 798, 0},
	}

	pickupDeliveries = [][]int{
		{1, 6},
		{2, 10},
		{4, 3},
		{5, 9},
		{7, 8},
		{15, 11},
		{13, 12},
		{16, 14},
	}
)

type DataModelVRPPD struct {
	distanceMatrix   [][]int64
	pickupDeliveries [][]int
	numVehicles      int
	depot            int
}

////////////////////////////////////////////////////////////////////////////////

func TestConstraintSolver_VRPPD(t *testing.T) {
	// Instantiate the data problem.
	data := DataModelVRPPD{
		distanceMatrix:   distanceMatrix,
		pickupDeliveries: pickupDeliveries,
		numVehicles:      4,
		depot:            0,
	}

	// Create Routing Index Manager
	starts := []int{0, 0, 0, 0}
	ends := []int{0, 0, 0, 0}
	manager := NewRoutingIndexManager(len(data.distanceMatrix), data.numVehicles,
		starts, ends)

	// Create Routing Model.
	routing := NewRoutingModel(manager, DefaultRoutingModelParameters())

	// Define cost of each arc.
	f := func(fromIndex, toIndex int64) int64 {
		// Convert from routing variable Index to distance matrix NodeIndex.
		fromNode := manager.IndexToNode(fromIndex)
		toNode := manager.IndexToNode(toIndex)
		return data.distanceMatrix[fromNode][toNode]
	}
	w := NewGoRoutingTransitCallback2Wrapper(f)
	defer w.Delete()
	transitCallbackIndex := routing.RegisterTransitCallback(w.Wrap())
	routing.SetArcCostEvaluatorOfAllVehicles(transitCallbackIndex)

	// Add Distance constraint.
	routing.AddDimension(transitCallbackIndex, // transit callback
		0,    // no slack
		3000, // vehicle maximum travel distance
		true, // start cumul to zero
		"Distance")
	distanceDimension := routing.GetMutableDimension("Distance")
	distanceDimension.SetGlobalSpanCostCoefficient(100)

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
				SwigcptrIntExpr(distanceDimension.CumulVar(pickupIndex).Swigcptr()),
				SwigcptrIntExpr(distanceDimension.CumulVar(deliveryIndex).Swigcptr())))
	}

	// Setting first solution heuristic.
	searchParameters := DefaultRoutingSearchParameters()
	searchParameters.FirstSolutionStrategy = FirstSolutionStrategy_PARALLEL_CHEAPEST_INSERTION

	// Solve the problem.
	solution := routing.SolveWithParameters(searchParameters)

	// Print solution on console.
	printSolutionVRPPD(t, data, manager, routing, solution)
}

////////////////////////////////////////////////////////////////////////////////

// Print the solution.
// data: Data of the problem.
// manager: Index manager used.
// routing: Routing solver used.
// solution: Solution found by the solver.
func printSolutionVRPPD(t *testing.T, data DataModelVRPPD, manager RoutingIndexManager,
	routing RoutingModel, solution Assignment) {
	var totalDistance int64
	for vehicleId := 0; vehicleId < data.numVehicles; vehicleId++ {
		index := routing.Start(vehicleId)
		t.Logf("Route for vehicle %v:", vehicleId)
		var routeDistance int64
		var route string
		for !routing.IsEnd(index) {
			route += fmt.Sprintf("%v -> ", manager.IndexToNode(index))
			previousIndex := index
			index = solution.Value(routing.NextVar(index))
			routeDistance += routing.GetArcCostForVehicle(previousIndex, index, int64(vehicleId))
		}
		t.Logf("%v%v", route, manager.IndexToNode(index))
		t.Logf("Distance of the route: %vm", routeDistance)
		totalDistance += routeDistance
	}
	t.Logf("Total distance of all routes: %vm", totalDistance)
	t.Log()
	t.Log("Advanced usage:")
	t.Logf("Problem solved in %vms", routing.Solver().WallTime())
}
