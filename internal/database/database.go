package database

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// ConnectDB производит подключение к базе данных
func ConnectDB() *sql.DB {
	godotenv.Load()
	db, err := sql.Open("postgres", getConfig())
	if err != nil {
		logrus.Fatal("can not access to db", err)
	}
	return db
}

// InsertUrlToDB позволяет произвести запись в базу данных
func InsertUrlToDB(shortUrl, longUrl string, db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf(
		"INSERT INTO urls (long, short) VALUES ('%s', '%s')",
		longUrl, shortUrl))
	if err != nil {
		return err
	}
	return nil
}

// CheckUrlInDB проверяет наличие записи в базе данных
func CheckUrlInDB(flag *bool, longUrl string, db *sql.DB) error {
	record, err := db.Query(fmt.Sprintf(
		"SELECT EXISTS(SELECT 1 FROM urls where long = '%s')", longUrl))
	if err != nil {
		return err
	}
	for record.Next() {
		err = record.Scan(flag)
		if err != nil {
			logrus.Fatal("can not check if data exists", err)
			return err
		}
	}
	return nil
}

// GetShortFromDB позволяет получить короткий URL из базы данных
func GetShortFromDB(longUrl string, db *sql.DB) (string, error) {

	var res string
	record, err := db.Query(fmt.Sprintf(
		"SELECT short FROM urls WHERE long = ('%s')", longUrl))
	if err != nil {
		return res, err
	}
	for record.Next() {
		err = record.Scan(&res)
		if err != nil {
			logrus.Fatal("can not select short url from db")
			return res, err
		}
	}
	return res, nil
}

// GetLongFromDB позволяет получить полный URL из базы данных
func GetLongFromDB(shortUrl string, db *sql.DB) (string, error) {

	var res string
	record, err := db.Query(fmt.Sprintf("SELECT long FROM urls WHERE short = ('%s')", shortUrl))
	if err != nil {
		return res, err
	}
	for record.Next() {
		err = record.Scan(&res)
		if err != nil {
			logrus.Fatal("can not select long url from db")
			return res, err
		}
	}
	return res, nil
}
