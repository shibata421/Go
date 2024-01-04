package main

import "fmt"

func main() {
	var numero int64 = 10000000000000
	fmt.Println(numero)

	var numero2 uint = 100
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000000000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234567.74
	fmt.Println(numeroReal3)


	var str string = "AKAKAKA"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B' // 66
	fmt.Println(char)


	var texto int
	fmt.Println(texto)


	var booleano1 bool = true
	fmt.Println(booleano1)
}
