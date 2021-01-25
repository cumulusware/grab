// Copyright (c) 2015-2021 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"log"

	"github.com/cumulusware/grab/internal/newapp"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new REST API app",
	Run: func(cmd *cobra.Command, args []string) {
		err := newapp.CreateApp()
		if err != nil {
			log.Fatal(err)
		}
	},
}
