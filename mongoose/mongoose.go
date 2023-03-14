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

	/*
		In addition to the standard connection format, MongoDB supports a DNS-constructed seed list. Using DNS to construct the available servers list allows more flexibility of deployment and the ability to change the servers in rotation without reconfiguring clients.

		In order to leverage the DNS seed list, use a connection string prefix of mongodb+srv rather than the standard mongodb. The +srv indicates to the client that the hostname that follows corresponds to a DNS SRV record. The driver or mongosh will then query the DNS for the record to determine which hosts are running the mongod instances.

		https://www.mongodb.com/docs/manual/reference/connection-string/#dns-seed-list-connection-format
	*/
	SRV bool
}

//ShortWaitTime Small Wait time
//MediumWaitTime Medium Wait Time
//LongWaitTime Long wait time
var (
	_mongo        Mongo
	connectionURL string = "mongodb://localhost:27017"
	dbName        string = "teamace"

	ShortWaitTime  time.Duration = 2
	MediumWaitTime time.Duration = 5
	LongWaitTime   time.Duration = 10
)

//InitiateDB This needs to be called if you are using some other than default DB
func InitiateDB(dbConnection DBConnection) {
	// fmt.Println(dbConnection.Port)
	if dbConnection.Port == 0 {
		dbConnection.Port = 27017
	}
	urlHeader := "mongodb://"
	if dbConnection.SRV {
		urlHeader = "mongodb+srv://"
	}

	if dbConnection.User == "" {
		connectionURL = urlHeader + dbConnection.Host
	} else {
		connectionURL = urlHeader + url.QueryEscape(dbConnection.User) + ":" + url.QueryEscape(dbConnection.Password) + "@" + dbConnection.Host
	}

	if !dbConnection.SRV {
		connectionURL += ":" + strconv.Itoa(dbConnection.Port)

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
