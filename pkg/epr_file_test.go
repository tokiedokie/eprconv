package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataPath = "../test/data/bes3tint.dta"
var cfgPath = "../test/data/bes3tint.dsc"
var testEprFile = newEprFile(dataPath, cfgPath)

func TestAsumeFormat(t *testing.T) {
	assert.Equal(t, brukerBES3T, asumeFormat("data.dta"))
}

func TestCreateAxisIDX(t *testing.T) {
	expect := []float64{0, 2, 4, 6}
	actual := createAxisIDX(4, 0, 6)
	assert.Equal(t, expect, actual)
}

func TestCfgMap(t *testing.T) {
	expect := "BIG"
	actual := testEprFile.cfg["BSEQ"]
	assert.Equal(t, expect, actual)
}

func TestGetData(t *testing.T) {
	testEprFile.cfg["BSEQ"] = "LIT"
	data1 := testEprFile.getData()
	testEprFile.cfg["BSEQ"] = "BIG"
	data2 := testEprFile.getData()
	assert.NotEqual(t, data1, data2)
}

func TestNumberOfData(t *testing.T) {
	data1 := len(testEprFile.getData())
	data2 := len(testEprFile.axes.x)
	assert.Equal(t, data1, data2)
}
