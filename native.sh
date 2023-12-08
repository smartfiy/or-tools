#!/usr/bin/env bash
set -xeuo pipefail

PROJECT=or-tools
TARGET=$(uname -m)

PROJECT_DIR=$(pwd -P)
BUILD_DIR=${PROJECT_DIR}/build/${TARGET}
CMAKE_DEFAULT_ARGS=(-G ${CMAKE_GENERATOR:-"Unix Makefiles"} -DBUILD_DEPS=ON -DBUILD_CXX=ON -DBUILD_GO=ON -DBUILD_GO_EXAMPLES=ON)

set -x
rm -rf "${BUILD_DIR}"
mkdir -p "${BUILD_DIR}"
cmake -S. -B"${BUILD_DIR}" "${CMAKE_DEFAULT_ARGS[@]}"
cmake --build "${BUILD_DIR}" --target all -j8 -v --verbose
set +x
