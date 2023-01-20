package database

import (
	"fmt"
	"github.com/juanesech/topo/constants"
	"github.com/juanesech/topo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoClient() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("%s://%s:%s@%s/", constants.DBSchema, constants.DBUser, constants.DBPass, constants.DBAddress)).
		SetServerAPIOptions(serverAPIOptions)

	rp, err := readpref.New(readpref.PrimaryMode)
	clientOptions.SetReadPreference(rp)
	clientOptions.SetDirect(true)

	client, err := mongo.NewClient(clientOptions)
	utils.CheckError(err)

	return client
}

func GetCollection(name string) *mongo.Collection {
	dbctx, dbclose := utils.GetCtx()
	client := GetMongoClient()
	defer dbclose()
	utils.CheckError(client.Connect(dbctx))

	return client.Database(constants.DBName).Collection(name)
}
