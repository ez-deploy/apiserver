PROTO_REPO=https://github.com/ez-deploy/protobuf.git
TEMP_DIR=./.IXZOCBSBADOASDVBSAD
GO_MOD_NAME=github.com/ez-deploy/apiserver

regen_all: clean gen_swagger gen_server
	@echo "regen_all over! write your code in handle package.\n"

gen_swagger:
	@./.scripts/gen_swagger.sh ${PROTO_REPO} ./gen

gen_server:
	@swagger generate server \
		--implementation-package=${GO_MOD_NAME}/handle \
		--principal=github.com/ez-deploy/apiserver/models.IdentityVerifyResp \
		-f ./.swagger/swagger.yaml
	@go mod tidy

clean: clean_server clean_swagger
	@echo "clear! to regenerate server, run: \n  make gen_swagger && make gen_server"

clean_server:
	@rm -rf cmd models restapi

clean_swagger:
	@rm -rf .swagger/models.swagger.yaml
	@rm -rf .swagger/swagger.yaml