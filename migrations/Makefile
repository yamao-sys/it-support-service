PATH := $(PATH):$(GOPATH)/bin
BOIL_VER := v4.16.2
BOIL_EXT_VER := v0.8.0

create-migration:
	sql-migrate new -env="mysql" ${FILENAME}

migrate:
	sql-migrate up -env="mysql"
