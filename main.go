package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexcoder04/iserv2go/iserv"
	"github.com/alexcoder04/iserv2go/iserv/iutils"
	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/joho/godotenv"
)

var (
	EnableEmail *bool = flag.Bool("enable-email", false, "whether to enable email module")
	EnableFiles *bool = flag.Bool("enable-files", false, "whether to enable files module")
	EnableWeb   *bool = flag.Bool("enable-web", false, "whether to enable web module")
)

func init() {
	godotenv.Load()
	flag.Parse()
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
		Password:  os.Getenv("ISERV_PASSWORD"),
		Username:  os.Getenv("ISERV_USERNAME"),
	}, &types.IServClientOptions{
		EnableModules: map[string]bool{
			"email": *EnableEmail,
			"files": *EnableFiles,
			"web":   *EnableWeb,
		},
	})
	if err != nil {
		Die("Cannot login: %s", err.Error())
	}
	defer client.Logout()

	// web
	if *EnableWeb {
		badges, err := client.Web.GetBadges()
		if err != nil {
			Warn("Cannot load badges: %s", err.Error())
		} else {
			fmt.Println("Badges:")
			for key, value := range badges {
				fmt.Printf("%s: %d\n", key, value)
			}
		}

		events, err := client.Web.GetUpcomingEvents(14)
		if err != nil {
			Warn("Cannot load upcoming events: %s", err.Error())
		} else {
			fmt.Printf("Events (%d):\n", len(events.Events))
			for _, e := range events.Events {
				fmt.Printf("%s on %s\n", e.Title, e.When)
			}
		}
	}

	// email
	if *EnableEmail {
		mailboxes, err := client.Email.ListMailboxes()
		if err != nil {
			Warn("Cannot load email mailboxes: %s", err.Error())
		}
		for _, m := range mailboxes {
			fmt.Printf(" * %s\n", m.Name)
		}

		messages, err := client.Email.ReadMailbox("INBOX", 10)
		if err != nil {
			Warn("Cannot read messages: %s", err.Error())
		}
		for _, m := range messages {
			fmt.Printf(" = '%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
		}

		myMail := fmt.Sprintf("%s@%s", os.Getenv("ISERV_USERNAME"), os.Getenv("ISERV_HOST"))
		m := types.EMail{
			Subject:    "Hello World",
			From:       myMail,
			To:         myMail,
			ToDispName: iutils.GetNameFromAddress(myMail),
			CCs:        []string{},
			Body:       "Hello World, it's me via iserv2go!",
		}
		err = client.Email.SendMail(m)
		if err != nil {
			Warn("Cannot send mail: %s", err.Error())
		}
		fmt.Println("Sent test message to myself!")
	}

	// files
	if *EnableFiles {
		files, err := client.Files.ReadDir("/Groups")
		if err != nil {
			Warn("Cannot read groups: %s", err.Error())
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}
}
