package web

import "io/ioutil"

func (c *IServWebClient) DoGetRequest(url string) ([]byte, error) {
	res, err := c.HttpClient.Get(c.IServUrl + url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
