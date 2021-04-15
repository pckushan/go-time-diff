package pkg

import (
	"math"
	"time"
)

const (
	WeekDays      = 7
	FortNightDays = 14
	QuarterMonths = 4
)

func Diff(a, b time.Time) (years, months, days, weeks, fortNights, quarters int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}

	y1, _, _ := a.Date()
	y2, _, _ := b.Date()

	days = 0
	weeks = 0
	fortNights = 0
	quarters = 0
	months = 0
	years = y2 - y1
	// needed to consider current year too
	years++

	startingMonth := a.Month()
	// loop day by day
	for day := a; !day.After(b); day = day.AddDate(0, 0, 1) {
		currentMonth := day.Month()

		if startingMonth != currentMonth {
			startingMonth = currentMonth
			months++
		}

		days++
	}

	weeks = int(math.Ceil(float64(days) / WeekDays))
	fortNights = int(math.Ceil(float64(days) / FortNightDays))
	quarters = int(math.Ceil(float64(months) / QuarterMonths))
	return years, months, days, weeks, fortNights, quarters
}
