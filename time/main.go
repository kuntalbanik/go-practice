package main

import (
	"fmt"
	"time"
)

func main() {

	// Formatted time print
	currentTime := time.Now()
	formattedTime := currentTime.Format("02-01-2006")
	fmt.Println(formattedTime)

	// Parsing time and print as per formatted time string
	format_str := "02/01/2006"
	date_str := "20/11/2025"
	formatted_time, _ := time.Parse(format_str, date_str)
	fmt.Println(formatted_time)

	// Add 1 more day
	newDay := currentTime.Add(48 * time.Hour)
	fmt.Println(newDay)

}
