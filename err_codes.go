package libdns

func FormatErrorCode(err int) string {
	switch err {
	case NoError:
		return "No Error"
	case FormErr:
		return "Format Error"
	case ServFail:
		return "Server Failure"
	case NXDomain:
		return "Non-Existent Domain"
	case NotImp:
		return "Not Implemented"
	case Refused:
		return "Query Refused"
	case YXDomain:
		return "Name Exists when it should not"
	case YXRRSet:
		return "RR Set Exists when it should not"
	case NXRRSet:
		return "RR Set that should exist does not"
	case NotAuth:
		return "Server Not Authoritative for zone"
	case NotZone:
		return "Name not contained in zone"
	case 11, 12, 13, 14, 15:
		return "available for assignment"
	case BADVERS:
		return "Bad OPT Version"
	case BADSIG:
		return "TSIG Signature Failure"
	case BADKEY:
		return "Key not recognized"
	case BADTIME:
		return "Signature out of time window"
	case BADMODE:
		return "Bad TKEY Mode"
	case BADNAME:
		return "Duplicate key name"
	case BADALG:
		return "Algorithm not supported"
	}
	if 3841 <= err && err <= 4095 {
		return "Private Use"
	}
	return "available for assignment"
}

const (
	NoError  = 0
	FormErr  = 1
	ServFail = 2
	NXDomain = 3
	NotImp   = 4
	Refused  = 5
	YXDomain = 6
	YXRRSet  = 7
	NXRRSet  = 8
	NotAuth  = 9
	NotZone  = 10
	BADVERS  = 16
	BADSIG   = 17
	BADKEY   = 18
	BADTIME  = 19
	BADMODE  = 20
	BADNAME  = 21
	BADALG   = 22
)
