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
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {

		var idlist []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)

			if err != nil {
				fmt.Println("Failed to convert arg", arg)
			} else {

				idlist = append(idlist, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong")
			return
		}

		for _, id := range idlist {

			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number")
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)
			if err != nil {
				fmt.Printf("Failed to Mark %d as completed", id)

			} else {
				fmt.Printf("Marked %d as completed", id)

			}

		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
