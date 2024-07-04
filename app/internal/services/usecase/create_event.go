package usecase

import (
	"celen/internal/database"
	"celen/internal/models"
)

type CreateEvent struct {
	eventRep *database.CacheEventRep
}

func NewCreateEvent(eventRep *database.CacheEventRep) *CreateEvent {
	return &CreateEvent{
		eventRep: eventRep,
	}
}

func (c *CreateEvent) Execute(event *models.Event) {
	c.eventRep.CreateEvent(event)
}
