package libdns

/*
   New OpCode assignments require an IETF Standards Action.
   Currently DNS OpCodes are assigned as follows:
*/
const (
	Query  = 0
	IQuery = 1
	Status = 2
	Notify = 4
	Update = 5
)

func FormatOpCode(oc int) string {
	switch oc {
	case Query:
		return "Query"
	case IQuery:
		return "Inverse Query"
	case Status:
		return "Status"
	case Notify:
		return "Notify"
	case Update:
		return "Update"
	case 3:
		fallthrough
	default:
		return "available for assignment"
	}
}
