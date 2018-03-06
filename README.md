# mongo

[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/yeeuu/mongo)
[![Build Status](https://travis-ci.org/yeeuu/mongo.svg?branch=master)](https://travis-ci.org/yeeuu/mongo)
[![codecov](https://codecov.io/gh/yeeuu/mongo/branch/master/graph/badge.svg)](https://codecov.io/gh/yeeuu/mongo)
[![Go Report Card](https://goreportcard.com/badge/github.com/yeeuu/mongo)](https://goreportcard.com/report/github.com/yeeuu/mongo)


Simple ODM wraps mgo with interface.

## TODO
 - [x] 0.1.0 release
 - [ ] Test cgo getoverage
 - [x] Travis-ci support

## Limits

Do not support complex mongo query for now.

More limits need to discovery :D

All API design may change in the future until reach 1.0

## Usage

```go
type user struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	LastIP   string        `bson:"last_ip"`
	LastTime time.Time     `bson:"last_time"`
}

type userSelector struct {
	ID   *bson.ObjectId   `bson:"_id"`
	IDs  *[]bson.ObjectId `bson:"_id"`
	Name *string          `bson:"name"`
}

func (us *userSelector) Database() string {
	return "test"
}

func (us *userSelector) Collection() string {
	return "users"
}

func main() {
	sess, err := mgo.DialWithTimeout("127.0.0.1", 2*time.Second)
	if err != nil {
		panic(err)
	}
	store := mongo.NewStorage(sess)
	err = store.Query(&userSelector{}).Insert(user{Name: "hello"})
	if err != nil {
		panic(err)
}
```
