package pkg

import (
	"encoding/binary"
	"fmt"
	"os"
)

func getMatrix() {
	
}

func readFile(filePath string, data interface{}) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error, cannot open a file")
		panic(err)
	}
	
	err = binary.Read(file, binary.BigEndian, data)
	if err != nil {
		fmt.Println("error, cannot read binary")
		panic(err)
	}
}
