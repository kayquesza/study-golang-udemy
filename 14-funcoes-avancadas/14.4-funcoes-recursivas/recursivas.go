package main

import "fmt"

// Função recursiva: que chama ela mesma

func fibonacci(posicao uint) uint {
	// Condição de parada (obrigatório; para evitar estouro de pilha[stack overflow])
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {
	fmt.Printf("\nFunções Rercusivas:\n")

	posicao := uint(10)

	for indice := uint(0); indice < posicao; indice++ {
		fmt.Println(fibonacci(indice))
	}
	fmt.Println(fibonacci(posicao))

}
