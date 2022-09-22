package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/juanesech/handleit/config"
	"github.com/juanesech/handleit/utils"
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
		Client: &http.Client{
			Timeout: time.Duration(30 * time.Second),
		},
	}

	resp := gitlab.Get(path)
	defer resp.Body.Close()

	projects := &[]Project{}
	utils.CheckError(json.NewDecoder(resp.Body).Decode(projects))

	return *projects
}
