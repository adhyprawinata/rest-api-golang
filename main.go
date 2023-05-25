package main

import (
	"database/sql"
	"log"
	"net/http"

	"rest-api/config"
	"rest-api/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Membuka koneksi ke database
	db, err := sql.Open("mysql", config.DBUsername+":"+config.DBPassword+"@tcp("+config.DBHost+":"+config.DBPort+")/"+config.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Memeriksa koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Menyiapkan router
	router := routes.SetupRoutes()

	// Menjalankan server HTTP
	log.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Server startup failed:", err)
	}
}
