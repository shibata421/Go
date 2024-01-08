package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	var numero1 int16 = 10
	var numero2 int16 = 2
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	var variavel string = "string"
	variavel2 := "string2"
	fmt.Println(variavel, variavel2)

	fmt.Println(1 > 2)
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	fmt.Println(true && true)
	fmt.Println(true || true)
	fmt.Println(!true)

	numero := 10
	numero++
	fmt.Println(numero)
	numero--
	fmt.Println(numero)

	numero += 1
	numero -= 1
	numero *= 1
	numero /= 1
	numero %= 1
	fmt.Println(numero)

}
