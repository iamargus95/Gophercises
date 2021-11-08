package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, birds := range birdsPerDay {
		count += birds
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	day := week*7 - 1
	count := 0
	for i := day; i >= day-6; i-- {
		count += birdsPerDay[i]
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, birds := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = birds + 1
		}
	}
	return birdsPerDay
}
