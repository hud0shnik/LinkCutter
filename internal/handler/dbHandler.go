package handler

import (
	"LinkCutter/internal/database"
	"LinkCutter/internal/models"
	"LinkCutter/internal/utilities"
	"database/sql"
	"errors"
	"net/http"
)

// DBHandler - роутер для PostgreSQL решения
func DBHandler(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB = database.ConnectDB()
	switch r.Method {
	case "POST":
		postToDB(w, r, db)
	case "GET":
		getFromDB(w, r, db)
	default:
		utilities.SendError(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}
}

// postToDB позволяет записать URL в базу данных
func postToDB(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var longUrl models.Url
	var result models.Url
	var exist bool

	// Запись тела реквеста в структуру
	err := utilities.DecodeJson(&longUrl, r.Body)
	if err != nil {
		utilities.SendError(w, err, http.StatusBadRequest)
		return
	}

	// Проверка на пустой URL
	if len(longUrl.Url) == 0 {
		utilities.SendResponse(w, result)
		return
	}

	// Проверка наличия URL в базе данных
	err = database.CheckUrlInDB(&exist, longUrl.Url, db)
	if err != nil {
		utilities.SendError(w, errors.New("db check for existence error"), http.StatusInternalServerError)
		return
	}

	// Если URL уже существует, запись его короткой версии в структуру
	if exist {
		var err error
		result.Url, err = database.GetShortFromDB(longUrl.Url, db)
		if err != nil {
			utilities.SendError(w, errors.New("getting short url from db error"), http.StatusBadRequest)
			return
		}
	} else {
		// Создание и вставка новой записи в базу данных
		result.Url = utilities.HashFunc(longUrl.Url)
		err := database.InsertUrlToDB(result.Url, longUrl.Url, db)
		if err != nil {
			utilities.SendError(w, errors.New("db adding data error"), http.StatusInternalServerError)
			return
		}
	}

	// Отправка ответа
	utilities.SendResponse(w, result)
}

// getFromDB позволяет получить полную запись из базы данных
func getFromDB(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var shortUrl models.Url
	var result models.Url

	// Запись тела реквеста в структуру
	err := utilities.DecodeJson(&shortUrl, r.Body)
	if err != nil {
		utilities.SendError(w, err, http.StatusBadRequest)
	}

	// Поиск полной записи по короткому URL
	result.Url, err = database.GetLongFromDB(shortUrl.Url, db)
	if err != nil {
		utilities.SendError(w, errors.New("getting long url from db error"), http.StatusBadRequest)
	}

	// Отправка результата
	utilities.SendResponse(w, result)
}
