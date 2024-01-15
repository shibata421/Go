package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // _ é para dizer que é um import implícito que será utilzado em outro lugar
)

func main() {
	db := criarConexao()

	linhas, erro := db.Query("SELECT * FROM usuarios")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()
}

func criarConexao() *sql.DB {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão com DB está aberta")
	return db
}
