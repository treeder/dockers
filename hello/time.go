package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	timesPer := 50

	// just to create this one container for docker start command
	run("docker run --name reuse treeder/hello:sh")
	cmds := []string{
		// "./hello.sh",
		// "docker run treeder/hello:sh",
		// "docker run --rm treeder/hello:sh",
		"docker start -a reuse",
	}
	for _, cmd := range cmds {
		fmt.Println("Running:", cmd)
		totalDuration, err := runTimes(timesPer, cmd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("avg:", (totalDuration / time.Duration(timesPer)))
	}
}
func runTimes(x int, s string) (time.Duration, error) {
	totalDuration := time.Duration(0)
	for i := 0; i < x; i++ {
		td, err := run(s)
		if err != nil {
			return totalDuration, err
		}
		totalDuration += td
	}
	return totalDuration, nil
}

func run(s string) (time.Duration, error) {
	cmd := exec.Command("time", strings.Split(s, " ")...)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	start := time.Now()
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	duration := time.Now().Sub(start)
	// fmt.Println("duration:", duration)
	// o := outb.String()
	// e := errb.String()
	// fmt.Println("out:", o, "err:", e)
	return duration, nil
}
