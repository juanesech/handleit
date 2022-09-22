package utils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Clone(folder, token, url string) {
	path := fmt.Sprintf("/tmp/%s", folder)
	CheckError(os.MkdirAll(path, os.ModePerm))
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: Logger.Out,
		Auth: &http.BasicAuth{
			Username: " ",
			Password: token,
		},
	})
	Logger.Infof("Cloning to %s", path)
	CheckError(err)
}
