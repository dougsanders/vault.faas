package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type Request struct {
	RepoURLHTTP string `json:"repoUrlHttp"`
	PathRegex   string `json:"pathRegex"`
}

type FileInfo struct {
	File     string `json:"file"`
	Contents string `json:"contents"`
}

type Response struct {
	Files []FileInfo `json:"files"`
}

// Handle a serverless request
func Handle(bytes []byte) string {
	req := &Request{}
	if err := json.Unmarshal(bytes, req); err != nil {
		log.Fatal(err)
	}

	usernameBytes, err := ioutil.ReadFile("/var/openfaas/secrets/username")
	if err != nil {
		log.Fatal(err)
	}
	passwordBytes, err := ioutil.ReadFile("/var/openfaas/secrets/password")
	if err != nil {
		log.Fatal(err)
	}

	// Pull repository
	o := &git.CloneOptions{
		URL: req.RepoURLHTTP,
		Auth: &http.BasicAuth{
			Username: string(usernameBytes),
			Password: string(passwordBytes),
		},
	}
	r, err := git.PlainClone("./temp", false, o)
	if err != nil {
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}

	ref, err := r.Head()
	if err != nil {
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}

	tree, err := commit.Tree()
	if err != nil {
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}

	response := &Response{
		Files: []FileInfo{},
	}
	tree.Files().ForEach(func(f *object.File) error {
		matched, _ := regexp.MatchString(req.PathRegex, f.Name)
		if matched {
			contents, err := f.Contents()
			if err != nil {
				log.Fatal(err)
			}

			file := FileInfo{
				File:     f.Name,
				Contents: contents,
			}
			response.Files = append(response.Files, file)
		}

		return nil
	})

	os.RemoveAll("./temp/")

	respBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", string(respBytes))
}
