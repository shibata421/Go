package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	adicao, menos := calculosMatematicos(1, 2)
	fmt.Println(adicao, menos)
}
