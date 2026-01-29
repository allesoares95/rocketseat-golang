package main

import (
	"fmt"
)

func main() {
	users := map[string]string{
		"nome":      "Alice",
		"cidade":    "Wonderland",
		"profissão": "Aventureira",
	}

	for value, _ := range users {
		fmt.Println(value)
	}
}
