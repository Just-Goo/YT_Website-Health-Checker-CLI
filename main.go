package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	flags := []cli.Flag{
		&cli.StringFlag{
			Name: "domain",
			Aliases: []string{"d"},
			Usage: "Domain name to check.", // E.g - google.com
			Required: true,
		},
		&cli.StringFlag{
			Name: "port",
			Aliases: []string{"p"},
			Usage: "Port number to check.",
			Required: false,
		},
	}

	action := func (c *cli.Context) error {
		port := c.String("port")
		if port == "" {
			port = "80" // Use the default port '80'
		}

		status := Check(c.String("domain"), port)
		fmt.Println(status)
		return nil

	}

	app := &cli.App{
		Name: "Health Checker",
		Usage: "A simple tool that checks if a website is running or is down" ,
		Flags: flags,
		Action: action,
	}

	err := app.Run(os.Args) // Run the app while passing in the 'os.Args' (OS arguments)
	if err != nil {
		log.Fatal(err)
	}

}