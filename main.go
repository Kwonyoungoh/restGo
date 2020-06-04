package main

import (
	"net/http"

	"testrest.com/testRest/myapp"
)

func main() {

	http.ListenAndServe(":2000", myapp.NewHttpHandler())
	
}