package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go callParallelActions(channel)

	fmt.Println(<-channel)
}

func displayWordWithAForIterator(word string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}

func callParallelActions(channel chan string) {
	go displayWordWithAForIterator("Go Routine")
	displayWordWithAForIterator("Just a normal Function")

	channel <- "end"
}
