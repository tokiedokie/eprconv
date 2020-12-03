package output

import (
	"fmt"
	"os"

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

	for i := 0; i < len(eprFile.Axes.X); i++ {
		fmt.Fprintf(file, "%e %e\n", eprFile.Axes.X[i], data[i])
	}
	return nil
}
