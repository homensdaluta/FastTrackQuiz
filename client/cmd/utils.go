package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type EventCode int

const (
	NOP                      = 0
	GAME_EXPIRED   EventCode = 11
	GAME_NOT_START EventCode = 22
	GAME_FINISH    EventCode = 33
	ONGOING_GAME   EventCode = 44
	NOT_FOUND      EventCode = 55
	NOT_OPTION     EventCode = 66
	INTERNAL_ERROR EventCode = 66
)

func getEndpoint(endpoint string) []byte {
	response, err := http.Get("http://localhost:8080/" + endpoint)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	/*
		fmt.Println(response)
		fmt.Println(string(responseData))
		var message ResponseQuestion
		_ = json.Unmarshal(responseData, &message)
		fmt.Println(message)*/

	return responseData

}

func displayEvent(errorCode EventCode) {
	switch errorCode {
	case GAME_EXPIRED:
		{
			fmt.Println("**************")
			fmt.Println("Your current game as expired. Your remaining questions have automatically been mark as wrong")
			fmt.Println("")
			fmt.Println("You can start a new game using the command 'FastTrackuser start'")
			fmt.Println("**************")
		}
	case GAME_NOT_START:
		{
			fmt.Println("**************")
			fmt.Println("You haven't started a game")
			fmt.Println("")
			fmt.Println("You can start a new game using the command 'FastTrackuser start'")
			fmt.Println("**************")
		}
	case GAME_FINISH:
		{
			fmt.Println("**************")
			fmt.Println("There is no active game")
			fmt.Println("")
			fmt.Println("You can start a new game using the command 'FastTrackuser start'")
			fmt.Println("**************")
		}
	case ONGOING_GAME:
		{
			fmt.Println("**************")
			fmt.Println("There is an active game going on")
			fmt.Println("")
			fmt.Println("You if you don't remember the question ples use the command 'FastTrackuser question'")
			fmt.Println("**************")
		}
	case NOT_FOUND:
		{
			fmt.Println("**************")
			fmt.Println("Couldn't find stats for the game you were looking for")
			fmt.Println("**************")
		}
	case NOT_OPTION:
		{
			fmt.Println("**************")
			fmt.Println("That is not a valid option for the question")
			fmt.Println("")
			fmt.Println("Please use one of the option provided")
			fmt.Println("**************")
		}
	default:
		{
			fmt.Println("**************")
			fmt.Println("Unknown Error")
			fmt.Println("**************")
		}
	}
}

func displayQuestion(question string, possibleAnwsers []string) {
	fmt.Println("**************")
	fmt.Println("")
	fmt.Println(question)
	fmt.Println("")
	for index, anwsers := range possibleAnwsers {
		fmt.Println("Option ", (index + 1), ": ", anwsers)
	}
	fmt.Println("")
	fmt.Println("Use the command 'FastTrackuser anwser (number of the option)' to answer the prompted question")
	fmt.Println("Example: 'FastTrackuser anwser 1'")
	fmt.Println("**************")
}

func displayCancel() {
	fmt.Println("**************")
	fmt.Println("")
	fmt.Println("Game was succefully aborted")
	fmt.Println("Your remaining questions have automatically been mark as wrong")
	fmt.Println("You can start a new game using the command 'FastTrackuser start'")
	fmt.Println("")
	fmt.Println("**************")
}

func displayLeatherBoard(leatherboard ResponseLeaderboard) {
	fmt.Println("**************")
	for _, test := range leatherboard.Success {
		fmt.Println("Game id:", test.GameID)
		fmt.Println("Game placed: ", test.Rank, " out of ", test.TotalNumberGames)
		fmt.Println("Total number of correct answers: ", test.TotalCorrect)
		fmt.Println("Game started on : ", test.StartTime)
		fmt.Println("Game had the duration of : ", test.Duration, " seconds")
		fmt.Println("")
		fmt.Println("----------------------------")
	}
	fmt.Println("**************")
}

func displayStats(stats ResponseGameStats) {
	fmt.Println("**************")
	fmt.Println("Stats for game id:", stats.Success.GameID)
	fmt.Println("")
	fmt.Println("All answers (in order): ", stats.Success.Anwser)
	fmt.Println("Total number of correct answers: ", stats.Success.TotalCorrect)
	fmt.Println("Game started on : ", stats.Success.StartTime)
	fmt.Println("Game had the duration of : ", stats.Success.Duration, " seconds")
	fmt.Println("Game placed: ", stats.Success.Rank, " out of ", stats.Success.TotalNumberGames)
	fmt.Println("You were top ", ((stats.Success.Rank * 100) / stats.Success.TotalNumberGames), "% of the world")
	fmt.Println("")
	fmt.Println("**************")
}
