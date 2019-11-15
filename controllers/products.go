package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/mayconb2/web-application-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SelectAllProdutos()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da Quantidade", err)
		}

		models.CreateNewProduct(nome, descricao, precoConvertido, quantidadeConvertida)

	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 204)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	product := models.ProductEdit(productID)
	temp.ExecuteTemplate(w, "Edit", product)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertedInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id para int: ", err)
		}

		precoConvertedFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float: ", err)
		}

		quantidadeConvertedInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade para int: ", err)
		}

		models.UpdateProduct(nome, descricao, idConvertedInt, quantidadeConvertedInt, precoConvertedFloat)
	}

	http.Redirect(w, r, "/", 301)
}
