package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startEndpoint() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/start", startGame)
	router.GET("/cancel", cancel)
	router.GET("/anwser/:option", anwser)
	router.GET("/question", question) //or help
	router.GET("/leaderboard", leaderboard)
	router.GET("/stats/:id", stats)
	router.Run("localhost:8080")
}

func startGame(c *gin.Context) {
	c.JSON(http.StatusOK, startQuiz())
}
func anwser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("option"))
	if err != nil {
		c.IndentedJSON(http.StatusOK, sendMessage(NOT_OPTION, gin.H{}))
	} else {
		c.IndentedJSON(http.StatusOK, anwserQuiz(id))
	}
}

func question(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getQuestion())
}

func cancel(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cancelCurrentGame())
}

func leaderboard(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getLeaderboard())
}

func stats(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusOK, sendMessage(NOT_FOUND, gin.H{}))
	} else {
		c.IndentedJSON(http.StatusOK, getGameStats(id))
	}
}
