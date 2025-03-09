#!/bin/bash

# テスト用の環境でAPIサーバを起動
docker-compose exec -d registration_api godotenv -f .env.test.local go run main.go

# テスト実行
docker-compose exec frontend pnpm test:e2e

# テスト用DBのリセット
docker-compose run --rm db mysql -h db -u root -p -e "DROP DATABASE it_support_test; CREATE DATABASE it_support_test;"
docker-compose run --rm migrations sh -c 'make migrate-test-local'
