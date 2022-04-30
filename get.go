package faunashared

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// Get retrieves single document from desired collection.
// The 'obj' parameter must be passed as a reference to the object to map to.
func Get(collection Collection, documentRef string, obj interface{}) error {
	res, err := db.Query(f.Get(f.Ref(f.Collection(collection), documentRef)))
	if err != nil {
		return err
	}

	return res.At(f.ObjKey("data")).Get(obj)
}

func GetFromIndex(collection Collection, index interface{}, obj interface{}) (f.Value, error) {
	res, err := db.Query(f.Get(f.Match(f.Index("installationID"))))
	if err != nil {
		return res, err
	}

	return res, res.At(f.ObjKey("data")).Get(obj)
}
