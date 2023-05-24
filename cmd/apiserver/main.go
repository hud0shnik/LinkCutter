package main

import (
	"LinkCutter/internal/handler"
	"LinkCutter/internal/models"
	"flag"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {

	// Настройка логгера
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Запись флага типа хранилища
	var dbFlag bool
	flag.BoolVar(&dbFlag, "db", false, "Run with DB postgres")
	flag.Parse()

	// Проверка флага места хранения ссылок
	if !dbFlag {
		logrus.Info("Start API with In-Memory storage")
		mapShort := make(map[string]models.Url)
		mapLong := make(map[string]models.Url)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handler.CacheHandler(w, r, mapShort, mapLong)
		})
	} else {
		logrus.Info("Start API with PostgreSQL storage")
		http.HandleFunc("/", handler.DBHandler)
	}

	// Запуск API
	logrus.Fatal(http.ListenAndServe("localhost:8080", nil))
}
