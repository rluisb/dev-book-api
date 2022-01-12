package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Will generate and return an instance of App
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CMD Application"
	app.Usage = "Find IPs and server names on the web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Find Ips addresses on the web",
			Flags: flags,
			Action: findIps,
		},
		{
			Name: "servers",
			Usage: "Find servers names on the web",
			Flags: flags,
			Action: findServers,
		},
	}

	return app
}

func findIps (c *cli.Context) {
	host := c.String("host")

	// package net call
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers (c *cli.Context) {
	host := c.String("host")

	// package net call
	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}