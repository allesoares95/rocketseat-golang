### Array (Tamanho fixo)

```go
var gavetas [2]string
gavetas[0] = "meias"
gavetas[1] = "camisas"
fmt.Println(gavetas) // Exibe [meias camisas]
fmt.Println(gavetas[0]) // Exibe "meias"
fmt.Println(gavetas[1]) // Exibe "camisas"

frutas := [3]string{"maçã", "banana", "laranja"} // declaração curta
fmt.Println(frutas) // Exibe [maçã banana laranja]

var numeros [5]int = [5]int{1, 2, 3, 4, 5}  // array de 5 inteiros
var notas [4]float64    // array de 4 notas decimais
notas[0] = 7.5
notas[1] = 8.0
notas[2] = 9.0
notas[3] = 6.5
fmt.Println(notas) // Exibe [7.5 8 9 6.5]
```