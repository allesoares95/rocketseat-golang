package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Saudacao() {
	p.Nome = "ale"
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Ana", Idade: 25}
	p1.Saudacao()

	fmt.Println(p1.Nome)
}
