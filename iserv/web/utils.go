package web

import (
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

func (c *IServWebClient) doGetRequest(url string) ([]byte, error) {
	res, err := c.httpClient.Get(c.iServUrl + url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func (c *IServWebClient) doGetRequestQueryDoc(url string) (*goquery.Document, error) {
	res, err := c.httpClient.Get(c.iServUrl + url)
	if err != nil {
		return &goquery.Document{}, err
	}
	defer res.Body.Close()

	return goquery.NewDocumentFromReader(res.Body)
}
