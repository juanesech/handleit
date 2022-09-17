package database

import (
	"github.com/ravendb/ravendb-go-client"
	log "github.com/sirupsen/logrus"
)

var DBAddress = "http://localhost:8083"

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
