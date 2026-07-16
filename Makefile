.PHONY: build up down logs restart run fmt vet test lint

build:
	docker compose build app

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f app

restart:
	docker compose restart app

run:
	go run ./cmd/app/

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

lint:
	golangci-lint run

