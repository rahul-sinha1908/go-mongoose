package mongoose

import (
	"context"
	"time"

	"github.com/rahul-sinha1908/go-mongoose/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

//InsertOne This will insert just one Data
func InsertOne(model interfaces.ModelInterface) (res *mongo.InsertOneResult, err error) {
	collection := Get().Database.Collection(model.GetName())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res, err = collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	return res, err
}

//InsertMany This will insert multiple Data
func InsertMany(collectionName string, models []interface{}) (res *mongo.InsertManyResult, err error) {
	collection := Get().Database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res, err = collection.InsertMany(ctx, models)
	if err != nil {
		return nil, err
	}
	return res, err
}
