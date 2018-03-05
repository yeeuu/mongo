package mongo

import (
	"fmt"
	"testing"
	"time"

	"github.com/globalsign/mgo"
)

func TestStorage(t *testing.T) {
	sess, err := mgo.DialWithTimeout("127.0.0.1", 2*time.Second)
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	sess.DB("test").DropDatabase()
	store := NewStorage(sess)
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
