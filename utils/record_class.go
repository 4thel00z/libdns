package utils

type RecordClass int

const (
	InternetClass = RecordClass(0x0001)
	CSNetClass    = RecordClass(0x0002)
	ChaosClass    = RecordClass(0x0003)
	HESIODClass   = RecordClass(0x0004)
	NoneClass     = RecordClass(0x00fe)
	AllClass      = RecordClass(0x00ff)
	AnyClass      = RecordClass(0x00ff)
)
