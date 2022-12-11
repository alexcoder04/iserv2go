package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/iserv2go/iserv/types"
	"golang.org/x/net/publicsuffix"
)

type WebClient struct {
	config     *types.ClientConfig
	httpClient *http.Client
}

func (c *WebClient) Login(config *types.ClientConfig) error {
	c.config = config

	// user agent
	if c.config.AgentString == "" {
		c.config.AgentString = "Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0"
	}

	// http client
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return err
	}
	c.httpClient = &http.Client{
		Jar: jar,
	}

	// load credentials if saved
	// TODO check whether login still valid
	if c.config.SaveSessions {
		err := c.loadCredentials()
		if err == nil {
			return nil
		}
		fmt.Printf("Failed to load cookies: %s, re-logging in\n", err)
	}

	// login data
	form := &url.Values{}
	form.Add("_username", c.config.Username)
	form.Add("_password", c.config.Password)
	if c.config.SaveSessions {
		form.Add("_remember_me", "on")
	}

	// login request
	req, err := http.NewRequest("POST", c.config.IServUrl()+"/auth/login", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", c.config.AgentString)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
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

func (c *WebClient) Logout() error {
	// save credentials if necessary
	if c.config.SaveSessions {
		err := c.saveCredentials()
		if err == nil {
			return nil
		}
		friendly.Warn("Failed to save cookies: %s, logging out\n", err)
	}

	req, err := http.NewRequest("POST", c.config.IServUrl()+"/auth/logout", strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", c.config.AgentString)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	return err
}
