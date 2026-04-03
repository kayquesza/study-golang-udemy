package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// "cli" é o pacote que usamos para a aplicação
// "app" é um tipo que está dentro do pacote
// 'Gerar' irá retornar a aplicação de CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"              // nome da aplicação
	app.Usage = "Busca IPs e Nomes de Servidor na Internet" // descrição da aplicação

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de enredeços na Internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "amazon.com",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
