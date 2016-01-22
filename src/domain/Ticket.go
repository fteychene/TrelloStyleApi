package domain

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Ticket struct {
	Id bson.ObjectId
	Name string
	Description string
	OpenDate time.Time
	AffectedUserId string
}