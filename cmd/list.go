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

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the tasks",

	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks")
			return
		}

		fmt.Println("You have Following tasks")
		for i, task := range tasks {
			fmt.Printf("%d . %s\n", i+1, task.Value)
		}

	},
}

func init() {

	RootCmd.AddCommand(listCmd)

}
