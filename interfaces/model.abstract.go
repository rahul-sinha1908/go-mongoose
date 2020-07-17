package interfaces

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModelAbstract This Model needs to be imported by all the models created
type ModelAbstract struct {
}

//CreateIndex This function would be used to create index for the table
func (a ModelAbstract) CreateIndex() {
	t := reflect.TypeOf(a)
	var indexes []interface{} = make([]interface{}, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(0)
		tagStr := field.Tag.Get("mson")
		tags := strings.Split(tagStr, ",")
		if len(tags) == 0 {
			continue
		}
		index := analyzeAndCreateTagIndex(tags)

		if index != nil {
			indexes = append(indexes, index)
		}
	}
}

func analyzeAndCreateTagIndex(tags []string) *interface{} {
	// TODO Do some stuff to analyze the tags

	return nil
}

//GetID Returns the Object ID
func (a ModelAbstract) GetID() primitive.ObjectID {
	panic("Needs to implement the GetID method")
}

//GetName Returns the collection Name
func (a ModelAbstract) GetName() string {
	panic("Needs to implement GetName method")
}
