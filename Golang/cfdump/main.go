package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/cloudfoundry-community/go-cfclient"
)

func main() {

	targetAPI := flag.String("api", "", "API URL")
	userName := flag.String("user", "", "Username")
	Password := flag.String("password", "", "Password")
	flag.Parse()
	
	if *targetAPI == "" {
		fmt.Println("\n\tERROR| Enter the target API URL.")
		fmt.Println("\tINFO| Usage: ./cfdump --api https://api.hostname.org --user admin --password changeit\n")
		os.Exit(1)
	}
	if *userName == "" {
		fmt.Println("\n\tERROR| Enter username for the target: ", *targetAPI)
		fmt.Println("\tINFO| Usage: ./cfdump --api https://api.hostname.org --user admin --password changeit\n")
		os.Exit(1)
	}
	if *Password == "" {
		fmt.Println("\n\tERROR| Enter password for the target: ", *targetAPI)
		fmt.Println("\tINFO| Usage: ./cfdump --api https://api.hostname.org --user admin --password changeit\n")
		os.Exit(1)
	}

	c := &cfclient.Config{
		ApiAddress: *targetAPI,
		Username:   *userName,
		Password:   *Password,
	}
	client, _ := cfclient.NewClient(c)
	apps, _ := client.ListApps()
	fmt.Println(apps)
}