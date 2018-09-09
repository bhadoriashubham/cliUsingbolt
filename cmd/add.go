// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"cli/db"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var add = &cobra.Command{
	Use:   "add",
	Short: "Adds a task that needs to be done",

	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := db.CreateTasks(task)

		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}

		fmt.Printf("Added %s task to your task list \n", task)

	},
}

func init() {
	RootCmd.AddCommand(add)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
