package pkg

import (
	"encoding/binary"
	"testing"
)

func TestReadFile(t *testing.T) {
	var data int32
	readFile("../test/data/bes3tint.dta", binary.BigEndian, &data)
	if data != -40 {
		t.Fatal()
	}
}
