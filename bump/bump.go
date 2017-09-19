package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bump"
	app.Usage = "bump it dawg!"
	app.Action = bump
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "filename",
			Value: "VERSION",
			Usage: "filename to look for version in",
		},
	}
	app.Run(os.Args)
}

func bump(c *cli.Context) error {
	arg := "patch"
	if len(os.Args) < 2 {
		// log.Fatal("Invalid arg")
	} else {
		arg = os.Args[1]
		arg = strings.ToLower(arg)
	}

	filename := c.String("filename")
	vbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	vs := strings.TrimSpace(string(vbytes))
	fmt.Println("Current version:", vs)

	v := semver.New(vs)
	switch arg {
	case "major":
		v.BumpMajor()
	case "minor":
		v.BumpMinor()
	case "patch":
		v.BumpPatch()
	default:
		log.Fatalln("Invalid arg:", arg)
	}
	fmt.Println("New version:", v)

	err = ioutil.WriteFile(filename, []byte(v.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
