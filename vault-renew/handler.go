package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Handle a serverless request
func Handle(req []byte) string {
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
		URL: "https://github.com/OneTechLP/devops.vault.git",
		Auth: &http.BasicAuth{
			Username: string(usernameBytes),
			Password: string(passwordBytes),
		},
		Progress: os.Stdout,
	}
	r, err := git.PlainClone("./", false, o)
	if err != nil {
		return fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	}

	// visit := func(path string, info os.FileInfo, err error) error {
	// 	if info.IsDir() {
	// 		fmt.Println("dir:  ", path)
	// 	} else {
	// 		fmt.Println("file: ", path)
	// 	}
	// 	return nil
	// }

	// err = filepath.Walk("./", visit)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := os.Stat("./devops.vault"); os.IsNotExist(err) {
	// 	return fmt.Sprintf("{\"message\": \"git repo exists\"}")
	// }

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

	tree.Files().ForEach(func(f *object.File) error {
		fmt.Printf("File %s    %s\n", f.Hash, f.Name)
		return nil
	})

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	return fmt.Sprintf("Hello, Go. You said: %s", string(req))
}
