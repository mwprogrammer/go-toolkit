package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/mwprogrammer/go-utilities/models"
)


func Connect(config models.DbConfig, logger *slog.Logger) (*sql.DB, bool) {

	is_success := true

	query := url.Values{}
	query.Add("app name", config.AppName)

	connection := &url.URL{

		Scheme: config.Scheme,
		User: url.UserPassword(config.Username, config.Password),
		Host: fmt.Sprintf("%s:%d", config.Host, config.Port),
		RawQuery: query.Encode(),

	}

	db, err := sql.Open(config.Scheme, connection.String())

	if err != nil {
		logger.Error(err.Error())
		is_success = false
	}

	return db, is_success
}

func GetRecords(query string, db *sql.DB, logger *slog.Logger) (*sql.Rows, bool) {

	is_success := true

	records, err := db.Query(query)

	if err != nil {
		logger.Error(err.Error())
		is_success = false
	}

	return records, is_success
}