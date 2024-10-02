package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func sendMessage(code EventCode, body any) gin.H {
	data := gin.H{
		"success":   body,
		"eventcode": code,
	}
	return data
}

func makeQuestionBody(question Question) gin.H {
	return gin.H{
		"question":         question.question,
		"possible_anwsers": question.possibleAnwsers,
	}
}

func makeStatsBody(stats Quiz, rank int) gin.H {
	return gin.H{
		"gameId":           stats.gameId,
		"anwser":           stats.anwser,
		"totalCorrect":     stats.totalCorrect,
		"startTime":        stats.startTime.Format(time.RFC1123),
		"duration":         time.Duration((stats.endTime.Sub(stats.startTime)).Seconds()),
		"rank":             rank + 1,
		"totalNumberGames": len(handRecords),
	}
}

func makeSingleLeaderboardStat(stats Quiz, rank int) gin.H {
	return gin.H{
		"gameId":           stats.gameId,
		"totalCorrect":     stats.totalCorrect,
		"startTime":        stats.startTime.Format(time.RFC1123),
		"duration":         time.Duration((stats.endTime.Sub(stats.startTime)).Seconds()),
		"rank":             rank + 1,
		"totalNumberGames": len(handRecords),
	}
}
