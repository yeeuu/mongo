package mongo

import "github.com/globalsign/mgo"

// QuerySet likes mgo.Query.
type QuerySet interface {
	All(result interface{}) error
	One(result interface{}) error
	Count() (int, error)
	Limit(n int) QuerySet
	Skip(n int) QuerySet
	Prefetch(p float64) QuerySet
	Sort(fields ...string) QuerySet
	Close()
}

type mongoQuerySet struct {
	sess  *mgo.Session
	query *mgo.Query
}

func (qs *mongoQuerySet) All(result interface{}) error {
	return qs.query.All(result)
}

func (qs *mongoQuerySet) One(result interface{}) error {
	return qs.query.One(result)
}

func (qs *mongoQuerySet) Count() (int, error) {
	return qs.query.Count()
}

func (qs *mongoQuerySet) Limit(n int) QuerySet {
	qs.query = qs.query.Limit(n)
	return qs
}

func (qs *mongoQuerySet) Skip(n int) QuerySet {
	qs.query = qs.query.Skip(n)
	return qs
}

func (qs *mongoQuerySet) Prefetch(p float64) QuerySet {
	qs.query = qs.query.Prefetch(p)
	return qs
}

func (qs *mongoQuerySet) Sort(fields ...string) QuerySet {
	qs.query = qs.query.Sort(fields...)
	return qs
}

func (qs *mongoQuerySet) Close() {
	qs.sess.Close()
}
