package internal

import (
	"fmt"
	"io/ioutil"
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

func getDataCfgPath(filePath string) (dataPath, cfgPath string) {
	if filepath.Ext(filePath) == "" {
		dir, _ := filepath.Split(filePath)
		fileInfos, _ := ioutil.ReadDir(dir)
		for _, fileInfo := range fileInfos {
			fmt.Println(fileInfo.Name())
		}

		// get file path with an extention
		// if no files found, then throw error
	}
	switch asumeFormat(filePath) {

	}

	return
}
