package handler

import (
	"LinkCutter/internal/models"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тесты для роутера c In-Memory решением
func Test_CacheHandler(t *testing.T) {
	testCases := []struct {
		description    string
		httpMethod     string
		requestBody    io.Reader
		expectedStatus int
	}{
		{
			description:    "invalid method",
			httpMethod:     "OPTIONS",
			requestBody:    strings.NewReader(`{}`),
			expectedStatus: http.StatusBadRequest,
		},
		{
			description:    "POST with empty JSON",
			httpMethod:     http.MethodPost,
			requestBody:    strings.NewReader(`{}`),
			expectedStatus: http.StatusOK,
		},
		{
			description:    "POST with valid JSON",
			httpMethod:     http.MethodPost,
			requestBody:    strings.NewReader(`{"url":"www.ozon.ru/"}`),
			expectedStatus: http.StatusOK,
		},
		{
			description:    "POST with invalid JSON",
			httpMethod:     http.MethodPost,
			requestBody:    strings.NewReader(`{"url":7}`),
			expectedStatus: http.StatusBadRequest,
		},
		{
			description:    "GET with empty JSON",
			httpMethod:     http.MethodGet,
			requestBody:    strings.NewReader(`{}`),
			expectedStatus: http.StatusOK,
		},
		{
			description:    "GET with invalid JSON",
			httpMethod:     http.MethodGet,
			requestBody:    strings.NewReader(`{"url":true}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	// Проверка кейсов
	shortMap := make(map[string]models.Url)
	longMap := make(map[string]models.Url)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CacheHandler(w, r, shortMap, longMap)
	})
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.httpMethod, "/", tc.requestBody)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedStatus, rec.Code)

		})

	}
}
