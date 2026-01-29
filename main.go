package main

import (
	"fmt"
)

func main() {
	dia := "sábado"

	switch dia {
	case "sábado", "domingo":
		fmt.Println("Final de semana!")
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia útil")
	default:
		fmt.Println("Dia inválido")
	}
}
