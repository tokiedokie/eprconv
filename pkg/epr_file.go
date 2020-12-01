package pkg

import (
	"bufio"
	"encoding/binary"
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
	cfgPath string
}

type eprFileInt8 struct {
	eprFile
}

type eprFileInt32 struct {
	eprFile
}

func newEprFileInt8(dataPath string) *eprFileInt8 {
	f := new(eprFileInt8)
	f.dataPath = dataPath
	return f
}

func newEprFileInt32(dataPath string, cfgPath string) *eprFileInt32 {
	f := new(eprFileInt32)
	f.dataPath = dataPath
	f.cfgPath = cfgPath
	return f
}

func (e *eprFileInt8) getData() []int8 {
	bufSize := e.dataSize()/1 // dividing by 4 because int8 has 1 byte.
	data := make([]int8, bufSize)
	readFile(e.dataPath, binary.BigEndian, &data)
	return data
}

func (e *eprFileInt32) getData() []int32 {
	bufSize := e.dataSize()/4 // dividing by 4 because int8 has 4 byte.
	data := make([]int32, bufSize)
	readFile(e.dataPath, binary.BigEndian, &data)
	return data
}

func (e *eprFile) getCfg() map[string]string {
	cfgFile, _ := os.Open(e.cfgPath)
	defer cfgFile.Close()
	
	scanner := bufio.NewScanner(cfgFile)

	cfgMap := make(map[string]string)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") || strings.HasPrefix(text, "*") { continue }
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
