package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time package in Golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// syntax for date, time, day
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// syntax for date creation
	createdDate := time.Date(2024, time.August, 12, 17, 8, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// syntax for creating in other OS: GOOS="windows" go build

}
