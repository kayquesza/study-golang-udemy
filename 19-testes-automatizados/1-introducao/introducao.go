package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

// O go tem sua própria biblioteca para realização de testes

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)

	tipoEndereco1 := enderecos.TipoDeEndereco("Praça das Rosas")
	fmt.Println(tipoEndereco1)
}
