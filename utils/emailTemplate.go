package utils

import (
	"fmt"
	"time"
)

func InviteeEmail(eventImageUrl, firstName, name, venue string, date *time.Time) string {
	return fmt.Sprintf(`
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Will Be There</title>
			<style>
				%s
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
	`, styles, eventImageUrl, firstName, name, date, venue)
}
func PendingInviteeEmail(eventImageUrl, firstName, name, venue, confirmLink string, date *time.Time) string {
	return fmt.Sprintf(`
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Will Be There</title>
			<style>
				%s
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
				<a href="%s">
					<button>Confirm you're attending or not</button>
				</a>
				<p>We look forward to seeing you at the event!</p>
				<p class="footer-text">Best regards,<br>Will Be There Team.</p>
			</div>
		</body>
	</html>
	`, styles, eventImageUrl, firstName, name, date, venue, confirmLink)
}

const styles = `body {{
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
				button {{
					color: #fff;
					background-color: #0D35FB;
					padding: 5px;
					border-radius: 3px;
				}}
				.footer-text {{
					color: #666666;
					margin-top: 20px;
				}}`