package mongoose

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//DeleteOne Deletes one Object
func DeleteOne(filter bson.M, collectionName string) (*mongo.DeleteResult, error) {
	// fmt.Println("Collection Name : ", b.GetName())
	collection := Get().Database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), MediumWaitTime*time.Second)

	return collection.DeleteOne(ctx, filter)
}
