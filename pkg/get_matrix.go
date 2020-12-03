package pkg

import (
	"encoding/binary"
	"os"
	"reflect"
)

func getMatrix(filePath string, byteOrder binary.ByteOrder, arrayType reflect.Type) []float64 {
	data := reflect.New(arrayType).Interface()
	readFile(filePath, byteOrder, data)
	return arrayInterfaceToFloat64(data)
}

func readFile(filePath string, byteOrder binary.ByteOrder, data interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Read(file, byteOrder, data)
	if err != nil {
		return err
	}
	return nil
}

func arrayInterfaceToFloat64(input interface{}) []float64 {
	array := reflect.Indirect(reflect.ValueOf(input))
	out := make([]float64, array.Len())
	for i := 0; i < array.Len(); i++ {
		out[i] = array.Index(i).Convert(reflect.TypeOf(float64(0))).Float()
	}
	return out
}
