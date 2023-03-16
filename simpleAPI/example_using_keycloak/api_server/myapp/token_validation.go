package myapp

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	IntrospectURL = "http://localhost:8080/auth/realms/myrealm/protocol/openid-connect/token/introspect"
)

func ValidateToken(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusBadRequest)
			return
		}

		token := authHeader[len("Bearer "):]

		client := &http.Client{}
		req, err := http.NewRequest("POST", IntrospectURL, nil)
		if err != nil {
			http.Error(w, "Failed to create introspection request", http.StatusInternalServerError)
			return
		}

		req.SetBasicAuth("client_id", "client_secret")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		q := req.URL.Query()
		q.Add("token", token)
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to introspect token", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		body := make(map[string]interface{})
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			http.Error(w, "Failed to decode introspection response", http.StatusInternalServerError)
			return
		}

		active, ok := body["active"].(bool)
		if !ok || !active {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		exp, ok := body["exp"].(float64)
		if !ok {
			http.Error(w, "Failed to parse token expiration", http.StatusInternalServerError)
			return
		}

		if int64(exp) < time.Now().Unix() {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
			return
		}

		inner.ServeHTTP(w, r)

	})
}
