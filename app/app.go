package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// build the command line app
func Create() *cli.App {

	application := cli.NewApp()
	application.Name = "command line app"
	application.Usage = "searching for IPs address and server names on the internet"
	flags := defaultFlags()
	application.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Searching Ips from internet",
			Flags:  flags,
			Action: searchingIps,
		},
		{
			Name:   "servers",
			Usage:  "Searching server name on the internet",
			Flags:  flags,
			Action: searchingServers,
		},
	}
	return application
}

func searchingIps(context *cli.Context) {
	host := searchingHost(context)
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchingServers(context *cli.Context) {
	host := searchingHost(context)
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}

func defaultFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}
}

func searchingHost(context *cli.Context) string {
	return context.String("host")
}
