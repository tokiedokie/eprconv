package pkg

import (
	"encoding/binary"
	"testing"
)

func TestReadOneByte(t *testing.T) {
	path := "../test/data/bes3tint.dta"
	var data int32
	readFile(path, binary.BigEndian, &data)
	if data != -40 {
		t.Fatal()
	}
}

func TestReadFile(t *testing.T) {
	path := "../test/data/bes3tint.dta"
	bufSize := fileSize(path)/4 // dividing by 4 because int32 has 4 bytes.
	data := make([]int32, bufSize)
	readFile(path, binary.BigEndian, &data)
}
