package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

var getLeaderboardCmd = &cobra.Command{
	Use:   "leaderboard",
	Short: "Command to the leadeboard",
	Long:  `
	This command provides a leaderboard of the game.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		response := getEndpoint("leaderboard")
		var question ResponseLeaderboard
		errss := json.Unmarshal(response, &question)
		if errss != nil {
			log.Fatal(errss)
		}

		if question.Error == NOP {
			displayLeatherBoard(question)
		} else {
			displayEvent(question.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(getLeaderboardCmd)
}
