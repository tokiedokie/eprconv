package main

import (
	"fmt"
	"os"

	"eprconv/internal"
	"eprconv/internal/args"
	"eprconv/pkg/output"
)

func main() {
	parsedArgs, err := args.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	epr, err := internal.EprLoad(parsedArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputFile := output.NewOutputFile(*epr, parsedArgs.OutputPath, "")
	err = output.Output(*outputFile)
}
