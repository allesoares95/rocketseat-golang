## Concorrência em Go

### O que é concorrência?
Concorrência é a capacidade de executar **múltiplas tarefas ao mesmo tempo**. Go foi projetado com concorrência em mente, tornando muito fácil escrever programas concorrentes.

### Conceitos principais
- **Goroutine** - Uma função que executa de forma concorrente (leve, não é uma thread do sistema)
- **Channel** - Canal de comunicação entre goroutines
- **WaitGroup** - Sincroniza a execução de múltiplas goroutines

---

## Goroutines

### Goroutine básica
```go
// Goroutine é criada com a palavra-chave 'go' antes da chamada da função
func main() {
	go dizOla()                              // Inicia goroutine (executa em paralelo)
	fmt.Println("Função main")               // Continua executando sem esperar

	time.Sleep(time.Second)                  // Espera 1 segundo (sem isso, main termina antes)
}

func dizOla() {
	fmt.Println("Olá da goroutine!")         // Executa em paralelo com main
}
```

### Múltiplas goroutines
```go
// Várias goroutines executando ao mesmo tempo
func main() {
	for i := 1; i <= 5; i++ {
		go imprimir(i)                       // Cria 5 goroutines
	}

	time.Sleep(time.Second)                  // Espera todas terminarem
}

func imprimir(n int) {
	fmt.Printf("Goroutine %d\n", n)          // A ordem pode variar a cada execução!
}
```

### Duas goroutines executando em paralelo
```go
// Exemplo de duas funções executando simultaneamente
func exibirMsg() {
	for i := 0; i < 5; i++ {
		fmt.Println("Executando goroutine")  // Imprime a cada 100ms
		time.Sleep(100 * time.Millisecond)   // Pausa de 100ms entre cada print
	}
}

func exibirWorld() {
	for i := 0; i < 5; i++ {
		fmt.Println("World")                 // Imprime a cada 150ms
		time.Sleep(150 * time.Millisecond)   // Pausa de 150ms entre cada print
	}
}

func main() {
	go exibirMsg()                           // Inicia goroutine 1 (não bloqueia)

	go exibirWorld()                         // Inicia goroutine 2 (não bloqueia)
	
	time.Sleep(1 * time.Second)              // Espera 1 segundo para as goroutines executarem
	// Sem esse Sleep, main terminaria e as goroutines seriam canceladas!
}
// Saída intercalada (ordem pode variar):
// Executando goroutine
// World
// Executando goroutine
// World
// Executando goroutine
// Executando goroutine
// World
// ...
```

### Goroutine com função anônima
```go
// Função anônima como goroutine
func main() {
	go func() {                              // Goroutine com função anônima
		fmt.Println("Goroutine anônima!")
	}()                                      // () executa a função imediatamente

	time.Sleep(time.Second)
}
```

---

## Channels (Comunicação)

### Channel básico
```go
// Channel permite comunicação entre goroutines
func main() {
	ch := make(chan string)                  // Cria channel do tipo string

	go func() {
		ch <- "Olá do channel!"              // Envia valor para o channel
	}()

	msg := <-ch                              // Recebe valor do channel (bloqueia até receber)
	fmt.Println(msg)                         // Imprime: "Olá do channel!"
}
```

### Channel com múltiplos valores
```go
// Enviando e recebendo múltiplos valores
func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i                          // Envia números 1 a 5
		}
		close(ch)                            // Fecha o channel (importante!)
	}()

	for num := range ch {                    // Recebe até o channel ser fechado
		fmt.Println(num)
	}
}
```

### Buffered Channel
```go
// Channel com buffer: pode armazenar valores sem bloquear
func main() {
	ch := make(chan int, 3)                  // Buffer de tamanho 3

	ch <- 1                                  // Não bloqueia (buffer tem espaço)
	ch <- 2
	ch <- 3
	// ch <- 4                               // Bloquearia! Buffer cheio

	fmt.Println(<-ch)                        // Imprime: 1
	fmt.Println(<-ch)                        // Imprime: 2
	fmt.Println(<-ch)                        // Imprime: 3
}
```

### Select (múltiplos channels)
```go
// Select escolhe qual channel está pronto
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "canal 1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "canal 2"
	}()

	select {                                 // Espera o primeiro que estiver pronto
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)                    // Este executa primeiro (50ms < 100ms)
	}
}
```

---


## WaitGroup (Sincronização)

### Usando WaitGroup para esperar goroutines
```go
// WaitGroup é a forma correta de esperar goroutines (em vez de time.Sleep)
func main() {
	var wg sync.WaitGroup                    // Cria WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)                            // Adiciona 1 ao contador (antes de criar a goroutine)
		go func(n int) {
			defer wg.Done()                  // Decrementa contador quando terminar
			fmt.Printf("Goroutine %d\n", n)
		}(i)                                 // Passa i como parâmetro (evita closure problem)
	}

	wg.Wait()                                // Bloqueia até contador chegar a 0
	fmt.Println("Todas as goroutines terminaram!")
}
```

---


## Exemplo prático: Worker Pool

```go
// Worker Pool: múltiplos workers processando jobs de uma fila
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {                  // Recebe jobs até o channel fechar
		fmt.Printf("Worker %d processando job %d\n", id, job)
		time.Sleep(time.Second)              // Simula trabalho
		results <- job * 2                   // Envia resultado
	}
}

func main() {
	jobs := make(chan int, 100)              // Channel de jobs
	results := make(chan int, 100)           // Channel de resultados

	// Cria 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Envia 9 jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)                              // Fecha channel de jobs

	// Coleta resultados
	for r := 1; r <= 9; r++ {
		fmt.Println("Resultado:", <-results)
	}
}
```

---

## Resumo de Concorrência

### Goroutines

| Sintaxe | Significado |
|---------|-------------|
| `go func()` | Inicia goroutine |
| `go func(){}()` | Goroutine com função anônima |

### WaitGroup

| Método | Uso |
|--------|-----|
| `wg.Add(n)` | Adiciona n ao contador |
| `wg.Done()` | Decrementa contador em 1 |
| `wg.Wait()` | Bloqueia até contador = 0 |

### Channels

| Sintaxe | Significado |
|---------|-------------|
| `make(chan T)` | Cria channel do tipo T |
| `make(chan T, n)` | Cria channel com buffer de tamanho n |
| `ch <- valor` | Envia valor para o channel |
| `<-ch` | Recebe valor do channel |
| `close(ch)` | Fecha o channel |
| `for v := range ch` | Itera até channel fechar |

### Select

| Sintaxe | Significado |
|---------|-------------|
| `select { case ... }` | Espera primeiro channel pronto |
| `default:` | Executa se nenhum channel pronto |

---

## Dicas importantes

⚠️ **Nunca use `time.Sleep` para sincronização** - Use WaitGroup ou channels

⚠️ **Sempre feche channels** quando não vai mais enviar valores

⚠️ **Cuidado com race conditions** - Use `go run -race main.go` para detectar

⚠️ **Passe variáveis de loop como parâmetro** para evitar closure problems
