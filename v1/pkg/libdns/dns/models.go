package dns

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/4thel00z/libdns/v1/pkg/libdns/utils"
	"github.com/DanielOaks/go-idn/idna2003"
	"io"
	"regexp"
	"strings"
)

type Response Query
type Query struct {
	ID     uint16 // An arbitrary 16 bit request identifier (same id is used in the response)
	QR     bool   // A 1 bit flat specifying whether this message is a query (0) or a response (1)
	Opcode uint8  // A 4 bit fields that specifies the query type; 0 (standard), 1 (inverse), 2 (status), 4 (notify), 5 (update)

	AA           bool // Authoritative answer
	TC           bool // 1 bit flag specifying if the message has been truncated
	RD           bool // 1 bit flag to specify if recursion is desired (if the DNS server we secnd out request to doesn't know the answer to our query, it can recursively ask other DNS servers)
	RA           bool // Recursive available
	Z            bool // Reserved for future use
	AD           bool
	CD           bool
	ResponseCode uint8

	QDCount uint16 // Number of entries in the question section
	ANCount uint16 // Number of answers
	NSCount uint16 // Number of authorities
	ARCount uint16 // Number of additional records

	Questions []Question
}

const (
	DomainPattern = "^(?:[_a-z0-9](?:[_a-z0-9-]{0,61}[a-z0-9]\\.)|(?:[0-9]+/[0-9]{2})\\.)+(?:[a-z](?:[a-z0-9-]{0,61}[a-z0-9])?)?$"
)

var matcher *regexp.Regexp
var calledInit bool

func Init() (err error) {
	if calledInit {
		return nil
	}
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

func DecodeQuery(payload []byte) (Query, error) {
	buf := bytes.NewReader(payload)
	var (
		ID           uint16
		secondRow    uint16
		QR           bool
		Opcode       uint8
		AA           bool // Authoritative answer
		TC           bool // 1 bit flag specifying if the message has been truncated
		RD           bool // 1 bit flag to specify if recursion is desired (if the DNS server we secnd out request to doesn't know the answer to our query, it can recursively ask other DNS servers)
		RA           bool // Recursive available
		Z            bool // Reserved for future use
		AD           bool
		CD           bool
		ResponseCode uint8

		QDCount uint16 // Number of entries in the question section
		ANCount uint16 // Number of answers
		NSCount uint16 // Number of authorities
		ARCount uint16 // Number of additional records

	)

	err := binary.Read(buf, binary.BigEndian, &ID)
	if err != nil {
		return Query{}, err
	}

	err = binary.Read(buf, binary.BigEndian, &secondRow)
	if err != nil {
		return Query{}, err
	}
	QR = utils.ReadQRBit(secondRow)
	Opcode = utils.ReadOpcode(secondRow)
	AA = utils.ReadAABit(secondRow)
	TC = utils.ReadTCBit(secondRow)
	RD = utils.ReadRDBit(secondRow)
	RA = utils.ReadRABit(secondRow)
	Z = utils.ReadZBit(secondRow)
	AD = utils.ReadADBit(secondRow)
	CD = utils.ReadCDBit(secondRow)
	ResponseCode = utils.ReadRCode(secondRow)

	err = binary.Read(buf, binary.BigEndian, &QDCount)
	if err != nil {
		return Query{}, err
	}
	err = binary.Read(buf, binary.BigEndian, &ANCount)
	if err != nil {
		return Query{}, err
	}
	err = binary.Read(buf, binary.BigEndian, &NSCount)
	if err != nil {
		return Query{}, err
	}
	err = binary.Read(buf, binary.BigEndian, &ARCount)
	if err != nil {
		return Query{}, err
	}

	return Query{
		ID:           ID,
		QR:           QR,
		Opcode:       Opcode,
		AA:           AA,
		TC:           TC,
		RD:           RD,
		RA:           RA,
		Z:            Z,
		AD:           AD,
		CD:           CD,
		ResponseCode: ResponseCode,
		QDCount:      QDCount,
		ANCount:      ANCount,
		NSCount:      NSCount,
		ARCount:      ARCount,
		//TODO: add Question parsing..
		Questions: nil,
	}, nil
}

func (q Query) Encode() ([]byte, error) {

	q.QDCount = uint16(len(q.Questions))

	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, q.ID)
	if err != nil {
		return nil, err
	}

	//QR|   Opcode  |AA|TC|RD|
	queryParams1 := byte(utils.B2i(q.QR)<<7 | int(q.Opcode)<<3 | utils.B2i(q.AA)<<1 | utils.B2i(q.RD))
	//RA| Z|AD|CD|   RCODE   |
	queryParams2 := byte(utils.B2i(q.RA)<<7 | utils.B2i(q.Z)<<6 | utils.B2i(q.AD)<<5 | utils.B2i(q.CD)<<4 | int(q.ResponseCode&0b00001111))

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

type Question struct {
	Domain string
	Type   uint16 // DNS Record type; look into utils.A, utils.NS, etc.
	Class  uint16 // DNS Class; look into
}

func (q Question) validate() error {
	domain := strings.ToLower(q.Domain)
	ascii, err := idna2003.ToASCII(domain)
	if err != nil {
		return err
	}
	if matcher == nil {
		return errors.New("Init() was not called!")
	}
	matched := matcher.MatchString(ascii)
	if !matched {
		return errors.New(fmt.Sprintf("did not match %s", ascii))
	}

	return nil
}

func DecodeQuestion(raw []byte) (Question, error) {
	var (
		q           = Question{}
		length      byte
		reader      io.Reader = bytes.NewReader(raw)
		domainParts           = []string{}
	)

	for {
		err := binary.Read(reader, binary.BigEndian, &length)
		if length == 0x0 {
			// Consumed the null byte
			break
		}
		if err == io.EOF {
			return Question{}, fmt.Errorf("unexpected EOF while decoding questions: %s %e", raw, err)
		}
		chars := make([]uint8, length)
		err = binary.Read(reader, binary.BigEndian, &chars)
		if err != nil {
			return Question{}, err
		}
		domainParts = append(domainParts, string(chars))
	}

	err := binary.Read(reader, binary.BigEndian, &q.Type)
	if err != nil {
		return Question{}, err
	}
	err = binary.Read(reader, binary.BigEndian, &q.Class)
	if err != nil {
		return Question{}, err
	}
	q.Domain = strings.Join(domainParts, ".")
	return q, nil
}

func (q Question) encode() ([]byte, error) {
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
func (q Question) Write(writer io.Writer) (int, error) {
	payload, err := q.encode()
	if err != nil {
		return -1, err
	}

	return writer.Write(payload)
}

func (q Query) Write(writer io.Writer) (int, error) {
	payload, err := q.Encode()
	if err != nil {
		return -1, err
	}

	return writer.Write(payload)
}
