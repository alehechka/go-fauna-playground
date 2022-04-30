package faunashared

import (
	"github.com/alehechka/go-fauna-playground/types"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func CreateTestData(testdata types.TestData) (types.TestData, error) {
	ref, value, err := Create(testData, testdata)
	if err != nil {
		return types.TestData{}, err
	}

	var t types.TestData
	if err := value.Get(&t); err != nil {
		return types.TestData{}, err
	}

	t.Ref = ref

	return t, nil
}

func GetTestData(documentRef string) (testdata types.TestData, err error) {
	err = Get(testData, documentRef, &testdata)
	testdata.Ref = documentRef

	return
}

func PaginateTestData(before, after string, size int) (testdata []types.TestData, nextRef, prevRef *string, err error) {
	data, nextRef, prevRef, err := Paginate(testData, before, after, size)
	if err != nil {
		return []types.TestData{}, nil, nil, err
	}

	return parseTestDataArr(data), nextRef, prevRef, nil
}

func UpdateTestData(documentRef string, testdata types.TestData) (types.TestData, error) {
	value, err := Update(testData, documentRef, testdata)
	if err != nil {
		return types.TestData{}, err
	}

	var t types.TestData
	if err := value.Get(&t); err != nil {
		return types.TestData{}, err
	}

	t.Ref = documentRef

	return t, nil
}

func DeleteTestData(documentRef string) (types.TestData, error) {
	value, err := Delete(testData, documentRef)
	if err != nil {
		return types.TestData{}, err
	}

	var t types.TestData
	if err := value.Get(&t); err != nil {
		return types.TestData{}, err
	}

	t.Ref = documentRef

	return t, nil
}

func parseTestDataArr(data map[string]f.Value) []types.TestData {
	testdataArr := make([]types.TestData, 0)
	for id, value := range data {
		var d types.TestData
		value.Get(&d)
		d.Ref = id

		testdataArr = append(testdataArr, d)
	}

	return testdataArr
}
