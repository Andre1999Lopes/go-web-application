package main

import (
	"database/sql"
	"go-web-application/routes"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=Lopes888*x host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
