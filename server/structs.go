package main

import "time"

type Quiz struct {
	gameId       int
	anwser       []int
	startTime    time.Time
	endTime      time.Time
	totalCorrect int
}

type Question struct {
	question        string
	possibleAnwsers []string
	answerIndex     int
}

type EventCode int

const (
	NOP                      = 0
	GAME_EXPIRED   EventCode = 11
	GAME_NOT_START EventCode = 22
	GAME_FINISH    EventCode = 33
	ONGOING_GAME   EventCode = 44
	NOT_FOUND      EventCode = 55
	NOT_OPTION     EventCode = 66
	INTERNAL_ERROR EventCode = 77
)
