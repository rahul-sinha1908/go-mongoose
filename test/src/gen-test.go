package test

import (
	"fmt"

	"github.com/rahul-sinha1908/go-mongoose/interfaces"
	"github.com/rahul-sinha1908/go-mongoose/mongoose"
	test "github.com/rahul-sinha1908/go-mongoose/test/models"
	"go.mongodb.org/mongo-driver/bson"
)

//RunTest Run Some Minor tests
func RunTest() {
	//mongoose.
	uM := test.UserModel{}
	allModels := make([]bson.M, 0)
	err := mongoose.FindAll(bson.M{}, test.UserModel{}, &allModels)

	sModel := make([]test.UserModel, len(allModels))
	mbytes, _ := bson.Marshal(allModels[0])
	bson.Unmarshal(mbytes, &sModel[0])

	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Total Length ", len(allModels), sModel[0].ID, sModel[0].GetID())
	fmt.Println(uM.GetID())
	fmt.Println(interfaces.GetName(uM))
}
