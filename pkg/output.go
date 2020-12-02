package pkg

import (
	"fmt"
	"os"
)

func Output(path string, eprFile eprFile) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := eprFile.getData()

	for i := 0; i < len(eprFile.axes.x); i++ {
		fmt.Fprintf(file, "%e %e\n", eprFile.axes.x[i], data[i])
	}
}
