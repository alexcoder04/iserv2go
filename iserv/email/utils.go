package email

import (
	"fmt"
	"strings"

	"github.com/alexcoder04/iserv2go/iserv/types"
)

func getToString(addr string, dispName string, ccs []string) string {
	var res string

	if dispName == "" {
		res = fmt.Sprintf("To: %s", addr)
	} else {
		res = fmt.Sprintf("To: %s <%s>", dispName, addr)
	}

	if len(ccs) > 0 {
		res += fmt.Sprintf("\r\nCc: %s", strings.Join(ccs, ", "))
	}

	return res
}

func buildMailBody(mail types.EMail) []byte {
	return []byte(
		strings.Join([]string{
			"MIME-version: 1.0",
			"Content-Type: text/plain; charset=utf-8",
			"From: " + mail.From,
			getToString(mail.To, mail.ToDispName, mail.CCs),
			"Subject: " + mail.Subject,
			"\r\n",
			mail.Body,
		}, "\r\n"),
	)
}
