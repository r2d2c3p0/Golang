package main

import (
	"flag"
	"fmt"
)

func main() {
	targetAPI := flag.String("api", "https://api.hostname.org", "API URL")
	userName := flag.String("user", "admin", "Username")
	Password := flag.String("password", "changeit", "Password")
	flag.Parse()
	c := &cfclient.Config{
		ApiAddress: *targetAPI,
		Username:   *userName,
		Password:   *Password,
	}
	client, _ := cfclient.NewClient(c)
	apps, _ := client.ListApps()
	fmt.Println(apps)
}
