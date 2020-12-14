package output

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"eprconv/pkg"
)

type OutputFile struct {
	path      string
	separater string
	eprFile   pkg.EprFile
}

// require EprFile in case outputPath is empty
func NewOutputFile(eprFile pkg.EprFile, outputPath, separater string) *OutputFile {
	of := new(OutputFile)

	of.eprFile = eprFile

	if outputPath == "" {
		filePathWithoutExt := strings.TrimSuffix(eprFile.CfgPath, filepath.Ext(eprFile.CfgPath))
		of.path = fmt.Sprint(filePathWithoutExt, ".txt")
	} else {
		of.path = outputPath
	}

	of.separater = getSeparater(outputPath)

	return of
}

func getSeparater(outputPath string) string {
	var separater string
	switch strings.ToLower(filepath.Ext(outputPath)) {
	case ".csv":
		separater = ","
	case ".tsv":
		separater = "\t"
	default:
		separater = " "
	}
	return separater
}

func Output(outputFile OutputFile) error {
	file, err := os.Create(outputFile.path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := outputFile.eprFile.GetData()
	if err != nil {
		return err
	}

	format := fmt.Sprintf("%%.8e%s%%.8e\n", outputFile.separater)
	for i := 0; i < len(outputFile.eprFile.Axes.X); i++ {
		fmt.Fprintf(file, format, outputFile.eprFile.Axes.X[i], data[i])
	}
	return nil
}
