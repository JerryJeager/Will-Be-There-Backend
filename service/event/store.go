package event

import (
	"context"

	"github.com/JerryJeager/will-be-there-backend/config"
	"github.com/JerryJeager/will-be-there-backend/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventStore interface {
	GetEvent(ctx context.Context, EventID uuid.UUID) (*Event, error)
	CreateEvent(ctx context.Context, Event *service.Event) error
	GetEvents(ctx context.Context, userID uuid.UUID) (*Events, error)

	CreateEventType(ctx context.Context, EventID uuid.UUID, eventType *service.EventType) (string, error)
	UpdateEventType(ctx context.Context, EventID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error)
	DeleteEventType(ctx context.Context, EventID, eventTypeID uuid.UUID) error
}

type EventRepo struct {
	client *gorm.DB
}

func NewEventRepo(client *gorm.DB) *EventRepo {
	return &EventRepo{client: client}
}

func (o *EventRepo) GetEvent(ctx context.Context, EventID uuid.UUID) (*Event, error) {
	var event Event
	query := config.Session.First(&event, "id = ?", EventID).WithContext(ctx)
	if query.Error != nil {
		return &Event{}, query.Error
	}
	return &event, nil
}

func (o *EventRepo) GetEvents(ctx context.Context, userID uuid.UUID) (*Events, error) {
	var events Events

	query := config.Session.WithContext(ctx).Model(Event{}).Where("user_id = ?", userID).Find(&events)

	if query.Error != nil {
		return nil, query.Error
	}

	return &events, nil
}

func (o *EventRepo) CreateEvent(ctx context.Context, event *service.Event) error {
	query := config.Session.Create(&event).WithContext(ctx)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (o *EventRepo) CreateEventType(ctx context.Context, EventID uuid.UUID, eventType *service.EventType) (string, error) {
	id := uuid.NewString()
	path := &[]string{id}
	updateEventType := `UPDATE events SET event_types = jsonb_insert(COALESCE(event_types, '{}'), ?, ?) WHERE id = ?`
	result := config.Session.Exec(updateEventType, path, *eventType, EventID).WithContext(ctx)

	if result.Error != nil{
		return "", result.Error
	}

	return id, nil
}

func (o *EventRepo) UpdateEventType(ctx context.Context, EventID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error) {
	path := &[]string{eventTypeID.String()}

	updateEventType := `UPDATE events SET event_types = jsonb_set(event_types, ?, ?, false) WHERE id = ?`
	result := config.Session.Exec(updateEventType, path, *eventType, EventID).WithContext(ctx)

	if result.Error != nil{
		return "", result.Error
	}

	return eventTypeID.String(), nil
}

func (o *EventRepo) DeleteEventType(ctx context.Context, EventID, eventTypeID uuid.UUID) error {
	udpateEventType := `UPDATE events SET event_types = jsonb_delete(event_types, ?) WHERE id = ?`
	result := config.Session.Exec(udpateEventType, eventTypeID, EventID).WithContext(ctx)

	if result.Error != nil{
		return result.Error
	}

	return nil
}
