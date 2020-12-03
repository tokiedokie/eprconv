package main

import (
	"eprconv/pkg"
	"eprconv/pkg/args"
	"fmt"
	"os"
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
	pkg.Output(*o, *epr)
}
