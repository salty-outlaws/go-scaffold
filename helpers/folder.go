package helpers

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// PathWalker - returns full path to folders and files
func PathWalker(path string, folderRunner func(path string) string, fileRunner func(path string) string) ([]string, []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	/*
		INFO
		When folderRunner or fileRunner is invoked, they might get renamed,
		so the renamed full path is expected as an output of folderRunner and fileRunner
		currently deletion of the same is not handled

	*/

	folderE := []string{}
	fileE := []string{}

	for _, file := range files {
		cPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			// run function on folder path
			cPath = folderRunner(cPath)
			// add folder to main list
			folderE = append(folderE, cPath)

			subFolders, subFiles := PathWalker(cPath, folderRunner, fileRunner)
			folderE = append(folderE, subFolders...)
			fileE = append(fileE, subFiles...)
		} else {
			// run function on file path
			cPath = fileRunner(cPath)
			// add file to main list
			fileE = append(fileE, cPath)
		}
	}
	return folderE, fileE
}

// PathRenamer renames files or folders if they contain the keyword by replacing the keyword with replacement
func PathRenamer(path, keyword, replacement string) string {
	_, lastPathPart := filepath.Split(path)
	if strings.Contains(lastPathPart, keyword) {
		to := filepath.Join(filepath.Dir(path), strings.ReplaceAll(lastPathPart, keyword, replacement))
		os.Rename(path, to)
		return to
	}
	return path
}

// ContentRenamer renames contents by replacing the keyword with replacement
func ContentRenamer(path, keyword, replacement string) error {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(strings.ReplaceAll(string(input), keyword, replacement)), 0644)
	if err != nil {
		return err
	}
	return nil
}
