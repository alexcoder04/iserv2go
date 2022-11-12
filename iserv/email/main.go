package email

import (
	"fmt"

	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/emersion/go-imap/client"
)

type IServEmailClient struct {
	config     *types.IServAccountConfig
	imapClient *client.Client
}

func (c *IServEmailClient) Login(config *types.IServAccountConfig) error {
	c.config = config

	conn, err := client.DialTLS(fmt.Sprintf("%s:993", c.config.IServHost), nil)
	if err != nil {
		fmt.Println("error dial tls")
		return err
	}
	c.imapClient = conn

	return c.imapClient.Login(c.config.Username, c.config.Password)
}

func (c *IServEmailClient) Logout() error {
	return c.imapClient.Logout()
}
