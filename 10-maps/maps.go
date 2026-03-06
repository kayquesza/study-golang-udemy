package main

import "fmt"

func main() {

	fmt.Println("Maps in Go")

	// Sintaxe: map[tipoDasChaves]tipoDosValores
	usuario := map[string]string{
		"nome":      "testeDeNome",
		"sobrenome": "testeDeSobrenome",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "fulano",
			"ultimo":   "ciclano",
		},
		"curso": {
			"nome":      "Engenharia de Software",
			"faculdade": "GRAN",
		},
	}

	fmt.Printf("\n\n%v\n\n", usuario2)
	delete(usuario2, "nome")
	fmt.Printf("\n\n%v\n\n", usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "escorpiao",
	}

	fmt.Printf("\n\n%v\n\n", usuario2)

}
