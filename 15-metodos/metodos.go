// Diferença de 'método' para 'função':
// Método está obrigatoriamente associado a alguma coisa
// Seja uma interface, uma struct...

// Métodos também podem ser definidos como 'ações' para a classe

package main

import "fmt"

func escrever() {
	fmt.Println("Escrevendo")
}

type usuario struct {
	nome  string
	idade uint8
}

// Por mais que não seja uma função, é declarado como uma
// Está dizendo: todos os usuários tem um método chamado 'salvar'
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário '%s' no banco de dados.\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Hello.")
	escrever()

	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
