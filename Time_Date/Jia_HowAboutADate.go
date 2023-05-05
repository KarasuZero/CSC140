package main

import (
	"fmt"
	"time"
)

type FixedHoliday struct {
	name  string
	month time.Month
	day   int
	year  int
}

type MovingHoliday struct {
	name      string
	month     time.Month
	dayOfWeek time.Weekday
	instance  int
	year      int
}

func newFixedHoliday(name string, month time.Month, day, year int) *FixedHoliday {
	return &FixedHoliday{name: name, month: month, day: day, year: year}
}

func newMovingHoliday(name string, month time.Month, dayOfWeek string, instance, year int) *MovingHoliday {
	return &MovingHoliday{
		name:      name,
		month:     month,
		dayOfWeek: parseWeekday(dayOfWeek),
		instance:  instance,
		year:      year,
	}
}

func parseWeekday(weekday string) time.Weekday {
	// turns string into time of the week
	switch weekday {
	case "sunday":
		return time.Sunday
	case "monday":
		return time.Monday
	case "tuesday":
		return time.Tuesday
	case "wednesday":
		return time.Wednesday
	case "thursday":
		return time.Thursday
	case "friday":
		return time.Friday
	case "saturday":
		return time.Saturday
	default:
		panic("invalid weekday")
	}
}

func (h *FixedHoliday) printHoliday() {
	t := time.Date(h.year, h.month, h.day, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%s is on %s, %s %02d in %d\n", h.name, t.Weekday(), t.Month(), t.Day(), h.year)
}

func (h *MovingHoliday) printHoliday() {
	// first day of the target month
	t := time.Date(h.year, h.month, 1, 0, 0, 0, 0, time.UTC)

	// formula I found on stack overflow and modify it to calculates how many days from the beginning of the month to target date
	daysToAdd := (int(h.dayOfWeek)-int(t.Weekday())+7)%7 + (h.instance-1)*7

	// if it's a week, land on the same weekday
	if daysToAdd == 0 {
		daysToAdd = 7
	}

	// add the days to the first day of the target month
	t = t.AddDate(0, 0, daysToAdd)
	fmt.Printf("%s is on %s, %s %02d in %d\n", h.name, t.Weekday(), t.Month(), t.Day(), h.year)
}

func main() {
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scanln(&year)

	newYearsDay := newFixedHoliday("New Year's Day", time.January, 1, year)
	independenceDay := newFixedHoliday("Independence Day", time.July, 4, year)
	halloween := newFixedHoliday("Halloween", time.October, 31, year)

	mothersDay := newMovingHoliday("Mother's Day", time.May, "sunday", 2, year)
	fathersDay := newMovingHoliday("Father's Day", time.June, "sunday", 3, year)
	thanksgiving := newMovingHoliday("Thanksgiving", time.November, "thursday", 4, year)

	newYearsDay.printHoliday()
	independenceDay.printHoliday()
	halloween.printHoliday()
	mothersDay.printHoliday()
	fathersDay.printHoliday()
	thanksgiving.printHoliday()
}
