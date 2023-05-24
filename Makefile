
# Запуск тестов
test:
	go test internal/handler -cover 

# Сборка API
build:
	go build -o ./bin/app ./cmd/apiserver/main.go

# Запись конфигов для БД
env:
	. initDBConfig.sh