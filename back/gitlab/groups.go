package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/juanesech/topo/config"
	"github.com/juanesech/topo/utils"
	log "github.com/sirupsen/logrus"
)

type Group struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Path string `json:"full_path"`
}

func (g Group) IsEmpty() bool {
	return reflect.DeepEqual(Group{}, g)
}

func GetGroup(source config.ModuleSource) Group {
	name := utils.GetLast(source.Group, "/")

	path := fmt.Sprintf("/api/v4/groups?search=%v", name)
	gitlab := &Gitlab{
		Url:   source.Address,
		Token: source.Auth,
		Client: &http.Client{
			Timeout: time.Duration(30 * time.Second),
		},
	}

	resp := gitlab.Get(path)
	respGroup := &[]Group{}
	utils.CheckError(json.NewDecoder(resp.Body).Decode(respGroup))

	var retGroup Group
	for _, rg := range *respGroup {
		if rg.Name == name && rg.Path == source.Group {
			retGroup = rg
		}
	}

	if retGroup.IsEmpty() {
		log.Errorf("Group %s not found", name)
	}

	return retGroup
}
