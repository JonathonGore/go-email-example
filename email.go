package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// MailClient is responsible for sending emails on behalf of the user.
type MailClient struct {
	client          *sendgrid.Client
	welcomeTemplate *template.Template
}

// NewMailClient creates a new client for the given API Key.
func NewMailClient(apikey string) (*MailClient, error) {
	templateName := "welcome.html"

	tmpl, err := template.New(templateName).ParseFiles(templateName)
	if err != nil {
		return nil, err
	}

	return &MailClient{
		client:          sendgrid.NewSendClient(apikey),
		welcomeTemplate: tmpl,
	}, nil
}

// renderVerificationTemplate consumes a name and renders the welcome template.
func (c *MailClient) renderWelcomeMessage(name string) (string, error) {
	data := map[string]string{"name": name}

	buf := &bytes.Buffer{}
	if err := c.welcomeTemplate.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// SendWelcomeEmail will send a welcome email to the provided email adress
// using the given name to greet the user.
func (c *MailClient) SendWelcomeEmail(name, email string) error {
	from := mail.NewEmail("Jack", "jack@example.com")
	to := mail.NewEmail(name, email)

	contentBody, err := c.renderWelcomeMessage(name)
	if err != nil {
		return err
	}

	content := &mail.Content{"text/html", contentBody}
	subject := "Welcome!"
	message := mail.NewV3MailInit(from, subject, to, content)
	response, err := c.client.Send(message)

	if err != nil {
		return err
	} else if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unable to send email: %v", response.Body)
	}

	return nil
}
