package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

var getQuestionCmd = &cobra.Command{
	Use:   "question",
	Short: "Command to get question",
	Long: `
	This command gives the question that you are currently on
	`,
	Run: func(cmd *cobra.Command, args []string) {
		
		response := getEndpoint("question")
		var question ResponseQuestion
		errss := json.Unmarshal(response, &question)
		if errss != nil {
			log.Fatal(errss)
		}

		if question.Error == 0 {
			displayQuestion(question.Success.Question, question.Success.PossibleAnwsers)
		} else {
			displayEvent(question.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(getQuestionCmd)
}
