package web

import (
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

func (c *WebClient) doGetRequest(url string) ([]byte, error) {
	res, err := c.httpClient.Get(c.config.IServUrl() + url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func (c *WebClient) doGetRequestQueryDoc(url string) (*goquery.Document, error) {
	res, err := c.httpClient.Get(c.config.IServUrl() + url)
	if err != nil {
		return &goquery.Document{}, err
	}
	defer res.Body.Close()

	return goquery.NewDocumentFromReader(res.Body)
}
