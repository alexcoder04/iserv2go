package files

import (
	"io/ioutil"
	"os"
)

func (c *IServFilesClient) ReadDir(path string) ([]os.FileInfo, error) {
	return c.davClient.ReadDir(path)
}

func (c *IServFilesClient) ReadFile(path string) ([]byte, error) {
	return c.davClient.Read(path)
}

func (c *IServFilesClient) DownloadFile(webdavpath string, localpath string) error {
	bytes, err := c.ReadFile(webdavpath)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(localpath, bytes, 0600)
}

func (c *IServFilesClient) WriteFile(path string, content []byte) error {
	return c.davClient.Write(path, content, 0600)
}

func (c *IServFilesClient) UploadFile(localpath string, webdavpath string) error {
	bytes, err := ioutil.ReadFile(localpath)
	if err != nil {
		return err
	}

	return c.WriteFile(webdavpath, bytes)
}
