package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Forma de criar usando Struct
type cachorro struct {
	Nome  string `json:"nome"` // `json:exemplo` mostra qual a chave se tornará dentro do json
	Raca  string `json:"raca"` // o nome não necessariamente precisa ser o mesmo
	Idade uint   `json:"idade"`
}

func main() {
	// json.Marshal() // converter um map ou struct para json
	// json.Unmarshal()() // transforma um json em um struct ou map

	c := cachorro{"Rex", "Bordercolleman", 02}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)                  // retorna apenas um slice de bytes
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // cria um buffer baseado no slice de bytes e exibe o JSON como texto

	// Forma de criar usando Map
	c2 := map[string]string{
		"nome": "tobi",
		"raca": "puddo",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println()
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

}
