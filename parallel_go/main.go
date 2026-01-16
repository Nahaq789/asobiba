package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	taskNames := []string{"task1", "task2", "task3", "task4", "task5"}

	for _, task := range taskNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			heavyTask(name)
		}(task)
	}

	wg.Wait()
	fmt.Println("all task complete")
}

func heavyTask(taskName string) {
	fmt.Printf("start task: %s \n", taskName)
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Printf("complete task: %s \n", taskName)
}
