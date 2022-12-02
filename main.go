package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/iserv2go/iserv"
	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/joho/godotenv"
)

var (
	VERSION    string = "unknown"
	COMMIT_SHA string = "unknown"

	Info        *bool = flag.Bool("info", false, "show program info")
	EnableEmail *bool = flag.Bool("enable-email", false, "whether to enable email module")
	EnableFiles *bool = flag.Bool("enable-files", false, "whether to enable files module")
	EnableWeb   *bool = flag.Bool("enable-web", false, "whether to enable web module")

	Args []string

	Client iserv.Client
)

var CommandsMap map[string]func([]string) = map[string]func([]string){
	"email.list_mailboxes":      CommandEmailListMailboxes,
	"email.read_mailbox":        CommandEmailReadMailbox,
	"email.send_mail":           CommandEmailSendMail,
	"web.get_badges":            CommandWebGetBadges,
	"web.get_notifications":     CommandWebGetNotifications,
	"web.get_upcoming_events":   CommandWebGetUpcomingEvents,
	"web.get_current_exercises": CommandWebGetCurrentExercises,
	"web.get_past_exercises":    CommandWebGetPastExercises,
	"files.ls":                  CommandFilesLs,
	"files.cat":                 CommandFilesCat,
	"files.download":            CommandFilesDownload,
	"files.upload":              CommandFilesUpload,
}

func init() {
	godotenv.Load()
	flag.Parse()
	Args = flag.Args()
}

func main() {
	if *Info {
		fmt.Printf("iserv2go %s (commit %s)\n", VERSION, COMMIT_SHA)
		return
	}

	Client = iserv.Client{}

	err := Client.Login(&types.AccountConfig{
		IServHost: os.Getenv("ISERV_HOST"),
		Password:  os.Getenv("ISERV_PASSWORD"),
		Username:  os.Getenv("ISERV_USERNAME"),
	}, &types.ClientOptions{
		EnableModules: map[string]bool{
			"email": *EnableEmail,
			"files": *EnableFiles,
			"web":   *EnableWeb,
		},
	})
	if err != nil {
		friendly.Die("Cannot login: %s", err.Error())
	}
	defer Client.Logout()

	if Args[0] != "" {
		if _, ok := CommandsMap[Args[0]]; ok {
			CommandsMap[Args[0]](Args[1:])
		} else {
			friendly.Die("Command '%s' not found", Args[0])
		}
	}
}
