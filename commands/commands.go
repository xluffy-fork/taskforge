package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	Root.AddCommand(new)
	Root.AddCommand(query)
	Root.AddCommand(complete)
	Root.AddCommand(edit)
	Root.AddCommand(next)
}

// Root is the root CLI command
var Root = &cobra.Command{
	Use:   "tsk",
	Short: "Manage your tasks",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		file := findConfigFile()

		var err error
		config, err = loadConfigFile(file)
		if err != nil {
			fmt.Println("ERROR Unable to load config file:", err)
			os.Exit(1)
		}
	},
}