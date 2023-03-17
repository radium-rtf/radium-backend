swag:
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag

migrate-create:
	migrate create -ext sql -dir migrations "$(migrate_name)"
.PHONY: migrate-create