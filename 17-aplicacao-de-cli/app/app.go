package app

import "github.com/urfave/cli"

// "cli" é o pacote que usamos para a aplicação
// "app" é um tipo que está dentro do pacote
// 'Gerar' irá retornar a aplicação de CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"              // nome da aplicação
	app.Usage = "Busca IPs e Nomes de Servidor na Internet" // descrição da aplicação

	return app
}
