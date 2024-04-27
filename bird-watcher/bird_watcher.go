package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekStart := (week - 1) * 7
	weekEnd := week * 7
	sum := 0
	for i := weekStart; i < weekEnd && i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
	// or return TotalBirdCount(birdsPerDay[weekStart:weeekEnd])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++ // Increment count of birds for every second day from the first one; adds the hidden bird
	}
	return birdsPerDay
}
