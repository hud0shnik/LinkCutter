package database

import (
	"fmt"
	"os"
)

// getConfig предоставляет доступ к конфигам PostgreSQL
func getConfig() string {

	return fmt.Sprintf(
		"host= %s port = %s user = %s password = %s dbname = %s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

}
