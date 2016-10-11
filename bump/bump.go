package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/coreos/go-semver/semver"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid arg")
	}
	arg := os.Args[1]
	if arg == "" {
		arg = "patch"
	} else {
		arg = strings.ToLower(arg)
	}

	filename := "VERSION"
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
}
