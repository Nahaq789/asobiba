package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	x := make(chan int)
	go func() {
		fmt.Println("ゴルーチン実行")

		for i := range 5 {
			fmt.Printf("%d秒経過\n", i+1)
			time.Sleep(1 * time.Second)
		}
		x <- 100
	}()

	for i := range 10 {
		fmt.Printf("%d秒\n", i+1)
		time.Sleep(1 * time.Second)
	}

	// チャネルに値が入るまで待機
	value := <-x
	fmt.Println(value)
	fmt.Println("finish")
}
