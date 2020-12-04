package internal

import (
	"eprconv/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsumeFormat(t *testing.T) {
	assert.Equal(t, brukerBES3T, asumeFormat("data.dta"))
}

func TestGetDataCfgPath(t *testing.T) {
	dp, cp, _ := getDataCfgPath("../test/data/bes3tint.dta")
	assert.Equal(t, dp, "../test/data/bes3tint.dta")
	assert.Equal(t, cp, "../test/data/bes3tint.dsc")
	dp, cp, _ = getDataCfgPath("../test/data/bes3tint")
	assert.Equal(t, dp, "../test/data/bes3tint.dta")
	assert.Equal(t, cp, "../test/data/bes3tint.dsc")
}

func TestAllTestData(t *testing.T) {
	var paths = []string{
		"../test/data/010_cutpp_10kfs",
		"../test/data/2010_06_25_IKKG_C95_2pESEEM",
		"../test/data/00012107",
		"../test/data/99090211",
		"../test/data/be3tintlit",
		"../test/data/bes3tint",
	}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			data, cfg, _ := getDataCfgPath(path)
			e, _ := pkg.NewEprFile(data, cfg)
			e.GetData()
		})
	}
}
