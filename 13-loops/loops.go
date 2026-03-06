package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\nLoops:\n\n")

	i := 0

	for i < 10 {
		i += 5
		fmt.Printf("Incrementando i atual: %d\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j += 5 {
		fmt.Printf("Incrementando j: %d\n", j)
		time.Sleep(time.Second)
	}

	// Iterar (array, slice, map...)

	nomes := [3]string{"Joao", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice := range nomes {
		fmt.Println(indice)
	}

	// Retornará os números de acordo com a tabela ASCII
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	// Trará as letras ao invés do código da tabla
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Kayque",
		"sobrenome": "Alekssey",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Não é possível fazer range em struct

	// Loop infinito:
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
