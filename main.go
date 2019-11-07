package main

import (
	"net/http"

	"github.com/mayconb2/web-application-go/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
