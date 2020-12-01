package pkg

import (
	"path/filepath"
	"strings"
)

type fileFormat int

const (
	brukerBES3T fileFormat = iota + 1
)

func EprLoad(filePath string) {
	switch asumeFormat(filePath) {
	case brukerBES3T:
		loadBrukerBES3T(filePath)
	default:
		panic("cannot load a epr file")
	}
}

func asumeFormat(filePath string) fileFormat {
	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".dta", ".dsc":
		return brukerBES3T
	}
	panic("format is not supported")
}

func loadBrukerBES3T(filePath string) {
	// fint8 := newEprFileInt8(filePath)
}
