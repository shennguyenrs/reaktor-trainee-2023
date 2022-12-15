package interfaces

import "encoding/xml"

type Drone struct {
	XMLName      xml.Name `xml:"drone"`
	SerialNumber string   `xml:"serialNumber"`
	Model        string   `xml:"model"`
	Manufacturer string   `xml:"manufacturer"`
	Mac          string   `xml:"mac"`
	IPv4         string   `xml:"ipv4"`
	IPv6         string   `xml:"ipv6"`
	Firmware     string   `xml:"firmware"`
	PositionY    float64  `xml:"positionY"`
	PositionX    float64  `xml:"positionX"`
	Altitude     float64  `xml:"altitude"`
}

