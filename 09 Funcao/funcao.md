### Função com retorno
```go
// Função que recebe parâmetros e retorna um valor
func main() {
	resultado := soma(3, 5)                  // Chama a função e armazena o retorno
	fmt.Println(resultado)                   // Imprime: 8
}

func soma(a, b int) int {                    // (a, b int) = dois parâmetros int, retorna int
	return a + b                             // Retorna a soma
}
```

### Função sem retorno
```go
// Função que não retorna valor (apenas executa)
func main() {
	soma(3, 5)                               // Chama a função
}

func soma(a, b int) {                        // Sem tipo de retorno = não retorna nada
	fmt.Println(a + b)                       // Imprime diretamente: 8
}
```

### Função anônima
```go
// Função sem nome, atribuída a uma variável
func main() {
	multiplica := func(x int) int {          // Declara função anônima e atribui à variável
		return x * 2
	}

	resultado := multiplica(5)               // Chama a função através da variável
	fmt.Println(resultado)                   // Imprime: 10
}
```

### Closure (função com acesso a variável externa)
```go
// Closure: função anônima que acessa variável do escopo externo
func main() {
	var fixo = 4                             // Variável do escopo externo
	multiplica := func(x int) int {
		return x * fixo                      // Acessa 'fixo' do escopo externo (closure)
	}

	resultado := multiplica(5)               // 5 * 4 = 20
	fmt.Println(resultado)                   // Imprime: 20
}
```

### Resumo de Funções

| Forma | Uso |
|-------|-----|
| `func nome(params) tipo {}` | Função com retorno |
| `func nome(params) {}` | Função sem retorno |
| `func(params) tipo {}` | Função anônima |
| `a, b int` | Parâmetros do mesmo tipo (forma curta) |
| `return valor` | Retorna um valor |