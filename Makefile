APP=./bin/app
IMG_NAME = ozon
IMG_VERSION = latest
SHELL := /bin/bash

# Запуск тестов
test:
	go test internal/handler -cover 

# Сборка API
build:
	go build -o $(APP) ./cmd/apiserver/main.go

# Запуск с Postgre
run_db:
	$(APP) -bd

# Запуск с In-Memory хранилищем
run_im:
	$(APP) -cache

# Сборка в контейнере с Postgres
docker_build_db:
	docker-compose build

# Запуск в контейнере с Postgres
docker_run_db:
	docker-compose up

# Сборка в контейнере с In-Memory хранилищем
docker_build_im:
	docker build -t $(IMG_NAME):$(IMG_VERSION) .

# Запуск в контейнере с In-Memory хранилищем
docker_run_im:
	docker run -d -p 8080:8080 $(IMG_NAME):$(IMG_VERSION)

# Запись конфигов для БД
env:
	./initDBConfig.sh