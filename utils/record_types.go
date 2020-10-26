package utils

type RecordType uint16

const (
	A     = RecordType(0x0001)
	NS    = RecordType(0x0002)
	MD    = RecordType(0x0003)
	MF    = RecordType(0x0004)
	CNAME = RecordType(0x0005)
	SOA   = RecordType(0x0006)
	MB    = RecordType(0x0007)
	MG    = RecordType(0x0008)
	MR    = RecordType(0x0009)
	NULL  = RecordType(0x000a)
	WKS   = RecordType(0x000b)
	PTR   = RecordType(0x000c)
	HINFO = RecordType(0x000d)
	MINFO = RecordType(0x000e)
	MX    = RecordType(0x000f)

	TXT     = RecordType(0x0010)
	RP      = RecordType(0x0011)
	AFSDB   = RecordType(0x0012)
	X25     = RecordType(0x0013)
	ISDN    = RecordType(0x0014)
	RT      = RecordType(0x0015)
	NSAP    = RecordType(0x0016)
	NSAPPTR = RecordType(0x0017)
	SIG     = RecordType(0x0018)
	KEY     = RecordType(0x0019)
	PX      = RecordType(0x001a)
	GPOS    = RecordType(0x001b)
	AAAA    = RecordType(0x001c)
	LOC     = RecordType(0x001d)
	NXT     = RecordType(0x001e)
	EID     = RecordType(0x001f)

	NIMLOC = RecordType(0x0020)
	SRV    = RecordType(0x0021)
	ATMA   = RecordType(0x0022)
	NAPTR  = RecordType(0x0023)
	KX     = RecordType(0x0024)
	CERT   = RecordType(0x0025)
	A6     = RecordType(0x0026)
	DNAME  = RecordType(0x0027)
	SINK   = RecordType(0x0028)
	OPT    = RecordType(0x0029)
	DS     = RecordType(0x002B)
	RRSIG  = RecordType(0x002E)
	NSEC   = RecordType(0x002F)

	DNSKEY = RecordType(0x0030)
	DHCID  = RecordType(0x0031)

	UINFO  = RecordType(0x0064)
	UID    = RecordType(0x0065)
	GID    = RecordType(0x0066)
	UNSPEC = RecordType(0x0067)

	ADDRS = RecordType(0x00f8)
	TKEY  = RecordType(0x00f9)
	TSIG  = RecordType(0x00fa)
	IXFR  = RecordType(0x00fb)
	AXFR  = RecordType(0x00fc)
	MAILB = RecordType(0x00fd)
	MAILA = RecordType(0x00fe)
	ANY   = RecordType(0x00ff)

	WINS  = RecordType(0xff01)
	WINSR = RecordType(0xff02)
)
