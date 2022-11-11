package iserv

import (
	"fmt"
	"io/ioutil"
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

func (c *IServFilesClient) DownloadFile(webdavpath string, localpath string) error {
	bytes, err := c.DavClient.Read(webdavpath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(localpath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *IServFilesClient) UploadFile(webdavpath string, localpath string) error {
	bytes, err := ioutil.ReadFile(localpath)
	if err != nil {
		return err
	}

	err = c.DavClient.Write(webdavpath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
