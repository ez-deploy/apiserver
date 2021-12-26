#!/bin/sh

TEMP_DIR=.XZCZPXIUCBOIAPDF

PROTO_REPO=$1
PROTO_DIRS="authority identity project"

SWAGGER_HOME=.swagger
GENERATE_TO=$2

OUT_OPTS="allow_delete_body=true,logtostderr=true,grpc_api_configuration=${SWAGGER_HOME}/service.yml:${GENERATE_TO}"

mkdir gen
echo "PROTO_REPO = ${PROTO_REPO}"

git clone ${PROTO_REPO} ${TEMP_DIR} 

for DIR in $PROTO_DIRS; do
    protoc -I ${TEMP_DIR} --swagger_out="$OUT_OPTS" \
        ${TEMP_DIR}/$DIR/*.proto
done

swagger mixin --format=yaml --keep-spec-order ./gen/*/* -o ${SWAGGER_HOME}/models.swagger.yaml --quiet

swagger mixin --format=yaml --keep-spec-order ${SWAGGER_HOME}/custom.swagger.yaml ${SWAGGER_HOME}/models.swagger.yaml -o ${SWAGGER_HOME}/swagger.yaml --quiet

# post.
rm -rf ${TEMP_DIR}
rm -rf gen