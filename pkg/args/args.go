package args

import "flag"

func Parse() (c, d, o *string){
	c = flag.String("c","", "path to config file")
	d = flag.String("d","", "path to data file")
	o = flag.String("o","", "path to output file")

	flag.Parse()

	return c, d, o
}