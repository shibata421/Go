package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Erro"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-feira"
		fallthrough
	case numero == 3:
		diaDaSemana = "Terça-feira"
		fallthrough
	case numero == 4:
		diaDaSemana = "Quarta-feira"
		fallthrough
	case numero == 5:
		diaDaSemana = "Quinta-feira"
		fallthrough
	case numero == 6:
		diaDaSemana = "Sexta-feira"
		fallthrough
	case numero == 7:
		diaDaSemana = "Sábado"
		fallthrough
	default:
		diaDaSemana = "Erro"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana2(3))
}
