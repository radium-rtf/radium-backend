swag radium:
	swag init -g ./httpserver/handlers/router.go --instanceName radium -dir ./internal/radium --output ./docs/radium
.PHONY: swag radium

swag wave:
	swag init -g ./httpserver/handlers/router.go --instanceName wave -dir ./internal/wave --output ./docs/wave
.PHONY: swag wave

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