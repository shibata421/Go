package main

import "fmt"

func inverterSinal(n *int) {
	*n = -*n
}

func main() {
	n := 1

	fmt.Println(&n)

	inverterSinal(&n)

	fmt.Println(n)
}
