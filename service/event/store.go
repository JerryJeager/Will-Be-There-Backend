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

	UpdateImageUrl(ctx context.Context, eventID uuid.UUID, imageUrl string) error

	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
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

// function to get event details that'll be emailed to guests
func GetMyEvent(EventID uuid.UUID) (*Event, error) {
	var event Event
	query := config.Session.First(&event, "id = ?", EventID)
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

func (o *EventRepo) UpdateImageUrl(ctx context.Context, eventID uuid.UUID, imageUrl string) error {
	query := config.Session.Model(&Event{}).Where("id = ?", eventID).Update("image_url", imageUrl)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (o *EventRepo) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	var event Event
	query := config.Session.WithContext(ctx).Where("id = ?", eventID).Delete(&event)

	if query.Error != nil {
		return query.Error
	}

	return nil
}
