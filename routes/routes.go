package routes

import (
	"net/http"

	"github.com/mayconb2/web-application-go/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
