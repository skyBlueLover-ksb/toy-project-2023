package main

import (
	"net/http"
	"simpleAPI/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())

}
