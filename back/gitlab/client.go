package gitlab

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

var (
	Client *resty.Client
)

func init() {
	Client = resty.New()
}

// Gitlab client struct
type Gitlab struct {
	Url   string
	Token string
}

// Get request to Gitlab instance
func (provider *Gitlab) Get(path string, resp ...interface{}) (response *resty.Response, err error) {
	rsp, err := Client.R().
		SetHeader("PRIVATE-TOKEN", provider.Token).
		SetResult(resp).
		Get(provider.Url + path)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// Post request to Gitlab instance
func (provider *Gitlab) Post(path string, reqBody interface{}) (response *resty.Response, err error) {
	rsp, err := Client.R().
		SetHeader("PRIVATE-TOKEN", provider.Token).
		SetBody(reqBody).
		Post(provider.Url + path)
	if err != nil {
		log.Error("API: ", err)
		return nil, err
	}

	return rsp, nil
}

// Put request to Gitlab instance
func (provider *Gitlab) Put(path string, reqBody interface{}) (response *resty.Response, err error) {
	rsp, err := Client.R().
		SetHeader("PRIVATE-TOKEN", provider.Token).
		SetBody(reqBody).
		Put(provider.Url + path)
	if err != nil {
		log.Fatal("API: ", err)
		return nil, err
	}

	return rsp, nil
}
