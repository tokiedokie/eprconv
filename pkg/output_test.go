package pkg

import "testing"

func TestOutput(t *testing.T) {
	output("../tmp", *testEprFile)
}
