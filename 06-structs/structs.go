package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Arquivo: Structs")

	// Formas de popular um struct
	//Forma 1
	var u usuario
	u.nome = "name"
	u.idade = 21
	fmt.Println(u)

	//Forma 2
	enderecoExemplo := endereco{"Rua exemplar", 2}
	u2 := usuario{"name", 21, enderecoExemplo}
	fmt.Println(u2)

	//Forma 3 - Passando apenas 1 campo
	u3 := usuario{nome: "name"}
	fmt.Println(u3)

	u4 := usuario{idade: 21}
	fmt.Println(u4)

}
