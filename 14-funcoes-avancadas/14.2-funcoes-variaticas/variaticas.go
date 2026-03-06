package main

import "fmt"

// Funções variáticas:
// função que pode receber 'n' paramêtros (não é necessário especificar, mesmo que seja nenhum)

// Sintaxe: irá receber de 1 a 'n' ints

// Nota: ao percorrer uma lista (slice, array ou maps) usando o "range"
// ele sempre duas coisas em cada volta do laço: o índice e o valor.

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

// O paramêtro variático precisa obrigatoriamente ser o último a ser mencionado.
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 100, 27)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo!", 10, 2, 3, 4, 12, 15)

}
