package main

import (
	"fmt"
	"time"
)

func timenew() {
	// Get current time
	currentTime := time.Now()

	// Print current time in default format
	fmt.Println("Current Time: ", currentTime)

	// You can also format the time output if needed
	fmt.Println("Formatted Time: ", currentTime.Format("2006-01-02 15:04:05"))
}
