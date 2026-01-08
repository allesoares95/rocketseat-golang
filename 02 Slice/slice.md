### Slice (Tamanho dinâmico)
```go
var gavetas []string
gavetas = append(gavetas, "meias", "camisas")
fmt.Println(gavetas)    // Exibe [meias camisas]

----------------------------------------------------------

var nomes []string = []string{"João", "Maria"}  // slice com declaração explícita
nomes = append(nomes, "Pedro")                  // adiciona "Pedro"1
fmt.Println(nomes)      // Exibe [João Maria Pedro]

----------------------------------------------------------

cores := []string{"Vermelho", "Verde", "Azul"}  // declaração curta
fmt.Println("Cores:", cores)    // Exibe [Vermelho Verde Azul]

cores = append(cores, "Amarelo")    // adiciona "Amarelo"
fmt.Println("Cores:", cores)    // Exibe [Vermelho Verde Azul Amarelo]

fmt.Println("Tamanho do slice:", len(cores)) // Exibe 4
fmt.Println(cores[2]) // Exibe "Azul" (acesso ao terceiro elemento)
fmt.Println(cores[:2]) // Exibe [Vermelho Verde]

----------------------------------------------------------

var gavetas []string
gavetas = append(gavetas, "Meias")
gavetas = append(gavetas, "Camisetas")
gavetas = append(gavetas, "Cuecas")

// slice[x:x] - fatia do slice do indice x até o indice x-1
// Exemplo: gavetas[1:2] - fatia do indice 1 até o indice 1 (2-1)
// Resultado: [Camisetas]
    
fmt.Println(gavetas[1:2])
```

----------------------------------------------------------

```go
var gavetas []string
gavetas = append(gavetas, "Meias")
gavetas = append(gavetas, "Camisetas")
gavetas = append(gavetas, "Cuecas")

// Removendo o último item
gavetas = gavetas[:2]

fmt.Println(gavetas)
```