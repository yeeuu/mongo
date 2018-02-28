package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

// Storage supply a interface for make Queryset.
type Storage interface {
	Query(interface{}) (QuerySet, error)
	Raw(interface{}, interface{}) (QuerySet, error)
	Insert(base interface{}, docs ...interface{}) error
	Update(filter, update interface{}) error
	UpdateAll(filter, update interface{}) (info *mgo.ChangeInfo, err error)
	Upsert(filter, update interface{}) (info *mgo.ChangeInfo, err error)
	Remove(filter interface{}) error
	RemoveAll(filter interface{}) (info *mgo.ChangeInfo, err error)
}

type mongoStorage struct {
	sess *mgo.Session
}

// NewStorage create a new mongodb storage.
func NewStorage(sess *mgo.Session) Storage {
	return &mongoStorage{
		sess: sess,
	}
}

func (ms *mongoStorage) realAction(action string, base interface{}, docs ...interface{}) (info *mgo.ChangeInfo, err error) {
	sess := ms.sess.Copy()
	defer sess.Close()
	c, err := getCollection(sess, base)
	if err != nil {
		err = errors.Wrap(err, "Get collection info failed")
		return
	}
	realQuery, err := structToBsonM(base)
	if err != nil {
		err = errors.Wrap(err, "Convert struct to query failed")
		return
	}
	switch action {
	case "insert":
		err = c.Insert(docs...)
	case "update":
		err = c.Update(realQuery, docs[0])
	case "updateAll":
		info, err = c.UpdateAll(realQuery, docs[0])
	case "upsert":
		info, err = c.Upsert(realQuery, docs[0])
	case "remove":
		err = c.Remove(realQuery)
	case "removeAll":
		info, err = c.RemoveAll(realQuery)
	}
	return
}

func (ms *mongoStorage) Insert(base interface{}, docs ...interface{}) error {
	_, err := ms.realAction("insert", base, docs...)
	return err
}

func (ms *mongoStorage) Update(filter, update interface{}) error {
	_, err := ms.realAction("update", filter, update)
	return err
}

func (ms *mongoStorage) UpdateAll(filter, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = ms.realAction("updateAll", filter, update)
	return
}

func (ms *mongoStorage) Upsert(filter, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = ms.realAction("upsert", filter, update)
	return
}

func (ms *mongoStorage) Remove(filter interface{}) error {
	_, err := ms.realAction("remove", filter, nil)
	return err
}

func (ms *mongoStorage) RemoveAll(filter interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = ms.realAction("removeAll", filter, nil)
	return
}

func (ms *mongoStorage) Query(query interface{}) (QuerySet, error) {
	var err error
	var c *mgo.Collection
	sess := ms.sess.Copy()
	qs := &mongoQuerySet{
		sess: sess,
	}
	c, err = getCollection(sess, query)
	if err != nil {
		err = errors.Wrap(err, "Get collection info failed")
		return qs, err
	}
	realQuery, err := structToBsonM(query)
	if err != nil {
		err = errors.Wrap(err, "Convert struct to query failed")
		return qs, err
	}
	qs.query = c.Find(realQuery)
	return qs, err
}

func (ms *mongoStorage) Raw(base, query interface{}) (QuerySet, error) {
	var err error
	var c *mgo.Collection
	sess := ms.sess.Copy()
	qs := &mongoQuerySet{
		sess: sess,
	}
	c, err = getCollection(sess, query)
	if err != nil {
		err = errors.Wrap(err, "Get collection info failed")
		return qs, err
	}
	qs.query = c.Find(query)
	return qs, err
}
