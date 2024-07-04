package services

import (
	"celen/internal/database"
	"celen/internal/services/usecase"
)

type Application struct {
	CreateEvent    *usecase.CreateEvent
	DeleteEvent    *usecase.DeleteEvent
	UpdateEvent    *usecase.UpdateEvent
	EventsForDay   *usecase.EventsForDay
	EventsForWeek  *usecase.EventsForWeek
	EventsForMonth *usecase.EventsForMonth
}

func NewApplication(eventRep *database.CacheEventRep) *Application {
	return &Application{
		CreateEvent:    usecase.NewCreateEvent(eventRep),
		DeleteEvent:    usecase.NewDeleteEvent(eventRep),
		UpdateEvent:    usecase.NewUpdateEvent(eventRep),
		EventsForDay:   usecase.NewEventsForDay(eventRep),
		EventsForWeek:  usecase.NewEventsForWeek(eventRep),
		EventsForMonth: usecase.NewEventsForMonth(eventRep),
	}
}
