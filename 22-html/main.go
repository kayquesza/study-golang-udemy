package main

import (
	"log"
	"net/http"
)

// Arquivo arquétipo para a próxima aula.

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de teste!"))
}

func main() {

	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
