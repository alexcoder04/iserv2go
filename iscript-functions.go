package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexcoder04/iserv2go/iserv"
	"github.com/alexcoder04/iserv2go/iserv/iutils"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func iScriptFunctionWebGetBadges(c *iserv.Client, s []string) {
	badges, err := c.Web.GetBadges()
	if err != nil {
		Warn("Cannot load badges: %s", err.Error())
		return
	}
	for key, value := range badges {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func iScriptFunctionWebGetNotifications(c *iserv.Client, s []string) {
	notifications, err := c.Web.GetNotifications()
	if err != nil {
		Warn("Cannot load notifications: %s", err.Error())
		return
	}
	if notifications.Status != "success" {
		Warn("IServ didn't return success: %s", notifications.Status)
		return
	}
	// TODO
	for _, n := range notifications.Data.Notifications {
		fmt.Printf("%v\n", n)
	}
}

func iScriptFunctionWebGetUpcomingEvents(c *iserv.Client, s []string) {
	events, err := c.Web.GetUpcomingEvents(30)
	if err != nil {
		Warn("Cannot load events: %s", err.Error())
		return
	}
	if len(events.Errors) > 0 {
		Warn("IServ returned error: %s", strings.Join(events.Errors, ", "))
		return
	}
	for _, e := range events.Events {
		fmt.Printf("%v - %v: %s\n", e.Start, e.End, e.Title)
	}
}

func iScriptFunctionWebGetCurrentExercises(c *iserv.Client, s []string) {
	exercises, err := c.Web.GetCurrentExercises()
	if err != nil {
		Warn("Cannot load exercises: %s", err.Error())
		return
	}
	for _, e := range exercises {
		fmt.Printf("%s: %s by %s\n", e.DueDate, e.Title, e.Teacher)
	}
}

func iScriptFunctionWebGetPastExercises(c *iserv.Client, s []string) {
	exercises, err := c.Web.GetPastExercises()
	if err != nil {
		Warn("Cannot load exercises: %s", err.Error())
		return
	}
	for _, e := range exercises {
		fmt.Printf("%s: %s by %s\n", e.DueDate, e.Title, e.Teacher)
	}
}

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
