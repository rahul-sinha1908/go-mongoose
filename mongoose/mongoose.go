package mongoose

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongo This is the Mongo struct
type Mongo struct {
	client   *mongo.Client
	Database *mongo.Database
	Err      error
}

var (
	_mongo Mongo
)

//Get This function will recieve the Mongo structure
func Get() Mongo {
	if _mongo.client == nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		_mongo.client, _mongo.Err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if _mongo.Err == nil {
			_mongo.Database = _mongo.client.Database("teamace")
			fmt.Print("Database Created Successfully\n")
		}
	}
	return _mongo
}
