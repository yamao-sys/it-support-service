#!/bin/bash

# シーディング
docker-compose run --rm business_api sh -c 'make test-seed-local'

# テスト用の環境でAPIサーバを起動
docker-compose exec -d registration_api godotenv -f .env.test.local go run main.go
docker-compose exec -d business_api godotenv -f .env.test.local go run main.go

# テスト実行
docker-compose exec frontend pnpm test:e2e

# テスト用DBのリセット
docker-compose run --rm db mysql -h db -u root -p -e "DROP DATABASE it_support_test; CREATE DATABASE it_support_test; DROP DATABASE it_support_business_test; CREATE DATABASE it_support_business_test;"
docker-compose run --rm migrations sh -c 'make migrate-test-local && make migrate-business-test-local'
