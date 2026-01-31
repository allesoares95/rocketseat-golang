package main

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua    string
	Cidade string
	Estado string
}

func main() {
	cliente1 := Cliente{
		Nome:  "Ana Silva",
		Idade: 28,
		Endereco: Endereco{
			Rua:    "Rua das Flores, 123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
		},
		Email: "ana_silva@teste.com",
	}

	fmt.Println(cliente1)

	cliente1.Endereco.Rua = "Rua centro, 321"

	fmt.Println(cliente1)
}
