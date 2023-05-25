package routes

import (
	"net/http"

	"rest-api/controllers"
)

func SetupRoutes() http.Handler {

	http.HandleFunc("/users", controllers.GetAllUsers)

	return http.DefaultServeMux
}
