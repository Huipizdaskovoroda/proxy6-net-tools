package functions

import (
	"fmt"
)

func Help() {
	fmt.Println("This script is used to check your account balance, check your account ID, check your current proxies list, check the price of a specific proxy, and get help for working with this script.")
	fmt.Println("Please note, before starting to use this script, you need to create a .env file in the same directory as the script and add your API token to API_TOKEN variable. You can get your API token from your account page on proxy6.net - https://proxy6.net/en/user/developers")
	fmt.Println("To check your account balance, enter 1")
	fmt.Println("To check your account ID, enter 2")
	fmt.Println("To check your current proxies list, enter 3")
	fmt.Println("To check the price of a specific proxy, enter 4. Please note, if you specify the wrong country, you will not get the answer you need")
	fmt.Println("To exit the program, enter 6") // Print the help message
}
