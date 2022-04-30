package faunashared

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func Create(collection Collection, obj interface{}) (string, f.Value, error) {
	res, err := db.Query(f.Create(
		f.Collection(collection),
		f.Obj{"data": obj},
	))
	if err != nil {
		return "", nil, err
	}

	value, err := res.At(f.ObjKey("data")).GetValue()
	if err != nil {
		return "", nil, err
	}

	var ref f.RefV
	if err := res.At(f.ObjKey("ref")).Get(&ref); err != nil {
		return "", nil, err
	}

	return ref.ID, value, nil
}
