package main

import (
	"fmt"
	"proxynet/functions"
)

func main() {
	var action int                                                                                                                                                                                                   // Declare a variable to store the user's choice
	fmt.Println("Welcome, traveler! This is a simple script made by Quanor to connect and retrieve information from proxy6.net!\nWhat would you like to do?\n1. Check account balance\n2. Check account's id\n3. Check your current proxies list\n4. Check the price of a specific proxy\n5. Get help for working with this script\n6. Exit") // Ask the user to choose an action
	fmt.Scanln(&action)
	// switch statement to choose the action forever until the user enters 0
	switch action {
	case 1:
		functions.CheckBalance() // Check the account balance
	case 2:
		functions.CheckID() // Check the account ID
	case 3:
		functions.CheckProxies() // Check the current proxies list
	case 4:
		functions.CheckCountCountry() // Check the price of a specific proxy
	case 5:
		functions.Help() // Get help for working with this script
	case 6:
		fmt.Println("Goodbye!") // Exit the program
		return
	}
}
