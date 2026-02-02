package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	p1 := Pessoa{Nome: "Ana"}
	p2 := Pessoa{Nome: "Bruno"}

	fmt.Println(p1)

	var p3 *Pessoa = &p1
	p3.Nome = "Carla"

	fmt.Println(p1)
	fmt.Println(p2)
}
