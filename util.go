package mongo

import (
	"reflect"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

func structToBsonM(s interface{}) (bson.M, error) {
	var err error
	r := bson.M{}
	v := reflect.ValueOf(s)
	switch v.Kind() {
	case reflect.Ptr:
	default:
		return r, errors.New("Query must be struct pointer")
	}
	if !v.IsNil() {
		u := v.Elem()
		var key string
		var value reflect.Value
		for i := 0; i < u.NumField(); i++ {
			value = u.Field(i)
			if value.IsNil() || !value.CanInterface() {
				continue
			}
			key = u.Type().Field(i).Tag.Get("bson")
			// drop support slice field query to support $in.
			if value.Elem().Type().Kind() == reflect.Slice {
				r[key] = bson.M{"$in": value.Elem().Interface()}
			} else {
				r[key] = value.Elem().Interface()
			}
		}
	}
	return r, err
}

func getCollection(sess *mgo.Session, query interface{}) (c *mgo.Collection, err error) {
	if mc, ok := query.(Collection); ok {
		c = sess.DB(mc.Database()).C(mc.Collection())
	} else {
		err = errors.New("Struct not support MongoCollection")
		return
	}
	return
}
