package output

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"eprconv/pkg"
)

func Output(path string, eprFile pkg.EprFile) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := eprFile.GetData()
	if err != nil {
		return err
	}

	separater := " "
	switch strings.ToLower(filepath.Ext(path)) {
	case ".csv":
		separater = ","
	case ".tsv":
		separater = "\t"
	}

	format := fmt.Sprintf("%%.8e%s%%.8e\n", separater)
	for i := 0; i < len(eprFile.Axes.X); i++ {
		fmt.Fprintf(file, format, eprFile.Axes.X[i], data[i])
	}
	return nil
}
