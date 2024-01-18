package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // operador ... cria o array de acordo com a quantidade de items
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("-------------------------------------")

	slice3 := make([]float32, 10, 11)
	// make cria um array de 15 posições e retorna um slice
	// com as primeiras 10 posições

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // 11
	fmt.Println(cap(slice3)) // 11

	slice3 = append(slice3, 5)
	fmt.Println(len(slice3)) // 12
	fmt.Println(cap(slice3)) // 24. O go, quando vê que vai extourar o slice, ele dobra o seu tamanho

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
