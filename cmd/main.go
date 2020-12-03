package main

import (
	"fmt"
	"os"

	"eprconv/internal/args"
	"eprconv/pkg"
	"eprconv/pkg/output"
)

func main() {
	c, d, o, err := args.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	epr, err := pkg.NewEprFile(*d, *c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output.Output(*o, *epr)
}
