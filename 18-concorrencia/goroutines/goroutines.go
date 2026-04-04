package main

// goroutines são funções ou métodos que podem ser chamados e não
// necessariamente precisam terminar para que o programa continue executando

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá, Mundo!") // 'go' = goroutine. executa a função em paralelo com a linha de código abaixo
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
