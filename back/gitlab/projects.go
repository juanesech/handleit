package gitlab

import (
	"encoding/json"
	"fmt"

	"github.com/juanesech/topo/config"
	"github.com/juanesech/topo/utils"
	log "github.com/sirupsen/logrus"
)

type Project struct {
	Name string `json:"path"`
	Id   int    `json:"id"`
	Url  string `json:"http_url_to_repo"`
}

func GetProjects(source config.ModuleSource, group int) []Project {
	path := fmt.Sprintf("/api/v4/groups/%v/projects", group)
	gitlab := &Gitlab{
		Url:   source.Address,
		Token: source.Auth,
	}

	projects := []Project{}
	res, err := gitlab.Get(path)
	utils.CheckError(err)
	_ = json.Unmarshal(res.Body(), &projects)
	log.Debug(res)

	return projects
}
