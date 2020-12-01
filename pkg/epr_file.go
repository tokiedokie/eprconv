package pkg

import (
	"bufio"
	"encoding/binary"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type eprFileMethod interface {
	getData()
	dataSize() int64
}

type eprFile struct {
	dataPath   string
	cfgPath    string
	cfgMap     map[string]string
	fileFormat fileFormat
	axes       axes
}

type axes struct {
	x []float64
	y []float64
	z []float64
}

func asumeFormat(filePath string) fileFormat {
	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".dta", ".dsc":
		return brukerBES3T
	}
	panic("format is not supported")
}

func getCfg(cfgPath string) map[string]string {
	cfgFile, _ := os.Open(cfgPath)
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

func createAxes(cfgMap map[string]string) axes {
	axes := new(axes)

	// maybe we should use float64 for `MIN` and `WID`
	xPts, _ := strconv.Atoi(cfgMap["XPTS"])
	xMin, _ := strconv.Atoi(cfgMap["XMIN"])
	xWid, _ := strconv.Atoi(cfgMap["XWID"])
	switch cfgMap["XTYP"] {
	case "IDX":
		axes.x = createAxisIDX(xPts, xMin, xWid)
	}

	return *axes
}

func createAxisIDX(points, min, width int) []float64 {
	abscissa := make([]float64, points)
	minFloat := float64(min)
	widthFloat := float64(width)
	for i := 0; i < points; i++ {
		abscissa[i] = minFloat + widthFloat/float64(points-1)*float64(i)
	}
	return abscissa
}

func newEprFile(dataPath string, cfgPath string) *eprFile {
	f := new(eprFile)
	f.dataPath = dataPath
	f.cfgPath = cfgPath
	f.cfgMap = getCfg(f.cfgPath)
	f.fileFormat = asumeFormat(f.dataPath)
	f.axes = createAxes(f.cfgMap)
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

	XPTS, ok := e.cfgMap["XPTS"]
	if !ok {
		panic("No XPTS in DSC file.")
	}
	xPoints, _ := strconv.Atoi(XPTS)

	// TODO: fix this if there is a better way
	switch e.cfgMap["IRFMT"] {
	case "C":
		data := make([]int8, xPoints)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "S":
		data := make([]int16, xPoints)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "I":
		data := make([]int32, xPoints)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "F":
		data := make([]float32, xPoints)
		getMatrix(e.dataPath, byteOrder, &data)
		return data
	case "D":
		data := make([]float32, xPoints)
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

func (e *eprFile) dataSize() int64 {
	info, err := os.Stat(e.dataPath)
	if err != nil {
		panic(err)
	}
	return info.Size()
}
