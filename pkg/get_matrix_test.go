package pkg

import (
	"encoding/binary"
	"testing"
)

func TestReadOneByte(t *testing.T) {
	var data int32
	readFile(dataPath, binary.BigEndian, &data)
	if data != -40 {
		t.Fatal()
	}
}

func TestReadFile(t *testing.T) {
	eprFile := newEprFile(dataPath, cfgPath)
	bufSize := eprFile.dataSize()/4
	data := make([]int32, bufSize)
	readFile(dataPath, binary.BigEndian, &data)
}
