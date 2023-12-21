if(NOT BUILD_GO_EXAMPLES)
	return()
endif()

file(GLOB_RECURSE GO_SRCS "constraintsolver/*.go")
foreach(SAMPLE IN LISTS GO_SRCS)
  add_go_example(${SAMPLE})
endforeach()
