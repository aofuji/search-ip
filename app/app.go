package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de busca"
	app.Usage = "Busca IPs e nomes de servidro na internet"

	return app
}
