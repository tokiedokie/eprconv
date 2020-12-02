package pkg

import (
	"bufio"
	"encoding/binary"
	"log"
	"os"
	"reflect"
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
	cfg        map[string]string
	fileFormat fileFormat
	axes       axes
}

type axes struct {
	x []float64
	y []float64
	z []float64
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

func NewEprFile(dataPath string, cfgPath string) *eprFile {
	f := new(eprFile)
	f.dataPath = dataPath
	f.cfgPath = cfgPath
	f.cfg = getCfg(f.cfgPath)
	f.fileFormat = asumeFormat(f.dataPath)
	f.axes = createAxes(f.cfg)
	return f
}

func (e *eprFile) getData() []float64 {
	var byteOrder binary.ByteOrder
	BSEQ, ok := e.cfg["BSEQ"]
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

	XPTS, ok := e.cfg["XPTS"]
	if !ok {
		panic("No XPTS in DSC file.")
	}
	xPoints, _ := strconv.Atoi(XPTS)

	var t reflect.Type
	switch e.cfg["IRFMT"] {
	case "C":
		t = reflect.TypeOf(int8(0))
	case "S":
		t = reflect.TypeOf(int16(0))
	case "I":
		t = reflect.TypeOf(int32(0))
	case "F":
		t = reflect.TypeOf(float32(0))
	case "D":
		t = reflect.TypeOf(float64(0))
	case "A":
		panic("Cannot read BES3T data in ASCII format!")
	case "0", "N":
		panic("No BES3T data!")
	default:
		panic("Unknown value for keyword IRFMT in .DSC file!")
	}

	return getMatrix(e.dataPath, byteOrder, reflect.ArrayOf(xPoints, t))
}

func (e *eprFile) dataSize() int64 {
	info, err := os.Stat(e.dataPath)
	if err != nil {
		panic(err)
	}
	return info.Size()
}
