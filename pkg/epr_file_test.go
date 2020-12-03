package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsumeFormat(t *testing.T) {
	assert.Equal(t, brukerBES3T, asumeFormat("data.dta"))
}

func TestCreateAxisIDX(t *testing.T) {
	expect := []float64{0, 2, 4, 6}
	actual := createAxisIDX(4, 0, 6)
	assert.Equal(t, expect, actual)
}

func TestCfgMap(t *testing.T) {
	var dataPath = "../test/data/bes3tint.dta"
	var cfgPath = "../test/data/bes3tint.dsc"
	var testEprFile, _ = NewEprFile(dataPath, cfgPath)
	expect := "BIG"
	actual := testEprFile.cfg["BSEQ"]
	assert.Equal(t, expect, actual)
}

func TestGetData(t *testing.T) {
	var dataPath = "../test/data/bes3tint.dta"
	var cfgPath = "../test/data/bes3tint.dsc"
	var testEprFile, _ = NewEprFile(dataPath, cfgPath)
	testEprFile.cfg["BSEQ"] = "LIT"
	data1, _ := testEprFile.GetData()
	testEprFile.cfg["BSEQ"] = "BIG"
	data2, _ := testEprFile.GetData()
	assert.NotEqual(t, data1, data2)
}

func TestNumberOfData(t *testing.T) {
	var dataPath = "../test/data/bes3tint.dta"
	var cfgPath = "../test/data/bes3tint.dsc"
	var testEprFile, _ = NewEprFile(dataPath, cfgPath)
	data, _ := testEprFile.GetData()
	data1 := len(data)
	data2 := len(testEprFile.Axes.X)
	assert.Equal(t, data1, data2)
}
