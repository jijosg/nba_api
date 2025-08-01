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
	"io"
	"net/http"
	"strings"

	"github.com/jijosg/nba_api/pkg/player"
	"github.com/spf13/cobra"

	// Import rootCmd from parent package
	. "github.com/jijosg/nba_api/cmd"
)

var playerID string
var playerName string
var showRecent bool

// playersCmd represents the players command
var playersCmd = &cobra.Command{
	Use:   "players",
	Short: "List the player info",
	Long:  `List the player info for a given player ID or search by player name`,
	Run: func(cmd *cobra.Command, args []string) {
		if playerName != "" {
			tr := &http.Transport{DisableCompression: true}
			client := &http.Client{Transport: tr}
			req, err := http.NewRequest("GET", "https://stats.nba.com/stats/commonallplayers", nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			q := req.URL.Query()
			q.Add("LeagueID", "00")
			q.Add("Season", "2023-24") // You may want to make this dynamic
			q.Add("IsOnlyCurrentSeason", "0")
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
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading body:", err)
				return
			}
			var data map[string]interface{}
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println("Error unmarshalling data:", err)
				return
			}
			resultSets, ok := data["resultSets"].([]interface{})
			if !ok || len(resultSets) == 0 {
				fmt.Println("No resultSets found")
				return
			}
			playersSet := resultSets[0].(map[string]interface{})
			headers := playersSet["headers"].([]interface{})
			rowSet := playersSet["rowSet"].([]interface{})
			fmt.Printf("%-10s %-25s %-10s\n", headers[0], headers[2], headers[3]) // ID, Name, Team
			for _, row := range rowSet {
				fields := row.([]interface{})
				name := fields[2].(string)
				if playerName == "" || (playerName != "" && (containsIgnoreCase(name, playerName))) {
					playerID := int(fields[0].(float64))                              // Convert player ID to integer
					fmt.Printf("%-10d %-25v %-10v\n", playerID, fields[2], fields[3]) // Use %d for integer formatting
				}
			}
			return
		}
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
		body, err := io.ReadAll(resp.Body)
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

		// Show recent games if flag is set
		if showRecent {
			reqGames, err := http.NewRequest("GET", "https://stats.nba.com/stats/playergamelog", nil)
			if err != nil {
				fmt.Println("Error creating request for recent games:", err)
				return
			}
			qGames := reqGames.URL.Query()
			qGames.Add("PlayerID", playerID)
			qGames.Add("Season", "2023-24") // You may want to make this dynamic
			qGames.Add("SeasonType", "Regular Season")
			qGames.Add("LeagueID", "00")
			reqGames.URL.RawQuery = qGames.Encode()

			reqGames.Header.Add("Accept-Encoding", "identity")
			reqGames.Header.Add("Accept", "application/json, text/plain, */*")
			reqGames.Header.Add("x-nba-stats-origin", "stats")
			reqGames.Header.Add("x-nba-stats-token", "true")
			reqGames.Header.Add("Host", "stats.nba.com")
			reqGames.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0")
			reqGames.Header.Add("Connection", "keep-alive")
			reqGames.Header.Add("Referer", "https://stats.nba.com/")
			reqGames.Header.Add("Pragma", "no-cache")
			reqGames.Header.Add("Cache-Control", "no-cache")
			reqGames.Header.Add("Accept-Language", "en-US,en;q=0.5")

			respGames, err := client.Do(reqGames)
			if err != nil {
				fmt.Println("Error fetching recent games:", err)
				return
			}
			defer respGames.Body.Close()
			bodyGames, err := io.ReadAll(respGames.Body)
			if err != nil {
				fmt.Println("Error reading recent games body:", err)
				return
			}
			var gamesData map[string]interface{}
			err = json.Unmarshal(bodyGames, &gamesData)
			if err != nil {
				fmt.Println("Error unmarshalling recent games data:", err)
				return
			}
			resultSetsGames, ok := gamesData["resultSets"].([]interface{})
			if !ok || len(resultSetsGames) == 0 {
				fmt.Println("No recent games found")
				return
			}
			gamesSet := resultSetsGames[0].(map[string]interface{})
			headersGames := gamesSet["headers"].([]interface{})
			rowSetGames := gamesSet["rowSet"].([]interface{})
			fmt.Println("\nRecent Games:")
			for _, h := range headersGames {
				fmt.Printf("%-12s", h)
			}
			fmt.Println()
			for i, row := range rowSetGames {
				if i >= 5 { // Show only the 5 most recent games
					break
				}
				fields := row.([]interface{})
				for _, f := range fields {
					fmt.Printf("%-12v", f)
				}
				fmt.Println()
			}
		}

	},
}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func init() {
	rootCmd.AddCommand(playersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	playersCmd.Flags().StringVarP(&playerID, "playerId", "p", "2544", "Player ID to get info")
	playersCmd.Flags().StringVarP(&playerName, "name", "n", "", "Player name to search for")
	playersCmd.Flags().BoolVarP(&showRecent, "recent", "r", false, "Show recent games for the player")
}
