package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// QuerySet likes mgo.Query.
type QuerySet interface {
	Insert(docs ...interface{}) error
	Update(update interface{}) error
	UpdateAll(update interface{}) (info *mgo.ChangeInfo, err error)
	Upsert(update interface{}) (info *mgo.ChangeInfo, err error)
	Remove() error
	RemoveAll() (info *mgo.ChangeInfo, err error)
	All(result interface{}) error
	One(result interface{}) error
	Count() (int, error)
	Limit(n int) QuerySet
	Skip(n int) QuerySet
	Prefetch(p float64) QuerySet
	Sort(fields ...string) QuerySet
}

type mongoQuerySet struct {
	sess       *mgo.Session
	database   string
	collection string
	query      bson.M
	err        error
	limit      *int
	skip       *int
	prefetch   *float64
	sort       []string
}

func (qs *mongoQuerySet) getSession() (*mgo.Collection, func()) {
	sess := qs.sess.Copy()
	return sess.DB(qs.database).C(qs.collection), func() {
		sess.Close()
	}
}

func (qs *mongoQuerySet) getQuery(c *mgo.Collection) *mgo.Query {
	q := c.Find(qs.query)
	if qs.limit != nil {
		q = q.Limit(*qs.limit)
	}
	if qs.skip != nil {
		q = q.Skip(*qs.skip)
	}
	if qs.prefetch != nil {
		q = q.Prefetch(*qs.prefetch)
	}
	if len(qs.sort) != 0 {
		q = q.Sort(qs.sort...)
	}
	return q
}

func (qs *mongoQuerySet) Insert(docs ...interface{}) error {
	c, close := qs.getSession()
	defer close()
	err := c.Insert(docs...)
	return err
}

func (qs *mongoQuerySet) Update(update interface{}) error {
	if qs.err != nil {
		return qs.err
	}
	c, close := qs.getSession()
	defer close()
	err := c.Update(qs.query, update)
	return err
}

func (qs *mongoQuerySet) UpdateAll(update interface{}) (info *mgo.ChangeInfo, err error) {
	if qs.err != nil {
		err = qs.err
		return
	}
	c, close := qs.getSession()
	defer close()
	info, err = c.UpdateAll(qs.query, update)
	return
}

func (qs *mongoQuerySet) Upsert(update interface{}) (info *mgo.ChangeInfo, err error) {
	if qs.err != nil {
		err = qs.err
		return
	}
	c, close := qs.getSession()
	defer close()
	info, err = c.Upsert(qs.query, update)
	return
}

func (qs *mongoQuerySet) Remove() error {
	if qs.err != nil {
		return qs.err
	}
	c, close := qs.getSession()
	defer close()
	err := c.Remove(qs.query)
	return err
}

func (qs *mongoQuerySet) RemoveAll() (info *mgo.ChangeInfo, err error) {
	if qs.err != nil {
		err = qs.err
		return
	}
	c, close := qs.getSession()
	defer close()
	info, err = c.RemoveAll(qs.query)
	return
}

func (qs *mongoQuerySet) All(result interface{}) error {
	if qs.err != nil {
		return qs.err
	}
	c, close := qs.getSession()
	defer close()
	return qs.getQuery(c).All(result)
}

func (qs *mongoQuerySet) One(result interface{}) error {
	if qs.err != nil {
		return qs.err
	}
	c, close := qs.getSession()
	defer close()
	return qs.getQuery(c).One(result)
}

func (qs *mongoQuerySet) Count() (int, error) {
	if qs.err != nil {
		return 0, qs.err
	}
	c, close := qs.getSession()
	defer close()
	return qs.getQuery(c).Count()
}

func (qs *mongoQuerySet) Limit(n int) QuerySet {
	qs.limit = &n
	return qs
}

func (qs *mongoQuerySet) Skip(n int) QuerySet {
	qs.skip = &n
	return qs
}

func (qs *mongoQuerySet) Prefetch(p float64) QuerySet {
	qs.prefetch = &p
	return qs
}

func (qs *mongoQuerySet) Sort(fields ...string) QuerySet {
	qs.sort = fields
	return qs
}
