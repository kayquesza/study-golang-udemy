package main

import "fmt"

func main() {
	// O que seria o 'worker pools'?
	// Uma fila de tarefas pra serem executadas e se tem vários
	// processos que pegam itens dessa fila de maneira independente

	tarefas := make(chan int, 30)
	resultados := make(chan int, 30)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 30; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 30; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

// '<- chan' significa: é um canal que apenas recebe dados
// 'chan<-' significa: um canal que apenas envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
