package pkg

import (
	"encoding/binary"
	"fmt"
	"os"
)

func readFile(filePath string, endian binary.ByteOrder, data interface{}) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error, cannot open a file")
		panic(err)
	}

	err = binary.Read(file, endian, data)
	if err != nil {
		fmt.Println("error, cannot read binary")
		panic(err)
	}

	file.Close()
}

func fileSize(filePath string) int64 {
	info, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}

	return info.Size()
}
