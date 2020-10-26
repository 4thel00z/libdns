package utils

import (
	"fmt"
	"net"
	"time"
)

type DNSServer string

func (d DNSServer) ToConnection(deadline int64) (net.Conn, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:53", string(d)))
	if err != nil {
		return nil, err
	}
	if err := conn.SetDeadline(time.Now().Add(time.Duration(deadline) * time.Second)); err != nil {
		return conn, nil

	}
	return conn, nil
}

const (
	GooglePrimary            = DNSServer("8.8.8.8")
	GoogleSecondary          = DNSServer("8.8.4.4")
	Quad9Primary             = DNSServer("9.9.9.9")
	Quad9Secondary           = DNSServer("149.112.112.112")
	OpenDNSHomePrimary       = DNSServer("208.67.222.222")
	OpenDNSHomeSecondary     = DNSServer("208.67.220.220")
	CloudflarePrimary        = DNSServer("1.1.1.1")
	CloudflareSecondary      = DNSServer("1.0.0.1")
	CleanBrowsingPrimary     = DNSServer("185.228.168.9")
	CleanBrowsingSecondary   = DNSServer("185.228.169.9")
	VerisignPrimary          = DNSServer("64.6.64.6")
	VerisignSecondary        = DNSServer("64.6.65.6")
	AlternateDNSPrimary      = DNSServer("198.101.242.72")
	AlternateDNSSecondary    = DNSServer("23.253.163.53")
	AdGuardDNSPrimary        = DNSServer("94.140.14.14")
	AdGuardDNSSecondary      = DNSServer("94.140.15.15")
	DNSWATCHPrimary          = DNSServer("84.200.69.80")
	DNSWATCHSecondary        = DNSServer("84.200.70.40")
	ComodoSecureDNSPrimary   = DNSServer("8.26.56.26")
	ComodoSecureDNSSecondary = DNSServer("8.20.247.20")
	CenturyLinkPrimary       = DNSServer("205.171.3.66")
	CenturyLinkSecondary     = DNSServer("205.171.202.166")
	SafeDNSPrimary           = DNSServer("195.46.39.39")
	SafeDNSSecondary         = DNSServer("195.46.39.40")
	OpenNICPrimary           = DNSServer("192.71.245.208")
	OpenNICSecondary         = DNSServer("94.247.43.254")
	DynPrimary               = DNSServer("216.146.35.35")
	DynSecondary             = DNSServer("216.146.36.36")
	FreeDNSPrimary           = DNSServer("45.33.97.5")
	FreeDNSSecondary         = DNSServer("37.235.1.177")
	YandexDNSPrimary         = DNSServer("77.88.8.8")
	YandexDNSSecondary       = DNSServer("77.88.8.1")
	UncensoredDNSPrimary     = DNSServer("91.239.100.100")
	UncensoredDNSSecondary   = DNSServer("89.233.43.71")
	HurricaneElectricPrimary = DNSServer("74.82.42.42")
	PuntCATPrimary           = DNSServer("109.69.8.51")
	NeustarPrimary           = DNSServer("156.154.70.5")
	NeustarSecondary         = DNSServer("156.154.71.5")
	FourthEstatePrimary      = DNSServer("45.77.165.194")
	FourthEstateSecondary    = DNSServer("45.32.36.36")
)
