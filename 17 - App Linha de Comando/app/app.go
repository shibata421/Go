package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	flags := []cli.Flag{cli.StringFlag{
		Name:  "host",
		Value: "devbook.com.br",
	}}

	app.Commands = []cli.Command{{
		Name:   "ip",
		Usage:  "Busca IPs de endereço na internet",
		Flags:  flags,
		Action: buscarIps,
	}, {
		Name:   "servidores",
		Usage:  "Busca servidores de endereço na internet",
		Flags:  flags,
		Action: buscarServidores,
	}}

	return app
}

func buscarIps(cliContext *cli.Context) {
	host := cliContext.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(cliContext *cli.Context) {
	host := cliContext.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
