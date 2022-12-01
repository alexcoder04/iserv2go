package main

import (
	"strings"

	"github.com/alexcoder04/iserv2go/iserv"
)

var CommandsMap map[string]func(*iserv.Client, []string) = map[string]func(*iserv.Client, []string){
	"email.list_mailboxes": iScriptFunctionEmailListMailboxes,
	"email.read_mailbox":   iScriptFunctionEmailReadMailbox,
	"email.send_mail":      iScriptFunctionEmailSendMail,
}

func runLine(client iserv.Client, line string) {
	parts := strings.Split(line, "|")

	var command string
	var args []string

	switch len(parts) {
	case 0:
		return
	case 1:
		command = parts[0]
		args = []string{}
	case 2:
		command = parts[0]
		args = strings.Split(parts[1], " ")
	}

	if _, ok := CommandsMap[command]; ok {
		CommandsMap[command](&client, args)
	} else {
		Warn("Command '%s' not found", command)
	}
}
