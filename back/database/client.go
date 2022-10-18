package database

import (
	"github.com/juanesech/topo/constants"
	"github.com/ravendb/ravendb-go-client"
	log "github.com/sirupsen/logrus"
)

func GetClient(databaseName string) *ravendb.DocumentStore {
	serverNodes := []string{constants.DBAddress}
	store := ravendb.NewDocumentStore(serverNodes, databaseName)
    dc := ravendb.NewDocumentConventions()
	dc.MaxNumberOfRequestsPerSession = 100
	store.SetConventions(dc)
	if err := store.Initialize(); err != nil {
		log.Error(err)
	}
	return store
}

var Client = GetClient(constants.DBName)
