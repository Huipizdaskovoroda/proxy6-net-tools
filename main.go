package main

import (
	"fmt"
	"proxynet/functions"
    "os"
    "os/exec"
    "runtime"
)

func clear(ostype string) {
    switch ostype {
        case "linux": {
            cmd := exec.Command("clear") //Linux example, its tested
            cmd.Stdout = os.Stdout
            cmd.Run()
        }
        case "windows": {
            cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
            cmd.Stdout = os.Stdout
            cmd.Run()
        }
    }
}

func main() {
	var action int                                                                                                                                                                                                   // Declare a variable to store the user's choice
	fmt.Println("Welcome, traveler! This is a simple script made by Quanor to connect and retrieve information from proxy6.net!\nWhat would you like to do?\n1. Check account balance\n2. Check account's id\n3. Check your current proxies list\n4. Check the price of a specific proxy\n5. Get help for working with this script\n6. Exit") // Ask the user to choose an action
	// switch statement to choose the action forever until the user enters 0
    for { 
    	fmt.Scanln(&action)
        clear(runtime.GOOS)
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
}
