package test

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserModel This is the model for the Users
type UserModel struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty" mson:"cunique"`
	FirebaseToken string             `json:"firebaseToken" bson:"firebaseToken" mson:"unique"`
	Test          interface{}        `json:"test" bson:"test" mson:"collection=UserModel"`
	Teams         interface{}        `json:"teams" bson:"teams" mson:"collection=UserModel"`
	MainProfile   string             `json:"mainProfile" bson:"mainProfile"`
	Phone         string             `json:"phone" bson:"phone" mson:"cunique"`
	SocialMedia   string             `json:"socialMedia" bson:"socialMedia"`
	Strengths     string             `json:"strengths" bson:"strengths"`
	Developments  string             `json:"developments" bson:"developments"`
	UserType      int                `json:"userType" bson:"userType"`
	Status        int                `json:"status" bson:"status"`
}
