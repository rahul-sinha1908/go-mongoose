package mongoose

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PopulateObject an Object
func PopulateObject(obj interface{}, field string, modelPtr interface{}) {
	t := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj).Elem().FieldByName(field)

	if t.Kind() != reflect.Ptr {
		panic("Model should be a Pointer")
	}
	if reflect.TypeOf(modelPtr).Kind() != reflect.Ptr {
		panic("Model should be a Pointer")
	}
	t = t.Elem()
	f, b := t.FieldByName(field)
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
func PopulateObjectArray(obj interface{}, field string, modelType interface{}) ([]interface{}, error) {
	t := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj).Elem().FieldByName(field)

	if t.Kind() != reflect.Ptr {
		panic("Object should be a Pointer")
	}
	if reflect.TypeOf(modelType).Kind() == reflect.Ptr {
		panic("The Type need not to be pointer")
	}
	t = t.Elem()
	f, b := t.FieldByName(field)
	if b == false {
		return nil, errors.New("No Field")
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

		data := make([]bson.M, 0)
		err := FindAll(bson.M{
			"_id": bson.M{
				"$in": objIds,
			},
		}, modelType, &data)

		if err != nil {
			return nil, err
		}

		// mType := reflect.TypeOf(modelType)

		for i := range data {
			bD := data[i]
			b, er := bson.Marshal(bD)
			if er != nil {
				return nil, er
			}
			er = bson.Unmarshal(b, &modelType)
			if er != nil {
				return nil, er
			}

		}

		// err := FindByObjectID(t1, modelPtr)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// val.Set(reflect.ValueOf(modelPtr).Elem())
	}

	return nil, nil
}
