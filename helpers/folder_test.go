package helpers

import (
	"os"
	"testing"
)

func TestPathWalker(t *testing.T) {
	os.Chdir("../project2")
	folders, files := PathWalker(".",
		func(path string) string { return PathRenamer(path, "go-scaffold-project-name", "project2") },
		func(path string) string {
			ContentRenamer(path, "go-scaffold-project-name", "project2")
			return PathRenamer(path, "go-scaffold-project-name", "project2")
		})
	_, _ = folders, files
}
