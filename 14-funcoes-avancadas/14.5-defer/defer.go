package main

import "fmt"

func funcao1() {
	fmt.Printf("Executando a Função 1\n")
}

func funcao2() {
	fmt.Printf("Executando a Função 2\n")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado.")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1()
	// Clausula: Defer = Adiar, até o último momento possível
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 7))
}
