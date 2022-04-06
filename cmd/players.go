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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jijosg/nba_api/pkg/player"
	"github.com/spf13/cobra"
)

var playerID string

// playersCmd represents the players command
var playersCmd = &cobra.Command{
	Use:   "players",
	Short: "List the player info",
	Long:  `List the player info for a given player ID`,
	Run: func(cmd *cobra.Command, args []string) {
		tr := &http.Transport{
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}

		req, err := http.NewRequest("GET", "https://stats.nba.com/stats/commonplayerinfo", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		q := req.URL.Query()
		q.Add("PlayerID", playerID)
		q.Add("LeagueID", "00")
		req.URL.RawQuery = q.Encode()

		req.Header.Add("Accept-Encoding", "identity")
		req.Header.Add("Accept", "application/json, text/plain, */*")
		req.Header.Add("x-nba-stats-origin", "stats")
		req.Header.Add("x-nba-stats-token", "true")
		req.Header.Add("Host", "stats.nba.com")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0")
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("Referer", "https://stats.nba.com/")
		req.Header.Add("Pragma", "no-cache")
		req.Header.Add("Cache-Control", "no-cache")
		req.Header.Add("Accept-Language", "en-US,en;q=0.5")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body:", err)
			return
		}
		var player player.PlayerInfo
		// unmarshall data as json
		err = json.Unmarshal(body, &player)
		if err != nil {
			fmt.Println("Error unmarshalling data:", player)
			return
		}
		for _, i := range player.ResultSets {
			if i.Name == "PlayerHeadlineStats" {
				for _, h := range i.Headers {
					fmt.Printf("%-14s", h)
				}
				fmt.Println()
				for _, j := range i.RowSet[0] {
					fmt.Printf("%-14v", j)
				}
				fmt.Println()
			}

		}

	},
}

func init() {
	listCmd.AddCommand(playersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	playersCmd.Flags().StringVarP(&playerID, "playerId", "p", "2544", "Player ID to get info")
}
