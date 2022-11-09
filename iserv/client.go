package iserv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

type IServClient struct {
	Config        *IServAccountConfig
	ClientOptions *IServClientOptions
	IServUrl      string
	HttpClient    *http.Client
	EmailClient   *IServEmailClient
}

func (c *IServClient) Login(ac *IServAccountConfig, cc *IServClientOptions) error {
	c.Config = ac
	c.ClientOptions = cc

	// full iserv url
	c.IServUrl = fmt.Sprintf("https://%s/iserv", c.Config.IServHost)

	// user agent
	if c.ClientOptions.AgentString == "" {
		c.ClientOptions.AgentString = "Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0"
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
	req, err := http.NewRequest("POST", c.IServUrl+"/auth/login", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", c.ClientOptions.AgentString)
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

	if c.ClientOptions.EnableEmail {
		c.EmailClient = &IServEmailClient{}
		err := c.EmailClient.Login(c.Config)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *IServClient) Logout() error {
	if c.ClientOptions.EnableEmail {
		err := c.EmailClient.Logout()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *IServClient) GetBadges() (map[string]int, error) {
	res, err := c.HttpClient.Get(c.IServUrl + "/app/navigation/badges")
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

//func (c *IServClient) GetExercises() ([]IServExercise, error) {
//	res, err := c.HttpClient.Get(c.Config.IServURL + "app/navigation/badges")
//	if err != nil {
//		return []IServExercise{}, err
//	}
//	defer res.Body.Close()
//
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		return []IServExercise{}, err
//	}
//	tasksTable := doc.Find("#crud-table")
//	fmt.Println(tasksTable)
//	for _, tr := range tasksTable.Filter("tr").Nodes {
//		fmt.Println(tr)
//	}
//	return []IServExercise{}, nil
//}
