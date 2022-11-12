package web

import "io/ioutil"

func (c *IServWebClient) doGetRequest(url string) ([]byte, error) {
	res, err := c.httpClient.Get(c.iServUrl + url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
