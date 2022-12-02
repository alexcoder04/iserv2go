
# iserv2go

[![License](https://img.shields.io/github/license/alexcoder04/iserv2go)](https://github.com/alexcoder04/iserv2go/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/alexcoder04/iserv2go)](https://github.com/alexcoder04/iserv2go/blob/main/go.mod)
[![Lines](https://img.shields.io/tokei/lines/github/alexcoder04/iserv2go?label=lines)](https://github.com/alexcoder04/iserv2go/pulse)
[![Release](https://img.shields.io/github/v/release/alexcoder04/iserv2go?display_name=tag&sort=semver)](https://github.com/alexcoder04/iserv2go/releases/latest)
[![Stars](https://img.shields.io/github/stars/alexcoder04/iserv2go)](https://github.com/alexcoder04/iserv2go/stargazers)
[![Contributors](https://img.shields.io/github/contributors-anon/alexcoder04/iserv2go)](https://github.com/alexcoder04/iserv2go/graphs/contributors)


Unofficial IServ Go library and CLI.

**Disclaimer 1**: I am **not** affiliated with the [IServ GmbH](https://iserv.eu/) in any way.

**Disclaimer 2**: This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the [GNU Affero General Public License](./LICENSE) for more details.

**Disclaimer 3**: Use it at YOUR OWN RISK!

## Use as CLI

**Work in progress**

`main.go` rather serves as a demo right now, however, it will be turned into a proper CLI app in the future.

## Use as Library

```sh
go get github.com/alexcoder04/iserv2go/iserv # in your project directory
```

### Example usage

```go
package main

import (
    "github.com/alexcoder04/iserv2go/iserv"
    "github.com/alexcoder04/iserv2go/iserv/types"
)

func main(){
    // create new client instance
    client := iserv.Client{}

    // login your client
    err := client.Login(&types.AccountConfig{
		IServHost: os.Getenv("ISERV_HOST"),
		Username:  os.Getenv("ISERV_USERNAME"),
		Password:  os.Getenv("ISERV_PASSWORD"),
	}, &types.ClientOptions{
		EnableModules: map[string]bool{
			"email": true,
			"files": false,
			"web":   false,
		},
	})
    if err != nil {
        panic("failed to login")
    }

    // don't forget to logout
    defer client.Logout()

    // get mails in INBOX
    messages, err := client.Email.ReadMailbox("INBOX", 10)
    if err != nil {
        return
    }
    // print them
    for _, m := range messages {
        fmt.Printf(" = '%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
    }
}
```

## Project Structure

The `iserv` folder contains the Go Library, the subfolders `email`, `files`, `web` are modules, which can be (de-)activated separately.
They contain each `user.go` files, which include all the functions meant to be used by end-user.

## IScript

IScript is a way to use the IServ API. Currently, it is used as an argument to the CLI:

```sh
iserv2go -enable-email -e "email.read_mailboxes"
```

### Syntax

```text
group.command|argument1 argument2 ...
```

### List of commands

|Command|Args|Description|
|---|---|---|
|`email.list_mailboxes`|none|get a list of mailboxes|
|`email.read_mailbox`|`mailbox path`|get last 50 messages from mailbox|
|`email.send_mail`|`recipient address`, `subject`, `body`|send email|
|`web.get_badges`|none|get badges (for modules on the nav bar left)|
|`web.get_notifications`|none|get unread notifications|
|`web.get_upcoming_events`|none|list of upcoming events|
|`web.get_current_exercises`|none|list of current exercises|
|`web.get_past_exercises`|none|list of past exercises|
