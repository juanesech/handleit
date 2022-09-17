package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/juanesech/handleit/config"
	"github.com/juanesech/handleit/utils"
	log "github.com/sirupsen/logrus"
)

type Group struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Path string `json:"path"`
}

func (g Group) IsEmpty() bool {
	return reflect.DeepEqual(Group{}, g)
}

func GetGroup(name string) Group {
	glconfig := config.GetSource("gitlab")
	path := fmt.Sprintf("/api/v4/groups?search=%v", name)
	gitlab := &Gitlab{
		Url:   glconfig.Address,
		Token: glconfig.Auth,
		Client: &http.Client{
			Timeout: time.Duration(30 * time.Second),
		},
	}

	resp := gitlab.Get(path)
	respGroup := &[]Group{}
	utils.CheckError(json.NewDecoder(resp.Body).Decode(respGroup))

	var retGroup Group
	for _, rg := range *respGroup {
		if rg.Name == name {
			retGroup = rg
		}
	}

	if retGroup.IsEmpty() {
		log.Error("Group not found")
	}

	return retGroup
}
