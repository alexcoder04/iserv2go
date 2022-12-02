package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexcoder04/iserv2go/iserv/iutils"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func iScriptFunctionWebGetBadges(s []string) {
	badges, err := Client.Web.GetBadges()
	if err != nil {
		Warn("Cannot load badges: %s", err.Error())
		return
	}
	for key, value := range badges {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func iScriptFunctionWebGetNotifications(s []string) {
	notifications, err := Client.Web.GetNotifications()
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

func iScriptFunctionWebGetUpcomingEvents(s []string) {
	events, err := Client.Web.GetUpcomingEvents(30)
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

func iScriptFunctionWebGetCurrentExercises(s []string) {
	exercises, err := Client.Web.GetCurrentExercises()
	if err != nil {
		Warn("Cannot load exercises: %s", err.Error())
		return
	}
	for _, e := range exercises {
		fmt.Printf("%s: %s by %s\n", e.DueDate, e.Title, e.Teacher)
	}
}

func iScriptFunctionWebGetPastExercises(s []string) {
	exercises, err := Client.Web.GetPastExercises()
	if err != nil {
		Warn("Cannot load exercises: %s", err.Error())
		return
	}
	for _, e := range exercises {
		fmt.Printf("%s: %s by %s\n", e.DueDate, e.Title, e.Teacher)
	}
}

func iScriptFunctionEmailListMailboxes(s []string) {
	mailboxes, err := Client.Email.ListMailboxes()
	if err != nil {
		Warn("Cannot load email mailboxes: %s", err.Error())
		return
	}
	for _, m := range mailboxes {
		fmt.Println(m.Name)
	}
}

func iScriptFunctionEmailReadMailbox(s []string) {
	var mb string
	if len(s) < 1 {
		mb = "INBOX"
	} else {
		mb = s[0]
	}
	messages, err := Client.Email.ReadMailbox(mb, 50)
	if err != nil {
		Warn("Cannot read messages: %s", err.Error())
		return
	}
	for _, m := range messages {
		fmt.Printf("'%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
	}
}

func iScriptFunctionEmailSendMail(s []string) {
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
	err := Client.Email.SendMail(m)
	if err != nil {
		Warn("Error sending mail: %s", err.Error())
	}
}
