package main

import "fmt"

func main() {

	// É uma função anônima
	func(texto string) {
		fmt.Println(texto) // Faz isso aqui
	}("Passando Parâmetro") // "()" Assim que acabar de declarar, já executa

	// Armazenando o retorno para ser usado
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
