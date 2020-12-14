package pkg

import (
	"bufio"
	"encoding/binary"
	"errors"
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

type EprFile struct {
	DataPath string
	CfgPath  string
	cfg      map[string]string
	Axes     axes
}

type axes struct {
	X []float64
	Y []float64
	Z []float64
}

func getCfg(cfgPath string) (map[string]string, error) {
	cfgFile, err := os.Open(cfgPath)
	if err != nil {
		return make(map[string]string), err
	}
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
			continue
		}
		cfgMap[kv[0]] = kv[1]
	}

	return cfgMap, nil
}

func createAxes(cfgMap map[string]string) (axes, error) {
	axes := new(axes)

	// maybe we should use float64 for `MIN` and `WID`
	xPts, err := strconv.Atoi(cfgMap["XPTS"])
	if err != nil {
		return *axes, errors.New("cannot find XPTS")
	}
	xMin, err := strconv.ParseFloat(cfgMap["XMIN"], 64)
	if err != nil {
		return *axes, errors.New("cannot find XMIN")
	}
	xWid, err := strconv.ParseFloat(cfgMap["XWID"], 64)
	if err != nil {
		return *axes, errors.New("cannot find XWID")
	}
	switch cfgMap["XTYP"] {
	case "IDX":
		axes.X = createAxisIDX(xPts, xMin, xWid)
	}

	return *axes, nil
}

func createAxisIDX(points int, min, width float64) []float64 {
	abscissa := make([]float64, points)
	for i := 0; i < points; i++ {
		abscissa[i] = min + width/float64(points-1)*float64(i)
	}
	return abscissa
}

func NewEprFile(dataPath string, cfgPath string) (*EprFile, error) {
	var err error
	f := new(EprFile)
	f.DataPath = dataPath
	f.CfgPath = cfgPath
	f.cfg, err = getCfg(f.CfgPath)
	if err != nil {
		return f, err
	}
	f.Axes, err = createAxes(f.cfg)
	if err != nil {
		return f, err
	}
	return f, nil
}

func (e *EprFile) GetData() ([]float64, error) {
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
	xPoints, err := strconv.Atoi(XPTS)
	if err != nil {
		return make([]float64, 0), err
	}

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

	return getMatrix(e.DataPath, byteOrder, reflect.ArrayOf(xPoints, t)), nil
}

func (e *EprFile) dataSize() (int64, error) {
	info, err := os.Stat(e.DataPath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
