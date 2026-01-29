### Range com slice
```go
// Range retorna índice e valor ao percorrer um slice
func main() {
    nums := []string{"alexandre", "bruno", "carlos", "daniel"}

    for key, value := range nums {              // key = índice (0, 1, 2, 3), value = valor
        fmt.Println(key, value)                 // Imprime: 0 alexandre, 1 bruno, 2 carlos, 3 daniel
    }
}
```

### Range com map (chave e valor)
```go
// Range em map retorna chave e valor
func main() {
    users := map[string]string{
        "nome":      "Alice",
        "cidade":    "Wonderland",
        "profissão": "Aventureira",
    }

    for key, value := range users {             // key = chave do map, value = valor
        fmt.Println(key, value)                 // Imprime: nome Alice, cidade Wonderland, etc.
    }
}
```

### Range ignorando a chave
```go
// Usa _ para ignorar a chave quando só precisa do valor
func main() {
    users := map[string]string{
        "nome":      "Alice",
        "cidade":    "Wonderland",
        "profissão": "Aventureira",
    }

    for _, value := range users {               // _ ignora a chave
        fmt.Println(value)                      // Imprime apenas: Alice, Wonderland, Aventureira
    }
}
```

### Range ignorando o valor
```go
// Usa _ para ignorar o valor quando só precisa da chave
func main() {
    users := map[string]string{
        "nome":      "Alice",
        "cidade":    "Wonderland",
        "profissão": "Aventureira",
    }

    for key, _ := range users {                 // _ ignora o valor (ou pode omitir: for key := range)
        fmt.Println(key)                        // Imprime apenas: nome, cidade, profissão
    }
}
```

### Range com string
```go
// Range em string retorna índice e rune (código Unicode)
func main() {
    texto := "Olá"

    for i, char := range texto {                // i = índice do byte, char = rune (caractere)
        fmt.Printf("%d: %c\n", i, char)         // Imprime: 0: O, 1: l, 2: á
    }
}
```

### Resumo do Range

| Forma | Uso |
|-------|-----|
| `for i, v := range slice {}` | Percorre slice com índice e valor |
| `for k, v := range map {}` | Percorre map com chave e valor |
| `for _, v := range coleção {}` | Ignora índice/chave, usa só valor |
| `for k := range coleção {}` | Usa só índice/chave, ignora valor |
| `for i, char := range string {}` | Percorre string caractere por caractere |