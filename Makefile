.PHONY: run
run:
	go run ./cmd/auth/main.go --config=./config/local.yaml

.PHONY: migrate
migrate:
	go run ./cmd/migrator --storage-path=./storage/auth.db --migrations-path=./migrations