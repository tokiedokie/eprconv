package main

import (
	"fmt"
	"os"

	"eprconv/internal/args"
	"eprconv/pkg"
	"eprconv/pkg/output"
)

func main() {
	parsedArgs, err := args.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	epr, err := pkg.EprLoad(parsedArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = output.Output(*parsedArgs.OutputPath, *epr)
}
