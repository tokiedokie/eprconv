package pkg

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	var data float64
	readFile("../test/data/bes3tint.dta", &data)
}
