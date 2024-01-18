package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	tipoEndereco := strings.ToLower(strings.Split(endereco, " ")[0])

	for _, tipo := range tiposValidos {
		if tipoEndereco == tipo {
			return cases.Title(language.BrazilianPortuguese).String(tipo)
		}
	}

	return "Tipo de endereço inválido"
}
