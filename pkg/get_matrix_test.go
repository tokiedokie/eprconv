package pkg

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataPath = "../test/data/bes3tint.dta"
var cfgPath = "../test/data/bes3tint.dsc"

func TestReadOneByte(t *testing.T) {
	var data int32
	readBinaryFile(dataPath, binary.BigEndian, &data)
	assert.Equal(t, data, int32(-40))
}

func TestReadFile(t *testing.T) {
	eprFile, _ := NewEprFile(dataPath, cfgPath)
	dataSize, _ := eprFile.dataSize()
	bufSize := dataSize / 4
	data := make([]int32, bufSize)
	readBinaryFile(dataPath, binary.BigEndian, &data)
}
