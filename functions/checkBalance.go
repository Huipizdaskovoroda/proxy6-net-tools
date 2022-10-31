package functions

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
   	err := godotenv.Load()
	if err != nil {
		panic("Failed to load ENV")
	}
}

func CheckBalance() {
	api_token := os.Getenv("API_TOKEN") // Get the API token from the ENV file

	resp, err := http.Get("https://proxy6.net/api/" + api_token)
	if err != nil {
		log.Fatal("Error while sending request to the API", err)
	} // Send a request to the API
	defer resp.Body.Close() // Close the connection
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	} // Read the response
	fmt.Println(string(body)) // Print the response

	var result map[string]interface{}                                                  // Create a map to store the response
	json.Unmarshal([]byte(body), &result)                                              // Unmarshal the response
	fmt.Printf("Your account balance is %v %v", result["balance"], result["currency"]) // Print the balance
}
