package main

import (
	"fmt"
)

func main() {
	players := map[string]int{
		"alexandre": 10,
	}

	value, ok := players["alexandre"]
	fmt.Println("pontos:", value, ok)
}
