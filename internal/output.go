package internal

import (
	"fmt"
	"os"
)

func Output(path string, eprFile eprFile) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := eprFile.getData()
	if err != nil {
		return err
	}

	for i := 0; i < len(eprFile.axes.x); i++ {
		fmt.Fprintf(file, "%e %e\n", eprFile.axes.x[i], data[i])
	}
	return nil
}
