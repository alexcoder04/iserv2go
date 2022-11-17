
# iserv2go

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
    client := iserv.IServClient{}

    // login your client
    err := client.Login(&types.IServAccountConfig{
		IServHost: os.Getenv("ISERV_HOST"),
		Username:  os.Getenv("ISERV_USERNAME"),
		Password:  os.Getenv("ISERV_PASSWORD"),
	}, &types.IServClientOptions{
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
