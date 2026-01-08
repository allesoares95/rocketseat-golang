package main

import "fmt"

func main() {
	var gavetas [2]string

	gavetas[0] = "Meias"
	gavetas[1] = "Cuecas"

	fmt.Println(gavetas)
	fmt.Println("Conteúdo da gaveta 1:", gavetas[0])
	fmt.Println("Conteúdo da gaveta 2:", gavetas[1])
}
