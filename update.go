package faunashared

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func Update(collection Collection, ref string, obj interface{}) (f.Value, error) {
	res, err := db.Query(
		f.Update(
			f.Ref(f.Collection(collection), ref),
			f.Obj{"data": obj},
		),
	)
	if err != nil {
		return nil, err
	}

	value, err := res.At(f.ObjKey("data")).GetValue()
	if err != nil {
		return nil, err
	}

	return value, nil
}
