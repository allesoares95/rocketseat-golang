### Método básico (value receiver)
```go
// Método é uma função associada a um tipo (struct)
type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Saudacao() {                 // (p Pessoa) = receiver, associa método à struct
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Ana", Idade: 25}     // Cria instância da struct
	p2 := Pessoa{Nome: "Bruno", Idade: 30}

	p1.Saudacao()                            // Chama o método na instância p1
	p2.Saudacao()                            // Chama o método na instância p2
}
```

### Value receiver (cópia - não modifica o original)
```go
// Com value receiver, o método recebe uma CÓPIA da struct
type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Saudacao() {                 // Recebe cópia de Pessoa
	p.Nome = "ale"                           // Modifica apenas a cópia local
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Ana", Idade: 25}
	p1.Saudacao()                            // Imprime: "ale" (cópia modificada)

	fmt.Println(p1.Nome)                     // Imprime: "Ana" (original inalterado)
}
```

### Pointer receiver (referência - modifica o original)
```go
// Com pointer receiver, o método recebe um PONTEIRO para a struct
type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Saudacao() {                // *Pessoa = ponteiro, modifica o original
	p.Nome = "ale"                           // Modifica o objeto original
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Ana", Idade: 25}
	p1.Saudacao()                            // Imprime: "ale"

	fmt.Println(p1.Nome)                     // Imprime: "ale" (original foi modificado!)
}
```

### Resumo de Métodos

| Forma | Uso |
|-------|-----|
| `func (t Tipo) Nome() {}` | Método com value receiver (cópia) |
| `func (t *Tipo) Nome() {}` | Método com pointer receiver (referência) |
| `instancia.Metodo()` | Chama o método na instância |

### Quando usar cada um?

| Receiver | Quando usar |
|----------|-------------|
| `(t Tipo)` | Quando NÃO precisa modificar a struct original |
| `(t *Tipo)` | Quando PRECISA modificar a struct original |
| `(t *Tipo)` | Quando a struct é grande (evita cópia) | 