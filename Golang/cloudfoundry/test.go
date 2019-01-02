package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	targetAPI := flag.String("api", "", "API URL")
	userName := flag.String("user", "", "Username")
	Password := flag.String("password", "", "Password")
	flag.Parse()
	if *targetAPI == "" {
		fmt.Println("ERROR| Enter the target API URL.")
		fmt.Println("INFO| Usage: ./cfdump --api https://api.hostname.org --user admin --password changeit")
		os.Exit(1)
	}
	if *userName == "" {
		fmt.Println("ERROR| Enter username for the target .", *targetAPI)
		fmt.Println("INFO| Usage: ./cfdump --api https://api.hostname.org --user admin --password changeit")
		os.Exit(1)
	}
	if *Password == "" {
		fmt.Println("ERROR| Enter password for the target .", *targetAPI)
		fmt.Println("INFO| Usage: ./cfdump --api https://api.hostname.org --user admin --password changeit")
		os.Exit(1)
	}
	fmt.Println("api is ", *targetAPI)
	fmt.Println("user is ", *userName)
	fmt.Println("password is ", *Password)

}
