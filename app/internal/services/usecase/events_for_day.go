package usecase

import (
	"celen/internal/database"
	"celen/internal/models"
	"time"
)

type EventsForDay struct {
	eventRep *database.CacheEventRep
}

func NewEventsForDay(eventRep *database.CacheEventRep) *EventsForDay {
	return &EventsForDay{
		eventRep: eventRep,
	}
}

func (c *EventsForDay) Execute(userID int) []models.Event {
	events := c.eventRep.GetEventsForDay(userID, time.Now())
	return events
}
