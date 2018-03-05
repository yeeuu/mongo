package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

// Storage supply a interface for make Queryset.
type Storage interface {
	Query(Collection) QuerySet
	Raw(Collection, bson.M) QuerySet
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

func (ms *mongoStorage) Query(query Collection) QuerySet {
	qs := &mongoQuerySet{
		sess:       ms.sess,
		database:   query.Database(),
		collection: query.Collection(),
	}
	realQuery, err := structToBsonM(query)
	qs.err = errors.Wrap(err, "convert to bson.M failed")
	qs.query = realQuery
	return qs
}

func (ms *mongoStorage) Raw(base Collection, query bson.M) QuerySet {
	qs := &mongoQuerySet{
		sess:       ms.sess,
		database:   base.Database(),
		collection: base.Collection(),
		query:      query,
	}
	return qs
}
