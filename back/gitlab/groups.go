package gitlab

import (
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
	name := utils.GetLast(source.Group, "/")

	path := fmt.Sprintf("/api/v4/groups?search=%v", name)
	gitlab := &Gitlab{
		Url:   source.Address,
		Token: source.Auth,
	}
	respGroup := &[]Group{}
	resp, err := gitlab.Get(path, respGroup)
	utils.CheckError(err)
	log.Debug(resp)

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
