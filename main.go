package main

import "fmt"

func main() {
	var gavetas []string
	gavetas = append(gavetas, "Meias")
	gavetas = append(gavetas, "Camisetas")
	gavetas = append(gavetas, "Cuecas")

	// Removendo o último item
	gavetas = gavetas[:2]

	fmt.Println(gavetas)
}
