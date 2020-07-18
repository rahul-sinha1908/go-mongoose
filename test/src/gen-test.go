package test

import (
	"fmt"

	"github.com/rahul-sinha1908/go-mongoose/mongoose"
	"github.com/rahul-sinha1908/go-mongoose/mutility"
	test "github.com/rahul-sinha1908/go-mongoose/test/models"
	"go.mongodb.org/mongo-driver/bson"
)

//RunTest Run Some Minor tests
func RunTest() {
	//mongoose.
	uM := test.UserModel{}
	allModels := make([]bson.M, 0)
	err := mongoose.FindOne(bson.M{
		"name": "Something here",
	}, &uM)

	if err != nil {
		fmt.Println("Error 1 ", err)
	}
	err = mongoose.FindAll(bson.M{}, test.UserModel{}, &allModels)
	if err != nil {
		fmt.Println("Error 2 ", err)
	}

	fmt.Println(allModels)
	sModel := make([]test.UserModel, len(allModels))
	mbytes, _ := bson.Marshal(allModels[0])
	bson.Unmarshal(mbytes, &sModel[0])

	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Total Length ", len(allModels), sModel[0].ID)
	fmt.Println(uM.ID, uM.Name)
	fmt.Println(mutility.GetName(uM))
}
