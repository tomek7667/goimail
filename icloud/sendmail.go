package icloud

import (
	"fmt"

	"gopkg.in/mail.v2"
)

// optional SendMail options to set
type SendMailOptions struct {
	// paths to attachments
	Attachments []string

	Cc []struct {
		Email string
		Name  string
	}

	// embed image and set email body to reference the embedded image
	// key is the path to the image and value is the image variable that can be used in email body
	//
	// e.g.: "/path/img.jpg" can be referenced in body:
	//
	// <img src="cid:img.jpg" alt="embedded image">
	EmbeddedImages []string

	MessageSettings []mail.MessageSetting

	// by default "text/html"
	BodyContentType string

	// another common practice is to attach a plain text version of your HTML message, just in case some of your recipients’ clients don’t support HTML.
	// fot this you can use alternatives array with "text/plain" as one of the body content-types.
	Alternatives []struct {
		BodyContentType string
		Body            string
	}
	FromTitle *string
}

func (c *Client) SendMail(subject, body string, options *SendMailOptions, to ...string) error {
	if len(to) == 0 {
		return fmt.Errorf("you need to specify at least one recipient")
	}
	if options == nil {
		options = &SendMailOptions{
			MessageSettings: []mail.MessageSetting{},
		}
	}
	m := mail.NewMessage(options.MessageSettings...)

	var from string
	if options.FromTitle != nil {
		from = fmt.Sprintf(`"%s" <%s>`, *options.FromTitle, c.SenderEmail)
	} else {
		from = c.SenderEmail
	}
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)

	if options.Cc != nil {
		for _, c := range options.Cc {
			m.SetAddressHeader("Cc", c.Email, c.Name)
		}
	}

	if options.Attachments != nil {
		for _, p := range options.Attachments {
			m.Attach(p)
		}
	}

	if options.EmbeddedImages != nil {
		for _, p := range options.EmbeddedImages {
			m.Embed(p)
		}
	}

	if options.BodyContentType != "" {
		m.SetBody(options.BodyContentType, body)
	} else {
		m.SetBody("text/html", body)
	}

	if options.Alternatives != nil {
		for _, a := range options.Alternatives {
			m.AddAlternative(a.BodyContentType, a.Body)
		}
	}

	err := c.Dialer.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("dialing SMTP server and sending the message failed: %w", err)
	}
	return nil
}
