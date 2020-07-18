package interfaces

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModelAbstract This Model needs to be imported by all the models created
type ModelAbstract struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

//CreateIndex This function would be used to create index for the table
func (a ModelAbstract) CreateIndex() {
	t := reflect.TypeOf(a)
	t.Name()
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
	return a.ID
}

//GetName Returns the collection Name
// func GetName(a ModelInterface) string {
// 	t := reflect.TypeOf(a)
// 	return t.Name()
// }

//GetName Returns the collection Name
func GetName(a interface{}) string {
	t := reflect.TypeOf(a)
	// fmt.Println(t.NumField())
	// fmt.Println(t.Kind())
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Ptr || t.Kind() == reflect.Array || t.Kind() == reflect.Map || t.Kind() == reflect.Chan {
		return t.Elem().Name()
	}
	return t.Name()
}
