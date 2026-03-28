package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Snow-kal/ci-demo/internal/greeting"
)

func main() {
	name := flag.String("name", "CI/CD Learner", "name to show in greeting")
	flag.Parse()

	message, err := greeting.Build(*name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(message)
}
