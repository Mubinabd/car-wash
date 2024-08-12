#!/bin/bash
CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/genproto
for x in $(find ${CURRENT_DIR}/quantum_submodule -type d); do
  if ls ${x}/*.proto 1> /dev/null 2>&1; then
    protoc -I=${x} -I=${CURRENT_DIR}/quantum_submodule -I /usr/local/go --go_out=${CURRENT_DIR} \
    --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
  fi
done