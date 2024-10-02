package cmd

type ResponseQuestion struct {
	Success struct {
		Question        string   `json:"question"`
		PossibleAnwsers []string `json:"possible_anwsers"`
	} `json:"success"`
	Error EventCode `json:"eventcode"`
}

type ResponseGameStats struct {
	Eventcode int `json:"eventcode"`
	Success   struct {
		Anwser           []int  `json:"anwser"`
		Duration         int    `json:"duration"`
		GameID           int    `json:"gameId"`
		Rank             int    `json:"rank"`
		StartTime        string `json:"startTime"`
		TotalCorrect     int    `json:"totalCorrect"`
		TotalNumberGames int    `json:"totalNumberGames"`
	} `json:"success"`
	Error EventCode `json:"eventcode"`
}

type ResponseLeaderboard struct {
	Eventcode int `json:"eventcode"`
	Success   []struct {
		Duration         int    `json:"duration"`
		GameID           int    `json:"gameId"`
		Rank             int    `json:"rank"`
		StartTime        string `json:"startTime"`
		TotalCorrect     int    `json:"totalCorrect"`
		TotalNumberGames int    `json:"totalNumberGames"`
	} `json:"success"`
	Error EventCode `json:"eventcode"`
}
