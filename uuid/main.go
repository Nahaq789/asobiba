package main

import (
	"fmt"

	"github.com/google/uuid"
)

const bufferSize = 100
const loopCount = 1000000
const tryCount = 100

func main() {
	for a := 0; a < tryCount; a++ {
		fmt.Printf("%v回目！\n", a+1)
		c := make(chan bool, bufferSize)

		for i := 0; i < bufferSize; i++ {
			go loop(c)
		}
		found := false
		for i := 0; i < bufferSize; i++ {
			if <-c {
				found = true
			}
		}

		if found {
			println("重複発見！！！！！！！！！！！！！！！！！！！！！！！！！！！！")
		} else {
			println("重複なし")
		}
	}
}

func loop(c chan bool) {
	seen := make(map[uuid.UUID]bool, loopCount)
	for i := 0; i < loopCount; i++ {
		id := uuid.New()
		if seen[id] {
			c <- true
			return
		}
		seen[id] = true
	}
	c <- false
}
