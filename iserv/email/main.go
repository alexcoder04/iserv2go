package email

import (
	"fmt"
	"net/smtp"

	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/emersion/go-imap/client"
)

type IServEmailClient struct {
	config     *types.IServAccountConfig
	imapClient *client.Client
	smtpAuth   smtp.Auth
}

func (c *IServEmailClient) Login(config *types.IServAccountConfig) error {
	c.config = config

	// imap client
	conn1, err := client.DialTLS(fmt.Sprintf("%s:993", c.config.IServHost), nil)
	if err != nil {
		return err
	}
	c.imapClient = conn1

	err = c.imapClient.Login(c.config.Username, c.config.Password)
	if err != nil {
		return err
	}

	// smtp client
	c.smtpAuth = smtp.PlainAuth("", c.config.Username, c.config.Password, c.config.IServHost)

	return nil
}

func (c *IServEmailClient) Logout() error {
	return c.imapClient.Logout()
}
