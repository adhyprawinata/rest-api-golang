package utils

import (
	"database/sql"
	"rest-api/config"
)

func OpenDB() (*sql.DB, error) {
	// Membuka koneksi ke database
	db, err := sql.Open("mysql", config.DBUsername+":"+config.DBPassword+"@tcp("+config.DBHost+":"+config.DBPort+")/"+config.DBName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
