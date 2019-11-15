package models

import "github.com/mayconb2/web-application-go/db"

type Produto struct {
	Id         int
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

		p.Id = id
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

func DeleteProduct(idProduct string) {
	db := db.ConectaDb()

	deleteProduct, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}
	deleteProduct.Exec(idProduct)
	defer db.Close()
}

func ProductEdit(id string) Produto {
	db := db.ConectaDb()

	dataBaseProduct, err := db.Query("select * from produtos where id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	produtcToUpdate := Produto{}

	for dataBaseProduct.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = dataBaseProduct.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtcToUpdate.Id = id
		produtcToUpdate.Nome = nome
		produtcToUpdate.Descricao = descricao
		produtcToUpdate.Preco = preco
		produtcToUpdate.Quantidade = quantidade
	}

	defer db.Close()
	return produtcToUpdate
}

func UpdateProduct(nome, descricao string, id, quantidade int, preco float64) {

	db := db.ConectaDb()

	updateProduct, err := db.Prepare("update produtos set nome=?, descricao=?, preco=?, quantidade=? where id= ? order by id asc")
	if err != nil {
		panic(err.Error)
	}

	updateProduct.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
