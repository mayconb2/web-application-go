package models

import "github.com/mayconb2/web-application-go/db"

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SelectAllProdutos() []Produto {

	db := db.ConectaDb()
	selectTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos

}

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaDb()

	insertDataIntoDB, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insertDataIntoDB.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
