package teams

import "time"

type ScoreboardResponse struct {
	Meta struct {
		Version int    `json:"version"`
		Request string `json:"request"`
		Time    string `json:"time"`
		Code    int    `json:"code"`
	} `json:"meta"`
	Scoreboard struct {
		GameDate   string `json:"gameDate"`
		LeagueID   string `json:"leagueId"`
		LeagueName string `json:"leagueName"`
		Games      []struct {
			GameID            string    `json:"gameId"`
			GameCode          string    `json:"gameCode"`
			GameStatus        int       `json:"gameStatus"`
			GameStatusText    string    `json:"gameStatusText"`
			Period            int       `json:"period"`
			GameClock         string    `json:"gameClock"`
			GameTimeUTC       time.Time `json:"gameTimeUTC"`
			GameEt            time.Time `json:"gameEt"`
			RegulationPeriods int       `json:"regulationPeriods"`
			IfNecessary       bool      `json:"ifNecessary"`
			SeriesGameNumber  string    `json:"seriesGameNumber"`
			SeriesText        string    `json:"seriesText"`
			HomeTeam          struct {
				TeamID            int         `json:"teamId"`
				TeamName          string      `json:"teamName"`
				TeamCity          string      `json:"teamCity"`
				TeamTricode       string      `json:"teamTricode"`
				Wins              int         `json:"wins"`
				Losses            int         `json:"losses"`
				Score             int         `json:"score"`
				Seed              interface{} `json:"seed"`
				InBonus           interface{} `json:"inBonus"`
				TimeoutsRemaining int         `json:"timeoutsRemaining"`
				Periods           []struct {
					Period     int    `json:"period"`
					PeriodType string `json:"periodType"`
					Score      int    `json:"score"`
				} `json:"periods"`
			} `json:"homeTeam"`
			AwayTeam struct {
				TeamID            int         `json:"teamId"`
				TeamName          string      `json:"teamName"`
				TeamCity          string      `json:"teamCity"`
				TeamTricode       string      `json:"teamTricode"`
				Wins              int         `json:"wins"`
				Losses            int         `json:"losses"`
				Score             int         `json:"score"`
				Seed              interface{} `json:"seed"`
				InBonus           interface{} `json:"inBonus"`
				TimeoutsRemaining int         `json:"timeoutsRemaining"`
				Periods           []struct {
					Period     int    `json:"period"`
					PeriodType string `json:"periodType"`
					Score      int    `json:"score"`
				} `json:"periods"`
			} `json:"awayTeam"`
			GameLeaders struct {
				HomeLeaders struct {
					PersonID    int         `json:"personId"`
					Name        string      `json:"name"`
					JerseyNum   string      `json:"jerseyNum"`
					Position    string      `json:"position"`
					TeamTricode string      `json:"teamTricode"`
					PlayerSlug  interface{} `json:"playerSlug"`
					Points      int         `json:"points"`
					Rebounds    int         `json:"rebounds"`
					Assists     int         `json:"assists"`
				} `json:"homeLeaders"`
				AwayLeaders struct {
					PersonID    int         `json:"personId"`
					Name        string      `json:"name"`
					JerseyNum   string      `json:"jerseyNum"`
					Position    string      `json:"position"`
					TeamTricode string      `json:"teamTricode"`
					PlayerSlug  interface{} `json:"playerSlug"`
					Points      int         `json:"points"`
					Rebounds    int         `json:"rebounds"`
					Assists     int         `json:"assists"`
				} `json:"awayLeaders"`
			} `json:"gameLeaders"`
			PbOdds struct {
				Team      interface{} `json:"team"`
				Odds      float64     `json:"odds"`
				Suspended int         `json:"suspended"`
			} `json:"pbOdds"`
		} `json:"games"`
	} `json:"scoreboard"`
}
