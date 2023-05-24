package handler

import (
	"LinkCutter/internal/models"
	"LinkCutter/internal/utilities"
	"errors"
	"net/http"
)

// CacheHandler - роутер для In-Memory решения
func CacheHandler(w http.ResponseWriter, r *http.Request, mapShort map[string]models.Url, mapLong map[string]models.Url) {
	switch r.Method {
	case "POST":
		postToCache(w, r, mapShort, mapLong)
	case "GET":
		getFromCache(w, r, mapShort)
	default:
		utilities.SendError(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}
}

// postToCache позволяет записать в память URL
func postToCache(w http.ResponseWriter, r *http.Request, mapShort map[string]models.Url, mapLong map[string]models.Url) {

	var longUrl models.Url
	var result models.Url

	// Запись тела реквеста в структуру
	err := utilities.DecodeJson(&longUrl, r.Body)
	if err != nil {
		utilities.SendError(w, err, http.StatusBadRequest)
		return
	}

	// Генерация короткого URL
	result.Url = utilities.HashFunc(longUrl.Url)
	mapLong[longUrl.Url] = result
	mapShort[result.Url] = longUrl

	// Отправка Результата
	utilities.SendResponse(w, result)
}

// getFromCache позволяет получить URL из памяти
func getFromCache(w http.ResponseWriter, r *http.Request, mapShort map[string]models.Url) {

	var shortUrl models.Url

	// Запись тела реквеста в структуру
	err := utilities.DecodeJson(&shortUrl, r.Body)
	if err != nil {
		utilities.SendError(w, err, http.StatusBadRequest)
	}

	// Отправка результата
	utilities.SendResponse(w, mapShort[shortUrl.Url])
}
