package pkg

import "testing"

func TestAsumeFormat(t *testing.T) {
	if brukerBES3T != asumeFormat("data.dta") {
		t.Fatal()
	}
}