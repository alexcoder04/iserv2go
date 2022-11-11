package iserv

import (
	"github.com/alexcoder04/iserv2go/iserv/email"
	"github.com/alexcoder04/iserv2go/iserv/files"
	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/alexcoder04/iserv2go/iserv/web"
)

type IServClient struct {
	Config        *types.IServAccountConfig
	ClientOptions *types.IServClientOptions

	EmailClient *email.IServEmailClient
	FilesClient *files.IServFilesClient
	WebClient   *web.IServWebClient
}

func (c *IServClient) Login(ac *types.IServAccountConfig, cc *types.IServClientOptions) error {
	c.Config = ac
	c.ClientOptions = cc

	// login modules

	if c.ClientOptions.EnableWeb {
		c.WebClient = &web.IServWebClient{}
		err := c.WebClient.Login(c.Config, c.ClientOptions.AgentString)
		if err != nil {
			return err
		}
	}

	if c.ClientOptions.EnableEmail {
		c.EmailClient = &email.IServEmailClient{}
		err := c.EmailClient.Login(c.Config)
		if err != nil {
			return err
		}
	}

	if c.ClientOptions.EnableFiles {
		c.FilesClient = &files.IServFilesClient{}
		err := c.FilesClient.Login(c.Config)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *IServClient) Logout() error {
	// logout modules

	if c.ClientOptions.EnableEmail {
		err := c.EmailClient.Logout()
		if err != nil {
			return err
		}
	}

	return nil
}
