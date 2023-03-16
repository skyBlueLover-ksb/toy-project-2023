package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

// Define constants for Keycloak server URL and credentials
const (
	keycloakURL    = "https://YOUR_KEYCLOAK_SERVER/auth/realms/YOUR_REALM/protocol/openid-connect/token/introspect"
	keycloakClient = "YOUR_CLIENT_ID"
	keycloakSecret = "YOUR_CLIENT_SECRET"
)

// Define edge site data (mock data for demonstration purpose only)
var edgeSiteData = []map[string]interface{}{
	{
		// ...
	},
}

// Define a struct to hold the introspection response data
type IntrospectionResponse struct {
	Active bool     `json:"active"`
	Roles  []string `json:"roles"` // Add a field to store roles from token claim
}

// Define a function to verify Bearer token using introspection
func verifyBearerToken(token string) (bool, []string, error) { // Modify return type to include roles
	// Create a HTTP client
	client := &http.Client{}

	// Prepare request body with required parameters
	body := fmt.Sprintf("client_id=%s&client_secret=%s&token=%s", keycloakClient, keycloakSecret, token)
	reqBody := bytes.NewBuffer([]byte(body))

	// Create a HTTP POST request to Keycloak server's introspect endpoint
	req, err := http.NewRequest("POST", keycloakURL, reqBody)
	if err != nil {
		return false, nil, err // Modify return value to include roles
	}

	// Set request header with content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send request and get response
	resp, err := client.Do(req)
	if err != nil {
		return false, nil, err // Modify return value to include roles
	}
	defer resp.Body.Close()

	// Read response body as bytes
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err // Modify return value to include roles
	}

	// Decode response body as JSON into IntrospectionResponse struct
	var respData IntrospectionResponse
	err = json.Unmarshal(respBodyBytes, &respData)
	if err != nil {
		return false, nil, err // Modify return value to include roles
	}

	// Return active field value and roles as verification result
	return respData.Active, respData.Roles, nil // Modify return value to include roles

}

// Define a handler function to verify authentication token and return edge site data using introspection logic
func handleEdgeSiteSelectionAndRouting(w http.ResponseWriter, r *http.Request) {
	// Get authentication token from request header
	token := r.Header.Get("Authorization")

	// Check if token is empty or not prefixed with "Bearer "
	if token == "" || len(token) < 7 || token[:7] != "Bearer " {
		// Return 401 Unauthorized error if token is invalid
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// Verify Bearer token using introspection logic
	tokenString := token[7:]
	isValidToken, roles, err := verifyBearerToken(tokenString) // Get roles from verification result

	// Check if there is an error during token verification
	if err != nil {
		// Return 500 Internal Server Error if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	// Check if token is valid
	if !isValidToken {
		// Return 401 Unauthorized error if token is not valid
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// Check if the user has the required role to access edge site data
	if !contains(roles, "edgeSiteAdmin") {
		// Return 403 Forbidden error if user does not have the required role
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
		return
	}

	// Marshal edge site data into JSON format
	edgeSiteDataBytes, err := json.Marshal(edgeSiteData)
	if err != nil {
		// Return 500 Internal Server Error if there is an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	// Write edge site data as response
	w.Header().Set("Content-Type", "application/json")
	w.Write(edgeSiteDataBytes)
}

// Define a helper function to check if a slice of strings contains a specific string
func contains(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

// Define the main function to start the server
func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter().StrictSlash(true)

	// Register the edge site selection and routing handler function to the router
	r.HandleFunc("/edge-site", handleEdgeSiteSelectionAndRouting)

	// Start the server and log any errors
	log.Fatal(http.ListenAndServe(":8080", r))
}
