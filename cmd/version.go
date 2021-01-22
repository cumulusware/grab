// Copyright (c) 2015-2021 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	ver = "0.1.0"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version of grab",
	Long:  `Print grab's version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Go REST API Builder v%s\n", ver)
	},
}
