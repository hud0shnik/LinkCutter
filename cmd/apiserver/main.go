package main

import (
	"LinkCutter/internal/handler"
	"LinkCutter/internal/models"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	// Настройка логгера и вывод записи о начале работы
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Info("Start API")

	// Проверка количества параметров
	if len(os.Args) == 1 {
		logrus.Fatal("Memory type not expected")
	}

	// Проверка флага места хранения ссылок
	if os.Args[1] == "Cache" {
		mapShort := make(map[string]models.Url)
		mapLong := make(map[string]models.Url)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			handler.CacheHandler(w, r, mapShort, mapLong)
		})
	} else {
		http.HandleFunc("/", handler.DBHandler)
	}

	// Запуск API
	logrus.Fatal(http.ListenAndServe("localhost:8080", nil))
}
