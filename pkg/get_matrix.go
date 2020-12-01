package pkg

import (
	"encoding/binary"
	"fmt"
	"os"
)

func getMatrix(filePath string, byteOrder binary.ByteOrder, data interface{}) {
	readFile(filePath, byteOrder, data)
}

func readFile(filePath string, byteOrder binary.ByteOrder, data interface{}) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error, cannot open a file")
		panic(err)
	}

	err = binary.Read(file, byteOrder, data)
	if err != nil {
		fmt.Println("error, cannot read binary")
		panic(err)
	}

	file.Close()
}
