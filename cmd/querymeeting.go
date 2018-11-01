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
	"agenda/entity"
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// For querying meetings 
var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting",
	Short: "query meetings with start and end time",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Query Meeting called")
		startTime, _ := cmd.Flags().GetString("starttime")
		endTime, _ := cmd.Flags().GetString("endtime")
		// Time can not be empty
		if startTime == "" || endTime == "" {
			fmt.Println("Please input the start and end time of the interval")
			return
		}
		// Make sure you have logged in
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("You should log in first")
		} else {
			if meetingList, flag := service.QueryMeeting(user.Name, startTime, endTime); flag != true {
				fmt.Println("Please input the correct data format like 2018-10-01/00:01")
			} else {
				fmt.Println("Meetings are as follow:")
				for _, oneMeeting := range meetingList {
					fmt.Println("****************")
					fmt.Println("Title: ", oneMeeting.Title)
					startTimeStr, _ := entity.DateToString(oneMeeting.StartDate)
					fmt.Println("Start Time", startTimeStr)
					endTimeStr, _ := entity.DateToString(oneMeeting.EndDate)
					fmt.Println("End Time", endTimeStr)
					fmt.Printf("Participator(s): ")
					for _, parts := range oneMeeting.Participators {
						fmt.Printf(parts, " ")
					}
					fmt.Printf("\n")
					fmt.Println("****************")
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(querymeetingCmd)

	// Here you will define your flags and configuration settings.
	querymeetingCmd.Flags().StringP("starttime", "s", "", "the start time of the meeting")
	querymeetingCmd.Flags().StringP("endtime", "e", "", "the end time of the meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
