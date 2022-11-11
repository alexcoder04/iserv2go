package email

import (
	"fmt"

	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/emersion/go-imap/client"
)

type IServEmailClient struct {
	Config     *types.IServAccountConfig
	ImapClient *client.Client
}

func (c *IServEmailClient) Login(config *types.IServAccountConfig) error {
	c.Config = config

	conn, err := client.DialTLS(fmt.Sprintf("%s:993", c.Config.IServHost), nil)
	if err != nil {
		fmt.Println("error dial tls")
		return err
	}
	c.ImapClient = conn

	err = c.ImapClient.Login(c.Config.Username, c.Config.Password)
	if err != nil {
		return err
	}

	return nil
}

func (c *IServEmailClient) Logout() error {
	return c.ImapClient.Logout()
}
