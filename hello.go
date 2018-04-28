package main

import (
	"fmt"
	"time"
)

var now = time.Now()

func TodayOfMonth() int {
	return now.Day()
}

func main() {
	fmt.Println(TodayOfMonth())
}
