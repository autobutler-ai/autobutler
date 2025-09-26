package types

import "fmt"

type Email struct {
	Subject string `json:"subject"`
	From    string `json:"from"`
	Body    string `json:"body"`
}

var mockEmails = []Email{
	{
		Subject: "Welcome to Autobutler",
		From:    "info@autobutler.ai",
		Body:    "Thank you for joining <strong>Autobutler</strong>. We're excited to have you on board!",
	},
	{
		Subject: "Your Invoice",
		From:    "billing@autobutler.ai",
		Body:    "Please find attached your invoice for this month. Contact us if you have any questions.",
	},
	{
		Subject: "Service Reminder",
		From:    "service@autobutler.ai",
		Body:    "This is a reminder for your upcoming car service appointment. Let us know if you need to reschedule.",
	},
}

func MockEmails() []Email {
	mockEmails = append(mockEmails, Email{
		Subject: fmt.Sprintf("New Email %d", len(mockEmails)),
		From:    "bingus@autobutler.ai",
		Body:    "This is a newly generated email for testing purposes",
	})
	return mockEmails
}
