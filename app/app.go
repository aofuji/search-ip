package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de busca"
	app.Usage = "Busca IPs e nomes de servidro na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPS de enderecos na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "teste.com.br",
				},
			},
			Action: buscaIps,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// go run main.go ip --host amazon.com.br
