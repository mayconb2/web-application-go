package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaDb() *sql.DB {

	fmt.Println("Conectando com DB")
	db, err := sql.Open("mysql", "userNadir:nadir1020@/nadir_store?interpolateParams=TRUE")
	if err != nil {
		panic(err.Error())
	}
	return db
}
