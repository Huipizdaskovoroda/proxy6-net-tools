package functions

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func CheckProxies() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load ENV")
	} // Load the ENV file

	api_token := os.Getenv("API_TOKEN") // Get the API token from the ENV file

	resp, err := http.Get("https://proxy6.net/api/" + api_token + "/getproxy")
	if err != nil {
		log.Fatal("Error while sending request to the API", err)
	} // Send a request to the API
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	} // Read the response
	fmt.Println(string(body)) // Print the response
}
