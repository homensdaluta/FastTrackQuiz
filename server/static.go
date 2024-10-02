package main

import "time"

const EXPIRETIME = 60

var QUESTION1 = Question{
	question:        "What color is the sky?",
	possibleAnwsers: []string{"Black", "Red", "Blue", "I am colorblind :("},
	answerIndex:     3,
}

var QUESTION2 = Question{
	question:        "Which animal has 4 legs?",
	possibleAnwsers: []string{"Worm", "Horse", "Human", "Pigeon"},
	answerIndex:     3,
}

var QUESTION3 = Question{
	question:        "How many legs does a spider have?",
	possibleAnwsers: []string{"4", "8", "1", "14"},
	answerIndex:     2,
}

var QUESTION4 = Question{
	question:        "If you freeze water, what do you get?",
	possibleAnwsers: []string{"Ice", "Wood", "Gold", "Water"},
	answerIndex:     1,
}

var QUESTION5 = Question{
	question:        "What is the largest mammal in the world?",
	possibleAnwsers: []string{"Fish", "Panda", "Dog", "Whale"},
	answerIndex:     4,
}

var GAME1 = Quiz{
	gameId:    1,
	anwser:    []int{1, 1, 1, 1, 2},
	startTime: time.Date(2024, 9, 20, 13, 47, 30, 651387237, time.UTC),
	endTime:   time.Date(2024, 9, 20, 13, 48, 58, 651387237, time.UTC),
}

var GAME2 = Quiz{
	gameId:    2,
	anwser:    []int{2, 2, 2, 2, 2},
	startTime: time.Date(2024, 9, 20, 13, 47, 30, 651387237, time.UTC),
	endTime:   time.Date(2024, 9, 20, 13, 48, 58, 651387237, time.UTC),
}

var GAME3 = Quiz{
	gameId:    3,
	anwser:    []int{3, 3, 3, 3, 3},
	startTime: time.Date(2024, 9, 20, 13, 47, 30, 651387237, time.UTC),
	endTime:   time.Date(2024, 9, 20, 13, 48, 58, 651387237, time.UTC),
}

var GAME4 = Quiz{
	gameId:    4,
	anwser:    []int{4, 4, 4, 4, 4},
	startTime: time.Date(2024, 9, 20, 13, 47, 30, 651387237, time.UTC),
	endTime:   time.Date(2024, 9, 20, 13, 48, 58, 651387237, time.UTC),
}
