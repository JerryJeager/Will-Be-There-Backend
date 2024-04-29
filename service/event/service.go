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

	UpdateImageUrl(ctx context.Context, eventID uuid.UUID, imageUrl string) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
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

func (o *EventServ) UpdateImageUrl(ctx context.Context, eventID uuid.UUID, imageUrl string) error{
	return o.repo.UpdateImageUrl(ctx, eventID, imageUrl)
}

func (o *EventServ) DeleteEvent (ctx context.Context, eventID uuid.UUID) error {
	return o.repo.DeleteEvent(ctx, eventID)
}