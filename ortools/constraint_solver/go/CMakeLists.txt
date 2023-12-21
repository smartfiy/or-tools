set_property(SOURCE routing.i PROPERTY CPLUSPLUS ON)
set_property(SOURCE routing.i PROPERTY SWIG_MODULE_NAME gowrap_constraint_solver)
set_property(SOURCE routing.i PROPERTY COMPILE_DEFINITIONS
  ${OR_TOOLS_COMPILE_DEFINITIONS} ABSL_MUST_USE_RESULT=)
set_property(SOURCE routing.i PROPERTY COMPILE_OPTIONS
  -package constraintsolver -use-shlib)
swig_add_library(go_constraint_solver
  TYPE OBJECT
  LANGUAGE go
  OUTPUT_DIR ./wrap
  SOURCES routing.i)

target_include_directories(go_constraint_solver PRIVATE ${GO_INCLUDE_DIRS})
set_target_properties(go_constraint_solver PROPERTIES
  SWIG_USE_TARGET_INCLUDE_DIRECTORIES ON
  POSITION_INDEPENDENT_CODE ON)
target_link_libraries(go_constraint_solver PRIVATE ortools::ortools)
