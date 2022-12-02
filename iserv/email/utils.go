package email

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/alexcoder04/friendly/v2"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func getToFieldString(addr string, dispName string, ccs []string) string {
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
			"From: " + friendly.GetFullNameFromMailAddress(mail.From) + "<" + mail.From + ">",
			getToFieldString(mail.To, mail.ToDispName, mail.CCs),
			"Subject: " + mail.Subject,
			"Content-Type: text/plain; charset=utf-8",
			"Content-Transfer-Encoding: base64",
			"\r\n",
			base64.StdEncoding.EncodeToString([]byte(mail.Body)),
		}, "\r\n"),
	)
}
