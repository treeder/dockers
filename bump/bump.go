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
		cli.StringFlag{
			Name:  "input",
			Usage: "use this if you want to pass in a string to pass, rather than read it from a file. Cannot be used with --filename.",
		},
		cli.BoolFlag{
			Name:  "extract",
			Usage: "this will just find the version and return it, does not modify anything. Safe operation.",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
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

	// check for `[bump X]` in input, user can pass in git commit messages to auto bump different versions
	if strings.Contains(arg, "[bump minor]") {
		arg = "minor"
	} else if strings.Contains(arg, "[bump major]") {
		arg = "major"
	}

	var err error
	var vbytes []byte
	filename := c.String("filename")
	if c.IsSet("input") {
		vbytes = []byte(c.String("input"))
	} else {
		vbytes, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	old, new, loc, err := getAndBump(vbytes, arg)
	if err != nil {
		return err
	}
	if c.Bool("extract") {
		fmt.Println(old)
		return nil
	}

	fmt.Fprintln(os.Stderr, "Old version:", old)
	fmt.Fprintln(os.Stderr, "New version:", new)
	if !c.IsSet("input") {
		// write file
		loc1 := loc[1]
		len1 := loc[1] - loc[0]
		// fmt.Println("len1:", len1, len(v.String()))
		if len(new) > len1 {
			loc1 += len(new) - len1
		}
		b := vbytes[:loc[0]]
		b = append(b, []byte(new)...)
		b = append(b, vbytes[loc1:]...)
		// fmt.Println("writing:", string(b))

		err = ioutil.WriteFile(filename, b, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(new) // write it to stdout so scripts can use it
	return nil
}

func getAndBump(vbytes []byte, part string) (old string, new string, loc []int, err error) {
	re := regexp.MustCompile(`(\d+\.)?(\d+\.)?(\*|\d+)`)
	loc = re.FindIndex(vbytes)
	// fmt.Println(loc)
	if loc == nil {
		return "", "", nil, fmt.Errorf("Did not find semantic version")
	}
	vs := string(vbytes[loc[0]:loc[1]])

	v := semver.New(vs)
	switch part {
	case "major":
		v.BumpMajor()
	case "minor":
		v.BumpMinor()
	default:
		v.BumpPatch()
	}

	return vs, v.String(), loc, nil
}
