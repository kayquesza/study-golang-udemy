package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`  // `json:exemplo` mostra qual a chave se tornará dentro do json
	Raca  string `json:"raca"`  // o nome não necessariamente precisa ser o mesmo
	Idade uint   `json:"idade"` // para deixar nulo ou vázio, substitur o conteúdo por `json:-`
}

func main() {
	cachorroEmJSON := `{"nome": "rex", "raca": "pudol", "idade": 3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	} // converte o string para um slice de byte
	fmt.Println(c)

	cachorro2EmJSON := `{"nome": "tobi", "raca": "doberman"}`
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
