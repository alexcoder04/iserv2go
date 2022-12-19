package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

// TODO use the friendly implmentation
func GetCacheFolder() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		home, err := os.UserHomeDir()
		if err != nil {
			return "."
		}
		dir = path.Join(home, ".cache")
	}

	ourDir := path.Join(dir, "iserv2go")

	err = os.MkdirAll(ourDir, 0700)
	if err != nil {
		return "."
	}

	return ourDir
}

func (c *WebClient) saveCredentials() error {
	url, err := url.Parse(c.config.IServUrl())
	if err != nil {
		return err
	}

	data, err := json.Marshal(c.httpClient.Jar.Cookies(url))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(GetCacheFolder(), "cookies.json"), data, 0600)
}

func (c *WebClient) loadCredentials() error {
	data, err := ioutil.ReadFile(path.Join(GetCacheFolder(), "cookies.json"))
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
