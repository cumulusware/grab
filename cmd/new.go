// Copyright (c) 2015-2021 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"log"
	"path/filepath"

	"github.com/cumulusware/grab/internal/newapp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new <app-name>",
	Short: "Create directory <app-name> and initialize as REST API",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}
