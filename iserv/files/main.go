package files

import (
	"fmt"

	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/studio-b12/gowebdav"
)

type IServFilesClient struct {
	config    *types.IServAccountConfig
	davClient *gowebdav.Client
}

func (c *IServFilesClient) Login(config *types.IServAccountConfig) error {
	c.config = config

	c.davClient = gowebdav.NewClient(fmt.Sprintf("https://webdav.%s", c.config.IServHost), c.config.Username, c.config.Password)
	return c.davClient.Connect()
}

func (c *IServFilesClient) Logout() error {
	return nil
}
