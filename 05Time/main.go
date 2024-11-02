package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study of go")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// always use 02 for date
	// 01 for month
	// 2006 for year
	// 15:04:05 for time stamp
	// Monday for day
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.October, 22, 16, 15, 24, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))
}
