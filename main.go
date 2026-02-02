package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(ch chan int) {
	for valor := range ch {
		fmt.Println(valor)
	}

	fmt.Println("consumer finalizado")
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)
	go consumer(ch)

	time.Sleep(1 * time.Second)
}
