PATH := $(PATH):$(GOPATH)/bin
BOIL_VER := v4.16.2
BOIL_EXT_VER := v0.8.0

test-local:
	godotenv -f /app/.env.test.local go test -v ./... -p 1 -shuffle=on

test-ci:
	godotenv -f .env.test go test -v ./... -p 1 -shuffle=on -coverprofile=coverage.txt

test-seed-ci:
	godotenv -f .env.test go run test/seeds/main.go

test-seed-local:
	godotenv -f .env.test.local go run test/seeds/main.go

prepare-sqlboiler:
	@go get -u github.com/tiendc/sqlboiler-extensions@$(BOIL_EXT_VER)
	@go install github.com/volatiletech/sqlboiler/v4@$(BOIL_VER)
	@go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@$(BOIL_VER)

gen-models:
	@sqlboiler mysql \
 		--templates $(GOPATH)/pkg/mod/github.com/volatiletech/sqlboiler/v4@$(BOIL_VER)/templates/main \
 		--templates $(GOPATH)/pkg/mod/github.com/tiendc/sqlboiler-extensions@$(BOIL_EXT_VER)/templates/boilv4/mysql

gen-registration-schema:
	make gen-schema SERVICE=registration

gen-business-schema:
	make gen-schema SERVICE=business

gen-schema:
	@oapi-codegen -config oapi_codegen_config.yaml -package ${SERVICE}api api/${SERVICE}/reference.yaml > api/${SERVICE}/reference.go
