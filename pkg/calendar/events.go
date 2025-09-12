package calendar

import "time"

type EventMap map[int][]*Event

func GetMonthEvents(now time.Time) (EventMap, error) {
	// Give placeholder events
	events := []*Event{
		NewEvent(
			"Meeting with Bingus",
			"Discuss project updates",
			time.Date(now.Year(), now.Month(), 10, 14, 0, 0, 0, time.UTC),
			time.Date(now.Year(), now.Month(), 10, 14, 1, 0, 0, time.UTC),
		),
		NewEvent(
			"Meeting with Bingus's dumb cat",
			"Discuss project updates",
			time.Date(now.Year(), now.Month(), 10, 14, 1, 0, 0, time.UTC),
			time.Date(now.Year(), now.Month(), 10, 15, 2, 0, 0, time.UTC),
		),
		NewAllDayEvent(
			"Conference",
			"Annual tech conference",
			time.Date(now.Year(), now.Month(), 20, 0, 0, 0, 0, time.UTC),
		),
	}
	monthEvents := make(EventMap, 0)
	for _, event := range events {
		day := event.StartTime.Day()
		if _, exists := monthEvents[day]; !exists {
			monthEvents[day] = []*Event{}
		}
		monthEvents[day] = append(monthEvents[day], event)
	}
	return monthEvents, nil
}
