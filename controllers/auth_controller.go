package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"rest-api/config"
	"rest-api/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Membuka koneksi ke database
	db, err := sql.Open("mysql", config.DBUsername+":"+config.DBPassword+"@tcp("+config.DBHost+":"+config.DBPort+")/"+config.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Melakukan query ke tabel user
	rows, err := db.Query("SELECT username, password FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Membaca hasil query dan menyimpannya dalam slice pengguna (users)
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Mengirim respon JSON dengan pengguna yang ditemukan
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
