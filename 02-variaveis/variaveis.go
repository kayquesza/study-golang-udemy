package main

import "fmt"

func main() {

	var variavel1 string = "Variável 1" // Declaração explícita de variável
	fmt.Println(variavel1)

	variavel2 := "Variável 2" // Declaração implícita de variável; é definida pelo valor recebido
	fmt.Println(variavel2)

	// Outra forma de declarar uma variável de maneira explícita
	var (
		variavel3 string = "Variável 3"
		variavel4 int    = 5
	)

	fmt.Println(variavel3, variavel4)

	// Outra forma de declarar uma variável de maneira implícita
	variavel5, variavel6 := "Variável 5", "12"
	fmt.Println(variavel5, variavel6)

	// "Constantes" são variáveis que não podem ser alteradas
	// Seguem o mesmo princípio.
	const constante1 = "Constante 1"
	fmt.Println(constante1)

	fmt.Printf("\n--------------------------------\n")
	// Inverter uma variável
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
