package mongoose

import (
	"reflect"
)

func generateTables(doc *interface{}) {
	// Get().Database.Collection("test", &options.CollectionOptions{})

	t := reflect.TypeOf(doc)
	field := t.Field(0)
	//Get().Database.Collection("").Indexes().CreateOne()
	println(field.Tag.Get("mson"))
}
