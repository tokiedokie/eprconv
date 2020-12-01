package pkg

import (
	"bufio"
	"encoding/binary"
	"log"
	"os"
	"strings"
)

type eprFileMethod interface {
	getData()
	getCfg()
	dataSize() int64
}

type eprFile struct {
	dataPath string
	cfgPath  string
	cfgMap   map[string]string
}

func newEprFile(dataPath string, cfgPath string) *eprFile {
	f := new(eprFile)
	f.dataPath = dataPath
	f.cfgPath = cfgPath
	f.cfgMap = f.getCfg()
	return f
}

func (e *eprFile) getData() interface{} {
	var byteOrder binary.ByteOrder
	BSEQ, ok := e.cfgMap["BSEQ"]
	if !ok {
		log.Println("Keyword BSEQ not found in .DSC file! Assuming BSEQ=BIG.")
		byteOrder = binary.BigEndian
	} else if BSEQ == "BIG" {
		byteOrder = binary.BigEndian
	} else if BSEQ == "LIT" {
		byteOrder = binary.LittleEndian
	} else {
		panic("Unknown value for keyword BSEQ in .DSC file!")
	}

	// TODO: fix this if there is a better way
	switch e.cfgMap["IRFMT"] {
	case "C":
		data := make([]int8, e.dataSize()/1)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "S":
		data := make([]int16, e.dataSize()/2)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "I":
		data := make([]int32, e.dataSize()/4)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "F":
		data := make([]float32, e.dataSize()/4)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "D":
		data := make([]float32, e.dataSize()/8)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "A":
		panic("Cannot read BES3T data in ASCII format!")
	case "0", "N":
		panic("No BES3T data!")
	default:
		panic("Unknown value for keyword IRFMT in .DSC file!")
	}
}

func (e *eprFile) getCfg() map[string]string {
	cfgFile, _ := os.Open(e.cfgPath)
	defer cfgFile.Close()

	scanner := bufio.NewScanner(cfgFile)

	cfgMap := make(map[string]string)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") || strings.HasPrefix(text, "*") {
			continue
		}
		kv := strings.Fields(text)
		if len(kv) < 2 {
			cfgMap[kv[0]] = ""
			continue
		}
		cfgMap[kv[0]] = kv[1]
	}

	return cfgMap
}

func (e *eprFile) dataSize() int64 {
	info, err := os.Stat(e.dataPath)
	if err != nil {
		panic(err)
	}
	return info.Size()
}
