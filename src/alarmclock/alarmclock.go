package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	currentTime := time.Now()
	currentInFormat := time.Time.Format(currentTime, "15.04.05:")
	fmt.Println("The time is", currentInFormat, text)
	time.Sleep(delay)
	Remind(text, delay)
}

func main() {
	//Remind("hej", 2*time.Second)

	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	Remind("Time to sleep", time.Minute)
}
