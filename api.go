package libdns

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/DanielOaks/go-idn/idna2003"
	"io"
	"regexp"
	"strings"
)

const (
	DomainPattern = "^(((?!-))(xn--|_{1,1})?[a-z0-9-]{0,61}[a-z0-9]{1,1}\\.)*(xn--)?([a-z0-9][a-z0-9\\-]{0,60}|[a-z0-9-]{1,30}\\.[a-z]{2,})$"
)

var matcher *regexp.Regexp
var calledInit bool

func Init() (err error) {
	matcher, err = regexp.Compile(DomainPattern)
	if err != nil {
		return err
	}
	calledInit = true
	return nil
}

/*
	Example header:
	AA AA - ID
	01 00 - Query parameters (QR | Opcode | AA | TC | RD | RA | Z | ResponseCode)
	00 01 - Number of questions
	00 00 - Number of answers
	00 00 - Number of authority records
	00 00 - Number of additional records
*/

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

type DNSQuery struct {
	ID     uint16 // An arbitrary 16 bit request identifier (same id is used in the response)
	QR     bool   // A 1 bit flat specifying whether this message is a query (0) or a response (1)
	Opcode uint8  // A 4 bit fields that specifies the query type; 0 (standard), 1 (inverse), 2 (status), 4 (notify), 5 (update)

	AA           bool  // Authoritative answer
	TC           bool  // 1 bit flag specifying if the message has been truncated
	RD           bool  // 1 bit flag to specify if recursion is desired (if the DNS server we secnd out request to doesn't know the answer to our query, it can recursively ask other DNS servers)
	RA           bool  // Recursive available
	Z            uint8 // Reserved for future use
	ResponseCode uint8

	QDCount uint16 // Number of entries in the question section
	ANCount uint16 // Number of answers
	NSCount uint16 // Number of authorities
	ARCount uint16 // Number of additional records

	Questions []DNSQuestion
}

func (q DNSQuery) Encode() ([]byte, error) {

	q.QDCount = uint16(len(q.Questions))

	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, q.ID)
	if err != nil {
		return nil, err
	}

	queryParams1 := byte(b2i(q.QR)<<7 | int(q.Opcode)<<3 | b2i(q.AA)<<1 | b2i(q.RD))
	queryParams2 := byte(b2i(q.RA)<<7 | int(q.Z)<<4)

	err = binary.Write(&buffer, binary.BigEndian, queryParams1)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, queryParams2)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, q.QDCount)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, q.ANCount)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, q.NSCount)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, q.ARCount)
	if err != nil {
		return nil, err
	}
	for _, question := range q.Questions {
		encoded, err := question.encode()
		if err != nil {
			return nil, err
		}
		buffer.Write(encoded)
	}

	return buffer.Bytes(), nil
}

/*
	Example Question:
	07 65 - 'example' has length 7, e
	78 61 - x, a
	6D 70 - m, p
	6C 65 - l, e
	03 63 - 'com' has length 3, c
	6F 6D - o, m
	00    - zero byte to end the QNAME
	00 01 - QTYPE
	00 01 - QCLASS
	76578616d706c6503636f6d0000010001
*/

type DNSQuestion struct {
	Domain string
	Type   uint16 // DNS Record type we are looking up; 1 (A record), 2 (authoritive name server)
	Class  uint16 // 1 (internet)
}

func (q DNSQuestion) validate() error {
	domain := strings.ToLower(q.Domain)
	ascii, err := idna2003.ToASCII(domain)
	if err != nil {
		return err
	}
	matched := matcher.MatchString(ascii)
	if !matched {
		return errors.New(fmt.Sprintf("did not match %s", ascii))
	}

	return nil
}

func (q DNSQuestion) encode() ([]byte, error) {
	var buffer bytes.Buffer
	err := q.validate()
	if err != nil {
		return nil, err
	}
	domainParts := strings.Split(q.Domain, ".")
	for _, part := range domainParts {
		if err := binary.Write(&buffer, binary.BigEndian, byte(len(part))); err != nil {
			return nil, err
		}

		for _, c := range part {
			if err := binary.Write(&buffer, binary.BigEndian, uint8(c)); err != nil {
				return nil, err
			}
		}
	}

	err = binary.Write(&buffer, binary.BigEndian, uint8(0))
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, q.Type)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, q.Class)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil

}
func (q DNSQuestion) Write(writer io.Writer) (int, error) {
	payload, err := q.encode()
	if err != nil {
		return -1, err
	}

	return writer.Write(payload)
}
