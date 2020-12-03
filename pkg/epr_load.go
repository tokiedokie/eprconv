package pkg

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type fileFormat int

const (
	brukerBES3T fileFormat = iota + 1
)

func asumeFormat(filePath string) fileFormat {
	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".dta", ".dsc":
		return brukerBES3T
	}
	panic("format is not supported")
}

func EprLoad(filePath string) {
}

func getDataCfgPath(filePath string) (dataPath, cfgPath string, err error) {
	if filepath.Ext(filePath) == "" {
		pattern := fmt.Sprint(filePath, ".*")
		matchedFiles, globErr := filepath.Glob(pattern)
		if globErr != nil {
			err = globErr
			return
		}
		// if there is no format(e.g. BrukerBES3T) that needs one file
		// we should update if statement below
		if len(matchedFiles) == 0 {  
			err = errors.New("No such file")
			return
		}
		for _, file := range matchedFiles {
			switch strings.ToLower(filepath.Ext(file)) {
			case ".dta":
				dataPath = file
			case ".dsc":
				cfgPath = file
			default:
				err = errors.New("file extension does not match")
				return
			}
		}
	}
	return
}
