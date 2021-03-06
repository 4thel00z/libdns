package dns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoInit(t *testing.T) {
	validQuestion := Question{
		Domain: "google.de",
		Type:   0x1,
		Class:  0x1,
	}

	_, err := validQuestion.encode()
	if err == nil {
		t.Fatal("Init() has to be called!")
	}
}

func TestQueryEncode(t *testing.T) {

	query := Query{
		ID:           0b1111111111111111,
		QR:           true,
		Opcode:       0,
		AA:           false,
		TC:           false,
		RD:           false,
		RA:           false,
		Z:            false,
		AD:           true,
		CD:           true,
		ResponseCode: 0b1111,
		QDCount:      63,
		ANCount:      1337,
		NSCount:      0,
		ARCount:      0b1111111111111111,
		Questions:    nil,
	}
	_, err := query.Encode()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeQuery(t *testing.T) {
	rawQuery := []byte{
		0b11111111, 0b11111111,
		//1:QR|4:Opcode|1:AA|1:TC|1:RD|1:RA|1:Z|1:AD|1:CD|4:RCODE
		0b00000000, 0b00111111,
		0b00000000, 0b00000000,
		0b00000101, 0b00111001,
		0b00000000, 0b00000000,
		0b11111111, 0b11111111,
	}

	dnsQuery, err := DecodeQuery(rawQuery)
	if err != nil {
		t.Fatal(err)
	}
	expected := Query{
		ID:           0b1111111111111111,
		QR:           false,
		Opcode:       0,
		AA:           false,
		TC:           false,
		RD:           false,
		RA:           false,
		Z:            false,
		AD:           true,
		CD:           true,
		ResponseCode: 0b1111,
		QDCount:      0,
		ANCount:      1337,
		NSCount:      0,
		ARCount:      0b1111111111111111,
		Questions:    nil,
	}

	assert.Equal(t, expected, dnsQuery)

	encoded, err := dnsQuery.Encode()
	if err != nil {
		t.Fatal(err)
	}
	query, err := DecodeQuery(encoded)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, query)

}

func TestDecodeQuestion(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatal(err)
	}
	question := Question{Domain: "ransomware.host", Type: 0x1, Class: 0x1}
	encoded, err := question.encode()
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeQuestion(encoded)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, question, decoded)
}
