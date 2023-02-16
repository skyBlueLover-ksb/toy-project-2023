package myapp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"sangbeom", "last_name":"kim","email":"sb.kim@kt.com" }`))

	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
}
