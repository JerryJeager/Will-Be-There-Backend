package event

import (
	"context"

	"github.com/JerryJeager/will-be-there-backend/service"
	"github.com/google/uuid"
) 

type EventSv interface {
	GetEvent(ctx context.Context, eventID uuid.UUID) (*Event, error)
	GetEvents(ctx context.Context, eventID uuid.UUID) (*Events, error)
	CreateEvent(ctx context.Context, Event *service.Event) (string, error)

	CreateEventType(ctx context.Context, eventID uuid.UUID, eventType *service.EventType) (string, error)
	UpdateEventType(ctx context.Context, eventID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error)
	DeleteEventType(ctx context.Context, eventID, eventTypeID uuid.UUID) error
}

type EventServ struct {
	repo EventStore
}

func NewEventService(repo EventStore) *EventServ {
	return &EventServ{repo: repo}
}

func (o *EventServ) GetEvent(ctx context.Context, eventID uuid.UUID) (*Event, error) {
	return o.repo.GetEvent(ctx, eventID)
}

func (o *EventServ) GetEvents(ctx context.Context, userID uuid.UUID) (*Events, error) {
	return o.repo.GetEvents(ctx, userID)
}


func (o *EventServ) CreateEvent(ctx context.Context, event *service.Event) (string, error) {
	id := uuid.New()

	event.ID = id
	if err := o.repo.CreateEvent(ctx, event); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (o *EventServ) CreateEventType(ctx context.Context, eventID uuid.UUID, eventType *service.EventType) (string, error) {
	return o.repo.CreateEventType(ctx, eventID, eventType)
}

func (o *EventServ) UpdateEventType(ctx context.Context, eventID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error) {
	return o.repo.UpdateEventType(ctx, eventID, eventTypeID, eventType)
}

func (o *EventServ) DeleteEventType(ctx context.Context, eventID, evenTypeID uuid.UUID) error {
	return o.repo.DeleteEventType(ctx, eventID, evenTypeID)
}
