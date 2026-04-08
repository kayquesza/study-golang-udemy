package main

import "fmt"

func main() {
	// canal := make(chan string) // Canal sem buffer
	canal := make(chan string, 2) // Canal com buffer: está dizendo que o canal tem uma capacidade de '2'
	canal <- "Olá, Mundo!"
	canal <- "Programando em Go!"
	// canal <- "Terceiro valor." // Dará um erro, excede o limite de '2'

	mensagem := <-canal
	mensagem2 := <-canal
	// mensagem3 := <-canal // Dará erro, excede o limite

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	// fmt.Println(mensagem3) // Erro

}
