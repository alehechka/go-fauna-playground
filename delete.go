package faunashared

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func Delete(collection Collection, ref string) (f.Value, error) {
	res, err := db.Query(
		f.Delete(f.Ref(f.Collection(collection), ref)),
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
