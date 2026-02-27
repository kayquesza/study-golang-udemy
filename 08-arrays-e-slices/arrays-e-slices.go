package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("Arrays e Slies:\n")

	// Arrays
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println(slice)

	// "TypeOf" devolve o tipo da variável
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// Adiciona um elemento a variável
	slice = append(slice, 11)
	fmt.Println(slice)

	// Uma fatia de um array que já existe. Conta a partir do primeiro (1) e para antes, assim que chegar no (3)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Printf("\n---------------------------")
	// Função Make: Irá alocar um espaço na memória para alguma ação feita no programa

	// Sintaxe: "slice := make([]<tipo>, <tamanho>, <capacidadeMáxima>)"
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // "len" = "length" para ver o tamanho
	fmt.Println(cap(slice3)) // "cap" = "capacity" para ver a capacidade

	// "Make" cria um array, nesse caso, de 11 posições, e retorna um slice que pega as 10 primeiras posições desse array.

	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// Estourar o tamanho do slice
	slice3 = append(slice3, 2)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	// Ele dobra o tamanho do array, nesse caso, para 24 de capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 2)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// Em resumo
	// Array: lista de tamanho fixo
	// Slice: lista sem tamanho fixo

}
