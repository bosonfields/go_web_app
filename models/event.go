package models

import "github.com/globalsign/mgo/bson"

type Event struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	EventId            int64         `json:"event_id" bson:"event_id"`
	EventName          string        `json:"event_name" bson:"event_name"`
	EventDateAndTime   int64         `json:"event_date_and_time" bson:"event_date_and_time"`
	Cast               []string      `json:"cast" bson:"cast"`
	EventAddress       Address       `json:"event_address" bson:"event_address"`
	EventType          string        `json:"event_type" bson:"event_type"`
	TicketType         []TicketType  `json:"ticket_type" bson:"ticket_type"`
	TicketingStartTime int64         `json:"ticketing_start_time" bson:"ticketing_start_time"`
	EventDescription   string        `json:"event_description" bson:"event_description"`
}
type Address struct {
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Zipcode string `json:"zipcode" bson:"zipcode"`
}
type TicketType struct {
	Level       string `json:"level" bson:"level"`
	Price       int    `json:"price" bson:"price"`
	TotalVolume int    `json:"total_volume" bson:"total_volume"`
}
