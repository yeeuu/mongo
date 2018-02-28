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
	err = store.Insert(&userSelector{}, user{Name: "hello"})
	if err != nil {
		panic(err)
	}
	var u user
	name := "hello"
	qs, err := store.Query(&userSelector{
		Name: &name,
	})
	if err != nil {
		panic(err)
	}
	err = qs.One(&u)
	if err != nil {
		panic(err)
	}
	qs.Close()
	fmt.Println(u)
	err = store.Insert(&userSelector{}, user{Name: "world"})
	qs, err = store.Query(&userSelector{})
	if err != nil {
		panic(err)
	}
	var users []user
	err = qs.All(&users)
	if err != nil {
		panic(err)
	}
	qs.Close()
	fmt.Println(users)
}
