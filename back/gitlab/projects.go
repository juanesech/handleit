package gitlab

import (
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

	projects := &[]Project{}
	resp, err := gitlab.Get(path, projects)
	utils.CheckError(err)
	log.Debug(resp)

	return *projects
}
