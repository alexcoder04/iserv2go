package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/alexcoder04/friendly/v2/ffiles"
)

func (c *WebClient) saveCredentials() error {
	url, err := url.Parse(c.config.IServUrl())
	if err != nil {
		return err
	}

	data, err := json.Marshal(c.httpClient.Jar.Cookies(url))
	if err != nil {
		return err
	}

	dir, err := ffiles.GetCacheDirFor("iserv2go")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(dir, "cookies.json"), data, 0600)
}

func (c *WebClient) loadCredentials() error {
	dir, err := ffiles.GetCacheDirFor("iserv2go")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path.Join(dir, "cookies.json"))
	if err != nil {
		return err
	}

	res := make([]*http.Cookie, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	url, err := url.Parse(c.config.IServUrl())
	if err != nil {
		return err
	}

	c.httpClient.Jar.SetCookies(url, res)

	return nil
}
