package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando aplicação...")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args) // parâmetro padrão para ser reconhecido e executado
	if erro != nil {
		log.Fatal(erro)
	}
}
