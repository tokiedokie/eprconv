package pkg

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

func loadBrukerBES3T(filePath string) {
	// fint8 := newEprFileInt8(filePath)
}
