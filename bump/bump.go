package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
	// fmt.Println("ARGS:", c.Args())
	if len(c.Args()) < 1 {
		// log.Fatal("Invalid arg")
	} else {
		arg = c.Args().First()
		arg = strings.ToLower(arg)
	}

	filename := c.String("filename")
	vbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(\d+\.)?(\d+\.)?(\*|\d+)`)
	loc := re.FindIndex(vbytes)
	// fmt.Println(loc)
	vs := string(vbytes[loc[0]:loc[1]])
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

	b := vbytes[:loc[0]]
	b = append(b, []byte(v.String())...)
	b = append(b, vbytes[loc[1]:]...)
	// fmt.Println(string(b))

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
