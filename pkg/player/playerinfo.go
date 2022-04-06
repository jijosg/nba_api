package player

type PlayerInfo struct {
	Resource   string `json:"resource"`
	Parameters []struct {
		PlayerID int    `json:"PlayerID,omitempty"`
		LeagueID string `json:"LeagueID,omitempty"`
	} `json:"parameters"`
	ResultSets []struct {
		Name    string          `json:"name"`
		Headers []string        `json:"headers"`
		RowSet  [][]interface{} `json:"rowSet"`
	} `json:"resultSets"`
}
