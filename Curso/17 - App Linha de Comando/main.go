package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	if erro := app.Gerar().Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
