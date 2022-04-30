package faunashared

import (
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// Paginate retrieves a list of objects from a provided collection using the provided pagination parameters.
func Paginate(collection Collection, before, after string, size int) (data map[string]f.Value, nextRef, prevRef *string, err error) {

	options := getPaginationOptions(collection, before, after, size)

	res, err := db.Query(f.Map(
		f.Paginate(
			f.Documents(f.Collection(collection)),
			options...),
		f.Lambda("X", f.Get(f.Var("X"))),
	))

	if err != nil {
		return nil, nil, nil, err
	}

	data, nextRef, prevRef = createMappedValues(res)

	return data, nextRef, prevRef, nil
}

// PaginateSize retrieves a list of objects from a provided collection with a max number of paginated results.
func PaginateSize(collection Collection, size int) (data map[string]f.Value, nextRef, prevRef *string, err error) {
	return Paginate(collection, "", "", size)
}

// PaginateAfter retrieves a list of objects from a provided collection with a max number of paginated results starting from provided after ref.
func PaginateAfter(collection Collection, after string, size int) (data map[string]f.Value, nextRef, prevRef *string, err error) {
	return Paginate(collection, "", after, size)
}

// PaginateBefore retrieves a list of objects from a provided collection with a max number of paginated results looking backwards from provided before ref.
func PaginateBefore(collection Collection, before string, size int) (data map[string]f.Value, nextRef, prevRef *string, err error) {
	return Paginate(collection, before, "", size)
}

func getPaginationOptions(collection Collection, before, after string, size int) []f.OptionalParameter {
	var options []f.OptionalParameter

	if size > 0 {
		options = append(options, f.Size(size))
	}

	if len(before) > 0 {
		options = append(options, f.Before(f.Ref(f.Collection(collection), before)))
	}

	if len(after) > 0 {
		options = append(options, f.After(f.Ref(f.Collection(collection), after)))
	}

	return options
}

func createMappedValues(res f.Value) (data map[string]f.Value, nextRef, prevRef *string) {
	valueMap := make(map[string]f.Value, 0)

	for i := 0; true; i++ {
		value, err := res.At(f.ObjKey("data").AtIndex(i).AtKey("data")).GetValue()
		if err != nil {
			break
		}

		var ref f.RefV
		res.At(f.ObjKey("data").AtIndex(i).AtKey("ref")).Get(&ref)

		valueMap[ref.ID] = value
	}

	return valueMap, getNextRef(res), getPrevRef(res)
}

func getNextRef(res f.Value) *string {
	var ref f.RefV
	if err := res.At(f.ObjKey("after").AtIndex(0)).Get(&ref); err != nil {
		return nil
	}

	return &ref.ID
}

func getPrevRef(res f.Value) *string {
	var ref f.RefV
	if err := res.At(f.ObjKey("before").AtIndex(0)).Get(&ref); err != nil {
		return nil
	}

	return &ref.ID
}
