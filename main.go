package main

import (
	"log"
	"path/filepath"

	"github.com/cumulusware/grab/newapp"
	"github.com/spf13/cobra"
)

func newApp(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Usage()
		log.Fatalln("Path for new site needs to be provided")
	}

	appPath, err := filepath.Abs(filepath.Clean(args[0]))
	if err != nil {
		cmd.Usage()
		log.Fatalln(err)
	}

	newapp.CreateApp(appPath)
}

func main() {

	var newCmd = &cobra.Command{
		Use:   "new <app-name>",
		Short: "Create directory <app-name> and initialize as REST API",
		Run:   newApp,
	}
	var rootCmd = &cobra.Command{Use: "grab"}
	rootCmd.AddCommand(newCmd)
	rootCmd.Execute()

}
