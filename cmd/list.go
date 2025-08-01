/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/jijosg/nba_api/pkg/teams"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all NBA teams",
	Long:  `List all NBA teams that exist today. This command does NOT list players. For player information, use the 'players' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		fmt.Printf("%-12s%-14s%-14s%-14s%-15s%-24s%s\n", "TEAM_ID", "ABBREVIATION", "NICKNAME", "YEAR_FOUNDED", "CITY", "FULLNAME", "STATE")
		for _, i := range teams.NBATeams {
			fmt.Printf("%-12d%-14s%-14s%-14d%-15s%-24s%s\n", i.Id, i.Abbreviation, i.Nickname, i.YearFounded, i.City, i.FullName, i.State)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
