package main

import (
	"eprconv/internal"
	"eprconv/internal/args"
	"fmt"
	"os"
)

func main() {
	c, d, o, err := args.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	epr, err := internal.NewEprFile(*d, *c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	internal.Output(*o, *epr)
}
