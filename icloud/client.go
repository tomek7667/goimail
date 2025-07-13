package icloud

import (
	"fmt"

	"gopkg.in/mail.v2"
)

type Client struct {
	Dialer      *mail.Dialer
	SenderEmail string

	transporter *MailerTransporter
}

type MailerTransporter struct {
	Host     string
	Port     int
	Secure   *bool
	User     string
	Password string
}

type NewCustomOptions struct {
	Transporter *MailerTransporter
	SenderEmail string
}

func NewCustom(opts *NewCustomOptions) (*Client, error) {
	c := &Client{
		transporter: opts.Transporter,
		SenderEmail: opts.SenderEmail,
	}
	c.Dialer = mail.NewDialer(
		c.transporter.Host,
		c.transporter.Port,
		c.transporter.User,
		c.transporter.Password,
	)
	if c.transporter.Secure == nil {
		c.Dialer.StartTLSPolicy = mail.OpportunisticStartTLS
	} else if !*c.transporter.Secure {
		c.Dialer.StartTLSPolicy = mail.NoStartTLS
	} else if *c.transporter.Secure {
		c.Dialer.StartTLSPolicy = mail.MandatoryStartTLS
	}
	closer, err := c.Dialer.Dial()
	if err != nil {
		return nil, fmt.Errorf("invalid transporter config. Dialing SMTP server failed: %w", err)
	}
	defer closer.Close()

	return c, nil
}

func New(icloudEmail, senderEmail, appSpecificPassword string) (*Client, error) {
	return NewCustom(&NewCustomOptions{
		Transporter: &MailerTransporter{
			Host:     "smtp.mail.me.com",
			Port:     587,
			User:     icloudEmail,
			Password: appSpecificPassword,
		},
		SenderEmail: senderEmail,
	})
}
