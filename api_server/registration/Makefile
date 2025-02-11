PATH := $(PATH):$(GOPATH)/bin
BOIL_VER := v4.16.2
BOIL_EXT_VER := v0.8.0

test-local:
	godotenv -f /app/.env.test.local go test -v ./... -p 1

test-ci:
	godotenv -f env.test go test -v ./... -p 1

prepare-sqlboiler:
	@go get -u github.com/tiendc/sqlboiler-extensions@$(BOIL_EXT_VER)
	@go install github.com/volatiletech/sqlboiler/v4@$(BOIL_VER)
	@go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@$(BOIL_VER)

gen-models:
	@sqlboiler mysql \
 		--templates $(GOPATH)/pkg/mod/github.com/volatiletech/sqlboiler/v4@$(BOIL_VER)/templates/main \
 		--templates $(GOPATH)/pkg/mod/github.com/tiendc/sqlboiler-extensions@$(BOIL_EXT_VER)/templates/boilv4/mysql

gen-schema:
	@oapi-codegen -config oapi_codegen_config.yaml -package ${DOMAIN} api/${DOMAIN}.yaml > api/generated/${DOMAIN}.go
