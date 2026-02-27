package main

import "fmt"

func main() {
	fmt.Printf("Ponteiros:\n")

	var variavel1 int = 1
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)
	// "variavel2" retornará o valor "1" porque, anteriormente, fez somente uma cópia do valor da "variavel1"

	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int // Utiliza-se do '*' para denominar um ponteiro

	// "var variavel3 int" guarda um *valor inteiro*
	// "var ponteiro *int" guarda o *endereço de memória* de um valor inteiro

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	// No endereço de memória "ponteiro" está salvo o valor de "variavel3" (no caso, 100)

	// Desreferenciação - Coloca-se um '*' na frente
	fmt.Println(variavel3, *ponteiro) // Lê o valor dentro do endereço de memória

	// O endereço de memória será o mesmo, mesmo se o valor da variável trocar
	variavel3 = 150
	fmt.Println(variavel3, ponteiro)

	// Substitui o valor, pois busca o valor atual de nova referência
	fmt.Println(variavel3, *ponteiro)

}
