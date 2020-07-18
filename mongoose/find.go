package mongoose

import (
	"context"
	"fmt"
	"time"

	"github.com/rahul-sinha1908/go-mongoose/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindOne Searches one object and returns its value
func FindOne(filter bson.M, b interfaces.ModelInterface) (err error) {
	// fmt.Println("Collection Name : ", b.GetName())
	collection := Get().Database.Collection(interfaces.GetName(b))
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	res := collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return res.Err()
	}

	err = res.Decode(b)
	if err != nil {
		return err
	}

	return nil
}

// FindByID Searches by ID
func FindByID(id string, b interfaces.ModelInterface) (err error) {
	// fmt.Println("Collection Name : ", b.GetName())
	collection := Get().Database.Collection(interfaces.GetName(b))
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res := collection.FindOne(ctx, bson.M{
		"_id": userID,
	})
	if res.Err() != nil {
		return res.Err()
	}
	err = res.Decode(b)
	if err != nil {
		return err
	}

	return nil
}

// FindByObjectID Searches by Object ID
func FindByObjectID(objectID primitive.ObjectID, b interfaces.ModelInterface) (err error) {
	// fmt.Println("Collection Name : ", b.GetName())
	collection := Get().Database.Collection(interfaces.GetName(b))
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	res := collection.FindOne(ctx, bson.M{
		"_id": objectID,
	})
	if res.Err() != nil {
		return res.Err()
	}
	err = res.Decode(b)
	if err != nil {
		return err
	}

	return nil
}

// FindAll Get All Docs
func FindAll(filter bson.M, model interfaces.ModelInterface, allModels *[]bson.M) error {
	fmt.Println("Find All Name ", interfaces.GetGenericName(model))
	collection := Get().Database.Collection(interfaces.GetGenericName(model))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	err = cur.All(ctx, allModels)
	if err != nil {
		return err
	}
	return nil
}
