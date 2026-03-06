package main

import "fmt"

// Retorno nomeado: irá retornar uma variável chamada "soma" e outra "subtracao"
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
	// Sem a necessidade de acresceta-los após o "return"
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
