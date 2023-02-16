package main

import (
	"net/http"
	"testAPI/myapp"
)

func main() {
	http.ListenAndServe(":80", myapp.NewHandler())
}
