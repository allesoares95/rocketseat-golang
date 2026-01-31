package main

import "fmt"

func main() {
	var fixo = 4
	multiplica := func(x int) int {
		return x * fixo
	}

	resultado := multiplica(5)
	fmt.Println(resultado)
}
