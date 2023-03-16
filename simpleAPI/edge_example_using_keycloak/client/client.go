// Import packages
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2" // Import oauth2 package
)

// Define API endpoint
const apiEndpoint = "https://api.camara.org/edge-site-selection-and-routing"

// Define Keycloak server endpoint
const keycloakEndpoint = "https://keycloak.camara.org/auth/realms/camara/protocol/openid-connect/token"

// Define OAuth 2.0 configuration
var oauthConfig = &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	Scopes:       []string{"openid", "profile", "email"},
	Endpoint: oauth2.Endpoint{
		TokenURL: keycloakEndpoint,
	},
}

// Define query parameters
var queryParams = map[string]string{
	"country":  "KR",
	"operator": "KT",
	"service":  "video-streaming",
}

// Main function
func main() {
	// Create a new HTTP client with OAuth 2.0 transport
	ctx := context.Background()

	// Get access token from Keycloak server using client credentials grant type
	token, err := oauthConfig.Exchange(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print access token for debugging purpose only (do not expose it to public)
	fmt.Println("Access token:", token.AccessToken)

	client := oauthConfig.Client(ctx, token)

	// Create a new HTTP request with query parameters and authentication token
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v) // Add each query parameter to the request URL
	}
	req.URL.RawQuery = q.Encode() // Encode the query parameters to URL format

	// Send HTTP request to API endpoint and get response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response body and print it
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body)) // Print the response body as a string

}
