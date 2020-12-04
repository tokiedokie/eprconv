package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDataCfgPath(t *testing.T) {
	dp, cp, _ := getDataCfgPath("../test/data/bes3tint.dta")
	assert.Equal(t, dp, "../test/data/bes3tint.dta")
	assert.Equal(t, cp, "../test/data/bes3tint.dsc")
	dp, cp, _ = getDataCfgPath("../test/data/bes3tint")
	assert.Equal(t, dp, "../test/data/bes3tint.dta")
	assert.Equal(t, cp, "../test/data/bes3tint.dsc")
}
