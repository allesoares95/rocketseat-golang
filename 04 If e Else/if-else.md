### Estrutura básica
```go
// Verifica as notas e retorna uma mensagem conforme a validação
func main() {
    nota := 96

    if nota >= 80 {  // Se nota for maior ou igual a 80
        fmt.Println("Parabéns! Você foi aprovado com distinção.")
    } else if nota >= 60 {  // Senão, se nota for maior ou igual a 60
        fmt.Println("Parabéns! Você foi aprovado.")
    } else {    // Senão (nota menor que 60)
        fmt.Println("Infelizmente, você reprovou.")
    }
}
```

### If com declaração curta (tratamento de erro)
```go
// Declaração de variável dentro do if - escopo limitado ao bloco
func main() {
    if err := thisIsAnError(); err != nil {  // Declara 'err' e verifica se não é nil
        fmt.Println(err.Error())  // Se houver erro, imprime a mensagem
    }
}

func thisIsAnError() error {
    return errors.New("this is an error")  // Retorna um erro
}
```

### If com verificação de existência em map (inline)
```go
// Verifica se a chave existe no map usando declaração curta no if
func main() {
    players := map[string]int{
        "alexandre": 10,
    }

    if value, ok := players["alexandre"]; ok {  // Declara value e ok, verifica se ok é true
        fmt.Println("pontos:", value, ok)  // Se a chave existir, imprime o valor
    }
    // value e ok só existem dentro do bloco if
}
```

### Verificação de existência em map (forma tradicional)
```go
// Mesma verificação, mas com variáveis declaradas fora do if
func main() {
    players := map[string]int{
        "alexandre": 10,
    }

    value, ok := players["alexandre"]  // Declara as variáveis
    fmt.Println("pontos:", value, ok)  // value=10, ok=true
    // value e ok continuam acessíveis aqui
}
```

### Resumo das formas de if

| Forma | Uso | Escopo da variável |
|-------|-----|-------------------|
| `if condição {}` | Verificação simples | N/A |
| `if x := valor; condição {}` | Declaração + verificação | Dentro do if/else |
| `if valor, ok := map[chave]; ok {}` | Verificar existência em map | Dentro do if/else |