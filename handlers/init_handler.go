package handlers

import (
	"fmt"
	"go-scaffold/constants"
	"go-scaffold/helpers"
	"os"

	"github.com/spf13/cobra"
)

var (
	flagRepo     string
	flagLanguage string
)

func InitHandlerFlags(c *cobra.Command) {
	c.Flags().StringVarP(&flagRepo, "repo", "r", "", "-r repo (orgname/reponame)")
	c.Flags().StringVarP(&flagLanguage, "lang", "l", "", "-l language (go, python, rust)")
}

func InitHandler(cmd *cobra.Command, args []string) {
	fmt.Println("flags", flagRepo, flagLanguage)
	fmt.Println("args", args[0])

	projectName := args[0]

	repoToClone := flagRepo
	if repoToClone == "" && flagLanguage != "" {
		if langTemplate, ok := constants.LanguageDefaultTemplates[flagLanguage]; ok {
			repoToClone = langTemplate
		}
	}

	if repoToClone != "" {
		err := helpers.CloneRepo(flagRepo, projectName)
		helpers.Must(err)
		helpers.PathWalker(projectName,
			func(path string) string { return helpers.PathRenamer(path, constants.TemplateName, projectName) },
			func(path string) string {
				helpers.ContentRenamer(path, constants.TemplateName, projectName)
				return helpers.PathRenamer(path, constants.TemplateName, projectName)
			})
		fmt.Println("initialized repo")
		return
	}
	fmt.Println("no repos or languages specified")
	os.Exit(1)
}
