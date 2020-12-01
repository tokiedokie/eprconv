package pkg

import "testing"

func TestAsumeFormat(t *testing.T) {
	if brukerBES3T != asumeFormat("data.dta") {
		t.Fatal()
	}
}

func TestLoadBrukerBES3T(t *testing.T) {
	path := "../test/data/bes3tint.dta"
	loadBrukerBES3T(path)
}
