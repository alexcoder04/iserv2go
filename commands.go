package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexcoder04/iserv2go/iserv/iutils"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func CommandFilesLs(s []string) {
	data, err := Client.Files.ReadDir(s[0])
	if err != nil {
		Die("Cannot ls: %s", err.Error())
	}

	for _, f := range data {
		fmt.Println(f.Name())
	}
}

func CommandFilesCat(s []string) {
	data, err := Client.Files.ReadFile(s[0])
	if err != nil {
		Die("Cannot cat: %s", err.Error())
	}

	fmt.Println(data)
}

func CommandFilesDownload(s []string) {
	err := Client.Files.DownloadFile(s[0], s[1])
	if err != nil {
		Die("Cannot download: %s", err.Error())
	}
}

func CommandFilesUpload(s []string) {
	err := Client.Files.UploadFile(s[0], s[1])
	if err != nil {
		Die("Cannot upload: %s", err.Error())
	}
}

func CommandWebGetBadges(s []string) {
	badges, err := Client.Web.GetBadges()
	if err != nil {
		Die("Cannot load badges: %s", err.Error())
	}

	for key, value := range badges {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func CommandWebGetNotifications(s []string) {
	notifications, err := Client.Web.GetNotifications()
	if err != nil {
		Die("Cannot load notifications: %s", err.Error())
	}

	if notifications.Status != "success" {
		Die("IServ didn't return success: %s", notifications.Status)
	}

	// TODO
	for _, n := range notifications.Data.Notifications {
		fmt.Printf("%v\n", n)
	}
}

func CommandWebGetUpcomingEvents(s []string) {
	events, err := Client.Web.GetUpcomingEvents(30)
	if err != nil {
		Die("Cannot load events: %s", err.Error())
	}

	if len(events.Errors) > 0 {
		Die("IServ returned error: %s", strings.Join(events.Errors, ", "))
	}

	for _, e := range events.Events {
		fmt.Printf("%v - %v: %s\n", e.Start, e.End, e.Title)
	}
}

func CommandWebGetCurrentExercises(s []string) {
	exercises, err := Client.Web.GetCurrentExercises()
	if err != nil {
		Die("Cannot load exercises: %s", err.Error())
	}

	for _, e := range exercises {
		fmt.Printf("%s: %s by %s\n", e.DueDate, e.Title, e.Teacher)
	}
}

func CommandWebGetPastExercises(s []string) {
	exercises, err := Client.Web.GetPastExercises()
	if err != nil {
		Die("Cannot load exercises: %s", err.Error())
	}

	for _, e := range exercises {
		fmt.Printf("%s: %s by %s\n", e.DueDate, e.Title, e.Teacher)
	}
}

func CommandEmailListMailboxes(s []string) {
	mailboxes, err := Client.Email.ListMailboxes()
	if err != nil {
		Die("Cannot load email mailboxes: %s", err.Error())
	}

	for _, m := range mailboxes {
		fmt.Println(m.Name)
	}
}

func CommandEmailReadMailbox(s []string) {
	var mb string
	if len(s) < 1 {
		mb = "INBOX"
	} else {
		mb = s[0]
	}

	messages, err := Client.Email.ReadMailbox(mb, 50)
	if err != nil {
		Die("Cannot read messages: %s", err.Error())
	}

	for _, m := range messages {
		fmt.Printf("'%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
	}
}

func CommandEmailSendMail(s []string) {
	if len(s) > 3 {
		Die("Not enough arguments")
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
		Die("Error sending mail: %s", err.Error())
	}
}
