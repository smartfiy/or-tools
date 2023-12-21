#!/usr/bin/env bash
set -xeuo pipefail

while [[ $# -gt 0 ]]; do
  case $1 in
    -a|--arm64)
      ARM64_TAR="$2"
      shift # past argument
      shift # past value
      ;;
    -x|--x86_64)
      X86_64_TAR="$2"
      shift # past argument
      shift # past value
      ;;
    -o|--out)
      OUT_TAR="$2"
      shift # past argument
      shift # past value
      ;;
    -*|--*)
      echo "Unknown option $1"
      exit 1
      ;;
  esac
done

echo "ARM64_TAR     = ${ARM64_TAR}"
echo "X86_64_TAR    = ${X86_64_TAR}"
echo "OUT_TAR       = ${OUT_TAR}"

PROJECT_DIR=$(pwd -P)
UNIVERSAL_DIR=${PROJECT_DIR}/universal
A_DIR=${UNIVERSAL_DIR}/a
X_DIR=${UNIVERSAL_DIR}/x
L_DIR=${UNIVERSAL_DIR}/lib
 
# create dir tree
mkdir -p "${UNIVERSAL_DIR}"
mkdir -p "${A_DIR}"
mkdir -p "${X_DIR}"
mkdir -p "${L_DIR}"

# extract
tar xzvf ${ARM64_TAR} --strip 1 -C ${A_DIR}
tar xzvf ${X86_64_TAR} --strip 1 -C ${X_DIR}
SYM_LINKS=$(find ${X_DIR} -mindepth 1 -type l)
X_COMBINE=$(find ${X_DIR} -mindepth 1 ! -type l)
A_COMBINE=$(find ${A_DIR} -mindepth 1 ! -type l)
SYM_ARR=(${SYM_LINKS})
X_ARR=(${X_COMBINE})
A_ARR=(${A_COMBINE})

# create universal files for each
for i in "${!X_ARR[@]}"; do
  NAME=$(basename "${X_ARR[$i]}")  
  lipo -create -output ${L_DIR}/${NAME} "${X_ARR[$i]}" "${A_ARR[$i]}"
done

# copy over symlinks
for i in "${SYM_ARR[@]}"; do
  cp -R "$i" ${L_DIR}
done

# create output tarball
tar czvf ${OUT_TAR} --no-same-owner -C ${UNIVERSAL_DIR} lib

# cleanup
rm -rf "${UNIVERSAL_DIR}"
