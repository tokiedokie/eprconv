package args

import (
	"errors"
	"flag"
)

type ParsedArgs struct {
	FilePath   string
	CfgPath    string
	DataPath   string
	OutputPath string
}

func Parse() (parsedArgs ParsedArgs, err error) {
	c := flag.String("c", "", "path to config file")
	d := flag.String("d", "", "path to data file")
	o := flag.String("o", "", "path to output file")

	flag.Parse()

	filePath := flag.Arg(0)
	if *c == "" && filePath == "" {
		err = errors.New("Configuration file must be specified")
	} else if *d == "" && filePath == "" {
		err = errors.New("Data file must be specified")
	}

	return ParsedArgs{
		FilePath:   filePath,
		CfgPath:    *c,
		DataPath:   *d,
		OutputPath: *o,
	}, err
}
