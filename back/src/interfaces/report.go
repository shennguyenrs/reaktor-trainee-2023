package interfaces

import "encoding/xml"

type Report struct {
	XMLName          xml.Name          `xml:"report"`
	DeviceInfomation DeviceInformation `xml:"deviceInformation"`
	Capture          Capture           `xml:"capture"`
}

