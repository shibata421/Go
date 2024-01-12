package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		w.Write([]byte("Oi!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
