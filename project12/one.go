package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time: ", currentTime)

	// formatted := currentTime.Format("02-01-2006 , Monday , 15:04:05 ")    // For 24 hr format
	formatted := currentTime.Format("02-01-2006 , Monday , 3:04 PM") // For AM/PM 12hr format
	fmt.Println("Formated Time: ", formatted)

	// To convert any date string into time format
	layoutStr := "2006-01-02" // dateStr is converted as the layout
	dateStr := "2002-03-08"
	formattedTime, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Formatted Time Is ", formattedTime)

	// Add one more day in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("New Date IS ", new_date)
	formattedNewDate := new_date.Format("02/01/2006 , Monday")
	fmt.Println("Formatted New Date ", formattedNewDate)
}
