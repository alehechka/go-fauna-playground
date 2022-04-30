package faunashared

import (
	"os"

	"github.com/fauna/faunadb-go/v4/faunadb"
)

var db *faunadb.FaunaClient

// Collection refers to a FaunaDB Collection name
type Collection string

func (c Collection) String() string {
	return string(c)
}

const (
	testData            Collection = "test-data"
	gitHubOAuth         Collection = "github_oauth"
	githubInstallations Collection = "github_installations"
)

// InitializeFaunaDBClient initializes a FaunaClient from available env secret
func InitializeFaunaDBClient() {
	db = faunadb.NewFaunaClient(os.Getenv("FAUNADB_CLIENT_SECRET"), faunadb.Endpoint("https://db.us.fauna.com"))
}
