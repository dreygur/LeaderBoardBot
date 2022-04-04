package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username,omitempty"`
	UserID     string             `bson:"user_id,omitempty"`
	NameHash   string             `bson:"name_hash,omitempty"`
	Level      int                `bson:"level,omitempty"`
	Points     int                `bson:"points,omitempty"`
	Invites    int                `bson:"invites,omitempty"`
	Activities Activity           `bson:"activities,omitempty"`
}

type Activity struct {
	Text     int `bson:"text,omitempty"`
	Reaction int `bson:"reaction,omitempty"`
	Voice    int `bson:"voice,omitempty"`
}
