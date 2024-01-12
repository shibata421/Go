package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Pessoa struct {
	Nome  string `json:"nomePessoa"`
	Idade uint
}

func main() {
	jsonStruct, erro := json.Marshal(Pessoa{Nome: "Jose", Idade: 27})

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(string(jsonStruct))

	pessoaMap := map[string]string{"nome": "Maria", "idade": "30"}
	jsonMap, _ := json.Marshal(pessoaMap)
	fmt.Println(bytes.NewBuffer(jsonMap))

	pessoaEmJson := `{"nomePessoa":"Joao", "idade":20}`
	var pessoaStruct Pessoa
	if erro := json.Unmarshal([]byte(pessoaEmJson), &pessoaStruct); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(pessoaStruct)
}
