package main

import (
	"net/http"
)

func main() {

	http.ListenAndServe(":2000", myapp.NewHTTPHandler)
	
}