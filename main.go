package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TwiN/go-color"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message = scanner.Text()
		break
	}

	fmt.Print("In how many minutes: ")
	fmt.Scan(&setTimeDurationToAlert)

	println("")
	println(color.Green+"Wait... you will be notified in", setTimeDurationToAlert, "minutes")
	time.Sleep(time.Duration(setTimeDurationToAlert) * time.Second)

	notify.Alert(Time(), "ALERT!", "Reminder: "+message, "path/to/icon.png")
}
