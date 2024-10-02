package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var anwserCmd = &cobra.Command{
	Use:   "anwser",
	Short: "Command to anwser the question",
	Long: `
	This command anwsers the question you are on.
	You can do this by giving the number of the option you want to anwser after the command
	Make sure that the game is running and the anwser is between the option shown.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		response := getEndpoint("anwser/" + args[0])
		var question ResponseQuestion
		errss := json.Unmarshal(response, &question)
		if errss != nil {
			log.Fatal(errss)
		}

		if question.Error == NOP {
			fmt.Println("**************")
			fmt.Println("Option ", args[0], " was succesfully submitted")
			fmt.Println("")
			fmt.Println("Next question:")
			displayQuestion(question.Success.Question, question.Success.PossibleAnwsers)
		} else if question.Error == GAME_FINISH {
			var stats ResponseGameStats
			errss := json.Unmarshal(response, &stats)
			if errss != nil {
				log.Fatal(errss)
			}
			fmt.Println("**************")
			fmt.Println("Game finish. Thank you for playing")
			fmt.Println("Below you can se the results of your game")
			fmt.Println("**************")
			displayStats(stats)
		} else {
			displayEvent(question.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(anwserCmd)
}
