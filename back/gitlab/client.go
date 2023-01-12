package gitlab

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juanesech/topo/utils"
	log "github.com/sirupsen/logrus"
)

// Gitlab client struct
type Gitlab struct {
	Url    string
	Token  string
	Client *http.Client
}

func (provider *Gitlab) request(req *http.Request) (response *http.Response, err error) {
	clonedReq := req.Clone(req.Context())
	if req.Body != nil {
		clonedReq.Body, err = req.GetBody()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	defer func() {
		if clonedReq.Body != nil {
			clonedReq.Body.Close()
		}
	}()

	// set headers common to all requests
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", provider.Token))
	response, err = provider.Client.Do(req)
	if err != nil {
		log.Fatal("request: ", err)
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Debug(string(bodyBytes))
		}
		// ask for a new refresh token
		response.Body.Close()
		response, err = provider.request(clonedReq)
		utils.CheckError(err)
	}

	return response, err
}

// Send a Get request to Gitlab instance
func (provider *Gitlab) Get(path string) (response *http.Response) {
	req, err := http.NewRequest(http.MethodGet, provider.Url+path, nil)
	if err != nil {
		return
	}
	rsp, err := provider.request(req)
	if err != nil {
		log.Fatal("API: ", err)
		return
	}

	return rsp
}

// Send a Post request to Gitlab instance
func (provider *Gitlab) Post(path string, reqBody []byte) (response *http.Response) {
	req, err := http.NewRequest(http.MethodPost, provider.Url+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	rsp, err := provider.request(req)
	if err != nil {
		log.Fatal("API: ", err)
		return
	}

	return rsp
}

// Send a Put request to Gitlab instance
func (provider *Gitlab) Put(path string, reqBody []byte) (response *http.Response) {
	req, err := http.NewRequest(http.MethodPut, provider.Url+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	rsp, err := provider.request(req)
	if err != nil {
		log.Fatal("API: ", err)
		return
	}

	return rsp
}
