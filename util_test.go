package mongo

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type user struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name"`
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

func Test_structToBsonM(t *testing.T) {
	type args struct {
		s interface{}
	}
	id := bson.NewObjectId()
	ids := []bson.ObjectId{bson.NewObjectId()}
	name := "hello"
	tests := []struct {
		name    string
		args    args
		want    bson.M
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "nil",
			args: args{
				nil,
			},
			want:    bson.M{},
			wantErr: true,
		},
		{
			name: "struct",
			args: args{
				user{},
			},
			want:    bson.M{},
			wantErr: true,
		},
		{
			name: "empty",
			args: args{
				&userSelector{},
			},
			want:    bson.M{},
			wantErr: false,
		},
		{
			name: "name",
			args: args{
				&userSelector{Name: &name},
			},
			want:    bson.M{"name": name},
			wantErr: false,
		},
		{
			name: "id",
			args: args{
				&userSelector{ID: &id},
			},
			want:    bson.M{"_id": id},
			wantErr: false,
		},
		{
			name: "ids",
			args: args{
				&userSelector{IDs: &ids},
			},
			want:    bson.M{"_id": bson.M{"$in": ids}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := structToBsonM(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("structToBsonM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structToBsonM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCollection(t *testing.T) {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		t.Skip("Cannot connect to mongodb.")
		return
	}
	defer sess.Close()
	type args struct {
		sess  *mgo.Session
		query interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantC   *mgo.Collection
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "nil",
			args: args{
				sess:  sess,
				query: nil,
			},
			wantC:   nil,
			wantErr: true,
		},
		{
			name: "struct",
			args: args{
				sess:  sess,
				query: user{},
			},
			wantC:   nil,
			wantErr: true,
		},
		{
			name: "struct pointer",
			args: args{
				sess:  sess,
				query: &user{},
			},
			wantC:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := getCollection(tt.args.sess, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("getCollection() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
