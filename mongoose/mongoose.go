package mongoose

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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

//DBConnection DB Connection Details
type DBConnection struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

var (
	_mongo        Mongo
	connectionURL string = "mongodb://localhost:27017"
	dbName        string = "teamace"
)

//InitiateDB This needs to be called if you are using some other than default DB
func InitiateDB(dbConnection DBConnection) {
	// fmt.Println(dbConnection.Port)
	if dbConnection.User == "" {
		connectionURL = "mongodb://" + dbConnection.Host + ":" + strconv.Itoa(dbConnection.Port)
	} else {
		connectionURL = "mongodb://" + url.QueryEscape(dbConnection.User) + ":" + url.QueryEscape(dbConnection.Password) + "@" + dbConnection.Host + ":" + strconv.Itoa(dbConnection.Port)
	}

	if dbConnection.Database != "" {
		dbName = dbConnection.Database
	}
}

//Get This function will recieve the Mongo structure
func Get() Mongo {
	if _mongo.client == nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		_mongo.client, _mongo.Err = mongo.Connect(ctx, options.Client().ApplyURI(connectionURL))
		if _mongo.Err == nil {
			_mongo.Database = _mongo.client.Database(dbName)
			fmt.Print("Database Created Successfully\n")
		} else {
			panic(_mongo.Err)
		}
	}
	return _mongo
}
