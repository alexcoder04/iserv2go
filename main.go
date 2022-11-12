package main

import (
	"fmt"
	"os"

	"github.com/alexcoder04/iserv2go/iserv"
	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

// include in friendly
func Die(message string, args ...any) {
	fmt.Printf("Fatal error: "+message+"\n", args...)
	os.Exit(1)
}

func Warn(message string, args ...any) {
	fmt.Printf("Warning: "+message+"\n", args...)
}

func main() {
	client := iserv.IServClient{}

	err := client.Login(&types.IServAccountConfig{
		IServHost: os.Getenv("ISERV_HOST"),
		Username:  os.Getenv("ISERV_USERNAME"),
		Password:  os.Getenv("ISERV_PASSWORD"),
	}, &types.IServClientOptions{
		EnableWeb:   true,
		EnableEmail: true,
		EnableFiles: true,
	})
	if err != nil {
		Die("Cannot login: %s", err.Error())
	}
	defer client.Logout()

	// web
	badges, err := client.WebClient.GetBadges()
	if err != nil {
		Warn("Cannot load badges: %s", err.Error())
	}
	fmt.Println("Badges:")
	for key, value := range badges {
		fmt.Printf("%s: %d\n", key, value)
	}

	events, err := client.WebClient.GetUpcomingEvents()
	if err != nil {
		Warn("Cannot load upcoming events: %s", err.Error())
	}
	fmt.Println("Events:")
	for _, e := range events.Events {
		fmt.Printf("%s on %s\n", e.Title, e.When)
	}

	// email
	mailboxes, err := client.EmailClient.ListMailboxes()
	if err != nil {
		Warn("Cannot load email mailboxes: %s", err.Error())
	}
	for _, m := range mailboxes {
		fmt.Printf(" * %s\n", m.Name)
	}

	messages, err := client.EmailClient.ReadMailbox("INBOX", 10)
	if err != nil {
		Warn("Cannot read messages: %s", err.Error())
	}
	for _, m := range messages {
		fmt.Printf(" = '%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
	}

	// files
	files, err := client.FilesClient.ReadDir("/Groups")
	if err != nil {
		Warn("Cannot read groups: %s", err.Error())
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
