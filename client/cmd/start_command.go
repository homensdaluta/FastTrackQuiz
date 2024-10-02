package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var startGameCmd = &cobra.Command{
	Use:   "start",
	Short: "Command to start the quiz",
	Long: `
	This command start the game.
	Afterwards you can use the "anwser" command to anwser the prompted question
	`,
	Run: func(cmd *cobra.Command, args []string) {
		response := getEndpoint("start")
		var question ResponseQuestion
		errss := json.Unmarshal(response, &question)
		if errss != nil {
			log.Fatal(errss)
		}

		if question.Error == 0 {
			fmt.Println("**************")
			fmt.Println("Welcome to the quiz")
			fmt.Println("")
			displayQuestion(question.Success.Question, question.Success.PossibleAnwsers)
		} else {
			displayEvent(question.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(startGameCmd)
}
