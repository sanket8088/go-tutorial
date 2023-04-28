package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `validate:"required,min=1" bson:"movie_name,omitempty" json:"movie_name,omitempty"`
	Watched bool               `bson:"watched,omitempty" json:"watched,omitempty"`
}
