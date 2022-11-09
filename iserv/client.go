package iserv

type IServClient struct {
	Config        *IServAccountConfig
	ClientOptions *IServClientOptions

	EmailClient *IServEmailClient
	FilesClient *IServFilesClient
	WebClient   *IServWebClient
}

func (c *IServClient) Login(ac *IServAccountConfig, cc *IServClientOptions) error {
	c.Config = ac
	c.ClientOptions = cc

	// login modules

	if c.ClientOptions.EnableWeb {
		c.WebClient = &IServWebClient{}
		err := c.WebClient.Login(c.Config, c.ClientOptions.AgentString)
		if err != nil {
			return err
		}
	}

	if c.ClientOptions.EnableEmail {
		c.EmailClient = &IServEmailClient{}
		err := c.EmailClient.Login(c.Config)
		if err != nil {
			return err
		}
	}

	if c.ClientOptions.EnableFiles {
		c.FilesClient = &IServFilesClient{}
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
