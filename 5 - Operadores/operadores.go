package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 4
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// Relacionais
	fmt.Println(1 > 2)  // Maior que
	fmt.Println(1 < 2)  // Menor que
	fmt.Println(1 == 2) // Igual a
	fmt.Println(1 <= 2) // Menor ou igual que
	fmt.Println(1 >= 2) // Maior ou igual que
	fmt.Println(1 != 2) // Diferente de

	// Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro) // Negação, responsável por inverter a resposta

	// Unários
	numero := 10
	numero++ // Aumenta em 1
	fmt.Println(numero)

	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero-- // Diminui em 1
	fmt.Println(numero)

	numero -= 4 // numero = numero - 4
	fmt.Println(numero)

	numero *= 3 // numero = numero * 3
	fmt.Println(numero)

	numero /= 2 // numero = numero / 2
	fmt.Println(numero)

	numero %= 2 // numero = numero % 2
	fmt.Println(numero)

	// Ternários
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
