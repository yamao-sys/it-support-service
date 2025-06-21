# CLAUDE.md

必ず日本語で回答してください。
This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an IT support service - a business matching platform built with Go backend services and Next.js frontend applications. The project uses schema-driven development with TypeSpec for API definitions and automatic code generation.

## Architecture

### Backend (Go)

- **Two separate services**: `registration` (port 8080) and `business` (port 8081)
- **Clean architecture**: handlers → services → models with validators and middlewares
- **Database**: GORM with MySQL, sql-migrate for migrations
- **Schema generation**: TypeSpec → OpenAPI → Go code via oapi-codegen
- **Testing**: Go testing with testify, uses txdb for database transactions

### Frontend (Next.js + Turborepo)

- **Monorepo structure**: Turborepo with pnpm workspaces
- **Two main apps**: `registration` (port 3000) and `business` (port 3002)
- **Schema generation**: TypeSpec → OpenAPI → TypeScript client via OpenAPI Generator
- **Testing**: Playwright for E2E tests

## Development Commands

### Local Development Setup

```bash
# Start all services (includes APIs, frontend, MySQL, fake GCS)
docker-compose up

# Frontend development (from frontend/it-support/)
pnpm dev                    # Start all apps in dev mode
pnpm dev:test              # Start apps in test mode for E2E testing
```

### Backend Development (from api-server/)

```bash
# Testing
make test-local            # Run tests locally
make test-ci              # Run tests with coverage
make test-seed-local      # Seed test data locally

# Code generation (after TypeSpec changes)
make gen-registration-schema  # Generate Go code for registration API
make gen-business-schema     # Generate Go code for business API
```

### Frontend Development (from frontend/it-support/)

```bash
# Building and linting
pnpm build                 # Build all apps
pnpm lint                  # Lint all code
pnpm format               # Format code with Prettier
pnpm check-types          # TypeScript type checking

# Testing
pnpm test:e2e             # Run all E2E tests
pnpm test:e2e:registration # Run registration E2E tests only
pnpm test:e2e:business    # Run business E2E tests only

# Schema generation (after TypeSpec changes)
pnpm gen:api-spec         # Generate OpenAPI from TypeSpec
pnpm gen:api-spec:watch   # Watch TypeSpec files and regenerate
pnpm gen:schema           # Generate TypeScript client from OpenAPI
pnpm diff:api-spec:check  # Check TypeSpec/OpenAPI sync
```

### Database Migrations (from migrations/)

```bash
# Create new migration
make create-migration FILENAME=your_migration_name

# Run migrations
make migrate              # Development environment
make migrate-test-local   # Test environment (local)
make migrate-test         # Test environment (CI)
```

## Schema-Driven Development Workflow

1. **Edit TypeSpec files** in `apps/{service}/api-spec/`
2. **Generate OpenAPI** with `pnpm gen:api-spec`
3. **Generate Go code** with `make gen-{service}-schema`
4. **Generate TypeScript client** with `pnpm gen:schema`
5. **Implement handlers/services** in Go backend
6. **Use generated client** in frontend applications

## Testing Strategy

### Backend Tests

- Unit tests for services and handlers
- Database transactions rolled back after each test
- Test factories in `api-server/test/factories/`
- Test seeding available via `make test-seed-local`

### Frontend Tests

- E2E tests using Playwright
- Tests run against built applications in test mode
- Japanese font support for UI testing
- Test artifacts saved on failure

## Project Structure

### Backend Services

```
api-server/
├── cmd/                   # Application entry points
├── {service}/             # Service-specific code
│   ├── handlers/         # HTTP handlers
│   ├── services/         # Business logic
│   ├── middlewares/      # HTTP middlewares
│   └── validators/       # Request validation
├── models/               # Database models (GORM)
├── api/{service}/        # Generated API code
└── test/                 # Test utilities and factories
```

### Frontend Applications

```
frontend/it-support/
├── apps/
│   ├── registration/     # Registration app (Next.js)
│   └── business/         # Business app (Next.js)
├── packages/
│   ├── ui/              # Shared UI components
│   └── {config}/        # Shared configurations
```

## Key Configuration Files

- `docker-compose.yaml` - Local development environment
- `turbo.json` - Turborepo task configuration
- `oapi_codegen_config.yaml` - Go code generation config
- `.air.{service}.toml` - Hot reload configuration for Go services

## Environment Variables

- Frontend apps use `REGISTRATION_API_ENDPOINT_URI` and `BUSINESS_API_ENDPOINT_URI`
- Backend services use various database and GCS environment variables
- Test environments have separate `.env.test*` files

## Common Development Tasks

### Adding a new API endpoint

1. Define in TypeSpec (`api-spec/routes/`)
2. Run `pnpm gen:api-spec` and `make gen-{service}-schema`
3. Implement handler, service, and validator
4. Add tests and run `make test-local`

### Adding a new database migration

1. Run `make create-migration FILENAME=descriptive_name`
2. Edit the generated SQL file in `migrations/ddl/`
3. Run `make migrate` to apply locally

### Running individual tests

```bash
# Backend: Run specific test file
go test -v ./business/services -run TestSpecificFunction

# Frontend: Run specific E2E test
npx playwright test --grep "test name"
```
