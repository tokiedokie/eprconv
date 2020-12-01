package pkg

import (
	"testing"
)

var dataPath = "../test/data/bes3tint.dta"
var cfgPath = "../test/data/bes3tint.dsc"
var testEprFile = newEprFile(dataPath, cfgPath)

func TestAsumeFormat(t *testing.T) {
	if brukerBES3T != asumeFormat("data.dta") {
		t.Fatal()
	}
}

func TestCfgMap(t *testing.T) {
	cfgmap := testEprFile.cfgMap
	if cfgmap["BSEQ"] != "BIG" {
		t.Fatal()
	}
}

func TestGetData(t *testing.T) {
	testEprFile.cfgMap["BSEQ"] = "LIT"
	data1 := testEprFile.getData().([]int32)
	testEprFile.cfgMap["BSEQ"] = "BIG"
	data2 := testEprFile.getData().([]int32)
	if data1[0] == data2[0] {
		t.Fatal()
	}
}
