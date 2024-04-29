package invitees

import (
	"context"
	"fmt"
	"os"

	"github.com/JerryJeager/will-be-there-backend/service/event"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type InviteeSv interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) (string, error)
	GetInvitees(ctx context.Context, eventID uuid.UUID) (*Invitees, error)
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

	if err := sendEmail(invitee); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (o *InviteeServ) GetInvitees(ctx context.Context, eventID uuid.UUID) (*Invitees, error) {
	return o.repo.GetInvitees(ctx, eventID)
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

func sendEmail(invitee *Invitee) error {
	event, err := event.GetMyEvent(invitee.EventID)

	if err != nil{
		return err
	}

	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", invitee.Email)
	m.SetAddressHeader("Cc", invitee.Email, email)
	m.SetHeader("Subject", "Will Be There")
	m.SetBody("text/html", fmt.Sprintf(`
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Will Be There</title>
			<style>
				body {{
					font-family: Arial, sans-serif;
					background-color: #f2f2f2;
					padding: 20px;
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
					margin: 0;
				}}
				.email-container {{
					max-width: 600px;
					width: 100%%;
					border-radius: 5px;
					padding: 20px;
					box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
				}}
				h2 {{
					color: #333333;
					margin-bottom: 20px;
				}}
				p {{
					color: #666666;
					margin-bottom: 10px;
				}}
				ul {{
					color: #666666;
					margin-bottom: 10px;
					padding-left: 20px;
				}}
				.footer-text {{
					color: #666666;
					margin-top: 20px;
				}}
			</style>
		</head>
		<body>
			<div class="email-container">
				<!-- Header -->
				<table width="100%%" cellpadding="0" cellspacing="0" border="0">
					<tr>
						<td align="center">
							<img src="%s" alt="Event Image" style="max-width: 200px;">
						</td>
					</tr>
				</table>
				<!-- Content -->
				<h2>You're Invited!</h2>
				<p>Dear %s,</p>
				<p>You're invited to an event, %s</p>
				<p>Event Details:</p>
				<ul>
					<li><strong>Date:</strong> %s</li>
					<li><strong>Location:</strong> %s</li>
				</ul>
				<p>We look forward to seeing you at the event!</p>
				<p class="footer-text">Best regards,<br>Will Be There Team.</p>
			</div>
		</body>
	</html>
	`, event.ImageUrl, invitee.FirstName, event.Name, event.Date, event.Venue))

	d := gomail.NewDialer("smtp.gmail.com", 587, emailUsername, emailPassword)

	// Send the email to invitee
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
