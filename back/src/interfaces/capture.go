package interfaces

import (
	"encoding/xml"
	"time"
)

type Capture struct {
	XMLName           xml.Name  `xml:"capture"`
	SnapshotTimestamp time.Time `xml:"snapshotTimestamp,attr"`
	Drones            []Drone   `xml:"drone"`
}

