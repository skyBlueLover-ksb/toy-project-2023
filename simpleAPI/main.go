package main

import (
	"net/http"
	"testAPI/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())

}
