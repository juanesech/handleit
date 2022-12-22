package module

import (
    db "github.com/juanesech/topo/database"
    "github.com/juanesech/topo/utils"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type ModuleResume struct {
	Name      string
	Providers []Provider
}

type Module struct {
	ID        string
	Name      string
	Variables []Variable
	Outputs   []Output
	Providers []Provider
}

func (m Module) WithID() Module {
    var modulefromdb bson.M

    dbctx, dbclose := utils.GetCtx()
    defer dbclose()

    coll := db.GetCollection("modules")
    filter := bson.D{{"name", m.Name}}
    findsrc := coll.FindOne(dbctx, filter).Decode(&modulefromdb)
    utils.CheckError(findsrc)
    m.ID = modulefromdb["_id"].(primitive.ObjectID).Hex()

    return m
}
type Variable struct {
	Name        string
	Type        string
	Description string
	Default     string
	Required    bool
}

type Output struct {
	Name        string
	Description string
}

type Provider struct {
	Source             string
	VersionConstraints []string
}

type ImportRequest struct {
	Name string `json:"name"`
}
