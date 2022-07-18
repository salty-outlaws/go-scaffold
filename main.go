package main

import (
	"fmt"
	"go-scaffold/handlers"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gaffold",
		Short: "Gaffold is a tool which can be used to create new projects with scaffoldings",
		Long:  `A Fast and Flexible Code Generator`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello!")
		},
	}

	initCommand = &cobra.Command{
		Use:   "init [project]",
		Short: "Initialize a new project in any project structure of choice",
		Run:   handlers.InitHandler,
		Args:  cobra.ExactArgs(1),
	}

	addCommand = &cobra.Command{
		Use:   "add",
		Short: "Adds a module to an existing project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("initializing new project")
		},
	}
)

func Execute() {
	rootCmd.AddCommand(initCommand)
	handlers.InitHandlerFlags(initCommand)

	rootCmd.AddCommand(addCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func main() {
	Execute()
}
