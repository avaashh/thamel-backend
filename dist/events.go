package dist

import (
	"thamel/types"
	"time"
)

// WeeklyEvents -> returns the events happening this week, scraped from https://events.grinnell.edu
func TodayWeeklyEvents() types.Response25 {
	currentTime := time.Now()

	// "2006-01-02" is the Go reference date
	formattedDate := currentTime.Format("20060102")
	return LiveEvents(formattedDate)
}

// WeeklyEvents -> returns the events happening today, scraped from https://events.grinnell.edu
func TodayDailyEvents() types.Response25 {
	currentTime := time.Now()

	// "2006-01-02" is the Go reference date
	formattedDate := currentTime.Format("20060102")
	return LiveEvents(formattedDate)
}
