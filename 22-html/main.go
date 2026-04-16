package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	// Referenciando todos os arquivos que são .html
	templates := template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{
			"João",
			"Joaopedro@gmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta :5000.")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
