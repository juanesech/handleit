package gitlab

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/juanesech/topo/config"
	"github.com/juanesech/topo/utils"
	log "github.com/sirupsen/logrus"
)

// Gitlab Group
type Group struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Path string `json:"full_path"`
}

func (g Group) IsEmpty() bool {
	return reflect.DeepEqual(Group{}, g)
}

func GetGroup(source config.ModuleSource) Group {
	path := fmt.Sprintf("/api/v4/groups/%v", source.Group)
	gitlab := &Gitlab{
		Url:   source.Address,
		Token: source.Auth,
	}
	respGroup := Group{}
	res, err := gitlab.Get(path)
	utils.CheckError(err)

	_ = json.Unmarshal(res.Body(), &respGroup)
	log.Debug("Response struct: ", respGroup)
	log.Debug("Response: ", res, res.StatusCode())

	if respGroup.IsEmpty() {
		log.Errorf("Group with id %v is empty", respGroup.ID)
	}

	return respGroup
}
