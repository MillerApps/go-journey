package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/_2/2006 15:04:05", date) // _ means to ignore (error in this case). The _2 the is used for preceding 0
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January _2, 2006 15:04:05", date)
	return t.Before(time.Now()) // Check to see if the time parsed is before the current time
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January _2, 2006 15:04:05", date)
	// Check if the hour is between 12 (inclusive) and 18 (exclusive).
	if int(t.Hour()) < 18 && int(t.Hour()) >= 12 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	t, _ := time.Parse("1/_2/2006 15:04:05", date)
	day := days[t.Weekday()]
	month := t.Month()
	newString := fmt.Sprintf("You have an appointment on %s, %s	 %v, %v, at %v:%v.", day, month, t.Day(), t.Year(), t.Hour(), t.Minute())

	return newString
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := int(time.Now().Year()) //Currnet year
	// Create a new time.Time object for this year's anniversary date,
	// specifically set to September 15 at midnight (00:00:00) and in UTC.
	t := time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
	return t
}
