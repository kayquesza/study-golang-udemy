package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
	// "default" é o que será executado caso não seja atendido nenhum dos casos
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // Claúsula que joga o código para dentro da próxima condição; sem sequer avaliar, ele só a considera instantaneamente.
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
		fallthrough
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Printf("\nSwitch:\n")

	dia := diaDaSemana(8)
	fmt.Println(dia)

	dia2 := diaDaSemana2(6)
	fmt.Println(dia2)

}
