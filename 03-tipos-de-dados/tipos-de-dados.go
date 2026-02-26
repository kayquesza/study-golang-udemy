package main

import (
	"errors"
	"fmt"
)

func main() {

	// NÚMEROS INTEIROS
	var numero int8 = 2
	fmt.Println(numero)

	// var numero2 uint8 = 200 // Erro; dará um overflow

	numero3 := 20000 // Utiliza a arquitetura do computador para determinar o tipo de dado
	fmt.Println(numero3)

	// var numero4 uint32 = -100000 // Erro; dará um overflow

	var numero5 rune = 12345 // "Rune" é exatamente igual ao "int32"
	fmt.Println(numero5)

	var numero6 byte = 123 // "Byte" é reconhecido como um "uint8"
	fmt.Println(numero6)

	// NÚMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	numeroReal2 := 1235.123
	fmt.Println(numeroReal2)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B' // Colocando entre 'aspas simples', trará o número da tabela ask que o mesmo corresponde.
	fmt.Println(char)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool // Por padrão, o valor é "False"
	fmt.Println(booleano2)

	// ERRO

	var err error
	fmt.Println(err) // Retornará "<nil>"

	var err1 error = errors.New("Erro Interno") // "errors." é uma biblioteca interna do Go
	fmt.Println(err1)

}
