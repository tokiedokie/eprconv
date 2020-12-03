package internal

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("", "test.txt")
	defer os.Remove(tmpFile.Name())

	Output(tmpFile.Name(), *testEprFile)
}
