package database

import (
	"os"

	"github.com/ravendb/ravendb-go-client"
	log "github.com/sirupsen/logrus"
)

var DBAddress = os.Getenv("DB_ADDRESS")

const Name = "topo"

func GetClient(databaseName string) *ravendb.DocumentStore {
	serverNodes := []string{DBAddress}
	store := ravendb.NewDocumentStore(serverNodes, databaseName)
	if err := store.Initialize(); err != nil {
		log.Fatal(err)
	}
	return store
}

var Client = GetClient(Name)
