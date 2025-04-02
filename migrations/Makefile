PATH := $(PATH):$(GOPATH)/bin
BOIL_VER := v4.16.2
BOIL_EXT_VER := v0.8.0

create-migration:
	sql-migrate new -env=development ${FILENAME}

migrate:
	godotenv -f .env sql-migrate up -env=development

migrate-registration-test-local:
	godotenv -f .env.registration.test.local sql-migrate up -env=test

migrate-registration-test:
	godotenv -f .env.registration.test sql-migrate up -env=test

migrate-business-test-local:
	godotenv -f .env.business.test.local sql-migrate up -env=test

migrate-business-test:
	godotenv -f .env.business.test sql-migrate up -env=test
