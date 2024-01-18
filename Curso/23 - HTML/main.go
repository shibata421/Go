package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", renderHome)

	porta := ":5000"

	fmt.Println("Escutando na porta " + porta)
	log.Fatal(http.ListenAndServe(porta, nil))
}

func renderHome(w http.ResponseWriter, r *http.Request) {
	u := usuario{Nome: "João", Email: "joao.pedro@email.com"}
	templates.ExecuteTemplate(w, "home.html", u)
}
