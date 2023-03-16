package myapp

// 필요한 패키지를 import
import (
	"encoding/json" // JSON 디코딩 패키지
	"net/http"      // HTTP 패키지
	"time"          // 시간 관련 패키지
)

// IntrospectURL
// introspection endpoint URL 상수 선언
const (
	IntrospectURL = "http://localhost:8080/auth/realms/myrealm/protocol/openid-connect/token/introspect"
)

// ValidateToken
// ValidateToken은 HTTP 요청에 대한 토큰 유효성을 검사하는 미들웨어를 반환합니다.
// 요청이 내부 핸들러에 전달되기 전에 토큰이 유효한지 확인합니다.
func ValidateToken(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청 헤더에서 Authorization 헤더에서 Bearer 토큰을 추출합니다.
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusBadRequest)
			return
		}
		token := authHeader[len("Bearer "):]

		// 토큰 유효성 검사를 위해 토큰을 Introspection Endpoint로 전송합니다.
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

		// Introspection Endpoint가 유효한 토큰을 반환하지 않으면 요청을 거부합니다.
		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Introspection Endpoint가 유효한 토큰을 반환하면, 토큰 유효성을 확인합니다.
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

		// 토큰의 만료 시간이 현재 시간보다 빠르면 요청을 거부합니다.
		if int64(exp) < time.Now().Unix() {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
			return
		}

		// 내부 HTTP 핸들러를 실행합니다.
		inner.ServeHTTP(w, r)

	})
}
