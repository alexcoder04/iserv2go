package files

import (
	"fmt"

	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/studio-b12/gowebdav"
)

type IServFilesClient struct {
	Config    *types.IServAccountConfig
	DavClient *gowebdav.Client
}

func (c *IServFilesClient) Login(config *types.IServAccountConfig) error {
	c.Config = config

	c.DavClient = gowebdav.NewClient(fmt.Sprintf("https://webdav.%s", c.Config.IServHost), c.Config.Username, c.Config.Password)
	return c.DavClient.Connect()
}
