package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contato struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Email   string             `json:"email" bson:"email"`
	Message string             `json:"message" bson:"message"`
}
