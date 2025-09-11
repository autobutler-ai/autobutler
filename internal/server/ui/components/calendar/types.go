package calendar

import "time"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type WeekMode int

const (
	WeekModeStandard WeekMode = iota // Week starts on Sunday
	WeekModeISO                      // Week starts on Monday
)

type CalendarView int

const (
	CalendarViewMonth CalendarView = iota
	CalendarViewWeek
	CalendarViewDay
)

type MonthInfo struct {
	StartOfMonth  time.Time
	LeadingDays   int
	TrailingDays  int
	MonthDays     int
	TotalDays     int
	WeeksToRender int
}

func NewMonthInfo(now time.Time, totalDays int, totalDaysInMonth int, leadingEmptyDays int) MonthInfo {
	return MonthInfo{
		StartOfMonth:  time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
		LeadingDays:   leadingEmptyDays,
		TrailingDays:  totalDays - (leadingEmptyDays + totalDaysInMonth),
		MonthDays:     totalDaysInMonth,
		TotalDays:     totalDays,
		WeeksToRender: totalDays / 7,
	}
}
