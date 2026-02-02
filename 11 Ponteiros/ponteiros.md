### O que é um ponteiro?
Um ponteiro é uma variável que armazena o **endereço de memória** de outra variável, em vez do valor em si.

### Operadores importantes
- `&` - Retorna o endereço de memória de uma variável
- `*` - Acessa o valor no endereço de memória (desreferenciação)

### Criando e usando ponteiros
```go
// Ponteiro armazena o endereço de memória de uma variável
func main() {
	x := 10                                  // Variável normal com valor 10
	var p *int = &x                          // p é um ponteiro para int, recebe endereço de x

	fmt.Println("Valor de x:", x)            // Imprime: 10
	fmt.Println("Endereço de x:", &x)        // Imprime: 0xc0000... (endereço na memória)
	fmt.Println("Valor de p:", p)            // Imprime: 0xc0000... (mesmo endereço)
	fmt.Println("Valor apontado por p:", *p) // Imprime: 10 (desreferencia o ponteiro)
}
```

### Modificando valor através do ponteiro
```go
// Ponteiros permitem modificar o valor original
func main() {
	x := 10
	p := &x                                  // p aponta para x

	fmt.Println("Antes:", x)                 // Imprime: 10

	*p = 20                                  // Modifica o valor no endereço (modifica x)

	fmt.Println("Depois:", x)                // Imprime: 20 (x foi modificado!)
}
```

### Ponteiros em funções (passagem por referência)
```go
// Sem ponteiro: função recebe CÓPIA (não modifica original)
func dobrarValor(n int) {
	n = n * 2                                // Modifica apenas a cópia local
}

// Com ponteiro: função recebe REFERÊNCIA (modifica original)
func dobrarPonteiro(n *int) {
	*n = *n * 2                              // Modifica o valor no endereço original
}

func main() {
	a := 10
	dobrarValor(a)                           // Passa cópia
	fmt.Println("Sem ponteiro:", a)          // Imprime: 10 (não mudou)

	b := 10
	dobrarPonteiro(&b)                       // Passa endereço
	fmt.Println("Com ponteiro:", b)          // Imprime: 20 (mudou!)
}
```

### Ponteiros com structs
```go
// Ponteiros são muito usados com structs para evitar cópias
type Pessoa struct {
	Nome  string
	Idade int
}

func aniversario(p *Pessoa) {                // Recebe ponteiro para Pessoa
	p.Idade++                                // Go permite p.Idade em vez de (*p).Idade
}

func main() {
	pessoa := Pessoa{Nome: "Ana", Idade: 25}
	fmt.Println("Antes:", pessoa.Idade)      // Imprime: 25

	aniversario(&pessoa)                     // Passa endereço da struct
	fmt.Println("Depois:", pessoa.Idade)     // Imprime: 26
}
```

### Criando ponteiro com new()
```go
// new() aloca memória e retorna um ponteiro
func main() {
	p := new(int)                            // Cria ponteiro para int (valor inicial: 0)
	fmt.Println("Valor inicial:", *p)        // Imprime: 0

	*p = 100                                 // Atribui valor
	fmt.Println("Novo valor:", *p)           // Imprime: 100
}
```

### Resumo de Ponteiros

| Sintaxe | Significado |
|---------|-------------|
| `*int` | Tipo: ponteiro para int |
| `&x` | Endereço de memória de x |
| `*p` | Valor no endereço apontado por p |
| `p := &x` | p recebe o endereço de x |
| `*p = 10` | Modifica o valor no endereço |
| `new(Tipo)` | Aloca memória e retorna ponteiro |

### Quando usar ponteiros?

| Situação | Usar ponteiro? |
|----------|----------------|
| Modificar variável em outra função | ✅ Sim |
| Struct grande (evitar cópia) | ✅ Sim |
| Variável pequena (int, bool) | ❌ Geralmente não |
| Slice, map, channel | ❌ Já são referências |

### Referência
- [GeeksforGeeks - Pointers in Golang](https://www.geeksforgeeks.org/pointers-in-golang/)