package calendar

import (
	"time"
)

var days = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func monthToInt(month time.Month) int {
	if month < time.January || month > time.December {
		return 0
	}
	return int(month)
}

func shortMonth(month time.Month) string {
	if month < time.January || month > time.December {
		return ""
	}
	return month.String()[:3]
}

func weekdayToString(day Weekday, mode WeekMode) string {
	if day < Sunday || day > Saturday {
		return ""
	}
	if mode == WeekModeISO {
		return days[(day+6)%7] // Shift so that Monday is 0
	}
	return days[day]
}

func weekdayToShortString(day Weekday, mode WeekMode) string {
	return weekdayToString(day, mode)[:3]
}

func getFirstDayOfMonth(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func getMonthInfo(now time.Time) MonthInfo {
	firstOfMonth := getFirstDayOfMonth(now)
	totalDaysInMonth := int(time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, now.Location()).Day())
	leadingEmptyDays := int(firstOfMonth.Weekday())
	totalDays := leadingEmptyDays + totalDaysInMonth
	if totalDays%7 != 0 {
		totalDays += 7 - (totalDays % 7) // Round up to the nearest week
	}
	return NewMonthInfo(now, totalDays, totalDaysInMonth, leadingEmptyDays)
}
