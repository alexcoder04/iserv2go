package files

import (
	"fmt"

	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/studio-b12/gowebdav"
)

type FilesClient struct {
	config    *types.AccountConfig
	davClient *gowebdav.Client
}

func (c *FilesClient) Login(config *types.AccountConfig) error {
	c.config = config

	c.davClient = gowebdav.NewClient(fmt.Sprintf("https://webdav.%s", c.config.IServHost), c.config.Username, c.config.Password)
	return c.davClient.Connect()
}

func (c *FilesClient) Logout() error {
	return nil
}
