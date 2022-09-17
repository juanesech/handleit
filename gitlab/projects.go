package gitlab

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"time"

// 	"github.com/juanesech/handleit/utils"
// )

// var logger = utils.Logger

// type Project struct {
// 	Name                string    `json:"name"`
// 	Description         string    `json:"description"`
// 	Id                  int       `json:"id"`
// 	NamespaceId         int       `json:"namespace_id"`
// 	Namespace           namespace `json:"namespace"`
// 	ProjectTemplateId   int       `json:"template_project_id"`
// 	UseCustomTemplate   bool      `json:"use_custom_template"`
// 	Visibility          string    `json:"visibility"`
// 	MergeMethod         string    `json:"merge_method"`
// 	MergeRequestEnabled bool      `json:"merge_requests_enabled"`
// 	Squash              string    `json:"squash_option"`
// 	SquashTemplate      string    `json:"squash_commit_template"`
// 	OnlyResolvedMR      bool      `json:"only_allow_merge_if_all_discussions_are_resolved"`
// 	RemoveSourceBranch  bool      `json:"remove_source_branch_after_merge"`
// 	Approvals           int       `json:"approvals_before_merge"`
// }

// type namespace struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// 	Path string `json:"path"`
// 	Kind string `json:"kind"`
// }

// func (p Project) SetApprovals() {
// 	path := fmt.Sprintf("/api/v4/projects/%v/approvals", p.Id)
// 	gitlab := &Gitlab{
// 		Url: constants.GitlabURL,
// 		Client: &http.Client{
// 			Timeout: time.Duration(30 * time.Second),
// 		},
// 	}

// 	requestBody, err := json.Marshal(map[string]interface{}{
// 		"reset_approvals_on_push":                    true,
// 		"merge_requests_author_approval":             false,
// 		"merge_requests_disable_committers_approval": true,
// 	})
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	resp := gitlab.Post(path, requestBody)
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	logger.Debug(string(body))

// }

// func (p Project) SetApprovalRules() {
// 	path := fmt.Sprintf("/api/v4/projects/%v/approval_rules", p.Id)
// 	gitlab := &Gitlab{
// 		Url: constants.GitlabURL,
// 		Client: &http.Client{
// 			Timeout: time.Duration(30 * time.Second),
// 		},
// 	}

// 	approvers := []int{constants.PlatformGorupId, p.Namespace.Id}
// 	requestBody, err := json.Marshal(map[string]interface{}{
// 		"id":                 p.Id,
// 		"name":               "required_approvals",
// 		"approvals_required": 1,
// 		"group_ids":          approvers,
// 	})
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	logger.Debug("Request: ", string(requestBody))

// 	resp := gitlab.Post(path, requestBody)
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	logger.Debug("Response: ", string(body))

// }

// func (p Project) SetupDeployKey() {
// 	path := fmt.Sprintf("/api/v4/projects/%v/deploy_keys/%v/enable", p.Id, constants.AutomationCIKey)
// 	logger.Info(logger.GetLevel())
// 	gitlab := &Gitlab{
// 		Url: constants.GitlabURL,
// 		Client: &http.Client{
// 			Timeout: time.Duration(30 * time.Second),
// 		},
// 	}

// 	requestBody, err := json.Marshal(map[string]interface{}{})
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	logger.Debug("Request: ", string(requestBody))

// 	resp := gitlab.Post(path, requestBody)
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	logger.Debug("Response: ", string(body))

// }

// func ListProjects() []Project {
// 	client := &http.Client{}
// 	path := "/api/v4/projects"

// 	p := &Project{}

// 	requestBody, err := json.Marshal(p)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	responseBody := bytes.NewBuffer(requestBody)

// 	req, err := http.NewRequest("GET", constants.GitlabURL+path, responseBody)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", constants.Token))
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	projects := &[]Project{}

// 	if err := json.NewDecoder(resp.Body).Decode(projects); err != nil {
// 		logger.Fatal(err)
// 	}

// 	return *projects
// }

// func NewProject(name string, groupId int, templateId int) Project {
// 	client := &http.Client{}
// 	path := "/api/v4/projects"
// 	project := &Project{
// 		Name:                name,
// 		NamespaceId:         groupId,
// 		ProjectTemplateId:   templateId,
// 		UseCustomTemplate:   true,
// 		Visibility:          "internal",
// 		MergeMethod:         "rebase_merge",
// 		MergeRequestEnabled: true,
// 		Squash:              "always",
// 		SquashTemplate:      "%{title} \n %{all_commits}",
// 		OnlyResolvedMR:      true,
// 		RemoveSourceBranch:  true,
// 		Approvals:           1,
// 	}

// 	checkProject := GetProject(name)

// 	if checkProject.Id != 0 {
// 		logger.Info(fmt.Sprintf("Project %v already exists", name))
// 		checkProject.SetApprovals()
// 		return checkProject
// 	}

// 	requestBody, err := json.Marshal(project)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	responseBody := bytes.NewBuffer(requestBody)

// 	req, err := http.NewRequest("POST", constants.GitlabURL+path, responseBody)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", constants.Token))

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	respProject := &Project{}

// 	if err := json.NewDecoder(resp.Body).Decode(respProject); err != nil {
// 		logger.Fatal(err)
// 	}

// 	respProject.SetApprovals()

// 	return *respProject
// }

// func GetProject(name string) Project {
// 	path := fmt.Sprintf("/api/v4/projects?search=%v", name)
// 	gitlab := &Gitlab{
// 		Url: constants.GitlabURL,
// 		Client: &http.Client{
// 			Timeout: time.Duration(30 * time.Second),
// 		},
// 	}

// 	resp := gitlab.Get(path)
// 	defer resp.Body.Close()

// 	resProject := &[]Project{}
// 	if err := json.NewDecoder(resp.Body).Decode(resProject); err != nil {
// 		logger.Fatal(err)
// 	}

// 	var retProject Project
// 	for _, rp := range *resProject {
// 		if rp.Name == name {
// 			retProject = rp
// 			logger.Debug(rp.Name)
// 		}
// 	}

// 	return retProject
// }

// func SetConfig(name string) {
// 	project := GetProject(name)
// 	logger.Debug(project)
// 	path := fmt.Sprintf("/api/v4/projects/%v", project.Id)
// 	gitlab := &Gitlab{
// 		Url: constants.GitlabURL,
// 		Client: &http.Client{
// 			Timeout: time.Duration(30 * time.Second),
// 		},
// 	}

// 	project.Visibility = "internal"
// 	project.MergeMethod = "rebase_merge"
// 	project.MergeRequestEnabled = true
// 	project.Squash = "always"
// 	project.SquashTemplate = "%{title}\n\n%{all_commits}"
// 	project.OnlyResolvedMR = true
// 	project.RemoveSourceBranch = true
// 	project.Approvals = 1

// 	project.SetApprovals()
// 	project.SetApprovalRules()
// 	project.SetupDeployKey()

// 	requestBody, err := json.Marshal(project)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	resp := gitlab.Put(path, requestBody)
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	logger.Debug(string(body))

// }
