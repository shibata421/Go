package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
	generica(1.2)

	mapa := map[interface{}]interface{}{
		1:        "String",
		"String": 2,
		false:    1.2,
		true:     1.222,
	}

	fmt.Println(mapa)
}
