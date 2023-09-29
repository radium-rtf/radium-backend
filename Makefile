swag:
	swag init -g ./internal/httpserver/handlers/router.go
.PHONY: swag

migrate-create:
	go run ./cmd/migrate db create_sql "$(name)"
.PHONY: migrate-create

migrate-run:
	go run ./cmd/migrate db migrate
.PHONY: migrate-run

migrate-init:
	go run ./cmd/migrate db init
.PHONY: migrate-init

migrate-rollback:
	go run ./cmd/migrate db rollback
.PHONY: migrate-rollback