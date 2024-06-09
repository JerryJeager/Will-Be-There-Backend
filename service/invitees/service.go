package invitees

import (
	"context"
	"fmt"
	"os"

	"github.com/JerryJeager/will-be-there-backend/service/event"
	"github.com/JerryJeager/will-be-there-backend/utils"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type InviteeSv interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) (string, error)
	CreateInviteeByEmail(ctx context.Context, invitee *InviteeByEmail) (string, error)
	GetInvitees(ctx context.Context, eventID uuid.UUID) (*Invitees, error)
	GetInviteeByID(ctx context.Context, InviteeID uuid.UUID) (*Invitee, error)
	UpdateInviteeStatus(ctx context.Context, inviteeID uuid.UUID, status *NewStatus) error
	UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error
	DeleteInvitee(ctx context.Context, inviteeID uuid.UUID) error
}

type InviteeServ struct {
	repo InviteeStore
}

func NewInviteeService(repo InviteeStore) *InviteeServ {
	return &InviteeServ{repo: repo}
}

func (o *InviteeServ) CreateInvitee(ctx context.Context, invitee *Invitee) (string, error) {
	id := uuid.New()
	invitee.ID = id
	if err := IsValidStatus(invitee.Status); err != nil {
		return "", err
	}

	if err := o.repo.CreateInvitee(ctx, invitee); err != nil {
		return "", err
	}

	if invitee.Status != ATTENDING {
		return id.String(), nil
	}

	if err := sendEmail(invitee, false); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (o *InviteeServ) CreateInviteeByEmail(ctx context.Context, invitee *InviteeByEmail) (string, error) {
	id := uuid.New()
	var newInvitee Invitee
	newInvitee.ID = id
	newInvitee.Status = PENDING
	newInvitee.Email = invitee.Email
	newInvitee.EventID = invitee.EventID
	if err := o.repo.CreateInvitee(ctx, &newInvitee); err != nil {
		return "", err
	}

	if err := sendEmail(&newInvitee, true); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (o *InviteeServ) GetInvitees(ctx context.Context, eventID uuid.UUID) (*Invitees, error) {
	return o.repo.GetInvitees(ctx, eventID)
}
func (o *InviteeServ)GetInviteeByID(ctx context.Context, InviteeID uuid.UUID) (*Invitee, error) {
	return o.repo.GetInviteeByID(ctx, InviteeID)
}

func (o *InviteeServ) UpdateInviteeStatus(ctx context.Context, inviteeID uuid.UUID, status *NewStatus) error {
	if err := IsValidStatus(status.Status); err != nil {
		return err
	}
	return o.repo.UpdateInviteeStatus(ctx, inviteeID, status)
}

func (o *InviteeServ) UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error {
	return o.repo.UpdateInvitee(ctx, inviteeID, invitee)
}

func (o *InviteeServ) DeleteInvitee(ctx context.Context, inviteeID uuid.UUID) error {
	return o.repo.DeleteInvitee(ctx, inviteeID)
}

func sendEmail(invitee *Invitee, isPending bool) error {
	event, err := event.GetMyEvent(invitee.EventID)

	if err != nil{
		return err
	}

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	confirmLink := fmt.Sprintf("https://will-be-there.vercel.app/invitation/%s?extras=0&guest=%s", invitee.EventID, invitee.ID)
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", invitee.Email)
	m.SetAddressHeader("Cc", invitee.Email, email)
	m.SetHeader("Subject", "Will Be There")
	if !isPending {
		m.SetBody("text/html", utils.InviteeEmail(event.ImageUrl, invitee.FirstName, event.Name, event.Venue, event.Date))
	}else{
		m.SetBody("text/html", utils.PendingInviteeEmail(event.ImageUrl, invitee.FirstName, event.Name, event.Venue, confirmLink, event.Date))
	}

	d := gomail.NewDialer("smtp.gmail.com", 587, emailUsername, emailPassword)

	// Send the email to invitee
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
