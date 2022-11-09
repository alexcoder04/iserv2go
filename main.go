package main

import (
	"fmt"
	"os"

	"github.com/alexcoder04/iserv2go/iserv"
)

func main() {
	client := iserv.IServClient{}

	err := client.Login(&iserv.IServAccountConfig{
		IServURL: os.Getenv("ISERV_URL"),
		Username: os.Getenv("ISERV_USERNAME"),
		Password: os.Getenv("ISERV_PASSWORD"),
	})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	badges, err := client.GetBadges()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	for key, value := range badges {
		fmt.Printf("%s: %d\n", key, value)
	}
}
