package dist

import (
	"fmt"
	"time"
)

func GetTimeOfDay(epochTime int64) string {
	// Convert epoch time to time.Time
	t := time.Unix(epochTime, 0)

	// Get hours, minutes, and seconds
	hour, min, sec := t.Clock()

	// Format the time of day
	timeOfDay := fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)

	return timeOfDay
}
