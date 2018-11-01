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
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// For quiting meetings
var quitmeetingCmd = &cobra.Command{
	Use:   "quitmeeting",
	Short: "quit meeting with title",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Quit Meeting called")
		queryTitle, _ := cmd.Flags().GetString("title")
		if queryTitle == "" {
			fmt.Println("Please input meeting title")
			return
		}
		// Make sure you have logged in
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("You should log in first")
		} else {
			if flag := service.QuitMeeting(user.Name, queryTitle); flag == false {
				fmt.Println("Make sure the meeting exit and you are in it")
			} else {
				fmt.Println("You have quited")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(quitmeetingCmd)

	// Here you will define your flags and configuration settings.
	quitmeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
