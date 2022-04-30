package types

type TestData struct {
	Item   string `fauna:"item" json:"item"`
	Number int    `fauna:"id" json:"number"`
	IsReal bool   `fauna:"isReal" json:"isReal"`
	Ref    string `fauna:"-" json:"-"`
}
