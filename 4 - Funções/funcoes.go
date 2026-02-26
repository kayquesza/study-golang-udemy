package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Recebe 2 paramêtros e dará 2 retornos, ao invés de somente 1
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// Ou
func calculosMatematicos1(n1, n2 int8) (soma int8, subtracao int8) {
	soma = n1 + n2
	subtracao = n1 - n2
	return // Retorno "limpo", ele já sabe o que retornar
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função 'f'")
	}

	// Maneira de chamar a função criada acima
	f()

	var f1 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	// Maneira de chamar a função criada acima
	resultado := f1("Texto da função f1")
	fmt.Println(resultado)

	// Função com 2 retornos
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Função com 2 retornos, mas buscando só 1 resultado
	resultadoSoma1, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma1)

	_, resultadoSubtracao1 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao1)

}
