#!/bin/sh

TEMP_DIR=.XZCZPXIUCBOIAPDF

PROTO_REPO=$1
PROTO_DIRS="authority identity project"

GENERATE_TO=$2

OUT_OPTS="allow_delete_body=true,logtostderr=true,grpc_api_configuration=service.yml:${GENERATE_TO}"

echo "PROTO_REPO = ${PROTO_REPO}"

git clone ${PROTO_REPO} ${TEMP_DIR} 

for DIR in $PROTO_DIRS; do
    protoc -I ${TEMP_DIR} --swagger_out="$OUT_OPTS" \
        ${TEMP_DIR}/$DIR/*.proto
done

# post.
rm -rf ${TEMP_DIR}