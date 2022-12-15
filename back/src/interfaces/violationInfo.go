package interfaces

import "time"

type ViolationInfo struct {
	Pilot           Pilot     `json:"pilot" bson:"pilot"`
	ClosestDistance float64   `json:"closestDistance" bson:"closestDistance"`
	CaptureTime     time.Time `json:"captureTime" bson:"captureTime"`
}
