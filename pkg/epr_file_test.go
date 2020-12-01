package pkg

import (
	"testing"
)

var dataPath = "../test/data/bes3tint.dta"
var cfgPath = "../test/data/bes3tint.dsc"
var testEprFile = newEprFileInt32(dataPath, cfgPath)

func TestGetCfg(t *testing.T) {
	cfgmap := testEprFile.getCfg()
	if cfgmap["BSEQ"] != "BIG" {
		t.Fatal()
	}
}
