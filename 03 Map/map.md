```go
var pessoa = map[string]int{}
pessoa["Alexandre"] = 30
pessoa["Joycy"] = 27

fmt.Println(pessoa) // Exibe map[Alexandre:30 Joycy:27]

var pessoa map[string]string = map[string]string{ // declaração explícita
    "nome":   "Carlos",
    "cidade": "São Paulo",
}   

usuario := map[string]int{              // declaração curta
    "idade":  35,
    "pontos": 1500,
}

var config map[string]string            // map vazio
config = make(map[string]string)        // inicializa o map
config["ambiente"] = "desenvolvimento"  // adiciona chave-valor
config["versao"] = "1.0.0"              // adiciona outra chave-valor
fmt.Println(config) // Exibe map[ambiente:desenvolvimento versao:1.0.0]

config["ambiente"] = "produção" // atualiza o valor da chave "ambiente"
fmt.Println(config) // Exibe map[ambiente:produção versao:1.0

config["debug"] = "true"    // adiciona nova chave-valor
fmt.Println(config) // Exibe map[ambiente:produção debug:true versao:1


// Verificando se uma chave existe no map
var pessoa = map[string]int{}
pessoa["alexandre"] = 30
pessoa["maria"] = 25


// Se a chave João existir, imprime a idade
if idade, ok := pessoa["joao"]; ok {
    fmt.Println("Idade de João:", idade, ok)
} else {
// Se a chave João não existir, imprime mensagem
    fmt.Println("João não está na lista.")
}

```