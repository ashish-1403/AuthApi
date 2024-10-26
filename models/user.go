package models

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

var userCollection *mongo.Collection

func SetDatabase(db *mongo.Database) {
	userCollection = db.Collection("users")
}
