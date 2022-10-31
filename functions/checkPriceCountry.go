package functions

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func CheckCountCountry() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load ENV")
	} // Load the ENV file

	var country string
	fmt.Println("Enter the country code (e.g. US, RU, DE, etc.)") // Ask the user to enter the country code

	fmt.Scanln(&country)

	country = strings.ToLower(country) // Convert the country code to lowercase, to avoid errors

	api_token := os.Getenv("API_TOKEN") // Get the API token from the ENV file

	resp, err := http.Get("https://proxy6.net/api/" + api_token + "/getcount?country=" + country)
	if err != nil {
		log.Fatal("Error while sending request to the API", err)
	} // Send a request to the API

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	} // Read the response

	var result map[string]interface{}     // Create a map to store the response
	json.Unmarshal([]byte(body), &result) // Unmarshal the response
	fmt.Println(result["count"])          // Print the response
	fmt.Println("Please note that if you specify the wrong country, you will not get the answer you need")

}
