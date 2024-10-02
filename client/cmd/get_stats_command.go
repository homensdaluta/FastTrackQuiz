package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

var getGameStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Command to get stats of the game",
	Long: `
	This command provides all the stats from a game. 
	Just need to specify the id of the game after the command
	`,
	Run: func(cmd *cobra.Command, args []string) {
		response := getEndpoint("stats/" + args[0])
		var stats ResponseGameStats
		errss := json.Unmarshal(response, &stats)
		if errss != nil {
			log.Fatal(errss)
		}
		if stats.Error == 0 {
			displayStats(stats)
		} else {
			displayEvent(stats.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(getGameStatsCmd)
}
