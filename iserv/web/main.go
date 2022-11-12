package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/alexcoder04/iserv2go/iserv/types"
	"golang.org/x/net/publicsuffix"
)

type IServWebClient struct {
	config      *types.IServAccountConfig
	agentString string
	iServUrl    string
	httpClient  *http.Client
}

func (c *IServWebClient) Login(config *types.IServAccountConfig, agentString string) error {
	c.config = config
	c.agentString = agentString

	// full iserv url
	c.iServUrl = fmt.Sprintf("https://%s/iserv", c.config.IServHost)

	// user agent
	if c.agentString == "" {
		c.agentString = "Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0"
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

	// login data
	form := &url.Values{}
	form.Add("_username", c.config.Username)
	form.Add("_password", c.config.Password)

	// login request
	req, err := http.NewRequest("POST", c.iServUrl+"/auth/login", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", c.agentString)
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
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

func (c *IServWebClient) Logout() error {
	req, err := http.NewRequest("POST", c.iServUrl+"/auth/logout", strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", c.agentString)
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
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
