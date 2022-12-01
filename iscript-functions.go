package main

import (
	"fmt"
	"os"

	"github.com/alexcoder04/iserv2go/iserv"
	"github.com/alexcoder04/iserv2go/iserv/iutils"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func iScriptFunctionEmailListMailboxes(c *iserv.Client, s []string) {
	mailboxes, err := c.Email.ListMailboxes()
	if err != nil {
		Warn("Cannot load email mailboxes: %s", err.Error())
		return
	}
	for _, m := range mailboxes {
		fmt.Println(m.Name)
	}
}

func iScriptFunctionEmailReadMailbox(c *iserv.Client, s []string) {
	var mb string
	if len(s) < 1 {
		mb = "INBOX"
	} else {
		mb = s[0]
	}
	messages, err := c.Email.ReadMailbox(mb, 50)
	if err != nil {
		Warn("Cannot read messages: %s", err.Error())
		return
	}
	for _, m := range messages {
		fmt.Printf("'%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
	}
}

func iScriptFunctionEmailSendMail(c *iserv.Client, s []string) {
	if len(s) > 3 {
		Warn("Not enough arguments")
		return
	}
	myMail := fmt.Sprintf("%s@%s", os.Getenv("ISERV_USERNAME"), os.Getenv("ISERV_HOST"))
	m := types.EMail{
		Subject:    s[1],
		From:       myMail,
		To:         s[0],
		ToDispName: iutils.GetNameFromAddress(s[0]),
		CCs:        []string{},
		Body:       s[2],
	}
	err := c.Email.SendMail(m)
	if err != nil {
		Warn("Error sending mail: %s", err.Error())
	}
}
