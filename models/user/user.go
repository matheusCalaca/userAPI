package model

import "time"

// User structure
type User struct {
	// ID        bson.ObjectId `bson:"_id"`
	Nome      string    `bson:"nome"`
	Endereço  string    `bson:"endereço"`
	Idade     int       `bson:"idade"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// Users list
type Users []User
