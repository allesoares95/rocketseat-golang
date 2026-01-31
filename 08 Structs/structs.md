### Struct básica
```go
// Struct define um tipo personalizado com campos nomeados
type Cliente struct {
	Nome     string                          // Campo do tipo string
	Idade    int                             // Campo do tipo int
	Endereco string                          // Campo do tipo string
	Email    string                          // Campo do tipo string
}

func main() {
	cliente1 := Cliente{                     // Cria uma instância com todos os campos
		Nome:     "Ana Silva",
		Idade:    28,
		Endereco: "Rua das Flores, 123",
		Email:    "ana_silva@teste.com",
	}

	cliente2 := Cliente{                     // Cria uma instância com alguns campos (Email fica "")
		Nome:     "Bruno Souza",
		Idade:    35,
		Endereco: "Avenida Central, 456",
	}

	fmt.Println(cliente1)                    // Imprime toda a struct
	fmt.Println(cliente2)                    // Imprime toda a struct

	fmt.Println(cliente1.Nome)               // Acessa campo específico: "Ana Silva"
	fmt.Println(cliente2.Nome)               // Acessa campo específico: "Bruno Souza"

	cliente2.Email = "bruno_souza@teste.com" // Modifica um campo da struct

	fmt.Println(cliente2)                    // Imprime struct atualizada
}
```

### Struct aninhada (composição)
```go
// Uma struct pode conter outra struct como campo
type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco                        // Campo do tipo Endereco (outra struct)
	Email    string
}

type Endereco struct {                       // Struct separada para endereço
	Rua    string
	Cidade string
	Estado string
}

func main() {
	cliente1 := Cliente{
		Nome:  "Ana Silva",
		Idade: 28,
		Endereco: Endereco{                  // Inicializa a struct aninhada
			Rua:    "Rua das Flores, 123",
			Cidade: "Rio de Janeiro",
			Estado: "RJ",
		},
		Email: "ana_silva@teste.com",
	}

	cliente2 := Cliente{
		Nome:  "Bruno Souza",
		Idade: 35,
		Endereco: Endereco{
			Rua:    "Avenida Central, 456",
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}

	fmt.Println(cliente1)                    // Imprime toda a struct (incluindo Endereco)
	fmt.Println(cliente2)

	fmt.Println(cliente1.Endereco.Cidade)    // Acessa campo da struct aninhada: "Rio de Janeiro"
	fmt.Println(cliente2.Endereco.Cidade)    // Acessa campo da struct aninhada: "São Paulo"
}
```

### Modificando struct aninhada
```go
// Modifica campos de structs aninhadas usando notação de ponto
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

	fmt.Println(cliente1)                    // Imprime antes da modificação

	cliente1.Endereco.Rua = "Rua centro, 321" // Modifica campo da struct aninhada

	fmt.Println(cliente1)                    // Imprime após a modificação
}
```

### Resumo de Structs

| Forma | Uso |
|-------|-----|
| `type Nome struct {}` | Define uma nova struct |
| `Nome{campo: valor}` | Cria instância com valores |
| `instancia.campo` | Acessa um campo |
| `instancia.campo = valor` | Modifica um campo |
| `instancia.struct.campo` | Acessa campo de struct aninhada |