package interfaces

import "time"

type Pilot struct {
	PilotID     string    `json:"pilotId" bson:"pilotId"`
	FirstName   string    `json:"firstName" bson:"firstName"`
	LastName    string    `json:"lastName" bson:"lastName"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	CreatedDt   time.Time `json:"createdDt" bson:"createdDt"`
	Email       string    `json:"email" bson:"email"`
}
