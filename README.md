
# iserv2go

Unofficial IServ Go library and CLI.

**Disclaimer 1**: I am not affiliated with the [IServ GmbH](https://iserv.eu/) in any way.

**Disclaimer 2**: This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

## 1. Use as CLI

**Work in progress**

## 2. Use as Library

```sh
go get github.com/alexcoder04/iserv2go/iserv
```

```go
package main

import "github.com/alexcoder04/iserv2go/iserv"

func main(){
    client := iserv.IServClient{}

    err := client.Login(&iserv.IServAccountConfig{
		IServHost: os.Getenv("ISERV_HOST"),
		Username:  os.Getenv("ISERV_USERNAME"),
		Password:  os.Getenv("ISERV_PASSWORD"),
	}, &iserv.IServClientOptions{
		EnableWeb:   false,
		EnableEmail: true,
		EnableFiles: false,
	})
    if err != nil {
        panic("failed to login")
    }
    defer client.Logout()

    messages, err := client.EmailClient.ReadMailbox("INBOX", 10)
    if err != nil {
        return
    }
    for _, m := range messages {
        fmt.Printf(" = '%s' from %s\n", m.Envelope.Subject, m.Envelope.Sender[0].Address())
    }
}
```
