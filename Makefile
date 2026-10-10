include .env
export 

export PROJECT_ROOT=$(shell pwd)

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

dashboard-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/dashboard/main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

lint:
	golangci-lint run

