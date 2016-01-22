package domain

import (
	"gopkg.in/mgo.v2/bson"
)

type Column struct {
	Id bson.ObjectId
	Name string
	Order int
	Tickets []Ticket
}

