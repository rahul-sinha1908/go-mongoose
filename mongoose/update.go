package mongoose

import (
	"context"
	"time"

	"github.com/rahul-sinha1908/go-mongoose/mutility"
	"go.mongodb.org/mongo-driver/bson"
)

//UpdateByID Updates by ID
func UpdateByID(model interface{}) error {
	collection := Get().Database.Collection(mutility.GetName(model))
	ctx, _ := context.WithTimeout(context.Background(), ShortWaitTime*time.Second)

	_, err := collection.UpdateOne(ctx, bson.M{
		"_id": mutility.GetID(model),
	}, model)

	if err != nil {
		return err
	}
	return nil
}
