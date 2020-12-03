package args

import (
	"errors"
	"flag"
)

func Parse() (c, d, o *string, err error) {
	c = flag.String("c", "", "path to config file")
	d = flag.String("d", "", "path to data file")
	o = flag.String("o", "", "path to output file")

	flag.Parse()

	if *c == "" {
		err = errors.New("Configuration file must be specified")
	} else if *d == "" {
		err = errors.New("Data file must be specified")
	} else if *o == "" {
		err = errors.New("Output file must be specified")
	}

	return c, d, o, err
}
