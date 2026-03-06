package main

import "fmt"

// Funções que referenciam variáveis que estão fora do seu corpo.

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {
	fmt.Printf("\nFunções Closure\n")

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
