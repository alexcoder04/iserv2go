package main

import (
	"fmt"
	"os"

	"github.com/alexcoder04/iserv2go/iserv"
)

func main() {
	client := iserv.IServClient{}

	err := client.Login(&iserv.IServAccountConfig{
		IServHost: os.Getenv("ISERV_HOST"),
		Username:  os.Getenv("ISERV_USERNAME"),
		Password:  os.Getenv("ISERV_PASSWORD"),
	}, &iserv.IServClientOptions{
		EnableWeb:   true,
		EnableEmail: true,
		EnableFiles: true,
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	defer client.Logout()

	// web
	badges, err := client.WebClient.GetBadges()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	for key, value := range badges {
		fmt.Printf("%s: %d\n", key, value)
	}

	// email
	mailboxes, err := client.EmailClient.ListMailboxes()
	if err != nil {
		return
	}
	for _, m := range mailboxes {
		fmt.Printf(" * %s\n", m.Name)
	}

	messages, err := client.EmailClient.ReadMailbox("INBOX/Wettbewerbe", 10)
	if err != nil {
		return
	}
	for _, m := range messages {
		fmt.Printf(" = '%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
	}

	// files
	files, err := client.FilesClient.ReadDir("/Groups")
	if err != nil {
		return
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
