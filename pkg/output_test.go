package pkg

import "testing"

func TestOutput(t *testing.T) {
	Output("../tmp", *testEprFile)
}
