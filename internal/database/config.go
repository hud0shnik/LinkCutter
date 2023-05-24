package database

import (
	"os"
)

// getConfig предоставляет доступ к конфигам PostgreSQL
func getConfig() string {

	return os.Getenv("DB_URL")

}
