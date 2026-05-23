package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicacao de linha de comandos pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	return app
}