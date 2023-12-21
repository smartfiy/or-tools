#!/usr/bin/env bash
set -xeuo pipefail

PROJECT=or-tools
TARGET=$(uname -m)

PROJECT_DIR=$(pwd -P)
BUILD_DIR=${PROJECT_DIR}/build/${TARGET}
CMAKE_DEFAULT_ARGS=(-G ${CMAKE_GENERATOR:-"Unix Makefiles"} -DBUILD_DEPS=ON -DBUILD_CXX=ON -DBUILD_GO=ON -DBUILD_GO_EXAMPLES=ON)

rm -rf "${BUILD_DIR}"
mkdir -p "${BUILD_DIR}"
cmake -S. -B"${BUILD_DIR}" "${CMAKE_DEFAULT_ARGS[@]}"
cmake --build "${BUILD_DIR}" --target all -j8 -v --verbose

# copy generated go sources to target dir
echo -n "Copying go sources to target..."
rm -rf ${PROJECT_DIR}/go
cp -r ${BUILD_DIR}/go ${PROJECT_DIR}

# make archive
INSTALL_GO_NAME=$(make print-INSTALL_GO_NAME 2> /dev/null | cut -d' ' -f3 | tr -d \')
LIBS=$(cd ${BUILD_DIR} && ls lib*/libgoortools.* lib*/libortools.*)
echo -n "Archiving..."
mkdir -p "${PROJECT_DIR}/export"
tar czvf export/${INSTALL_GO_NAME}.tar.gz --no-same-owner -C ${BUILD_DIR} ${LIBS}
