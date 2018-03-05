package main

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/yeeuu/mongo"
)

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
	var u user
	name := "hello"
	err = store.Query(&userSelector{
		Name: &name,
	}).One(&u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	store.Query(&userSelector{}).Insert(user{Name: "world"})
	var users []user
	err = store.Query(&userSelector{}).All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}
