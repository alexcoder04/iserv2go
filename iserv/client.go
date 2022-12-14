package iserv

import (
	"github.com/alexcoder04/iserv2go/iserv/email"
	"github.com/alexcoder04/iserv2go/iserv/files"
	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/alexcoder04/iserv2go/iserv/web"
)

type Client struct {
	Config *types.ClientConfig

	Email *email.EmailClient
	Files *files.FilesClient
	Web   *web.WebClient
}

func (c *Client) Login(conf *types.ClientConfig) error {
	c.Config = conf

	// login modules
	for key, val := range c.Config.EnableModules {
		if key == "web" && val {
			c.Web = &web.WebClient{}
			err := c.Web.Login(c.Config)
			if err != nil {
				return err
			}
		}

		if key == "files" && val {
			c.Files = &files.FilesClient{}
			err := c.Files.Login(c.Config)
			if err != nil {
				return err
			}
		}

		if key == "email" && val {
			c.Email = &email.EmailClient{}
			err := c.Email.Login(c.Config)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) Logout() error {
	// logout modules
	for key, val := range c.Config.EnableModules {
		if key == "web" && val {
			err := c.Web.Logout()
			if err != nil {
				return err
			}
		}

		if key == "files" && val {
			err := c.Files.Logout()
			if err != nil {
				return err
			}
		}

		if key == "email" && val {
			err := c.Email.Logout()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
