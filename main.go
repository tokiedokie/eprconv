package main

import (
	"eprconv/pkg"
	"flag"
)

func main() {
	var (
		c = flag.String("c","", "path to config file")
		d = flag.String("d","", "path to data file")
		o = flag.String("o","", "path to output file")
	)
	flag.Parse()
	
	epr := pkg.NewEprFile(*d, *c)
	pkg.Output(*o, *epr)
}
