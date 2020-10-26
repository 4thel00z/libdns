package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToConnection(t *testing.T) {
	testDNSServer(t, GooglePrimary)
	testDNSServer(t, GoogleSecondary)
	testDNSServer(t, Quad9Primary)
	testDNSServer(t, Quad9Secondary)
	testDNSServer(t, Quad9Secondary)
	testDNSServer(t, OpenDNSHomePrimary)
	testDNSServer(t, OpenDNSHomeSecondary)
	testDNSServer(t, CloudflarePrimary)
	testDNSServer(t, CloudflareSecondary)
	testDNSServer(t, CleanBrowsingPrimary)
	testDNSServer(t, CleanBrowsingSecondary)
	testDNSServer(t, VerisignPrimary)
	testDNSServer(t, VerisignSecondary)
	testDNSServer(t, AlternateDNSPrimary)
	testDNSServer(t, AlternateDNSSecondary)
	testDNSServer(t, AdGuardDNSPrimary)
	testDNSServer(t, AdGuardDNSSecondary)
	testDNSServer(t, DNSWATCHPrimary)
	testDNSServer(t, DNSWATCHSecondary)
	testDNSServer(t, ComodoSecureDNSPrimary)
	testDNSServer(t, ComodoSecureDNSSecondary)
	testDNSServer(t, ComodoSecureDNSSecondary)
	testDNSServer(t, CenturyLinkPrimary)
	testDNSServer(t, CenturyLinkSecondary)
	testDNSServer(t, SafeDNSPrimary)
	testDNSServer(t, SafeDNSSecondary)
	testDNSServer(t, OpenNICPrimary)
	testDNSServer(t, OpenNICSecondary)
	testDNSServer(t, DynPrimary)
	testDNSServer(t, DynSecondary)
	testDNSServer(t, FreeDNSPrimary)
	testDNSServer(t, FreeDNSSecondary)
	testDNSServer(t, YandexDNSPrimary)
	testDNSServer(t, YandexDNSSecondary)
	testDNSServer(t, UncensoredDNSPrimary)
	testDNSServer(t, UncensoredDNSSecondary)
	testDNSServer(t, HurricaneElectricPrimary)
	testDNSServer(t, PuntCATPrimary)
	testDNSServer(t, NeustarPrimary)
	testDNSServer(t, NeustarSecondary)
	testDNSServer(t, FourthEstatePrimary)
	testDNSServer(t, FourthEstateSecondary)
}

func testDNSServer(t *testing.T, d DNSServer) {
	connection, err := d.ToConnection()
	assert.Nil(t, err)
	defer connection.Close()
}
