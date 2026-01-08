package main

import "fmt"

func main() {
	var pessoas = map[string]int{}
	pessoas["Alice"] = 30
	pessoas["Bob"] = 25
	pessoas["Charlie"] = 35

	if idade, ok := pessoas["Bob"]; ok {
		fmt.Println("Idade de Bob:", idade, ok)
	} else {
		fmt.Println("Bob não encontrado")
	}

	delete(pessoas, "Bob")
	fmt.Println("Map após remover Bob:", pessoas)
}
