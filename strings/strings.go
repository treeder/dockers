package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "strings"
	app.Usage = "strings yo"
	app.Action = cli.ShowAppHelp
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filename",
			Value: "VERSION",
			Usage: "filename to look for version in",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "contains",
			Usage:  "returns true if first argument contains the second argument",
			Action: contains,
		},
	}
	app.Run(os.Args)
}

func bump(c *cli.Context) error {
	arg := "patch"
	fmt.Println("ARGS:", c.Args())
	if len(c.Args()) < 1 {
		log.Fatalln("Invalid arg")
	} else {
		arg = c.Args().First()
		arg = strings.ToLower(arg)
	}

	return nil
}

func contains(c *cli.Context) error {
	if len(c.Args()) < 2 {
		return errors.New("contains requires 2 parameters, the string and the string to find")
	}
	s := c.Args().Get(0)
	find := c.Args().Get(1)
	if strings.Contains(s, find) {
		write("true")
		return nil
	}
	write("false")
	return nil
}
func write(s string) {
	fmt.Print(s)
}
