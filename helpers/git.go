package helpers

import (
	"gopkg.in/src-d/go-git.v4"
)

func CloneRepo(repo string, projectName string) error {
	_, err := git.PlainClone(projectName, false, &git.CloneOptions{
		URL: "https://github.com/" + repo,
	})
	return err
}
