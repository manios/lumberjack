package main

import (
	"fmt"
	"strconv"
	"time"
)

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Truncate(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

func main() {
	chicago, err := time.LoadLocation("Europe/Athens")
	if err != nil {
		fmt.Println(err)
		return
	}

	lastLogLine := time.Date(2022, 8, 23, 23, 59, 0, 0, chicago)
	now := time.Date(2022, 8, 24, 0, 1, 0, 0, chicago)

	fmt.Println(Truncate(now))
	fmt.Printf("Last log line is at: %s\n", lastLogLine)
	fmt.Printf("Now() is  at: %s\n", now)
	fmt.Printf("Is now start of day after lastLogLine start of day ? %s\n", strconv.FormatBool(StartOfDay(now).After(StartOfDay(lastLogLine))))

}
