package mongo

// Collection to show database and collection.
type Collection interface {
	Database() string
	Collection() string
}
