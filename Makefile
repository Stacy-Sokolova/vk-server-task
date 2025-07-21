include .env

build:
	docker-compose build app

run:
	docker-compose up app

migrate:
	migrate -path ./schema -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${SSLMODE}' up