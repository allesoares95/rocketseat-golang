### Estrutura básica
```go
// Switch compara um valor com múltiplos casos possíveis
func main() {
	fmt.Println("Quando é sabado?")
	today := time.Now().Weekday() // Pega o dia da semana atual

	switch time.Saturday { // Compara com sábado
	case today + 0: // Se sábado for hoje
		fmt.Println("Hoje!")
	case today + 1: // Se sábado for amanhã
		fmt.Println("Amanhã!")
	case today + 2: // Se sábado for em 2 dias
		fmt.Println("Em dois dias!")
	default: // Se nenhum caso anterior for verdadeiro
		fmt.Println("Ainda falta muito tempo para sábado.")
	}
}
```

### Switch com expressão direta
```go
// Switch avaliando uma variável diretamente
func main() {
	dia := 3

	switch dia { // Avalia o valor de 'dia'
	case 1: // Se dia == 1
		fmt.Println("Domingo")
	case 2: // Se dia == 2
		fmt.Println("Segunda")
	case 3: // Se dia == 3
		fmt.Println("Terça")
	default: // Qualquer outro valor
		fmt.Println("Outro dia")
	}
}
```

### Switch sem expressão (como if/else)
```go
// Switch sem expressão - avalia condições booleanas
func main() {
	nota := 85

	switch { // Sem expressão após switch
	case nota >= 90: // Se nota >= 90
		fmt.Println("A")
	case nota >= 80: // Se nota >= 80
		fmt.Println("B")
	case nota >= 70: // Se nota >= 70
		fmt.Println("C")
	default: // Se nenhuma condição for verdadeira
		fmt.Println("Reprovado")
	}
}
``` 

### Switch com múltiplos valores no case
```go
// Um case pode ter múltiplos valores separados por vírgula
func main() {
	dia := "sábado"

	switch dia {
	case "sábado", "domingo": // Se dia for sábado OU domingo
		fmt.Println("Final de semana!")
	case "segunda", "terça", "quarta", "quinta", "sexta": // Dias úteis
		fmt.Println("Dia útil")
	default:
		fmt.Println("Dia inválido")
	}
}
```

### Resumo do Switch

| Forma | Uso |
|-------|-----|
| `switch valor {}` | Compara valor com os cases |
| `switch {}` | Avalia condições booleanas (como if/else) |
| `case a, b, c:` | Múltiplos valores no mesmo case |
| `default:` | Executado se nenhum case for verdadeiro |