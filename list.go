package faunashared

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// List retrieves a list of objects from a provided collection.
func List(collection Collection) (data map[string]f.Value, err error) {
	data, _, _, err = Paginate(collection, "", "", 0)
	return
}
