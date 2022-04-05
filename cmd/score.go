/*
Copyright © 2022 NAME HERE <EMAIL SetRESS>

Licensed under the Apache License, Version 2.0 (the"License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an"AS IS"BASIS,
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

	"github.com/jijosg/nba_api/pkg/teams"
	"github.com/spf13/cobra"
)

// scoreCmd represents the score command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "view today's score",
	Long:  `View today's score.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get data from api
		tr := &http.Transport{
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}

		req, err := http.NewRequest("GET", "https://cdn.nba.com/static/json/liveData/scoreboard/todaysScoreboard_00.json", nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Accept-Encoding", "identity")
		req.Header.Add("Content-Type", "application/json; charset=UTF-8")
		req.Header.Set("Host", "cdn.nba.com")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Accept-Language", "en-US,en;q=0.5")
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

		// UnMarshall json response to ScoreboardResponse object
		var scoreboardResponse teams.ScoreboardResponse
		err = json.Unmarshal(body, &scoreboardResponse)
		if err != nil {
			fmt.Println("Error unmarshalling json:", err)
			return
		}
		// Print scoreboard
		if len(scoreboardResponse.Scoreboard.Games) == 0 {
			fmt.Println("No games today :(")
			return
		} else {
			fmt.Printf("%-3s %-13s %-13s %-11s %-9s\n", "No.", "Home Team", "Away Team", "Game Status", "Scores")
			for i, s := range scoreboardResponse.Scoreboard.Games {
				fmt.Printf("%-3d %-13s %-13s %-11s %-4d-%4d\n", i+1, s.HomeTeam.TeamName, s.AwayTeam.TeamName, s.GameStatusText, s.HomeTeam.Score, s.AwayTeam.Score)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scoreCmd.PersistentFlags().String("foo","","A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scoreCmd.Flags().BoolP("toggle","t", false,"Help message for toggle")
}
