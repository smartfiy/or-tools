# ---------- Go support using SWIG ----------
.PHONY: help_go # Generate list of Go targets with descriptions.
help_go:
	@echo Use one of the following Go targets:
ifeq ($(SYSTEM),win)
	@$(GREP) "^.PHONY: .* #" $(CURDIR)/makefiles/Makefile.go.mk | $(SED) "s/\.PHONY: \(.*\) # \(.*\)/\1\t\2/"
	@echo off & echo(
else
	@$(GREP) "^.PHONY: .* #" $(CURDIR)/makefiles/Makefile.go.mk | $(SED) "s/\.PHONY: \(.*\) # \(.*\)/\1\t\2/" | expand -t24
	@echo
endif

# Check for required build tools
GO_BIN = $(shell $(WHICH) go)
GO_PATH = $(shell [ -z "${GOPATH}" ] || echo $(GOPATH))
GO_OR_TOOLS_NATIVE_LIB := goortools
PROTOC_GEN_GO = $(shell $(WHICH) protoc-gen-go)
PKG_ROOT = github.com/airspacetechnologies/or-tools/go

HAS_GO = true
ifndef GO_BIN
HAS_GO =
endif
ifndef GO_PATH
HAS_GO =
endif

# Main target
.PHONY: go # Build Go OR-Tools.
.PHONY: test_go # Test Go OR-Tools using various examples.
ifndef HAS_GO
go: detect_go
check_go: go
test_go: go
package_go: go
else
go: go_pimpl
check_go: check_go_pimpl
test_go: test_go_pimpl
package_go: go_pimpl
BUILT_LANGUAGES +=, Golang
endif

# Go gen dirs (go only wants go source in its tree)
$(GEN_DIR)/ortools/go/constraintsolver:
	-$(MKDIR_P) $(GEN_DIR)/ortools/go/constraintsolver
$(GEN_DIR)/ortools/go/util:
	-$(MKDIR_P) $(GEN_DIR)/ortools/go/util

#################################################################
## Protobuf: generate Go files from proto specification files. ##
#################################################################
.PHONY: proto
proto: \
 $(GEN_DIR)/ortools/go/constraintsolver \
 $(GEN_DIR)/ortools/go/constraintsolver/gowrap_constraint_solver.go \
 $(GEN_DIR)/ortools/go/constraintsolver/search_limit.pb.go \
 $(GEN_DIR)/ortools/go/constraintsolver/solver_parameters.pb.go \
 $(GEN_DIR)/ortools/go/constraintsolver/routing_parameters.pb.go \
 $(GEN_DIR)/ortools/go/constraintsolver/routing_enums.pb.go \
 $(GEN_DIR)/ortools/go/util \
 $(GEN_DIR)/ortools/go/util/optional_boolean.pb.go

$(GEN_DIR)/ortools/go/constraintsolver/search_limit.pb.go: \
 $(SRC_DIR)/ortools/constraint_solver/search_limit.proto \
 | $(GEN_DIR)/ortools/go/constraintsolver
	$(PROTOC) --proto_path=$(SRC_DIR) \
 --go_out=$(GEN_DIR) \
 --go_opt=paths=source_relative \
 $(SRC_DIR)$Sortools$Sconstraint_solver$Ssearch_limit.proto
	$(RENAME) $(GEN_DIR)/ortools/constraint_solver/search_limit.pb.go \
	$(GEN_DIR)/ortools/go/constraintsolver

$(GEN_DIR)/ortools/go/constraintsolver/solver_parameters.pb.go: \
 $(SRC_DIR)/ortools/constraint_solver/solver_parameters.proto \
 | $(GEN_DIR)/ortools/go/constraintsolver
	$(PROTOC) --proto_path=$(SRC_DIR) \
 --go_out=$(GEN_DIR) \
 --go_opt=paths=source_relative \
 $(SRC_DIR)$Sortools$Sconstraint_solver$Ssolver_parameters.proto
	$(RENAME) $(GEN_DIR)/ortools/constraint_solver/solver_parameters.pb.go \
	$(GEN_DIR)/ortools/go/constraintsolver

$(GEN_DIR)/ortools/go/constraintsolver/routing_parameters.pb.go: \
 $(SRC_DIR)/ortools/constraint_solver/routing_parameters.proto \
 | $(GEN_DIR)/ortools/go/constraintsolver
	$(PROTOC) --proto_path=$(SRC_DIR) \
 --go_out=$(GEN_DIR) \
 --go_opt=paths=source_relative \
 $(SRC_DIR)$Sortools$Sconstraint_solver$Srouting_parameters.proto
	$(RENAME) $(GEN_DIR)/ortools/constraint_solver/routing_parameters.pb.go \
	$(GEN_DIR)/ortools/go/constraintsolver

$(GEN_DIR)/ortools/go/constraintsolver/routing_enums.pb.go: \
 $(SRC_DIR)/ortools/constraint_solver/routing_enums.proto \
 | $(GEN_DIR)/ortools/go/constraintsolver
	$(PROTOC) --proto_path=$(SRC_DIR) \
 --go_out=$(GEN_DIR) \
 --go_opt=paths=source_relative \
 $(SRC_DIR)$Sortools$Sconstraint_solver$Srouting_enums.proto
	$(RENAME) $(GEN_DIR)/ortools/constraint_solver/routing_enums.pb.go \
	$(GEN_DIR)/ortools/go/constraintsolver

# $(GEN_DIR)/ortools/go/sat/cp_model.pb.go: \
#  $(SRC_DIR)/ortools/sat/cp_model.proto \
#  | $(GEN_DIR)/ortools/go/sat
# 	$(PROTOC) --proto_path=$(SRC_DIR) \
#  --go_out=$(GEN_DIR) \
#  --go_opt=paths=source_relative \
#  $(SRC_DIR)$Sortools$Ssat$Scp_model.proto

# $(GEN_DIR)/ortools/sat/sat_parameters.pb.go: \
#  $(SRC_DIR)/ortools/sat/sat_parameters.proto \
#  | $(GEN_DIR)/ortools/sat
# 	$(PROTOC) --proto_path=$(SRC_DIR) \
#  --go_out=$(GEN_DIR) \
#  --go_opt=paths=source_relative \
#  $(SRC_DIR)$Sortools$Ssat$Ssat_parameters.proto

$(GEN_DIR)/ortools/go/util/optional_boolean.pb.go: \
 $(SRC_DIR)/ortools/util/optional_boolean.proto \
 | $(GEN_DIR)/ortools/go/util
	$(PROTOC) --proto_path=$(SRC_DIR) \
 --go_out=$(GEN_DIR) \
 --go_opt=paths=source_relative \
 $(SRC_DIR)$Sortools$Sutil$Soptional_boolean.proto
	$(RENAME) $(GEN_DIR)/ortools/util/optional_boolean.pb.go \
	$(GEN_DIR)/ortools/go/util

#######################################################################################
# Swig: Generate Go wrapper files from swig specification files and C wrapper files. ##
#######################################################################################
# $(GEN_DIR)/ortools/go/linear_solver/linear_solver_go_wrap.cc: \
#  $(SRC_DIR)/ortools/linear_solver/go/linear_solver.i \
#  $(SRC_DIR)/ortools/base/base.i \
#  $(SRC_DIR)/ortools/util/go/proto.i \
#  $(GLOP_DEPS) \
#  $(LP_DEPS) \
#  | $(GEN_DIR)/ortools/go/linear_solver
# 	$(SWIG_BINARY) $(SWIG_INC) -I$(INC_DIR) -c++ -go -cgo -intgosize 64 \
#  -o $(GEN_PATH)$Sortools$Slinear_solver$Slinear_solver_go_wrap.cc \
#  -module gowrap_linear_solver \
#  -package linearsolver \
#  -outdir $(GEN_PATH)$Sortools$Slinear_solver \
#  $(SRC_DIR)$Sortools$Slinear_solver$Sgo$Slinear_solver.i

# $(OBJ_DIR)/swig/linear_solver_go_wrap.$O: \
#  $(GEN_DIR)/ortools/go/linear_solver/linear_solver_go_wrap.cc \
#  | $(OBJ_DIR)/swig
# 	$(CCC) $(CFLAGS) \
#  -c $(GEN_PATH)$Sortools$Slinear_solver$Slinear_solver_go_wrap.cc \
#  $(OBJ_OUT)$(OBJ_DIR)$Sswig$Slinear_solver_go_wrap.$O

$(GEN_DIR)/ortools/go/constraintsolver/gowrap_constraint_solver.go: \
 $(GEN_DIR)/ortools/constraint_solver/constraint_solver_go_wrap.cc

$(GEN_DIR)/ortools/constraint_solver/constraint_solver_go_wrap.cc: \
 $(SRC_DIR)/ortools/constraint_solver/go/routing.i \
 $(SRC_DIR)/ortools/constraint_solver/go/constraint_solver.i \
 $(SRC_DIR)/ortools/base/base.i \
 $(SRC_DIR)/ortools/util/go/proto.i \
 $(CP_DEPS) \
 | $(GEN_DIR)/ortools/go/constraintsolver
	$(SWIG_BINARY) $(SWIG_INC) -I$(INC_DIR) -c++ -go -cgo -intgosize 64 \
 -o $(GEN_PATH)$Sortools$Sconstraint_solver$Sconstraint_solver_go_wrap.cc \
 -module gowrap_constraint_solver \
 -package constraintsolver \
 -use-shlib \
 -soname $(LIB_PREFIX)$(GO_OR_TOOLS_NATIVE_LIB).$(SWIG_GO_LIB_SUFFIX) \
 -outdir $(GEN_PATH)$Sortools$Sgo$Sconstraintsolver \
 $(SRC_DIR)$Sortools$Sconstraint_solver$Sgo$Srouting.i

$(OBJ_DIR)/swig/constraint_solver_go_wrap.$O: \
 $(GEN_DIR)/ortools/constraint_solver/constraint_solver_go_wrap.cc \
 | $(OBJ_DIR)/swig
	$(CCC) $(CFLAGS) \
 -c $(GEN_PATH)$Sortools$Sconstraint_solver$Sconstraint_solver_go_wrap.cc \
 $(OBJ_OUT)$(OBJ_DIR)$Sswig$Sconstraint_solver_go_wrap.$O

# $(GEN_DIR)/ortools/go/algorithms/knapsack_solver_go_wrap.cc: \
#  $(SRC_DIR)/ortools/algorithms/go/knapsack_solver.i \
#  $(SRC_DIR)/ortools/base/base.i \
#  $(SRC_DIR)/ortools/util/go/proto.i \
#  $(SRC_DIR)/ortools/algorithms/knapsack_solver.h \
#  | $(GEN_DIR)/ortools/go/algorithms
# 	$(SWIG_BINARY) $(SWIG_INC) -I$(INC_DIR) -c++ -go -cgo -intgosize 64 \
#  -o $(GEN_PATH)$Sortools$Sgo$Salgorithms$Sknapsack_solver_go_wrap.cc \
#  -module gowrap_algorithms \
#  -package algorithms \
#  -outdir $(GEN_PATH)$Sortools$Sgo$Salgorithms \
#  $(SRC_DIR)$Sortools$Salgorithms$Sgo$Sknapsack_solver.i


# $(OBJ_DIR)/swig/knapsack_solver_go_wrap.$O: \
#  $(GEN_DIR)/ortools/go/algorithms/knapsack_solver_go_wrap.cc \
#  | $(OBJ_DIR)/swig
# 	$(CCC) $(CFLAGS) \
#  -c $(GEN_PATH)$Sortools$Sgo$Salgorithms$Sknapsack_solver_go_wrap.cc \
#  $(OBJ_OUT)$(OBJ_DIR)$Sswig$Sknapsack_solver_go_wrap.$O

# $(GEN_DIR)/ortools/go/graph/graph_go_wrap.cc: \
#  $(SRC_DIR)/ortools/graph/go/graph.i \
#  $(SRC_DIR)/ortools/base/base.i \
#  $(SRC_DIR)/ortools/util/go/proto.i \
#  $(GRAPH_DEPS) \
#  | $(GEN_DIR)/ortools/go/graph
# 	$(SWIG_BINARY) $(SWIG_INC) -I$(INC_DIR) -c++ -go -cgo -intgosize 64 \
#  -o $(GEN_PATH)$Sortools$Sgo$Sgraph$Sgraph_go_wrap.cc \
#  -module gowrap_graph \
#  -package graph \
#  -outdir $(GEN_PATH)$Sortools$Sgo$Sgraph \
#  $(SRC_DIR)$Sortools$Sgraph$Sgo$Sgraph.i

# $(OBJ_DIR)/swig/graph_go_wrap.$O: \
#  $(GEN_DIR)/ortools/go/graph/graph_go_wrap.cc \
#  | $(OBJ_DIR)/swig
# 	$(CCC) $(CFLAGS) \
#  -c $(GEN_PATH)$Sortools$Sgo$Sgraph$Sgraph_go_wrap.cc \
#  $(OBJ_OUT)$(OBJ_DIR)$Sswig$Sgraph_go_wrap.$O

# $(GEN_DIR)/ortools/go/sat/sat_go_wrap.cc: \
#  $(SRC_DIR)/ortools/base/base.i \
#  $(SRC_DIR)/ortools/sat/go/sat.i \
#  $(SRC_DIR)/ortools/sat/swig_helper.h \
#  $(SAT_DEPS) \
#  | $(GEN_DIR)/ortools/go/sat
# 	$(SWIG_BINARY) $(SWIG_INC) -I$(INC_DIR) -c++ -go -cgo -intgosize 64 \
#  -o $(GEN_PATH)$Sortools$Sgo$Ssat$Ssat_go_wrap.cc \
#  -module gowrap_sat \
#  -package sat \
#  -outdir $(GEN_PATH)$Sortools$Sgo$Ssat \
#  $(SRC_DIR)$Sortools$Ssat$Sgo$Ssat.i

# $(OBJ_DIR)/swig/sat_go_wrap.$O: \
#  $(GEN_DIR)/ortools/go/sat/sat_go_wrap.cc \
#  | $(OBJ_DIR)/swig
# 	$(CCC) $(CFLAGS) \
#  -c $(GEN_PATH)$Sortools$Ssat$Sgo$Ssat_go_wrap.cc \
#  $(OBJ_OUT)$(OBJ_DIR)$Sswig$Ssat_go_wrap.$O

# $(GEN_DIR)/ortools/go/util/sorted_interval_list_go_wrap.cc: \
#  $(SRC_DIR)/ortools/base/base.i \
#  $(SRC_DIR)/ortools/util/go/sorted_interval_list.i \
#  $(SRC_DIR)/ortools/util/go/proto.i \
#  $(UTIL_DEPS) \
#  |  $(GEN_DIR)/ortools/go/util
# 	$(SWIG_BINARY) $(SWIG_INC) -I$(INC_DIR) -c++ -go -cgo -intgosize 64 \
#  -o $(GEN_PATH)$Sortools$Sutil$Sgo$Ssorted_interval_list_go_wrap.cc \
#  -module gowrap_util \
#  -package util \
#  -outdir $(GEN_PATH)$Sortools$Sgo$Sutil \
#  $(SRC_DIR)$Sortools$Sutil$Sgo$Ssorted_interval_list.i

# $(OBJ_DIR)/swig/sorted_interval_list_go_wrap.$O: \
#  $(GEN_DIR)/ortools/go/util/sorted_interval_list_go_wrap.cc \
#  | $(OBJ_DIR)/swig
# 	$(CCC) $(CFLAGS) \
#  -c $(GEN_PATH)$Sortools$Sgo$Sutil$Ssorted_interval_list_go_wrap.cc \
#  $(OBJ_OUT)$(OBJ_DIR)$Sswig$Ssorted_interval_list_go_wrap.$O

##########################################
# Swig: Generate native wrapper binary. ##
##########################################
$(LIB_DIR)/$(LIB_PREFIX)$(GO_OR_TOOLS_NATIVE_LIB).$(SWIG_GO_LIB_SUFFIX): \
 $(OR_TOOLS_LIBS) \
 $(OBJ_DIR)/swig/constraint_solver_go_wrap.$O \
 | $(LIB_DIR)
	$(DYNAMIC_LD) \
 $(LD_OUT)$(LIB_DIR)$S$(LIB_PREFIX)$(GO_OR_TOOLS_NATIVE_LIB).$(SWIG_GO_LIB_SUFFIX) \
 $(OBJ_DIR)$Sswig$Sconstraint_solver_go_wrap.$O \
 $(OR_TOOLS_LNK) \
 $(GO_OR_TOOLS_LDFLAGS)

###################
##  Go SOURCE  ##
###################
.PHONY: go_pimpl
go_pimpl: $(GEN_DIR)/ortools/go

$(GEN_DIR)/ortools/go/go.mod:
	cd $(GEN_DIR)/ortools/go && \
	$(GO_BIN) mod init $(PKG_ROOT)

$(GEN_DIR)/ortools/go: \
 $(LIB_DIR)/$(LIB_PREFIX)$(GO_OR_TOOLS_NATIVE_LIB).$(SWIG_GO_LIB_SUFFIX) \
 proto \
 $(GEN_DIR)/ortools/go/go.mod
	cd $(GEN_DIR)/ortools/go && \
	$(GO_BIN) build ./...

ifeq ($(SOURCE_SUFFIX),.go) # Those rules will be used if SOURCE contain a .go file
.PHONY: build # Build a Go program.
build: $(SOURCE) $(LIB_DIR)/$(LIB_PREFIX)$(GO_OR_TOOLS_NATIVE_LIB).$(SWIG_GO_LIB_SUFFIX)
	$(GO_BIN) build $(SOURCE_PATH)

.PHONY: run # Run a Go program.
run: build
	cd $(GEN_DIR)/ortools/go && \
	$(GO_BIN) run $(SOURCE_PATH) $(ARGS)
endif

#############################
##  Go Examples/Samples  ##
#############################
.PHONY: check_go_pimpl
check_go_pimpl:

.PHONY: copy_tests
copy_tests:
	$(COPYREC) $(GO_EX_DIR)/* $(GEN_DIR)/ortools/go/

# Relevant Dylib/MacOS/rpath issue: https://github.com/golang/go/issues/36572
.PHONY: test_go_pimpl
test_go_pimpl: go_pimpl copy_tests
	cd $(GEN_DIR)/ortools/go && \
	$(GO_BIN) test -exec "env LD_LIBRARY_PATH=$(OR_ROOT_FULL)/lib" ./... -race -v

################
##  Cleaning  ##
################
.PHONY: clean_go # Clean Go output from previous build.
clean_go:
	-$(DELREC) $(GEN_PATH)$Sortools$Sgo
	-$(DEL) $(GEN_PATH)$Sortools$Salgorithms$S*.go
	-$(DEL) $(GEN_PATH)$Sortools$Salgorithms$S*go_wrap*
	-$(DEL) $(GEN_PATH)$Sortools$Sconstraint_solver$S*.go
	-$(DEL) $(GEN_PATH)$Sortools$Sconstraint_solver$S*go_wrap*
	-$(DEL) $(GEN_PATH)$Sortools$Sgraph$S*.go
	-$(DEL) $(GEN_PATH)$Sortools$Sgraph$S*go_wrap*
	-$(DEL) $(GEN_PATH)$Sortools$Slinear_solver$S*.go
	-$(DEL) $(GEN_PATH)$Sortools$Slinear_solver$S*go_wrap*
	-$(DEL) $(GEN_PATH)$Sortools$Ssat$S*.go
	-$(DEL) $(GEN_PATH)$Sortools$Ssat$S*go_wrap*
	-$(DEL) $(GEN_PATH)$Sortools$Sutil$S*.go
	-$(DEL) $(GEN_PATH)$Sortools$Sutil$S*go_wrap*
	-$(DEL) $(OBJ_DIR)$Sswig$S*_go_wrap*
	-$(DEL) $(LIB_DIR)$S$(LIB_PREFIX)$(GO_OR_TOOLS_NATIVE_LIB)*.$(SWIG_GO_LIB_SUFFIX)
	$(GO_BIN) clean -cache

###############
##  INSTALL  ##
###############
.PHONY: install_go # Install Go OR-Tools to $(DESTDIR)$(prefix)
install_go: install_cc go_pimpl
	$(COPY) $(LIB_DIR)$S$(LIB_PREFIX)$(GO_OR_TOOLS_NATIVE_LIB).$L "$(DESTDIR)$(prefix)$Slib"

#############
##  DEBUG  ##
#############
.PHONY: detect_go # Show variables used to build Go OR-Tools.
detect_go:
	@echo Relevant info for the Go build:
	@echo These must resolve to proceed
	@echo "  Install go: https://golang.org/doc/install"
	@echo "  Install protoc-gen-go: 'go get -u github.com/golang/protobuf/protoc-gen-go'"
	@echo GO_BIN = $(GO_BIN)
	@echo GO_PATH = $(GO_PATH)
	@echo PROTOC_GEN_GO = $(PROTOC_GEN_GO)
	@echo GO_OR_TOOLS_NATIVE_LIB = $(GO_OR_TOOLS_NATIVE_LIB)
ifeq ($(SYSTEM),win)
	@echo off & echo(
else
	@echo
endif
