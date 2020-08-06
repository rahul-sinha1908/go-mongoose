package mongoose

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PopulateObject an Object
func PopulateObject(objPtr interface{}, fieldName string, modelPtr interface{}) {
	t := reflect.TypeOf(objPtr)
	if t.Kind() != reflect.Ptr {
		panic("Model should be a Pointer")
	}
	if reflect.TypeOf(modelPtr).Kind() != reflect.Ptr {
		panic("Model should be a Pointer")
	}

	val := reflect.ValueOf(objPtr).Elem().FieldByName(fieldName)

	t = t.Elem()
	f, b := t.FieldByName(fieldName)
	if b == false {
		return
	}

	tags := strings.Split(f.Tag.Get("mson"), ",")
	for i := range tags {
		tag := strings.Split(tags[i], "=")
		if len(tag) != 2 {
			return
		}
		if tag[0] != "collection" {
			return
		}

		t1 := val.Interface().(primitive.ObjectID)

		err := FindByObjectID(t1, modelPtr)
		if err != nil {
			fmt.Println(err)
			return
		}

		val.Set(reflect.ValueOf(modelPtr).Elem())
	}
}

//PopulateObjectArray Populates the Object Array
func PopulateObjectArray(obj interface{}, field string, modelArrPtr interface{}) error {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		return errors.New("Object should be a Pointer")
	}
	if reflect.TypeOf(modelArrPtr).Kind() == reflect.Ptr {
		return errors.New("The Type need not to be pointer")
	}

	val := reflect.ValueOf(obj).Elem().FieldByName(field)

	t = t.Elem()
	f, b := t.FieldByName(field)
	if b == false {
		return errors.New("No Field")
	}

	tags := strings.Split(f.Tag.Get("mson"), ",")
	for i := range tags {
		tag := strings.Split(tags[i], "=")
		if len(tag) != 2 {
			continue
		}
		if tag[0] != "collection" {
			continue
		}

		objIds := val.Interface().([]primitive.ObjectID)

		err := FindAll(bson.M{
			"_id": bson.M{
				"$in": objIds,
			},
		}, modelArrPtr)

		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(modelArrPtr).Elem())
	}

	return nil
}
