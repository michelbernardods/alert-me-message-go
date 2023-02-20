package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/martinlindhe/notify"
)

func Time() string {
	now := time.Now()
	Date := now.AddDate(0, 0, -7).Format("2006-01-02")
	Time := time.Now().Format("02/Jan/2006")
	var dateTime []string
	dateTime = append(dateTime, Date, Time)
	dateTimeFormat := strings.Join(dateTime, " | ")
	return dateTimeFormat
}

func main() {
	var message string
	var setTimeDurationToAlert int

	fmt.Print("Notification message: ")
	fmt.Scan(&message)

	fmt.Print("In how many minutes:: ")
	fmt.Scan(&setTimeDurationToAlert)

	time.Sleep(time.Duration(setTimeDurationToAlert) * time.Second)
	notify.Alert(Time(), "ALERT!", "Reminder: "+message, "path/to/icon.png")
}
