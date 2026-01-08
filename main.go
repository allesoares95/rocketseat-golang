package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello = "Olá, Mundo!"
	var question string = " Como você está?"

	var meet = hello + question
	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet))           // Converte para maiúsculas
	fmt.Println(strings.Contains(meet, "Mundo")) // Verifica se contém "Mundo"
}
