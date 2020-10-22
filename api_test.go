package libdns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoInit(t *testing.T) {
	validQuestion := DNSQuestion{
		Domain: "google.de",
		Type:   0x1,
		Class:  0x1,
	}

	_, err := validQuestion.encode()
	if err == nil {
		t.Fatal("Init() has to be called!")
	}
}

func TestDecodeDNSQuery(t *testing.T) {
	toDecode := []byte{
		0b11111111, 0b11111111,
		//1:QR|4:Opcode|1:AA|1:TC|1:RD|1:RA|1:Z|1:AD|1:CD|4:RCODE
		0b00000000, 0b00111111,
		0b00000000, 0b00111111,
		0b00000101, 0b00111001,
		0b00000000, 0b00000000,
		0b11111111, 0b11111111,
	}

	dnsQuery, err := DecodeDNSQuery(toDecode)
	if err != nil {
		t.Fatal(err)
	}
	expected := DNSQuery{
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
		QDCount:      63,
		ANCount:      1337,
		NSCount:      0,
		ARCount:      0b1111111111111111,
		Questions:    nil,
	}

	assert.Equal(t, expected, dnsQuery)
	/*
	TODO: this needs to work, but Encode is broken..

	encoded, err := dnsQuery.Encode()
	if err != nil {
		t.Fatal(err)
	}
	query, err := DecodeDNSQuery(encoded)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, query)*/

}
