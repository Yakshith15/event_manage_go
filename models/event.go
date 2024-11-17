package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Venue string `json:"venue" bson:"venue"`
	Date string `json:"date" bson:"date"`
	TicketPrice float64 `json:"ticket_price" bson:"ticket_price"`
	TicketsAvailable int `json:"tickets_available" bson:"tickets_available"`
}