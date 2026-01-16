package main

import (
	"fmt"
	"math/rand"
)

const LOSE_PROBABILITY = 0.3513513513513514
const RETRY_COUNT = 50

func main() {
	var results []Result

	for i := 0; i < 10000; i++ {
		result := simulate()
		results = append(results, result)
	}

	r := NewResults(results)

	fmt.Printf("win count: %d\n", r.calcTotalWin())
	fmt.Printf("lose count: %d\n", r.calcTotalLose())
	fmt.Printf("max continuous lose: %d\n", r.calcAvgContinuLose())
}

func simulate() Result {
	loseCount := 0
	winCount := 0
	continuousLose := 0
	maxContinuousLose := 0

	for i := 0; i < RETRY_COUNT; i++ {
		randomValue := rand.Float64()
		if randomValue < LOSE_PROBABILITY {
			loseCount++
			continuousLose++

			if continuousLose > maxContinuousLose {
				maxContinuousLose = continuousLose
			}
		} else {
			winCount++
			continuousLose = 0
		}
	}

	result := NewResult(winCount, loseCount, maxContinuousLose)
	return result
}
