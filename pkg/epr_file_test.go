package pkg

import (
	"testing"
)

var dataPath = "../test/data/bes3tint.dta"
var cfgPath = "../test/data/bes3tint.dsc"
var testEprFile = newEprFile(dataPath, cfgPath)

func TestGetCfg(t *testing.T) {
	cfgmap := testEprFile.getCfg()
	if cfgmap["BSEQ"] != "BIG" {
		t.Fatal()
	}
}

func TestGetData(t *testing.T) {
	testEprFile.getData()
}
