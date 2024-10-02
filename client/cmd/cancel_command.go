package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Command to cancel the quiz",
	Long:  `
	This command cancel a game if it is already a game is already running.
	If you cancel, all the remaining awnser will be set to a wrong anwser.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		response := getEndpoint("cancel")
		var question ResponseQuestion
		errss := json.Unmarshal(response, &question)
		if errss != nil {
			log.Fatal(errss)
		}

		if question.Error == 0 {
			displayCancel()
		} else {
			displayEvent(question.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)
}
