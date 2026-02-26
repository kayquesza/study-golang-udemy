package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("teste@teste.com") // Irá retornar <nil>
	fmt.Println(erro)

	erro1 := checkmail.ValidateFormat("testeteste") // Irá retornar um erro ("Invalid format")
	fmt.Println(erro1)
}
