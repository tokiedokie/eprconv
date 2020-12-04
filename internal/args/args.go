package args

import (
	"errors"
	"flag"
)

type ParsedArgs struct {
	CfgPath string
	DataPath string
	OutputPath string
}

func Parse() (parsedArgs ParsedArgs, err error) {
	c := flag.String("c", "", "path to config file")
	d := flag.String("d", "", "path to data file")
	o := flag.String("o", "", "path to output file")

	flag.Parse()

	if *c == "" {
		err = errors.New("Configuration file must be specified")
	} else if *d == "" {
		err = errors.New("Data file must be specified")
	} else if *o == "" {
		err = errors.New("Output file must be specified")
	}

	return ParsedArgs{
		CfgPath: *c,
		DataPath: *d,
		OutputPath: *o,
	}, err
}
