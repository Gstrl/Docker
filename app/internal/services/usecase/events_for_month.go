package usecase

import (
	"celen/internal/database"
	"celen/internal/models"
	"time"
)

type EventsForMonth struct {
	eventRep *database.CacheEventRep
}

func NewEventsForMonth(eventRep *database.CacheEventRep) *EventsForMonth {
	return &EventsForMonth{
		eventRep: eventRep,
	}
}

func (c *EventsForMonth) Execute(userID int) []models.Event {
	events := c.eventRep.GetEventsForMonth(userID, time.Now())
	return events
}
