package main

// Canal tem duas operações: enviar um dado e receber um dado
// Essas operações bloqueiam a operação do programa

import (
	"fmt"
	"time"
)

func main() {

	// Como criar um canal?
	// canal := make(chan tipo)
	// Só poder enviar e receber strings nesse caso

	canal := make(chan string)

	go escrever("Olá, Mundo", canal)

	fmt.Println("Depois da função 'escrever' ser executada.")

	//	for { // 'for' infitino. irá ficar esperando sempre que uma mensagem seja enviada
	//		mensagem := <-canal // o canal está esperando receber um valor (pela ordem da sintaxe feita)
	//		fmt.Println(mensagem)
	// }

	// extra:
	// 'deadlock' não é pego em compilação, apenas em execução

	// for { // dentro desse 'for', é avaliado se o canal ainda está aberto através do segundo parâmetro ('aberto')
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	fmt.Println(mensagem)
	// }

	// Basicamente:
	// Pra cada mensagem que for recebida no canal, enquanto estiver aberto, será printado na tela
	// Faz a mesma coisa que o código anterior
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	// Assim que identifica que o canal fechou, ele para a execução

	fmt.Println("Fim do Programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// "canal" é o nome da variável
		// sintaxe: canal <- texto
		canal <- texto // o canal está mandando um valor para dentro do 'canal'
		time.Sleep(time.Second)
	}

	close(canal) // Depois do loop (no caso, de 5x), isso fecha o canal, sem receber nem enviar dados
}
