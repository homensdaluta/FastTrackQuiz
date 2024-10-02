package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var currentGame Quiz
var gameActive = false

func startQuiz() gin.H {
	if gameActive {
		return sendMessage(ONGOING_GAME, gin.H{})
	}

	currentGame = Quiz{
		gameId:    len(handRecords) + 1,
		anwser:    []int{},
		startTime: time.Now(),
	}
	gameActive = true
	return sendMessage(NOP, makeQuestionBody(gameQuestion[0]))
}

func anwserQuiz(answer int) gin.H {
	if !gameActive {
		return sendMessage(GAME_NOT_START, gin.H{})
	}

	if !isExpire() {
		cancelCurrentGame()
		return sendMessage(GAME_EXPIRED, gin.H{})
	}

	if !(answer > 0 && answer < 5) {
		return sendMessage(NOT_OPTION, gin.H{})
	}  

	answerQuestion(answer)

	if len(gameQuestion) > len(currentGame.anwser) {
		return sendMessage(NOP, makeQuestionBody(gameQuestion[len(currentGame.anwser)]))
	} else {
		submitGame(currentGame)
		id := currentGame.gameId
		currentGame = Quiz{}
		gameActive = false
		return getGameStats(id)
	}
}

func cancelCurrentGame() gin.H {
	if !gameActive {
		return sendMessage(GAME_FINISH, gin.H{})
	}
	currentGame.endTime = time.Now()
	n := len(currentGame.anwser)

	for n < len(gameQuestion) {
		currentGame.anwser = append(currentGame.anwser, -1)
		n = len(currentGame.anwser)
	}
	submitGame(currentGame)
	currentGame = Quiz{}
	gameActive = false
	return sendMessage(NOP, gin.H{})
}

func getQuestion() gin.H {
	if !gameActive {
		return sendMessage(GAME_NOT_START, gin.H{})
	}
	if !isExpire() {
		cancelCurrentGame()
		return sendMessage(GAME_EXPIRED, gin.H{})
	}

	return sendMessage(NOP, makeQuestionBody(gameQuestion[len(currentGame.anwser)]))
}

func isExpire() bool {
	fmt.Println(currentGame.startTime)
	return time.Now().Before(currentGame.startTime.Add(time.Duration(time.Duration(EXPIRETIME * 1000000000))))
}
