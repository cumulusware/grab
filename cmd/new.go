// Copyright (c) 2015-2017 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package cmd

import (
	"log"
	"path/filepath"

	newapp "github.com/cumulusware/grab/generators/new"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Initialize new REST API in <app-name>",
	Long: `Create a new directory named <app-name> and initialize a new
REST API project`,
	Run: newApp,
}

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
func init() {
	RootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
