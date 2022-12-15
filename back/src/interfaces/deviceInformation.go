package interfaces

import (
	"encoding/xml"
	"time"
)

type DeviceInformation struct {
	XMLName          xml.Name  `xml:"deviceInformation"`
	DeviceID         string    `xml:"deviceId,attr"`
	ListenRange      int       `xml:"listenRange"`
	DeviceStarted    time.Time `xml:"deviceStarted"`
	UptimeSeconds    int       `xml:"uptimeSeconds"`
	UpdateIntervalMs int       `xml:"updateIntervalMs"`
}

