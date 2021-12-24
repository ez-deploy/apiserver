PROTO_REPO=https://github.com/ez-deploy/protobuf.git
TEMP_DIR=./.IXZOCBSBADOASDVBSAD

gen_swagger:
	rm -rf ./gen
	mkdir gen
	./.scripts/gen_swagger.sh ${PROTO_REPO} ./gen
