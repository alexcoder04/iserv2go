package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/iserv2go/iserv"
	"github.com/alexcoder04/iserv2go/iserv/types"
	"github.com/joho/godotenv"
)

var (
	VERSION    string = "unknown"
	COMMIT_SHA string = "unknown"

	Info         *bool = flag.Bool("info", false, "show program info")
	EnableEmail  *bool = flag.Bool("enable-email", false, "enable email module")
	EnableFiles  *bool = flag.Bool("enable-files", false, "enable files module")
	EnableWeb    *bool = flag.Bool("enable-web", false, "enable web module")
	SaveSessions *bool = flag.Bool("save-sessions", false, "save login credentials on disk for subsequent logins")
	Interactive  *bool = flag.Bool("interactive", false, "interactive session")

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

func RunCommand(cmd string, args []string) error {
	if _, ok := CommandsMap[cmd]; ok {
		CommandsMap[cmd](args)
		return nil
	}
	return fmt.Errorf("command not found: '%s'", cmd)
}

func main() {
	if *Info {
		fmt.Printf("iserv2go %s (commit %s)\n", VERSION, COMMIT_SHA)
		return
	}

	Client = iserv.Client{}

	err := Client.Login(&types.ClientConfig{
		IServHost: os.Getenv("ISERV_HOST"),
		Password:  os.Getenv("ISERV_PASSWORD"),
		Username:  os.Getenv("ISERV_USERNAME"),

		EnableModules: map[string]bool{
			"email": *EnableEmail,
			"files": *EnableFiles,
			"web":   *EnableWeb,
		},
		SaveSessions: *SaveSessions,
	})
	if err != nil {
		friendly.Die("Cannot login: %s", err.Error())
	}
	defer Client.Logout()

	if *Interactive {
		for {
			inp, err := friendly.Input("iserv> ")
			if err != nil {
				friendly.Die("Failed to read command: %s", err.Error())
			}
			cmdline := strings.Split(strings.TrimSpace(inp), " ")
			if cmdline[0] == "exit" {
				return
			}
			err = RunCommand(cmdline[0], cmdline[1:])
			if err != nil {
				friendly.Warn("Command '%s' not found", cmdline[0])
			}
		}
	}

	if Args[0] != "" {
		err := RunCommand(Args[0], Args[1:])
		if err != nil {
			friendly.Die("Command '%s' not found", Args[0])
		}
	}
}
