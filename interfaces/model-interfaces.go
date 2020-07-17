package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

//ModelInterface This contains the interface of models
type ModelInterface interface {
	GetID() primitive.ObjectID
	GetName() string
	CreateIndex()
}
