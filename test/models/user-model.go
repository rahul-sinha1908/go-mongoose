package test

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserModel This is the model for the Users
type UserModel struct {
	ID            primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string               `json:"name,omitempty" bson:"name,omitempty"`
	Password      string               `json:"password,omitempty" bson:"password,omitempty"`
	Email         string               `json:"email,omitempty" bson:"email,omitempty"`
	FirebaseToken string               `json:"firebaseToken" bson:"firebaseToken"`
	Teams         []primitive.ObjectID `json:"teams" bson:"teams"`
	MainProfile   string               `json:"mainProfile" bson:"mainProfile"`
	Phone         string               `json:"phone" bson:"phone"`
	SocialMedia   string               `json:"socialMedia" bson:"socialMedia"`
	Strengths     string               `json:"strengths" bson:"strengths"`
	Developments  string               `json:"developments" bson:"developments"`
	UserType      int                  `json:"userType" bson:"userType"`
	Status        int                  `json:"status" bson:"status"`
}
