package main

import (
	"html/template"
	"net/http"
	"github.com/mayconb2/web-application-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html")){
	allProducts := models.
	
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
