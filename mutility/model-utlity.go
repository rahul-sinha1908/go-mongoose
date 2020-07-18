package mutility

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

//GetID Returns the Object ID
func GetID(a interface{}) primitive.ObjectID {
	// t := reflect.TypeOf(a)
	tv := reflect.ValueOf(a)
	mVal := tv.FieldByName("ID")
	if reflect.Value.IsZero(mVal) {
		return primitive.NilObjectID
	}
	aV := mVal.Interface().(primitive.ObjectID)
	return aV
}

//CreateIndex This function would be used to create index for the table
func CreateIndex(a interface{}) {
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
