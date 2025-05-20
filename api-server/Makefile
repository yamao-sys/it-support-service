PATH := $(PATH):$(GOPATH)/bin

test-local:
	godotenv -f /app/.env.test.local go test -v ./... -p 1 -shuffle=on

test-ci:
	godotenv -f .env.test go test -v ./... -p 1 -shuffle=on -coverprofile=coverage.txt

test-seed-ci:
	godotenv -f .env.test go run test/seeds/main.go

test-seed-local:
	godotenv -f .env.test.local go run test/seeds/main.go

gen-registration-schema:
	make gen-schema SERVICE=registration

gen-business-schema:
	make gen-schema SERVICE=business

gen-schema:
	@oapi-codegen -config oapi_codegen_config.yaml -package ${SERVICE}api api/${SERVICE}/schema/openapi.yaml > api/${SERVICE}/openapi.go
