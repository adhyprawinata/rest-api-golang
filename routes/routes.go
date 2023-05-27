package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"rest-api/controllers"
	"rest-api/utils"
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	router.Use(utils.CheckAPIKey)
	router.NotFoundHandler = http.HandlerFunc(utils.NotFoundHandler)

	router.Use(utils.Logging) // Tambahkan middleware pencatatan

	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")

	return router
}
