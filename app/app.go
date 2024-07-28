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

	application.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Searching Ips from internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: searchingIps,
		},
	}
	return application
}

func searchingIps(context *cli.Context) {
	host := context.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
