package v1

import (
	"bufio"
	"github.com/4thel00z/libdns/v1/pkg/libdns/dns"
	"github.com/4thel00z/libdns/v1/pkg/libdns/utils"
)

func SimpleQueryOnce(d utils.DNSServer, domain string, t utils.RecordType, c utils.RecordClass, timeout int64) (dns.Response, error) {
	err := dns.Init()
	if err != nil {
		return dns.Response{}, err
	}
	conn, err := d.ToConnection(timeout)
	if err != nil {
		return dns.Response{}, err
	}

	q := dns.Question{
		Domain: domain,
		Type:   uint16(t),
		Class:  uint16(c),
	}

	query := dns.Query{
		ID:        0xAAAA,
		RD:        true,
		Questions: []dns.Question{q},
	}

	n, err := query.Write(conn)
	if err != nil {
		return dns.Response{}, err
	}

	encodedAnswer := make([]byte, n)
	if _, err := bufio.NewReader(conn).Read(encodedAnswer); err != nil {
		return dns.Response{}, err
	}

	decodeQuery, err := dns.DecodeQuery(encodedAnswer)
	return dns.Response(decodeQuery), err
}
