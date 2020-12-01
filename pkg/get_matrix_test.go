package pkg

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadOneByte(t *testing.T) {
	var data int32
	readFile(dataPath, binary.BigEndian, &data)
	assert.Equal(t, data, int32(-40))
}

func TestReadFile(t *testing.T) {
	eprFile := newEprFile(dataPath, cfgPath)
	bufSize := eprFile.dataSize() / 4
	data := make([]int32, bufSize)
	readFile(dataPath, binary.BigEndian, &data)
}
