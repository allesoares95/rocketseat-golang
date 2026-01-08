# Golang

## Características Principais

- **Compilada**: Go compila diretamente para código de máquina
- **Tipagem estática**: Tipos são verificados em tempo de compilação
- **Garbage Collection**: Gerenciamento automático de memória
- **Concorrência nativa**: Goroutines e channels facilitam programação paralela
- **Sintaxe simples**: Fácil de aprender e ler

## Instalação

1. Acesse [golang.org/dl](https://golang.org/dl/)
2. Baixe a versão para seu sistema operacional
3. Siga as instruções de instalação

Verifique a instalação:
```bash
go version
```

## Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Para executar:
```bash
go run main.go
```

## Estrutura Básica

- `package main`: Define o pacote principal executável
- `import`: Importa pacotes externos
- `func main()`: Função de entrada do programa

## Comandos Úteis

| Comando | Descrição |
|---------|-----------|
| `go run` | Compila e executa o código |
| `go build` | Compila o código e gera um executável |
| `go fmt` | Formata o código |
| `go mod init` | Inicializa um módulo Go |
| `go get` | Baixa dependências |

## Por que usar Go?

1. **Performance**: Próxima de C/C++
2. **Simplicidade**: Sintaxe limpa e minimalista
3. **Concorrência**: Excelente para sistemas distribuídos
4. **Ferramentas**: Ótimo ecossistema de ferramentas integradas
5. **Deploy simples**: Gera um único binário estático
