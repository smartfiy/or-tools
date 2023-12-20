# Constraint Solver Go examples
The following examples showcase how to use the `constraintsolver` go pkg.

## Examples list
- Routing examples:
  - vrppd_test.go Capacitated Vehicle Routing Problem with Pickup and Delivery.
  - vrptw_test.go Capacitated Vehicle Routing Problem with Time Windows.
  - pdptw_test.go Pickup and Delivery Problem with Time Windows.

# Execution
Running the examples will involve building them, then running them.
You can run the following command from the **top** directory after building Go
examples via cmake:
```shell
make go_constraintsolver_vrppd_test
make go_constraintsolver_vrptw_test
make go_constraintsolver_pdptw_test
```
