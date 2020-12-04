package internal

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"eprconv/internal/args"
	"eprconv/pkg"
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

func EprLoad(parsedArgs args.ParsedArgs) (*pkg.EprFile, error) {
	if parsedArgs.CfgPath != "" && parsedArgs.DataPath != "" {
		return pkg.NewEprFile(parsedArgs.DataPath, parsedArgs.CfgPath)
	}
	return nil, nil
}

func getDataCfgPath(filePath string) (dataPath, cfgPath string, err error) {
	filePathWithoutExt := strings.TrimSuffix(filePath, filepath.Ext(filePath))
	pattern := fmt.Sprint(filePathWithoutExt, ".*")
	matchedFiles, globErr := filepath.Glob(pattern)
	if globErr != nil {
		err = globErr
		return
	} else if len(matchedFiles) == 0 {
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
	return
}
