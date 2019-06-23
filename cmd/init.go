/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates the necessary emojis needed in your slack workspace",
	Run: func(cmd *cobra.Command, args []string) {
		runInit(cmd, args)
	},
}


func runInit(cmd *cobra.Command, args []string) {
	unzipArchive(args)
}

func unzipArchive(args []string) {
	r, err := zip.OpenReader(args[0])
	if err != nil {
		panic(err)
	}

	defer r.Close()
	for _, f := range r.File {
		match, err := regexp.MatchString(`.*/g\d{1,3}\.png|.*/t.png`, f.Name)
		if err != nil {
			panic(err)
		}

		if match {
			fmt.Printf("Found file %s in zip\n", f.Name)
		}
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
