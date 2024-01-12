package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEspero    string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{enderecoInserido: "Avenida Paulista", retornoEspero: "Avenida"},
		{enderecoInserido: "Rua test", retornoEspero: "Rua"},
		{enderecoInserido: "Estrada test", retornoEspero: "Estrada"},
		{enderecoInserido: "Rodovia test", retornoEspero: "Rodovia"},
		{enderecoInserido: "Praça test", retornoEspero: "Tipo de endereço inválido"},
	}

	for _, test := range cenariosDeTeste {
		if test.retornoEspero != TipoDeEndereco(test.enderecoInserido) {
			t.Errorf("O tipo recebido é diferente do esperado. Esperava %s", test.retornoEspero)
		}
	}

}
