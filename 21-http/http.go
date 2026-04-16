package main

import (
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de teste!"))
}

func main() {

	// HTTP: É um protocolo de comunicação - Base da comunicação web
	// Cliente: Faz a requisição - Servidor: processa requisição e envia resposta

	// Rotas: Identificar a mensagem que está sendo enviada e qual processamento o servidor precisará fazer com base nela
	// URI: Identificador do  recurso
	// Método: GET, POST, PUT, DELETE
	// GET: Buscar dados de um recurso
	// POST: Cadastrar dados
	// PUT: Atualizações; atualizar dados
	// DELETE: Apagar dados

	// Sintaxe: http.HandleFunc("<URI>, <FUNÇÃO PARA PROCESSAR A REQUISIÇÃO>")
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá, Mundo!")) // A mensagem obrigatoriamente precisa estar em um slice de byte
	})

	// A assinatura do método (w http.ResponseWriter, r *http.Request) continua igual
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários!")) // A mensagem obrigatoriamente precisa estar em um slice de byte
	})

	// Maneira "correta" de se criar (mais organizado):
	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
