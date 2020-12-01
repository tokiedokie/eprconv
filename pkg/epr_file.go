package pkg

import (
	"encoding/binary"
	"os"
)

type eprFileMethod interface {
	getData()
	getCfg()
	dataSize() int64
}

type eprFile struct {
	filePath string
}

type eprFileInt8 struct {
	eprFile
}

type eprFileInt32 struct {
	eprFile
}

func newEprFileInt8(filePath string) *eprFileInt8 {
	f := new(eprFileInt8)
	f.filePath = filePath
	return f
}

func newEprFileInt32(filePath string) *eprFileInt32 {
	f := new(eprFileInt32)
	f.filePath = filePath
	return f
}

func (e *eprFileInt8) getData() []int8 {
	bufSize := e.dataSize()/1 // dividing by 4 because int8 has 1 byte.
	data := make([]int8, bufSize)
	readFile(e.filePath, binary.BigEndian, &data)
	return data
}

func (e *eprFileInt32) getData() []int32 {
	bufSize := e.dataSize()/4 // dividing by 4 because int8 has 4 byte.
	data := make([]int32, bufSize)
	readFile(e.filePath, binary.BigEndian, &data)
	return data
}

func (e *eprFile) getCfg() {
	// cfgMap := make(map[string]string)
}

func (e *eprFile) dataSize() int64 {
	info, err := os.Stat(e.filePath)
	if err != nil {
		panic(err)
	}
	return info.Size()
}
