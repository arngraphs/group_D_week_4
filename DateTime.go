package main

import (
	"fmt"
	"time"
)

func timenew() {

	currentTime := time.Now()

	fmt.Println("Current Time by Charith: ", currentTime)

	fmt.Println("Formatted Time by Charith: ", currentTime.Format("2006-01-02 15:04:05"))
}
