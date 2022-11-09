package iserv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

type IServClient struct {
	Config     *IServAccountConfig
	HttpClient *http.Client
}

func (c *IServClient) Login(config *IServAccountConfig) error {
	c.Config = config

	// full iserv url
	if !strings.HasPrefix(c.Config.IServURL, "https://") {
		c.Config.IServURL = "https://" + c.Config.IServURL
	}
	if !strings.HasSuffix(c.Config.IServURL, "/") {
		c.Config.IServURL += "/"
	}
	if !strings.HasSuffix(c.Config.IServURL, "iserv/") {
		c.Config.IServURL += "iserv/"
	}

	// user agent
	if c.Config.AgentString == "" {
		c.Config.AgentString = "Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0"
	}

	// http client
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return err
	}
	c.HttpClient = &http.Client{
		Jar: jar,
	}

	// login data
	form := &url.Values{}
	form.Add("_username", c.Config.Username)
	form.Add("_password", c.Config.Password)
	form.Add("_remember_me", "on")

	// login request
	req, err := http.NewRequest("POST", c.Config.IServURL+"auth/login", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", c.Config.AgentString)
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *IServClient) GetBadges() (map[string]int, error) {
	res, err := c.HttpClient.Get(c.Config.IServURL + "app/navigation/badges")
	if err != nil {
		return map[string]int{}, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return map[string]int{}, err
	}

	resData := map[string]int{}
	err = json.Unmarshal(data, &resData)
	return resData, err
}
