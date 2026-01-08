## Variáveis em Go

### Variáveis e Constantes
```go
// Variáveis (podem mudar)
var nome string = "João"
idade := 25  // declaração curta

// Constantes (não mudam)
const PI = 3.14159
const VERSAO = "1.0.0"
```

```go
package main

import "fmt"

func main() {
    var texto string = "olá"
    fmt.Println(texto);

    texto = "tchau"
    fmt.Println(texto);
}
```

### Int (Números inteiros)
```go
var idade int = 30  
var contador int32 = 2  //armazenamento de 32bits
var indice int8 = 1     //armazenamento de 8bits 

// - **int8**: de -128 a 127
// - **int16**: de -32.768 a 32.767
// - **int32**: de -2.147.483.648 a 2.147.483.647
// - **int64**: de -9.223.372.036.854.

var numero int = 42            // inteiro padrão (32 ou 64 bits, depende do sistema)
idade := 25                    // declaração curta
var pequeno int8 = 127         // -128 a 127
var grande int64 = 9999999999  // números muito grandes
```

### Float (Números decimais)
```go
var floatNumber float32 = 1.1   // float32 (menos precisão, menos memória)
var pi float64 = 3.14           // float64 (padrão, mais precisão)
var raio float64 = 2.5          
var area = pi * raio * raio     // área de um círculo

altura := 1.75                  // declaração curta
```

### Bool (Verdadeiro/Falso)
```go
var ativo bool = true
logado := false                // declaração curta
var disponivel bool            // false (valor padrão)

var maior bool = 10 > 5   // comparação
var menor bool = 3 < 1     // comparação
var igual bool = 2 == 2    // comparação
var diferente bool = 5 != 3 // comparação
```

### String (Texto)
```go
var nome string = "Maria"
sobrenome := "Silva"           // declaração curta
var vazio string               // "" (string vazia)
mensagem := `Texto com
múltiplas linhas`              // string literal

var saudacao string = "Olá, " + nome + " " + sobrenome // concatenação
fmt.Println(saudacao) // Exibe "Olá, Maria Silva"

var hello string = "Olá, Mundo!"
var frase string = "Como vai?"

var meet = hello + frase // concatenação de strings
fmt.Println(meet)   // Exibe "Olá, Mundo!Como vai?"
```