package main

import (
	"database/sql"

	db2 "github.com/batistondeoliveira/fullcycle_arquitetura_hexagonal/adapters/db"
	"github.com/batistondeoliveira/fullcycle_arquitetura_hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemplo", 30)

	productService.Enable(product)

}
