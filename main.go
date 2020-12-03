package main

import (
	"eprconv/pkg"
	"eprconv/pkg/args"
)

func main() {
	c, d, o := args.Parse()
	epr, _ := pkg.NewEprFile(*d, *c)
	pkg.Output(*o, *epr)
}
