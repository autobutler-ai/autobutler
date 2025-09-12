package calendar

import (
	"autobutler/pkg/rand"
	"time"
)

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

type Event struct {
	ID          string
	Title       string
	Description string
	StartTime   time.Time
	EndTime     *time.Time
	AllDay      bool
	Location    string
}

func NewEvent(title string, description string, startTime, endTime time.Time) *Event {
	return &Event{
		ID:          rand.ID(),
		Title:       title,
		Description: description,
		StartTime:   startTime,
		EndTime:     &endTime,
	}
}

func NewAllDayEvent(title string, description string, startTime time.Time) *Event {
	return &Event{
		ID:          rand.ID(),
		Title:       title,
		Description: description,
		StartTime:   startTime,
		AllDay:      true,
	}
}
