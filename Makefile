include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d dashboard-postgres dashboard-redis

env-down:
	@docker compose down dashboard-postgres dashboard-redis

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность потери данных. [y/N]:" ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down dashboard-postgres dashboard-redis postgres-port-forwarder redis-port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata ${PROJECT_ROOT}/out/redisdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

postgres-port-forward:
	@docker compose up -d postgres-port-forwarder

postgres-port-close:
	@docker compose down postgres-port-forwarder

redis-port-forward:
	@docker compose up -d redis-port-forwarder

redis-port-close:
	@docker compose down redis-port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1;
	fi; \
	docker compose run --rm dashboard-migrate \
		create \
		-ext sql \
		-dir=/migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \

	@docker compose run --rm dashboard-migrate \
		-path=/migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@dashboard-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

logs-cleanup:
	@read -p "Очистить все log файлы? Опасность потери логов. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Файлы логов очищены"; \
	else \
		echo "Очистка логов отменена"; \
	fi

dashboard-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost REDIS_ADDR=localhost:6379 && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/dashboard/main.go

dashboard-deploy:
	@docker compose up -d --build dashboard

dashboard-undeploy:
	@docker compose down dashboard

ps:
	@docker compose ps


