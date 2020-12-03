package output

import (
	"io/ioutil"
	"os"
	"testing"

	"eprconv/pkg"
)

func TestOutput(t *testing.T) {
	var dataPath = "../../test/data/bes3tint.dta"
	var cfgPath = "../../test/data/bes3tint.dsc"
	var testEprFile, _ = pkg.NewEprFile(dataPath, cfgPath)
	tmpFile, _ := ioutil.TempFile("", "test.txt")
	defer os.Remove(tmpFile.Name())

	Output(tmpFile.Name(), *testEprFile)
}
