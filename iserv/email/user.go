package email

import "github.com/emersion/go-imap"

func (c *IServEmailClient) ListMailboxes() ([]imap.MailboxInfo, error) {
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.imapClient.List("", "*", mailboxes)
	}()

	res := []imap.MailboxInfo{}
	for m := range mailboxes {
		res = append(res, *m)
	}

	if err := <-done; err != nil {
		return []imap.MailboxInfo{}, err
	}

	return res, nil
}

func (c *IServEmailClient) ReadMailbox(name string, limit uint32) ([]imap.Message, error) {
	mbox, err := c.imapClient.Select(name, false)
	if err != nil {
		return []imap.Message{}, err
	}

	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > (limit - 1) {
		from = mbox.Messages - limit
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.imapClient.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	res := []imap.Message{}
	for msg := range messages {
		res = append(res, *msg)
	}

	if err := <-done; err != nil {
		return []imap.Message{}, err
	}

	return res, nil
}
