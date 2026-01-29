## For em Go

### Estrutura básica
```go
// Sintaxe: for inicialização; condição; incremento { }
func main() {
    soma := 0

    for i := 0; i < 10; i++ {                   // i começa em 0, executa enquanto i < 10, incrementa i
        fmt.Println(i)                          // Imprime o valor de i (0 a 9)
        soma += i                               // Acumula o valor de i na soma
    }

    println(soma)                               // Imprime a soma total (45)
}
```

### For como while (apenas condição)
```go
// Go não tem while, usa for apenas com condição
func main() {
    soma := 0

    for soma < 20 {                             // Executa enquanto soma for menor que 20
        fmt.Println(soma)                       // Imprime o valor atual da soma
        soma += 2                               // Incrementa soma de 2 em 2
    }

    println(soma)                               // Imprime 20 (quando sai do loop)
}
```

### For percorrendo slice (forma tradicional)
```go
// Percorre um slice usando índice
func main() {
    nums := []int{1, 2, 3, 4, 5}

    for i := 0; i < len(nums); i++ {            // i vai de 0 até o tamanho do slice - 1
        fmt.Println(nums[i])                    // Acessa o elemento pelo índice
    }
}
```

### For com range (forma idiomática)
```go
// Forma preferida em Go para percorrer coleções
func main() {
    nums := []int{1, 2, 3, 4, 5}

    for i, num := range nums {                  // i = índice, num = valor
        fmt.Println(i, num)                     // Imprime índice e valor
    }

    for _, num := range nums {                  // _ ignora o índice
        fmt.Println(num)                        // Imprime apenas o valor
    }

    for i := range nums {                       // Omite o valor
        fmt.Println(i)                          // Imprime apenas o índice
    }
}
```

### For infinito
```go
// Loop infinito - precisa de break para sair
func main() {
    for {                                       // Sem condição = executa para sempre
        fmt.Println("infinito")
        // break                                // Descomente para sair do loop
    }
}
```

### Resumo do For

| Forma | Uso |
|-------|-----|
| `for i := 0; i < n; i++ {}` | Loop tradicional com contador |
| `for condição {}` | Loop tipo while |
| `for i, v := range slice {}` | Percorre slice com índice e valor |
| `for {}` | Loop infinito |
| `break` | Sai do loop |
| `continue` | Pula para próxima iteração |