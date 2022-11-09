package iserv

import (
	"fmt"
	"os"

	"github.com/studio-b12/gowebdav"
)

type IServFilesClient struct {
	Config    *IServAccountConfig
	DavClient *gowebdav.Client
}

func (c *IServFilesClient) Login(config *IServAccountConfig) error {
	c.Config = config

	c.DavClient = gowebdav.NewClient(fmt.Sprintf("https://webdav.%s", c.Config.IServHost), c.Config.Username, c.Config.Password)
	return c.DavClient.Connect()
}

func (c *IServFilesClient) ReadDir(path string) ([]os.FileInfo, error) {
	return c.DavClient.ReadDir(path)
}
