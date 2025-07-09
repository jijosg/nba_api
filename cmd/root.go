/*
Copyright © 2022 Jijo Sunny

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
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jijosg/nba_api/pkg/teams"
)

// Version information
var (
	version = "1.0.0"
	commit  = "unknown"
	date    = "unknown"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nba_api",
	Short: "NBA API CLI tool for retrieving NBA related data",
	Long: `A command line interface for the NBA API that allows users to
fetch information about NBA teams, players, and game scores. 

You can use this tool to search for team information, look up player stats,
view today's scores, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if version flag is provided
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Printf("nba_api version %s (commit: %s, built on: %s)\n", 
				version, commit, date)
			return
		}
		
		// If no arguments are provided and team flag is default, show help
		if len(args) == 0 && viper.GetString("team") == "LAL" && 
		   !cmd.Flags().Changed("team") {
			fmt.Println("NBA API CLI - A tool for accessing NBA statistics")
			fmt.Println("\nAvailable Commands:")
			for _, command := range cmd.Commands() {
				fmt.Printf("  %-12s %s\n", command.Name(), command.Short)
			}
			fmt.Println("\nUse \"nba_api [command] --help\" for more information about a command.")
			return
		}
		
		team := viper.GetString("team")
		foundTeam, found := findTeam(team)
		
		if found {
			fmt.Printf("Team Information for %s:\n", foundTeam.Abbreviation)
			fmt.Printf("  Full Name: %s\n", foundTeam.FullName)
			fmt.Printf("  City: %s\n", foundTeam.City)
			fmt.Printf("  State: %s\n", foundTeam.State)
			fmt.Printf("  Founded: %d\n", foundTeam.YearFounded)
			
			if len(foundTeam.ChampionshipYear) > 0 {
				fmt.Printf("  Championships: %d (%v)\n", len(foundTeam.ChampionshipYear), foundTeam.ChampionshipYear)
			} else {
				fmt.Printf("  Championships: None\n")
			}
		} else {
			fmt.Printf("Team '%s' not found. Use 'nba_api list' to see all available teams.\n", team)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// findTeamByAbbreviation searches for a team by its abbreviation
// and returns the team and whether it was found
func findTeamByAbbreviation(abbr string) (teams.Teams, bool) {
	for _, team := range teams.NBATeams {
		if team.Abbreviation == abbr {
			return team, true
		}
	}
	return teams.Teams{}, false
}

// findTeam searches for a team by abbreviation, nickname, or full name
// and returns the team and whether it was found
func findTeam(query string) (teams.Teams, bool) {
	// First try exact match on abbreviation
	team, found := findTeamByAbbreviation(query)
	if found {
		return team, true
	}
	
	// Then try case-insensitive match on nickname or full name
	query = strings.ToLower(query)
	for _, team := range teams.NBATeams {
		if strings.ToLower(team.Nickname) == query || 
		   strings.ToLower(team.FullName) == query ||
		   strings.Contains(strings.ToLower(team.FullName), query) {
			return team, true
		}
	}
	
	return teams.Teams{}, false
}

// getTeamAbbreviations returns all team abbreviations for auto-completion
func getTeamAbbreviations() []string {
	abbreviations := make([]string, len(teams.NBATeams))
	for i, team := range teams.NBATeams {
		abbreviations[i] = team.Abbreviation
	}
	return abbreviations
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nba_api.yaml)")
	
	// Add team flag for quick team lookups
	rootCmd.Flags().StringP("team", "n", "LAL", "NBA team abbreviation, nickname, or name (e.g., LAL, Lakers)")
	
	// Setup auto-completion for team names
	rootCmd.RegisterFlagCompletionFunc("team", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getTeamAbbreviations(), cobra.ShellCompDirectiveNoFileComp
	})
	
	// Add version flag
	rootCmd.Flags().BoolP("version", "v", false, "Display version information")
	
	// Bind flags to viper
	viper.BindPFlag("team", rootCmd.Flags().Lookup("team"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".nba_api" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nba_api")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
