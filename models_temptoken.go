package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// TempToken is a database entry for a temporary token
type TempToken struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"user_id,omitempty"`
	Token      string             `bson:"token" json:"token"`
	Expiration int64              `bson:"expiration" json:"expiration"`
	Purpose    string             `bson:"purpose" json:"purpose"`
	UserID     string             `bson:"user_id" json:"user_id`
	Info       string             `bson:"info" json:"info"`
	InstanceID string             `bson:"instance_id" json:"instance_id"`
}

func (t TempToken) ToAPI() {

}
