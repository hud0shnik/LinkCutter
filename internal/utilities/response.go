package utilities

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

// SendError обрабатывает ошибки
func SendError(w http.ResponseWriter, err error, code int) {
	logrus.Println(err.Error(), code)
	http.Error(w, strconv.Itoa(code), code)
}

// SendResponse отправляет ответ пользователю в формате json
func SendResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		SendError(w, errors.New("encoding error"), http.StatusInternalServerError)
	}
}
