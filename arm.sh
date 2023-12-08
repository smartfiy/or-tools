#!/usr/bin/env bash

set -xeuo pipefail

#./tools/cross_compile.sh --help

export PROJECT=or-tools
#export PROJECT=glop
#export TARGET=x86_64
export TARGET=aarch64-unknown-linux-gnu

# for m1 mac
# NOTE: aarch64 and arm64 are equivalent
export GOOS=darwin
export GOARCH=arm64

#./tools/cross_compile.sh toolchain
./tools/cross_compile.sh build
# ./tools/cross_compile.sh qemu
#./tools/cross_compile.sh test
