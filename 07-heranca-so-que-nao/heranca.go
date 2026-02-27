package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // Sem passar um tipo // Pega todos os campos da strut "pessoa" e joga para a struct "estudante"
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"name", "lastname", 22, 182}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia de Software", "GRAN"}
	fmt.Println(e1)

	e2 := estudante{p1, "Engenharia de Software", "GRAN"}
	fmt.Println(e2.nome) // Para pegar apenas o nome do estudante em questão

}
