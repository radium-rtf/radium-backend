swag:
	swag init -g ./internal/httpserver/handlers/router.go
.PHONY: swag

migrate-create:
	migrate create -ext sql -dir migrations "$(migrate_name)"
.PHONY: migrate-create