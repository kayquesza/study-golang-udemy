package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	// "*numero" = Buscar no endereço de memória o valor do número
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido) // -20
	fmt.Println(numero)          // 20, pois o anterior foi enviado somente uma cópia do número

	novoNumero := 40
	fmt.Println(novoNumero)
	// & é utilizado para que reconheça que é um ponteiro
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
