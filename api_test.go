package libdns

import (
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
