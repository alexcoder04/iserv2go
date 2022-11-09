package iserv

import (
	"fmt"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type IServEmailClient struct {
	Config     *IServAccountConfig
	ImapClient *client.Client
}

func (c *IServEmailClient) Login(config *IServAccountConfig) error {
	c.Config = config

	conn, err := client.DialTLS(fmt.Sprintf("%s:993", c.Config.IServHost), nil)
	if err != nil {
		fmt.Println("error dial tls")
		return err
	}
	c.ImapClient = conn

	err = c.ImapClient.Login(c.Config.Username, c.Config.Password)
	if err != nil {
		return err
	}

	return nil
}

func (c *IServEmailClient) ListMailboxes() ([]imap.MailboxInfo, error) {
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.ImapClient.List("", "*", mailboxes)
	}()

	res := []imap.MailboxInfo{}
	for m := range mailboxes {
		res = append(res, *m)
	}

	if err := <-done; err != nil {
		return []imap.MailboxInfo{}, err
	}

	return res, nil
}

func (c *IServEmailClient) ReadMailbox(name string, limit uint32) ([]imap.Message, error) {
	mbox, err := c.ImapClient.Select(name, false)
	if err != nil {
		return []imap.Message{}, err
	}

	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > (limit - 1) {
		from = mbox.Messages - limit
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.ImapClient.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	res := []imap.Message{}
	for msg := range messages {
		res = append(res, *msg)
	}

	if err := <-done; err != nil {
		return []imap.Message{}, err
	}

	return res, nil
}

func (c *IServEmailClient) Logout() error {
	return c.ImapClient.Logout()
}
