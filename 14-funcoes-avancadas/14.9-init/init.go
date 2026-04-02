package main

import "fmt"

var n int

// Função "init" é executada antes da função "main"
// Pode-se ter 1 função init por arquivo, diferente da main que é por pacote
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
