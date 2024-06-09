package invitees

import (
	"context"
	"errors"

	"github.com/JerryJeager/will-be-there-backend/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InviteeStore interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) error
	GetInvitees(ctx context.Context, weddingID uuid.UUID) (*Invitees, error)
	GetInviteeByID(ctx context.Context, InviteeID uuid.UUID) (*Invitee, error)
	UpdateInviteeStatus(ctx context.Context, inviteeID uuid.UUID, status *NewStatus) error
	UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error
	DeleteInvitee(ctx context.Context, inviteeID uuid.UUID) error
}

type InviteeRepo struct {
	client *gorm.DB
}

func NewInviteeRepo(client *gorm.DB) *InviteeRepo {
	return &InviteeRepo{client: client}
}

func (o *InviteeRepo) CreateInvitee(ctx context.Context, invitee *Invitee) error {
	result := config.Session.Create(&invitee).WithContext(ctx)

	if result.Error != nil {
		return errors.New("failed to add guest")
	}

	return nil
}

func (o *InviteeRepo) GetInvitees(ctx context.Context, eventID uuid.UUID) (*Invitees, error) {
	var invitees Invitees

	query := config.Session.WithContext(ctx).Model(Invitee{}).Where("event_id = ?", eventID).Find(&invitees)

	if query.Error != nil {
		return nil, query.Error
	}

	return &invitees, nil
}

func (o *InviteeRepo) GetInviteeByID(ctx context.Context, InviteeID uuid.UUID) (*Invitee, error) {
	var invitee Invitee

	query := config.Session.WithContext(ctx).Model(Invitee{}).Where("id = ?", InviteeID).Find(&invitee)

	if query.Error != nil {
		return nil, query.Error
	}

	return &invitee, nil
}

func (o *InviteeRepo) UpdateInviteeStatus(ctx context.Context, inviteeID uuid.UUID, status *NewStatus) error {

	query := config.Session.Model(Invitee{}).WithContext(ctx).Where("id = ?", inviteeID).Update("status", status.Status)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (o *InviteeRepo) UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error {
	query := config.Session.Model(Invitee{}).WithContext(ctx).Where("id = ? ", inviteeID).Updates(Invitee{FirstName: invitee.FirstName, LastName: invitee.LastName, PlusOnes: invitee.PlusOnes, Status: invitee.Status})

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (o *InviteeRepo) DeleteInvitee(ctx context.Context, inviteeID uuid.UUID) error {
	var invitee Invitee
	query := config.Session.WithContext(ctx).Where("id = ?", inviteeID).Delete(&invitee)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
