package main

import (
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

var gameQuestion []Question
var handRecords []Quiz

func startDB() {
	gameQuestion = []Question{QUESTION1, QUESTION2, QUESTION3, QUESTION4, QUESTION5}
	handRecords = []Quiz{}
}

func submitGame(game Quiz) {
	game.endTime = time.Now()
	handRecords = append(handRecords, game)
	if len(handRecords) > 1 {
		sort.Slice(handRecords, func(i, j int) bool {
			if handRecords[i].totalCorrect == handRecords[j].totalCorrect {
				return handRecords[i].gameId < handRecords[j].gameId
			} else {
				return handRecords[i].totalCorrect > handRecords[j].totalCorrect
			}
		})
	}
}

func answerQuestion(answer int) {
	currentQuestionIndex := len(currentGame.anwser)
	if answer == gameQuestion[currentQuestionIndex].answerIndex {
		currentGame.totalCorrect++
	}
	currentGame.anwser = append(currentGame.anwser, answer)
}

func getLeaderboard() gin.H {
	jsonRecord := []gin.H{}
	for index, game := range handRecords {
		jsonRecord = append(jsonRecord, makeSingleLeaderboardStat(game, index))
	}
	return sendMessage(NOP, jsonRecord)
}

func getGameStats(id int) gin.H {
	for index, value := range handRecords {
		if value.gameId == id {
			return sendMessage(GAME_FINISH, makeStatsBody(value, index))
		}
	}
	return sendMessage(NOT_FOUND, gin.H{})
}
