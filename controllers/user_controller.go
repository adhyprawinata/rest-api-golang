package controllers

import (
	"log"
	"net/http"

	"rest-api/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Membuka koneksi ke database
	db, err := utils.OpenDB()
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
	var users []map[string]interface{}
	for rows.Next() {
		var username, password string
		if err := rows.Scan(&username, &password); err != nil {
			log.Fatal(err)
		}
		user := map[string]interface{}{
			"username": username,
			"password": password,
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Mengirim respon JSON dengan pengguna yang ditemukan
	utils.SendJSONResponse(w, utils.Response{
		Status:  http.StatusOK,
		Message: "Successfully",
		Data:    users,
	}, http.StatusOK)
	return
}
