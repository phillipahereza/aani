package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("\nuse %q\n", "aani [port number]")
	}
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("error: missing port number\n")
		flag.Usage()
		os.Exit(1)
	}

	command := exec.Command("lsof", "-i", fmt.Sprintf("tcp:%s", flag.Arg(0)))
	command.Stdout = os.Stdout
	_ = command.Start()
}
